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

Based on fake_runtime_service.go and cri_stats_provider_test.go
from Kubernetes project.
Original copyright notice follows:

Copyright 2017 The Kubernetes Authors.

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
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"

	runtimeapi "github.com/Mirantis/criproxy/pkg/runtimeapi/v1_8"
	"golang.org/x/net/context"
)

var (
	version = "0.1.0"

	FakeRuntimeName  = "fakeRuntime"
	FakePodSandboxIP = "192.168.192.168"
)

func BuildContainerName(metadata *runtimeapi.ContainerMetadata, sandboxID string) string {
	// include the sandbox ID to make the container ID unique.
	return fmt.Sprintf("%s_%s_%d", sandboxID, metadata.Name, metadata.Attempt)
}

func BuildSandboxName(metadata *runtimeapi.PodSandboxMetadata) string {
	return fmt.Sprintf("%s_%s_%s_%d", metadata.Name, metadata.Namespace, metadata.Uid, metadata.Attempt)
}

type FakePodSandbox struct {
	// PodSandboxStatus contains the runtime information for a sandbox.
	runtimeapi.PodSandboxStatus
}

type FakeContainer struct {
	// ContainerStatus contains the runtime information for a container.
	runtimeapi.ContainerStatus

	// the sandbox id of this container
	SandboxID string
}

type FakeRuntimeServer struct {
	sync.Mutex
	streamUrl          string
	journal            Journal
	CurrentTime        int64
	FakeStatus         *runtimeapi.RuntimeStatus
	Containers         map[string]*FakeContainer
	Sandboxes          map[string]*FakePodSandbox
	FakeContainerStats map[string]*runtimeapi.ContainerStats
}

func (r *FakeRuntimeServer) SetFakeSandboxes(sandboxes []*FakePodSandbox) {
	r.Lock()
	defer r.Unlock()

	r.Sandboxes = make(map[string]*FakePodSandbox)
	for _, sandbox := range sandboxes {
		r.Sandboxes[sandbox.Id] = sandbox
	}
}

func (r *FakeRuntimeServer) SetFakeContainers(containers []*FakeContainer) {
	r.Lock()
	defer r.Unlock()

	r.Containers = make(map[string]*FakeContainer)
	for _, c := range containers {
		r.Containers[c.Id] = c
	}
}

func (r *FakeRuntimeServer) SetFakeContainerStats(containerStats []*runtimeapi.ContainerStats) {
	r.Lock()
	defer r.Unlock()

	r.FakeContainerStats = make(map[string]*runtimeapi.ContainerStats)
	for _, s := range containerStats {
		r.FakeContainerStats[s.Attributes.Id] = s
	}
}

func NewFakeRuntimeServer(journal Journal, streamUrl string) *FakeRuntimeServer {
	ready := true
	runtimeReadyStr := runtimeapi.RuntimeReady
	networkReadyStr := runtimeapi.NetworkReady
	return &FakeRuntimeServer{
		streamUrl:   streamUrl,
		journal:     journal,
		CurrentTime: time.Now().UnixNano(),
		FakeStatus: &runtimeapi.RuntimeStatus{
			Conditions: []*runtimeapi.RuntimeCondition{
				{
					Type:   runtimeReadyStr,
					Status: ready,
				},
				{
					Type:   networkReadyStr,
					Status: ready,
				},
			},
		},
		Containers: make(map[string]*FakeContainer),
		Sandboxes:  make(map[string]*FakePodSandbox),
	}
}

func (r *FakeRuntimeServer) Version(ctx context.Context, in *runtimeapi.VersionRequest) (*runtimeapi.VersionResponse, error) {
	r.journal.Record("Version")

	return &runtimeapi.VersionResponse{
		Version:           version,
		RuntimeName:       FakeRuntimeName,
		RuntimeVersion:    version,
		RuntimeApiVersion: version,
	}, nil
}

func (r *FakeRuntimeServer) Status(ctx context.Context, in *runtimeapi.StatusRequest) (*runtimeapi.StatusResponse, error) {
	r.journal.Record("Status")
	return &runtimeapi.StatusResponse{Status: r.FakeStatus}, nil
}

