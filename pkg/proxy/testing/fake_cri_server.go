/*
Copyright 2016 Mirantis

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

package testing

import (
	"log"
	"net"
	"os"
	"syscall"

	"google.golang.org/grpc"

	"github.com/Mirantis/criproxy/pkg/runtimeapis"
	v1_10 "github.com/Mirantis/criproxy/pkg/runtimeapis/v1_10"
	v1_9 "github.com/Mirantis/criproxy/pkg/runtimeapis/v1_9"
)

type FakeCriServer interface {
	Serve(addr string, readyCh chan struct{}) error
	Stop()
	SetFakeImages(images []string)
	SetFakeImageSize(size uint64)
	SetFakeContainerStats(containerId, containerName, imageFsUUID string) interface{}
	SetFakeFilesystemUsage(imageFsUUID string) interface{}
	CurrentTime() int64
}

type fakeCriServerBase struct {
	server *grpc.Server
}

func newFakeCriServerBase() *fakeCriServerBase {
	return &fakeCriServerBase{grpc.NewServer()}
}

func (s *fakeCriServerBase) Serve(addr string, readyCh chan struct{}) error {
	if err := syscall.Unlink(addr); err != nil && !os.IsNotExist(err) {
		return err
	}
	ln, err := net.Listen("unix", addr)
	if err != nil {
		return err
	}
	defer ln.Close()
	if readyCh != nil {
		close(readyCh)
	}
	return s.server.Serve(ln)
}

func (s *fakeCriServerBase) Stop() {
	s.server.GracefulStop()
}

type FakeCriServer19 struct {
	*fakeCriServerBase
	*FakeRuntimeServer19
	*FakeImageServer19
}

var _ FakeCriServer = &FakeCriServer19{}

func NewFakeCriServer19(journal Journal, streamUrl string) FakeCriServer {
	s := &FakeCriServer19{
		fakeCriServerBase:   newFakeCriServerBase(),
		FakeRuntimeServer19: NewFakeRuntimeServer19(NewPrefixJournal(journal, "runtime/"), streamUrl),
		FakeImageServer19:   NewFakeImageServer19(NewPrefixJournal(journal, "image/")),
	}
	v1_9.RegisterRuntimeServiceServer(s.server, s)
	v1_9.RegisterImageServiceServer(s.server, s)
	return s
}

func (s *FakeCriServer19) SetFakeContainerStats(containerId, containerName, imageFsUUID string) interface{} {
	r := MakeFakeContainerStats19(containerId, &v1_9.ContainerMetadata{
		Name: "container1",
	}, imageFsUUID)
	s.FakeRuntimeServer19.SetFakeContainerStats([]*v1_9.ContainerStats{r})
	return r
}

func (s *FakeCriServer19) SetFakeFilesystemUsage(imageFsUUID string) interface{} {
	r := MakeFakeImageFsUsage19(imageFsUUID)
	s.FakeImageServer19.SetFakeFilesystemUsage([]*v1_9.FilesystemUsage{r})
	return r
}

func (s *FakeCriServer19) CurrentTime() int64 {
	return s.FakeRuntimeServer19.CurrentTime
}

type FakeCriServer110 struct {
	*fakeCriServerBase
	*FakeRuntimeServer110
	*FakeImageServer110
}

var _ FakeCriServer = &FakeCriServer110{}

func NewFakeCriServer110(journal Journal, streamUrl string) FakeCriServer {
	s := &FakeCriServer110{
		fakeCriServerBase:    newFakeCriServerBase(),
		FakeRuntimeServer110: NewFakeRuntimeServer110(NewPrefixJournal(journal, "runtime/"), streamUrl),
		FakeImageServer110:   NewFakeImageServer110(NewPrefixJournal(journal, "image/")),
	}
	v1_10.RegisterRuntimeServiceServer(s.server, s)
	v1_10.RegisterImageServiceServer(s.server, s)
	return s
}

func (s *FakeCriServer110) SetFakeContainerStats(containerId, containerName, imageFsUUID string) interface{} {
	r := MakeFakeContainerStats110(containerId, &v1_10.ContainerMetadata{
		Name: "container1",
	}, imageFsUUID)
	s.FakeRuntimeServer110.SetFakeContainerStats([]*v1_10.ContainerStats{r})
	out, err := runtimeapis.Downgrade(r)
	if err != nil {
		log.Fatalf("Downgrade %T: %v", r, err)
	}
	return out
}

func (s *FakeCriServer110) SetFakeFilesystemUsage(imageFsUUID string) interface{} {
	r := MakeFakeImageFsUsage110(imageFsUUID)
	s.FakeImageServer110.SetFakeFilesystemUsage([]*v1_10.FilesystemUsage{r})
	out, err := runtimeapis.Downgrade(r)
	if err != nil {
		log.Fatalf("Downgrade %T: %v", r, err)
	}
	return out
}

func (s *FakeCriServer110) CurrentTime() int64 {
	return s.FakeRuntimeServer110.CurrentTime
}
