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
	"log"
	"strings"
	"sync"
	"time"

	runtimeapis "github.com/Mirantis/criproxy/pkg/runtimeapis"
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
	versionRequestMethod = "RuntimeService/Version"
)

var errNotConnected = errors.New("not connected")
var errOldConnection = errors.New("the request was made on an old closed connection")

type client interface {
	getID() string
	isPrimary() bool
	currentState() clientState
	connect() chan error
	stop()
	handleError(err error, tolerateDisconnect bool) error
	imageName(unprefixedName string) string
	augmentId(id string) string
	annotationsMatch(annotations map[string]string) bool
	idPrefixMatches(id string) (bool, string)
	imageMatches(imageName string) (bool, string)
	addPrefix(criObject CRIObject) CRIObject
	invoke(ctx context.Context, method string, req, resp CRIObject) (CRIObject, error)
	invokeWithErrorHandling(ctx context.Context, method string, req, resp CRIObject) (CRIObject, error)
}

type clientProbeFunc func(conn *grpc.ClientConn, connectionTimeout time.Duration) error

type clientConnection struct {
	sync.Mutex
	addr              string
	conn              *grpc.ClientConn
	probe             clientProbeFunc
	state             clientState
	connectionTimeout time.Duration
	connectErrChs     []chan error
}

func newClientConnection(addr string, connectionTimeout time.Duration) *clientConnection {
	return &clientConnection{
		addr:              addr,
		connectionTimeout: connectionTimeout,
	}
}

func (c *clientConnection) currentState() clientState {
	c.Lock()
	defer c.Unlock()
	return c.state
}