func (r *FakeRuntimeServer) RunPodSandbox(ctx context.Context, in *runtimeapi.RunPodSandboxRequest) (*runtimeapi.RunPodSandboxResponse, error) {
	r.Lock()
	defer r.Unlock()

	r.journal.Record("RunPodSandbox")

	// PodSandboxID should be randomized for real container runtime, but here just use
	// fixed name from BuildSandboxName() for easily making fake sandboxes.
	config := in.GetConfig()
	podSandboxID := BuildSandboxName(config.Metadata)
	r.Sandboxes[podSandboxID] = &FakePodSandbox{
		PodSandboxStatus: runtimeapi.PodSandboxStatus{
			Id:        podSandboxID,
			Metadata:  config.Metadata,
			State:     runtimeapi.PodSandboxState_SANDBOX_READY,
			CreatedAt: r.CurrentTime,
			Network: &runtimeapi.PodSandboxNetworkStatus{
				Ip: FakePodSandboxIP,
			},
			Labels:      config.Labels,
			Annotations: config.Annotations,
		},
	}

	return &runtimeapi.RunPodSandboxResponse{PodSandboxId: podSandboxID}, nil
}

func (r *FakeRuntimeServer) StopPodSandbox(ctx context.Context, in *runtimeapi.StopPodSandboxRequest) (*runtimeapi.StopPodSandboxResponse, error) {
	r.Lock()
	defer r.Unlock()

	r.journal.Record("StopPodSandbox")

	notReadyState := runtimeapi.PodSandboxState_SANDBOX_NOTREADY
	if s, ok := r.Sandboxes[in.PodSandboxId]; ok {
		s.State = notReadyState
	} else {
		return nil, fmt.Errorf("pod sandbox %s not found", in.PodSandboxId)
	}

	return &runtimeapi.StopPodSandboxResponse{}, nil
}

func (r *FakeRuntimeServer) RemovePodSandbox(ctx context.Context, in *runtimeapi.RemovePodSandboxRequest) (*runtimeapi.RemovePodSandboxResponse, error) {
	r.Lock()
	defer r.Unlock()

	r.journal.Record("RemovePodSandbox")

	// Remove the pod sandbox
	delete(r.Sandboxes, in.PodSandboxId)

	return &runtimeapi.RemovePodSandboxResponse{}, nil
}

func (r *FakeRuntimeServer) PodSandboxStatus(ctx context.Context, in *runtimeapi.PodSandboxStatusRequest) (*runtimeapi.PodSandboxStatusResponse, error) {
	r.Lock()
	defer r.Unlock()

	r.journal.Record("PodSandboxStatus")

	s, ok := r.Sandboxes[in.PodSandboxId]
	if !ok {
		return nil, fmt.Errorf("pod sandbox %q not found", in.PodSandboxId)
	}

	return &runtimeapi.PodSandboxStatusResponse{Status: &s.PodSandboxStatus}, nil
}

func (r *FakeRuntimeServer) ListPodSandbox(ctx context.Context, in *runtimeapi.ListPodSandboxRequest) (*runtimeapi.ListPodSandboxResponse, error) {
	r.Lock()
	defer r.Unlock()

	r.journal.Record("ListPodSandbox")

	var ids []string
	for id, _ := range r.Sandboxes {
		ids = append(ids, id)
	}
	sort.Strings(ids)

	filter := in.GetFilter()
	result := make([]*runtimeapi.PodSandbox, 0)
	for _, id := range ids {
		s := r.Sandboxes[id]
		if filter != nil {
			if filter.Id != "" && filter.Id != id {
				continue
			}
			if filter.State != nil && filter.GetState().State != s.State {
				continue
			}
			if filter.LabelSelector != nil && !filterInLabels(filter.LabelSelector, s.GetLabels()) {
				continue
			}
		}

		result = append(result, &runtimeapi.PodSandbox{
			Id:          s.Id,
			Metadata:    s.Metadata,
			State:       s.State,
			CreatedAt:   s.CreatedAt,
			Labels:      s.Labels,
			Annotations: s.Annotations,
		})
	}

	return &runtimeapi.ListPodSandboxResponse{Items: result}, nil
}

func (r *FakeRuntimeServer) CreateContainer(ctx context.Context, in *runtimeapi.CreateContainerRequest) (*runtimeapi.CreateContainerResponse, error) {
	r.Lock()
	defer r.Unlock()

	r.journal.Record("CreateContainer")

	// ContainerID should be randomized for real container runtime, but here just use
	// fixed BuildContainerName() for easily making fake containers.
	config := in.GetConfig()
	containerID := BuildContainerName(config.Metadata, in.PodSandboxId)
	r.Containers[containerID] = &FakeContainer{
		ContainerStatus: runtimeapi.ContainerStatus{
			Id:          containerID,
			Metadata:    config.Metadata,
			Image:       config.Image,
			ImageRef:    config.Image.Image,
			CreatedAt:   r.CurrentTime,
			State:       runtimeapi.ContainerState_CONTAINER_CREATED,
			Labels:      config.Labels,
			Annotations: config.Annotations,
		},
		SandboxID: in.PodSandboxId,
	}

	return &runtimeapi.CreateContainerResponse{ContainerId: containerID}, nil
}

