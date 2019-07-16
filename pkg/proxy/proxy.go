/*
Copyright 2018 Mirantis

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package proxy

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/ghodss/yaml"
	"github.com/golang/glog"
	digest "github.com/opencontainers/go-digest"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	criErrorLogLevel   = 2
	criRequestLogLevel = 3
	criNoisyLogLevel   = 4
	criListLogLevel    = 5
)

// RuntimeProxy is a gRPC implementation of internalapi.RuntimeService.
type RuntimeProxy struct {
	criVersion   CRIVersion
	streamUrl    url.URL
	conn         *grpc.ClientConn
	clients      []client
	methodPrefix string
	images       map[string]string
}

var _ Interceptor = &RuntimeProxy{}

type methodHandler func(r *RuntimeProxy, ctx context.Context, method string, req, resp CRIObject) (interface{}, error)

type dispatchItem struct {
	handler  methodHandler
	logLevel glog.Level
}

// NewRuntimeProxy creates a new internalapi.RuntimeService.
func NewRuntimeProxy(criVersion CRIVersion, addrs []string, connectionTimout time.Duration, streamUrl *url.URL) (*RuntimeProxy, error) {
	if len(addrs) == 0 {
		return nil, errors.New("no sockets specified to connect to")
	}

	r := &RuntimeProxy{
		criVersion:   criVersion,
		streamUrl:    *streamUrl,
		methodPrefix: fmt.Sprintf("/%s.", criVersion.ProtoPackage()),
		images:       make(map[string]string),
	}
	for _, addr := range addrs {
		r.clients = append(r.clients, newAutoClient(criVersion, addr, connectionTimout))
	}
	if !r.clients[0].isPrimary() {
		return nil, errors.New("the first client should be primary (no id)")
	}
	for _, client := range r.clients[1:] {
		if client.isPrimary() {
			return nil, errors.New("only the first client should be primary (no id)")
		}
	}

	return r, nil
}

// Register implements Register method of the Interceptor interface.
func (r *RuntimeProxy) Register(s *grpc.Server) {
	r.criVersion.Register(s)
}

// Stop implements Stop method of the Interceptor interface.
func (r *RuntimeProxy) Stop() {
	for _, client := range r.clients {
		client.stop()
	}
}

// Match implements Match method of the Interceptor interface.
func (r *RuntimeProxy) Match(fullMethod string) bool {
	lastDot := strings.LastIndex(fullMethod, ".")
	if lastDot < 0 {
		return false
	}
	return fullMethod[:lastDot+1] == r.methodPrefix
}

// Intercept implements Intercept method of the Interceptor interface.
func (r *RuntimeProxy) Intercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var err error
	defer func() {
		if err != nil {
			glog.V(criErrorLogLevel).Infof("FAIL: %s(): %v", info.FullMethod, err)
		}
	}()
	if !strings.HasPrefix(info.FullMethod, r.methodPrefix) {
		err = fmt.Errorf("bad method prefix in %q (expected to start with %q)", info.FullMethod, r.methodPrefix) // make it logged in defer
		return nil, err
	}

	method := info.FullMethod[len(r.methodPrefix):]
	dispatchItem, found := dispatchTable[method]
	if !found {
		err = fmt.Errorf("no handler for method %q", method) // make it logged in defer
		return nil, err
	}
	if glog.V(dispatchItem.logLevel) {
		glog.Infof("ENTER: %s():\n%s", info.FullMethod, dump(req))
	}
	wrappedReq, wrappedResp, err := r.criVersion.WrapObject(req)
	if err != nil {
		return nil, err
	}
	resp, err := dispatchItem.handler(r, ctx, info.FullMethod, wrappedReq, wrappedResp)
	if err != nil {
		return nil, err
	}
	if wrappedResp, ok := resp.(CRIObject); ok {
		resp = wrappedResp.Unwrap()
	}
	if glog.V(dispatchItem.logLevel) {
		glog.Infof("LEAVE: %s():\n%s", info.FullMethod, dump(resp))
	}
	return resp, nil
}

func (r *RuntimeProxy) getImageNameById(imageId string) string {
	return r.images[imageId]
}

func (r *RuntimeProxy) setImageNameById(imageId, imageName string, overwrite bool) {
	if _, ok := r.images[imageId]; !ok || overwrite {
		r.images[imageId] = imageName
	}
}

func (r *RuntimeProxy) deleteImageNameById(imageId string) {
	delete(r.images, imageId)
}

func (r *RuntimeProxy) primaryClient() (client, error) {
	if err := <-r.clients[0].connect(); err != nil {
		return nil, err
	}
	return r.clients[0], nil
}

func (r *RuntimeProxy) clientForAnnotations(annotations map[string]string) (client, error) {
	for _, client := range r.clients {
		if client.annotationsMatch(annotations) {
			if err := <-client.connect(); err != nil {
				return nil, err
			}
			return client, nil
		}
	}
	return nil, fmt.Errorf("criproxy: unknown runtime: %q", annotations[targetRuntimeAnnotationKey])
}

func (r *RuntimeProxy) clientForId(id string) (client, string, error) {
	client := r.clients[0]
	unprefixed := id
	for _, c := range r.clients[1:] {
		if ok, unpref := c.idPrefixMatches(id); ok {
			c.connect()
			if c.currentState() != clientStateConnected {
				return nil, "", fmt.Errorf("CRI proxy: target runtime is not available")
			}
			client = c
			unprefixed = unpref
			break
		}
	}
	if err := <-client.connect(); err != nil {
		return nil, "", err
	}
	return client, unprefixed, nil
}

func (r *RuntimeProxy) clientForImage(image string, noErrorIfNotConnected bool) (client, string, error) {
	client := r.clients[0]
	unprefixed := image
	for _, c := range r.clients[1:] {
		if ok, unpref := c.imageMatches(image); ok {
			c.connect()
			// don't wait for additional runtimes
			if c.currentState() != clientStateConnected {
				if noErrorIfNotConnected {
					return nil, "", nil
				}
				return nil, "", fmt.Errorf("CRI proxy: target runtime is not available")
			}
			client = c
			unprefixed = unpref
			break
		}
	}
	if err := <-client.connect(); err != nil {
		return nil, "", err
	}
	return client, unprefixed, nil
}

func (r *RuntimeProxy) fixStreamingUrl(url string) string {
	// The URLs provided by dockershim in k8s 1.11+ look like this:
	// //[::]:35057/cri/exec/tb8rgDBh
	// These can be passed as-is to the client because they
	// include the port.
	// In k8s 1.10-, the following URLs are passed:
	// /cri/exec/94B_NhGa
	// These need to be replaced to make exec/attach work with
	// dockershim.
	if strings.HasPrefix(url, "/") && !strings.Contains(url, ":") {
		u := r.streamUrl
		u.Path = url
		return u.String()
	}
	return url
}

func (r *RuntimeProxy) passToPrimary(ctx context.Context, method string, req, resp CRIObject) (interface{}, error) {
	client, err := r.primaryClient()
	if err != nil {
		return nil, err
	}
	return client.invokeWithErrorHandling(ctx, method, req, resp)
}

func (r *RuntimeProxy) updateRuntimeConfig(ctx context.Context, method string, req, resp CRIObject) (interface{}, error) {
	var errs []string
	for _, client := range r.clients {
		if client.currentState() != clientStateConnected {
			// This does nothing if the state is clientStateConnecting,
			// otherwise it tries to connect asynchronously
			client.connect()
			continue
		}

		_, err := client.invoke(ctx, method, req, resp)
		if err != nil {
			errs = append(errs, client.handleError(err, false).Error())
		}
	}

	if errs != nil {
		return nil, errors.New(strings.Join(errs, "\n"))
	}

	return resp, nil
}

func (r *RuntimeProxy) listObjects(ctx context.Context, method string, req, resp CRIObject) (interface{}, error) {
	out := resp.(ObjectList)
	clients := r.clients
	var singleClient client
	useSingleClient := false
	if in, ok := req.(IdFilterObject); ok && in.IdFilter() != "" {
		var unprefixed string
		var err error
		singleClient, unprefixed, err = r.clientForId(in.IdFilter())
		if err != nil {
			return nil, err
		}
		in.SetIdFilter(unprefixed)
		useSingleClient = true
	}

	if in, ok := req.(PodSandboxIdFilterObject); ok && in.PodSandboxIdFilter() != "" {
		anotherClient, unprefixed, err := r.clientForId(in.PodSandboxIdFilter())
		if err != nil {
			return nil, err
		}
		if anotherClient != nil {
			in.SetPodSandboxIdFilter(unprefixed)
			if singleClient == nil {
				singleClient = anotherClient
			} else if singleClient != anotherClient {
				// different id prefixes for sandbox & container
				out.SetItems(nil)
				return resp, nil
			}
		}
		useSingleClient = true
	}

	if in, ok := req.(ImageFilterObject); ok && in.ImageFilter() != "" {
		anotherClient, unprefixed, err := r.clientForImage(in.ImageFilter(), true)
		if err != nil {
			return nil, err
		}
		if anotherClient != nil {
			in.SetImageFilter(unprefixed)
			if singleClient == nil {
				singleClient = anotherClient
			} else if singleClient != anotherClient {
				// this should not really happen because list requests presently
				// don't filter by image and pod / container id at the same time,
				// but let's be sage here
				out.SetItems(nil)
				return resp, nil
			}
		}
		useSingleClient = true
	}

	if useSingleClient {
		if singleClient != nil {
			clients = []client{singleClient}
		} else {
			// The target client is offline
			out.SetItems(nil)
			return resp, nil
		}
	}

	var items []CRIObject
	for _, client := range clients {
		if client.currentState() != clientStateConnected {
			// This does nothing if the state is clientStateConnecting,
			// otherwise it tries to connect asynchronously
			client.connect()
			continue
		}

		out.SetItems(nil)
		_, err := client.invoke(ctx, method, req, resp)
		if err != nil {
			// if the runtime server is gone, let's just skip it
			err = client.handleError(err, true)
			if err != nil {
				// for more serious errors, log a warning but don't
				// block the other runtimes by making List* fail
				glog.Warningf("List request failed for runtime %q: %v", client.getID(), err)
			}
		}
		for _, item := range out.Items() {
			items = append(items, client.addPrefix(item))
		}
	}

	out.SetItems(items)
	return resp, nil

}

func (r *RuntimeProxy) invokePodSandboxMethod(ctx context.Context, method string, req, resp CRIObject) (client, error) {
	in := req.(PodSandboxIdObject)
	client, unprefixed, err := r.clientForId(in.PodSandboxId())
	if err != nil {
		return nil, err
	}
	in.SetPodSandboxId(unprefixed)
	_, err = client.invokeWithErrorHandling(ctx, method, req, resp)
	return client, err
}

func (r *RuntimeProxy) invokeContainerMethod(ctx context.Context, method string, req, resp CRIObject) (client, error) {
	in := req.(ContainerIdObject)
	client, unprefixed, err := r.clientForId(in.ContainerId())
	if err != nil {
		return nil, err
	}
	in.SetContainerId(unprefixed)

	_, err = client.invokeWithErrorHandling(ctx, method, req, resp)
	return client, err

}

func (r *RuntimeProxy) runPodSandbox(ctx context.Context, method string, req, resp CRIObject) (interface{}, error) {
	client, err := r.clientForAnnotations(req.(RunPodSandboxRequest).GetAnnotations())
	if err != nil {
		return nil, err
	}
	if _, err = client.invokeWithErrorHandling(ctx, method, req, resp); err == nil {
		out := resp.(RunPodSandboxResponse)
		out.SetPodSandboxId(client.augmentId(out.PodSandboxId()))
	}
	return resp, err
}

func (r *RuntimeProxy) handlePodSandbox(ctx context.Context, method string, req, resp CRIObject) (interface{}, error) {
	_, err := r.invokePodSandboxMethod(ctx, method, req, resp)
	if err == nil {
		if out, ok := resp.(UrlObject); ok {
			out.SetUrl(r.fixStreamingUrl(out.Url()))
		}
	}
	return resp, err
}

func (r *RuntimeProxy) podSandboxStatus(ctx context.Context, method string, req, resp CRIObject) (interface{}, error) {
	client, err := r.invokePodSandboxMethod(ctx, method, req, resp)
	if err != nil {
		return nil, err
	}
	if status := resp.(PodSandboxStatusResponse).Status(); status != nil {
		status.SetId(client.augmentId(status.Id()))
	}
	return resp, nil
}

func (r *RuntimeProxy) createContainer(ctx context.Context, method string, req, resp CRIObject) (interface{}, error) {
	in := req.(CreateContainerRequest)
	client, unprefixed, err := r.clientForId(in.PodSandboxId())
	if err != nil {
		return nil, err
	}
	in.SetPodSandboxId(unprefixed)

	if in.Image() == "" {
		return nil, errors.New("criproxy: no image specified")
	}

	// don't prefix image digests
	if _, err := digest.Parse(in.Image()); err != nil {
		imageClient, unprefixedImage, err := r.clientForImage(in.Image(), false)
		if err != nil {
			return nil, err
		}
		if imageClient != client {
			return nil, fmt.Errorf("criproxy: image %q is for a wrong runtime", in.Image())
		}
		in.SetImage(unprefixedImage)
	} else {
		// Image is a digest like
		// sha256:6a92cd1fcdc8d8cdec60f33dda4db2cb1fcdcacf3410a8e05b3741f44a9b5998.
		// Look up and set the name of the image instead, so the secondary runtime
		// can also use it.
		imageName := r.getImageNameById(in.Image())
		if imageName != "" {
			in.SetImage(imageName)
		}
	}

	_, err = client.invokeWithErrorHandling(ctx, method, req, resp)
	if err != nil {
		return nil, err
	}

	out := resp.(CreateContainerResponse)
	out.SetContainerId(client.augmentId(out.ContainerId()))
	return out, nil
}

func (r *RuntimeProxy) handleContainer(ctx context.Context, method string, req, resp CRIObject) (interface{}, error) {
	_, err := r.invokeContainerMethod(ctx, method, req, resp)
	if err == nil {
		if out, ok := resp.(UrlObject); ok {
			out.SetUrl(r.fixStreamingUrl(out.Url()))
		}
	}
	return resp, err
}

func (r *RuntimeProxy) containerStatus(ctx context.Context, method string, req, resp CRIObject) (interface{}, error) {
	client, err := r.invokeContainerMethod(ctx, method, req, resp)
	if err != nil {
		return nil, err
	}
	if status := resp.(ContainerStatusResponse).Status(); status != nil {
		status.SetId(client.augmentId(status.Id()))
		status.SetImage(client.imageName(status.Image()))
	}
	return resp, nil
}

func (r *RuntimeProxy) containerStats(ctx context.Context, method string, req, resp CRIObject) (interface{}, error) {
	client, err := r.invokeContainerMethod(ctx, method, req, resp)
	if err != nil {
		return nil, err
	}
	if stats := resp.(ContainerStatsResponse).Stats(); stats != nil {
		stats.SetId(client.augmentId(stats.Id()))
	}
	return resp, nil
}

func (r *RuntimeProxy) handleImage(ctx context.Context, method string, req, resp CRIObject) (interface{}, error) {
	in := req.(ImageObject)
	client, unprefixed, err := r.clientForImage(in.Image(), true)
	if client == nil {
		// the client is offline
		return resp, nil
	}
	in.SetImage(unprefixed)

	imageName := in.Image()

	_, err = client.invokeWithErrorHandling(ctx, method, req, resp)
	if err != nil {
		return nil, err
	}

	if out, ok := resp.(ImageStatusResponse); ok && out.Image() != nil {
		// ImageStatus
		img := out.Image().(Image)
		if len(img.RepoDigests()) > 0 {
			imageName = img.RepoDigests()[0]
		}
		r.setImageNameById(img.Id(), imageName, true)
		out.SetImage(client.addPrefix(out.Image()).(Image))
		return resp, err
	}

	if out, ok := resp.(ImageObject); ok {
		// PullImage
		r.setImageNameById(out.Image(), imageName, false)
		out.SetImage(client.imageName(out.Image()))
		return resp, err
	}

	// RemoveImage
	r.deleteImageNameById(in.Image())
	return resp, err
}

var dispatchTable = map[string]dispatchItem{
	"RuntimeService/Version":                  {(*RuntimeProxy).passToPrimary, criNoisyLogLevel},
	"RuntimeService/Status":                   {(*RuntimeProxy).passToPrimary, criNoisyLogLevel},
	"RuntimeService/UpdateRuntimeConfig":      {(*RuntimeProxy).updateRuntimeConfig, criRequestLogLevel},
	"RuntimeService/RunPodSandbox":            {(*RuntimeProxy).runPodSandbox, criRequestLogLevel},
	"RuntimeService/ListPodSandbox":           {(*RuntimeProxy).listObjects, criListLogLevel},
	"RuntimeService/StopPodSandbox":           {(*RuntimeProxy).handlePodSandbox, criRequestLogLevel},
	"RuntimeService/RemovePodSandbox":         {(*RuntimeProxy).handlePodSandbox, criRequestLogLevel},
	"RuntimeService/PodSandboxStatus":         {(*RuntimeProxy).podSandboxStatus, criNoisyLogLevel},
	"RuntimeService/CreateContainer":          {(*RuntimeProxy).createContainer, criRequestLogLevel},
	"RuntimeService/ListContainers":           {(*RuntimeProxy).listObjects, criListLogLevel},
	"RuntimeService/ListContainerStats":       {(*RuntimeProxy).listObjects, criListLogLevel},
	"RuntimeService/StartContainer":           {(*RuntimeProxy).handleContainer, criRequestLogLevel},
	"RuntimeService/StopContainer":            {(*RuntimeProxy).handleContainer, criRequestLogLevel},
	"RuntimeService/RemoveContainer":          {(*RuntimeProxy).handleContainer, criRequestLogLevel},
	"RuntimeService/ContainerStatus":          {(*RuntimeProxy).containerStatus, criNoisyLogLevel},
	"RuntimeService/ContainerStats":           {(*RuntimeProxy).containerStats, criNoisyLogLevel},
	"RuntimeService/UpdateContainerResources": {(*RuntimeProxy).handleContainer, criRequestLogLevel},
	"RuntimeService/ExecSync":                 {(*RuntimeProxy).handleContainer, criRequestLogLevel},
	"RuntimeService/Exec":                     {(*RuntimeProxy).handleContainer, criRequestLogLevel},
	"RuntimeService/Attach":                   {(*RuntimeProxy).handleContainer, criRequestLogLevel},
	"RuntimeService/ReopenContainerLog":       {(*RuntimeProxy).handleContainer, criRequestLogLevel},
	"RuntimeService/PortForward":              {(*RuntimeProxy).handlePodSandbox, criRequestLogLevel},
	"ImageService/ListImages":                 {(*RuntimeProxy).listObjects, criListLogLevel},
	"ImageService/ImageStatus":                {(*RuntimeProxy).handleImage, criNoisyLogLevel},
	"ImageService/PullImage":                  {(*RuntimeProxy).handleImage, criRequestLogLevel},
	"ImageService/RemoveImage":                {(*RuntimeProxy).handleImage, criRequestLogLevel},
	"ImageService/ImageFsInfo":                {(*RuntimeProxy).listObjects, criRequestLogLevel},
}

var replaceRx = regexp.MustCompile(`\(\*(v1alpha2.\w+)\)\(0x[0-9a-f]+\)`)
var rmRx = regexp.MustCompile(`(?: \(string\))? \(len=\d+(?: cap=\d+)?\)`)

func dump(o interface{}) string {
	out, err := yaml.Marshal(o)
	if err != nil {
		return fmt.Sprintf("<Error marshalling %T: %v>", o, err)
	}
	return string(out)
}