func (c *clientConnection) connectNonLocked() chan error {
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
			if err == nil && c.probe != nil {
				err = c.probe(conn, c.connectionTimeout)
				if err != nil {
					conn.Close()
				}
			}
			return err
		}); err != nil {
			glog.Errorf("Failed to connect to the socket: %v", err)
			err = fmt.Errorf("failed to connect to the socket: %v", err)
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

func (c *clientConnection) connect() chan error {
	c.Lock()
	defer c.Unlock()
	return c.connectNonLocked()
}

func (c *clientConnection) stopNonLocked() {
	if c.conn == nil {
		return
	}
	if err := c.conn.Close(); err != nil {
		glog.Errorf("Failed to close gRPC connection: %v", err)
	}
	c.conn = nil
	c.state = clientStateOffline
}

func (c *clientConnection) stop() {
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
func (c *clientConnection) handleError(err error, tolerateDisconnect bool) error {
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

type clientBase struct {
	id string
}

func (c *clientBase) getID() string { return c.id }

func (c *clientBase) isPrimary() bool {
	return c.id == ""
}

func (c *clientBase) imageName(unprefixedName string) string {
	if c.isPrimary() {
		return unprefixedName
	}
	return c.id + "/" + unprefixedName
}

func (c *clientBase) augmentId(id string) string {
	if !c.isPrimary() {
		return c.id + "__" + id
	}
	return id
}

func (c *clientBase) annotationsMatch(annotations map[string]string) bool {
	targetRuntime, found := annotations[targetRuntimeAnnotationKey]
	if c.isPrimary() {
		return !found
	}
	return found && targetRuntime == c.id
}

func (c *clientBase) idPrefixMatches(id string) (bool, string) {
	switch {
	case c.isPrimary():
		return true, id
	case strings.HasPrefix(id, c.id+"__"):
		return true, id[len(c.id)+2:]
	default:
		return false, ""
	}
}

func (c *clientBase) imageMatches(imageName string) (bool, string) {
	switch {
	case c.isPrimary():
		return true, imageName
	case strings.HasPrefix(imageName, c.id+"/"):
		return true, imageName[len(c.id)+1:]
	default:
		return false, ""
	}
}

func (c *clientBase) prefixSandbox(unprefixedSandbox PodSandbox) PodSandbox {
	if c.isPrimary() {
		return unprefixedSandbox
	}
	sandbox := unprefixedSandbox.Copy()
	sandbox.SetId(c.augmentId(unprefixedSandbox.Id()))
	return sandbox
}

func (c *clientBase) prefixContainer(unprefixedContainer Container) Container {
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

func (c *clientBase) prefixContainerStats(unprefixedStats ContainerStats) ContainerStats {
	if c.isPrimary() {
		return unprefixedStats
	}
	stats := unprefixedStats.Copy()
	stats.SetId(c.augmentId(unprefixedStats.Id()))
	return stats
}

func (c *clientBase) prefixImage(unprefixedImage Image) Image {
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
	// repo digests may or may not include the image name
	newRepoDigests := make([]string, len(image.RepoDigests()))
	for n, digest := range image.RepoDigests() {
		p := strings.Index(digest, "@")
		if p > 0 {
			newRepoDigests[n] = c.imageName(digest)
		} else {
			newRepoDigests[n] = digest
		}
	}
	image.SetRepoDigests(newRepoDigests)
	return image
}

func (c *clientBase) addPrefix(criObject CRIObject) CRIObject {
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

type apiClient struct {
	clientBase
	*clientConnection
	criVersion CRIVersion
}

var _ client = &apiClient{}

func newApiClient(criVersion CRIVersion, clientConn *clientConnection, id string) *apiClient {
	return &apiClient{
		clientBase:       clientBase{id},
		criVersion:       criVersion,
		clientConnection: clientConn,
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

type upgradingClient struct {
	client
	legacyVersion CRIVersion
	newVersion    CRIVersion
}

var _ client = &upgradingClient{}

func newUpgradingClient(next client, legacyVersion UpgradableCRIVersion) *upgradingClient {
	return &upgradingClient{
		client:        next,
		legacyVersion: legacyVersion,
		newVersion:    legacyVersion.UpgradesTo(),
	}
}

func (c *upgradingClient) addPrefix(o CRIObject) CRIObject {
	return c.downgradeCRIObject(c.client.addPrefix(c.upgradeCRIObject(o)))
}

func (c *upgradingClient) invoke(ctx context.Context, method string, req, resp CRIObject) (CRIObject, error) {
	method = strings.Replace(method, "runtime.", "runtime.v1alpha2.", 1)
	r, err := c.client.invoke(ctx, method, c.upgradeCRIObject(req), c.upgradeCRIObject(resp))
	if err != nil {
		return nil, err
	}
	return c.downgradeCRIObjectTo(r, resp), err
}

func (c *upgradingClient) invokeWithErrorHandling(ctx context.Context, method string, req, resp CRIObject) (CRIObject, error) {
	method = strings.Replace(method, "runtime.", "runtime.v1alpha2.", 1)
	r, err := c.client.invokeWithErrorHandling(ctx, method, c.upgradeCRIObject(req), c.upgradeCRIObject(resp))
	if err != nil {
		return nil, err
	}
	return c.downgradeCRIObjectTo(r, resp), nil
}

func (c *upgradingClient) upgradeCRIObject(o CRIObject) CRIObject {
	upgraded, err := runtimeapis.Upgrade(o.Unwrap())
	if err != nil {
		log.Panicf("Couldn't upgrade %T: %v", o.Unwrap(), err)
	}
	r, _, err := c.newVersion.WrapObject(upgraded)
	if err != nil {
		log.Panicf("Error wrapping upgraded object %T: %v", upgraded, err)
	}
	return r
}

func (c *upgradingClient) downgradeCRIObject(o CRIObject) CRIObject {
	downgraded, err := runtimeapis.Downgrade(o.Unwrap())
	if err != nil {
		log.Panicf("Couldn't downgrade %T: %v", o.Unwrap(), err)
	}
	r, _, err := c.legacyVersion.WrapObject(downgraded)
	if err != nil {
		log.Panicf("Error wrapping downgraded object %T: %v", downgraded, err)
	}
	return r
}

func (c *upgradingClient) downgradeCRIObjectTo(o CRIObject, resp CRIObject) CRIObject {
	downgraded, err := runtimeapis.Downgrade(o.Unwrap())
	if err != nil {
		log.Panicf("Couldn't downgrade %T: %v", o.Unwrap(), err)
	}
	resp.Wrap(downgraded)
	return resp
}

// autoClient detects server version and chooses upgradingClient
// or plain apiClient depending on it
type autoClient struct {
	clientBase
	*clientConnection
	proxyCRIVersion CRIVersion
	next            client
}

var _ client = &autoClient{}

func newAutoClient(proxyCRIVersion CRIVersion, addr string, connectionTimeout time.Duration) *autoClient {
	id := ""
	parts := strings.SplitN(addr, ":", 2)
	if len(parts) == 2 {
		id, addr = parts[0], parts[1]
	}
	conn := newClientConnection(addr, connectionTimeout)
	c := &autoClient{
		clientBase:       clientBase{id},
		clientConnection: conn,
		proxyCRIVersion:  proxyCRIVersion,
	}
	conn.probe = c.checkConnection
	return c
}

func (c *autoClient) checkVersion(criVersion CRIVersion, conn *grpc.ClientConn, connectionTimeout time.Duration) error {
	ctx, _ := context.WithTimeout(context.Background(), connectionTimeout)
	pReq, pResp := criVersion.ProbeRequest()
	reqMethod := fmt.Sprintf("/%s.%s", criVersion.ProtoPackage(), versionRequestMethod)
	return grpc.Invoke(ctx, reqMethod, pReq, pResp, conn)
}

func (c *autoClient) checkConnection(conn *grpc.ClientConn, connectionTimeout time.Duration) error {
	upgrade := []bool{false}
	toTry := []CRIVersion{c.proxyCRIVersion}
	upgradableVersion, upgradable := c.proxyCRIVersion.(UpgradableCRIVersion)
	if upgradable {
		upgrade = []bool{true, false}
		toTry = []CRIVersion{upgradableVersion.UpgradesTo(), c.proxyCRIVersion}
	}

	var err error
	for n, v := range toTry {
		if err = c.checkVersion(v, conn, connectionTimeout); err == nil {
			var next client = newApiClient(v, c.clientConnection, c.id)
			if upgrade[n] {
				next = newUpgradingClient(next, upgradableVersion)
			}
			c.next = next
			break
		}
	}
	return err
}

func (c *autoClient) getNext() (client, error) {
	c.Lock()
	defer c.Unlock()
	if c.state != clientStateConnected {
		return nil, errNotConnected
	}
	return c.next, nil
}

func (c *autoClient) invoke(ctx context.Context, method string, req, resp CRIObject) (CRIObject, error) {
	next, err := c.getNext()
	if err != nil {
		return nil, err
	}
	return next.invoke(ctx, method, req, resp)
}

func (c *autoClient) invokeWithErrorHandling(ctx context.Context, method string, req, resp CRIObject) (CRIObject, error) {
	next, err := c.getNext()
	if err != nil {
		return nil, err
	}
	return next.invokeWithErrorHandling(ctx, method, req, resp)
}

// TODO: handle grpc's ClientTransport.Error() to reconnect