func (r *FakeRuntimeServer) StartContainer(ctx context.Context, in *runtimeapi.StartContainerRequest) (*runtimeapi.StartContainerResponse, error) {
	r.Lock()
	defer r.Unlock()

	r.journal.Record("StartContainer")

	c, ok := r.Containers[in.ContainerId]
	if !ok {
		return nil, fmt.Errorf("container %s not found", in.ContainerId)
	}

	// Set container to running.
	c.State = runtimeapi.ContainerState_CONTAINER_RUNNING
	c.StartedAt = r.CurrentTime

	return &runtimeapi.StartContainerResponse{}, nil
}

func (r *FakeRuntimeServer) StopContainer(ctx context.Context, in *runtimeapi.StopContainerRequest) (*runtimeapi.StopContainerResponse, error) {
	r.Lock()
	defer r.Unlock()

	r.journal.Record("StopContainer")

	c, ok := r.Containers[in.ContainerId]
	if !ok {
		return nil, fmt.Errorf("container %q not found", in.ContainerId)
	}

	// Set container to exited state.
	c.State = runtimeapi.ContainerState_CONTAINER_EXITED
	c.FinishedAt = r.CurrentTime

	return &runtimeapi.StopContainerResponse{}, nil
}

func (r *FakeRuntimeServer) RemoveContainer(ctx context.Context, in *runtimeapi.RemoveContainerRequest) (*runtimeapi.RemoveContainerResponse, error) {
	r.Lock()
	defer r.Unlock()

	r.journal.Record("RemoveContainer")

	// Remove the container
	delete(r.Containers, in.ContainerId)

	return &runtimeapi.RemoveContainerResponse{}, nil
}

func (r *FakeRuntimeServer) ListContainers(ctx context.Context, in *runtimeapi.ListContainersRequest) (*runtimeapi.ListContainersResponse, error) {
	r.Lock()
	defer r.Unlock()

	r.journal.Record("ListContainers")

	var ids []string
	for id, _ := range r.Containers {
		ids = append(ids, id)
	}
	sort.Strings(ids)

	filter := in.GetFilter()
	result := make([]*runtimeapi.Container, 0)
	for _, id := range ids {
		s := r.Containers[id]
		if filter != nil {
			if filter.Id != "" && filter.Id != s.Id {
				continue
			}
			if filter.PodSandboxId != "" && filter.PodSandboxId != s.SandboxID {
				continue
			}
			if filter.State != nil && filter.GetState().State != s.State {
				continue
			}
			if filter.LabelSelector != nil && !filterInLabels(filter.LabelSelector, s.GetLabels()) {
				continue
			}
		}

		result = append(result, &runtimeapi.Container{
			Id:           s.Id,
			CreatedAt:    s.CreatedAt,
			PodSandboxId: s.SandboxID,
			Metadata:     s.Metadata,
			State:        s.State,
			Image:        s.Image,
			ImageRef:     s.ImageRef,
			Labels:       s.Labels,
			Annotations:  s.Annotations,
		})
	}

	return &runtimeapi.ListContainersResponse{Containers: result}, nil
}

func (r *FakeRuntimeServer) ContainerStatus(ctx context.Context, in *runtimeapi.ContainerStatusRequest) (*runtimeapi.ContainerStatusResponse, error) {
	r.Lock()
	defer r.Unlock()

	r.journal.Record("ContainerStatus")

	c, ok := r.Containers[in.ContainerId]
	if !ok {
		return nil, fmt.Errorf("container %q not found", in.ContainerId)
	}

	return &runtimeapi.ContainerStatusResponse{Status: &c.ContainerStatus}, nil
}

func (r *FakeRuntimeServer) UpdateContainerResources(ctx context.Context, in *runtimeapi.UpdateContainerResourcesRequest) (*runtimeapi.UpdateContainerResourcesResponse, error) {
	r.Lock()
	defer r.Unlock()

	r.journal.Record("UpdateContainerResources")

	if _, ok := r.Containers[in.ContainerId]; !ok {
		return nil, fmt.Errorf("container %q not found", in.ContainerId)
	}

	return &runtimeapi.UpdateContainerResourcesResponse{}, nil
}

