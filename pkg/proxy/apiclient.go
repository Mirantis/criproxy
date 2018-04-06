/*
Copyright 2017 Mirantis

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
	"strings"
	"sync"
	"time"

	"github.com/docker/distribution/digest"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/Mirantis/criproxy/pkg/utils"
)

type clientState int

const (
	targetRuntimeAnnotationKey = "kubernetes.io/target-runtime"
	clientStateOffline         = clientState(iota)
	clientStateConnecting
	clientStateConnected
	versionRequestMethod = "/runtime.v1alpha2.RuntimeService/Version"
)

var errNotConnected = errors.New("not connected")
var errOldConnection = errors.New("the request was made on an old closed connection")

type apiClient struct {
	sync.Mutex
	criVersion        CRIVersion
	conn              *grpc.ClientConn
	addr              string
	id                string
	connectionTimeout time.Duration
	state             clientState
	connectErrChs     []chan error
}

func newApiClient(criVersion CRIVersion, addr string, connectionTimeout time.Duration) *apiClient {
	id := ""
	parts := strings.SplitN(addr, ":", 2)
	if len(parts) == 2 {
		id, addr = parts[0], parts[1]
	}
	return &apiClient{
		criVersion:        criVersion,
		addr:              addr,
		id:                id,
		connectionTimeout: connectionTimeout,
	}
}

func (c *apiClient) isPrimary() bool {
	return c.id == ""
}

func (c *apiClient) currentState() clientState {
	c.Lock()
	defer c.Unlock()
	return c.state
}

func (c *apiClient) connectNonLocked() chan error {
	if c.state == clientStateConnected {
		errCh := make(chan error, 1)
		errCh <- nil
		return errCh
	}

	errCh := make(chan error, 1)
	c.connectErrChs = append(c.connectErrChs, errCh)
	if c.state == clientStateConnecting {
		return errCh
	}

	c.state = clientStateConnecting
	go func() {
		glog.V(1).Infof("Connecting to runtime service %s", c.addr)
		var conn *grpc.ClientConn
		if err := utils.WaitForSocket(c.addr, -1, func() error {
			var err error
			conn, err = grpc.Dial(c.addr, grpc.WithInsecure(), grpc.WithTimeout(c.connectionTimeout), grpc.WithDialer(utils.Dial))
			if err == nil {
				ctx, _ := context.WithTimeout(context.Background(), c.connectionTimeout)
				pReq, pResp := c.criVersion.ProbeRequest()
				if err = grpc.Invoke(ctx, versionRequestMethod, pReq, pResp, conn); err != nil {
					conn.Close()
				}
			}
			return err
		}); err != nil {
			glog.Errorf("Failed to find the socket: %v", err)
			err = fmt.Errorf("failed to find the socket: %v", err)
			for _, ch := range c.connectErrChs {
				ch <- err
			}
			return
		}

		c.Lock()
		defer c.Unlock()
		glog.V(1).Infof("Connected to runtime service %s", c.addr)
		c.state = clientStateConnected
		c.conn = conn

		for _, ch := range c.connectErrChs {
			ch <- nil
		}
		c.connectErrChs = nil
	}()
	return errCh
}

func (c *apiClient) connect() chan error {
	c.Lock()
	defer c.Unlock()
	return c.connectNonLocked()
}

func (c *apiClient) stopNonLocked() {
	if c.conn == nil {
		return
	}
	if err := c.conn.Close(); err != nil {
		glog.Errorf("Failed to close gRPC connection: %v", err)
	}
	c.conn = nil
	c.state = clientStateOffline
}

func (c *apiClient) stop() {
	c.Lock()
	defer c.Unlock()
	c.stopNonLocked()
}

// handleError checks whether an error returned by grpc call has
// 'Unavailable' code in which case it disconnects from the client and
// starts trying to reestablish the connection. In case if
// tolerateDisconnect is true, it also returns nil in this case. In
// other cases, including non-'Unavailable' errors, it returns the
// original err value
func (c *apiClient) handleError(err error, tolerateDisconnect bool) error {
	if grpc.Code(err) == codes.Unavailable {
		c.Lock()
		defer c.Unlock()
		c.stopNonLocked()
		c.connectNonLocked()

		if tolerateDisconnect {
			return nil
		}
	}
	return fmt.Errorf("%q: %v", c.addr, err)
}

func (c *apiClient) imageName(unprefixedName string) string {
	if c.isPrimary() {
		return unprefixedName
	}
	return c.id + "/" + unprefixedName
}

func (c *apiClient) augmentId(id string) string {
	if !c.isPrimary() {
		return c.id + "__" + id
	}
	return id
}

func (c *apiClient) annotationsMatch(annotations map[string]string) bool {
	targetRuntime, found := annotations[targetRuntimeAnnotationKey]
	if c.isPrimary() {
		return !found
	}
	return found && targetRuntime == c.id
}

func (c *apiClient) idPrefixMatches(id string) (bool, string) {
	switch {
	case c.isPrimary():
		return true, id
	case strings.HasPrefix(id, c.id+"__"):
		return true, id[len(c.id)+2:]
	default:
		return false, ""
	}
}

func (c *apiClient) imageMatches(imageName string) (bool, string) {
	switch {
	case c.isPrimary():
		return true, imageName
	case strings.HasPrefix(imageName, c.id+"/"):
		return true, imageName[len(c.id)+1:]
	default:
		return false, ""
	}
}

func (c *apiClient) prefixSandbox(unprefixedSandbox PodSandbox) PodSandbox {
	if c.isPrimary() {
		return unprefixedSandbox
	}
	sandbox := unprefixedSandbox.Copy()
	sandbox.SetId(c.augmentId(unprefixedSandbox.Id()))
	return sandbox
}

func (c *apiClient) prefixContainer(unprefixedContainer Container) Container {
	if c.isPrimary() {
		return unprefixedContainer
	}
	container := unprefixedContainer.Copy()
	container.SetId(c.augmentId(unprefixedContainer.Id()))
	container.SetPodSandboxId(c.augmentId(unprefixedContainer.PodSandboxId()))
	// don't prefix digests
	if _, err := digest.ParseDigest(unprefixedContainer.Image()); err != nil {
		container.SetImage(c.imageName(unprefixedContainer.Image()))
	}
	return container
}

func (c *apiClient) prefixContainerStats(unprefixedStats ContainerStats) ContainerStats {
	if c.isPrimary() {
		return unprefixedStats
	}
	stats := unprefixedStats.Copy()
	stats.SetId(c.augmentId(unprefixedStats.Id()))
	return stats
}

func (c *apiClient) prefixImage(unprefixedImage Image) Image {
	if c.isPrimary() {
		return unprefixedImage
	}
	image := unprefixedImage.Copy()
	// only prefix image id if it's not a digest
	// so we don't get prefix/sha256:... which doesn't make sense
	if _, err := digest.ParseDigest(image.Id()); err != nil {
		image.SetId(c.imageName(image.Id()))
	}
	newRepoTags := make([]string, len(image.RepoTags()))
	for n, tag := range image.RepoTags() {
		newRepoTags[n] = c.imageName(tag)
	}
	image.SetRepoTags(newRepoTags)
	return image
}

func (c *apiClient) addPrefix(criObject CRIObject) CRIObject {
	switch o := criObject.(type) {
	case PodSandbox:
		return c.prefixSandbox(o)
	case Container:
		return c.prefixContainer(o)
	case ContainerStats:
		return c.prefixContainerStats(o)
	case Image:
		return c.prefixImage(o)
	default:
		return o
	}
}

func (c *apiClient) getConn() (*grpc.ClientConn, error) {
	c.Lock()
	defer c.Unlock()
	if c.state != clientStateConnected {
		return nil, errNotConnected
	}
	return c.conn, nil
}

func (c *apiClient) invoke(ctx context.Context, method string, req, resp CRIObject) (CRIObject, error) {
	conn, err := c.getConn()
	if err != nil {
		return nil, err
	}

	if err = grpc.Invoke(ctx, method, req.Unwrap(), resp.Unwrap(), conn); grpc.Code(err) == codes.Unavailable {
		c.Lock()
		defer c.Unlock()
		if conn != c.conn {
			// do not close the current connection if the request is related
			// to a previously closed one
			err = errOldConnection
		}
	}
	return resp, err
}

func (c *apiClient) invokeWithErrorHandling(ctx context.Context, method string, req, resp CRIObject) (CRIObject, error) {
	err := grpc.Invoke(ctx, method, req.Unwrap(), resp.Unwrap(), c.conn)
	if err != nil {
		err = c.handleError(err, false)
	}
	return resp, err
}

// TODO: handle grpc's ClientTransport.Error() to reconnect