func (r *FakeRuntimeServer) ExecSync(ctx context.Context, in *runtimeapi.ExecSyncRequest) (*runtimeapi.ExecSyncResponse, error) {
	r.journal.Record("ExecSync")
	return &runtimeapi.ExecSyncResponse{Stdout: nil, Stderr: nil, ExitCode: int32(0)}, nil
}

func (r *FakeRuntimeServer) Exec(ctx context.Context, in *runtimeapi.ExecRequest) (*runtimeapi.ExecResponse, error) {
	r.journal.Record("Exec")
	return &runtimeapi.ExecResponse{Url: r.streamUrl}, nil
}

func (r *FakeRuntimeServer) Attach(ctx context.Context, in *runtimeapi.AttachRequest) (*runtimeapi.AttachResponse, error) {
	r.journal.Record("Attach")
	return &runtimeapi.AttachResponse{Url: r.streamUrl}, nil
}

func (r *FakeRuntimeServer) PortForward(ctx context.Context, in *runtimeapi.PortForwardRequest) (*runtimeapi.PortForwardResponse, error) {
	r.journal.Record("PortForward")
	return &runtimeapi.PortForwardResponse{Url: r.streamUrl}, nil
}

func (r *FakeRuntimeServer) UpdateRuntimeConfig(ctx context.Context, in *runtimeapi.UpdateRuntimeConfigRequest) (*runtimeapi.UpdateRuntimeConfigResponse, error) {
	r.journal.Record("UpdateRuntimeConfig")
	return &runtimeapi.UpdateRuntimeConfigResponse{}, nil
}

func (r *FakeRuntimeServer) ListContainerStats(ctx context.Context, in *runtimeapi.ListContainerStatsRequest) (*runtimeapi.ListContainerStatsResponse, error) {
	r.Lock()
	defer r.Unlock()

	r.journal.Record("ListContainerStats")

	var result []*runtimeapi.ContainerStats
	for _, c := range r.Containers {
		if in.Filter != nil {
			if in.Filter.Id != "" && in.Filter.Id != c.Id {
				continue
			}
			if in.Filter.PodSandboxId != "" && in.Filter.PodSandboxId != c.SandboxID {
				continue
			}
			if in.Filter.LabelSelector != nil && !filterInLabels(in.Filter.LabelSelector, c.GetLabels()) {
				continue
			}
		}
		s, found := r.FakeContainerStats[c.Id]
		if !found {
			continue
		}
		result = append(result, s)
	}

	return &runtimeapi.ListContainerStatsResponse{Stats: result}, nil
}

func (r *FakeRuntimeServer) ContainerStats(ctx context.Context, in *runtimeapi.ContainerStatsRequest) (*runtimeapi.ContainerStatsResponse, error) {
	r.Lock()
	defer r.Unlock()

	r.journal.Record("ContainerStats")

	s, found := r.FakeContainerStats[in.ContainerId]
	if !found {
		return nil, fmt.Errorf("no stats for container %q", in.ContainerId)
	}
	return &runtimeapi.ContainerStatsResponse{Stats: s}, nil
}

func MakeFakeContainerStats(id string, metadata *runtimeapi.ContainerMetadata, imageFsUUID string) *runtimeapi.ContainerStats {
	return &runtimeapi.ContainerStats{
		Attributes: &runtimeapi.ContainerAttributes{
			Id:       id,
			Metadata: metadata,
		},
		Cpu: &runtimeapi.CpuUsage{
			Timestamp:            time.Now().UnixNano(),
			UsageCoreNanoSeconds: &runtimeapi.UInt64Value{Value: rand.Uint64()},
		},
		Memory: &runtimeapi.MemoryUsage{
			Timestamp:       time.Now().UnixNano(),
			WorkingSetBytes: &runtimeapi.UInt64Value{Value: rand.Uint64()},
		},
		WritableLayer: &runtimeapi.FilesystemUsage{
			Timestamp:  time.Now().UnixNano(),
			StorageId:  &runtimeapi.StorageIdentifier{Uuid: imageFsUUID},
			UsedBytes:  &runtimeapi.UInt64Value{Value: rand.Uint64()},
			InodesUsed: &runtimeapi.UInt64Value{Value: rand.Uint64()},
		},
	}
}
