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
	"google.golang.org/grpc"

	runtimeapi "github.com/Mirantis/criproxy/pkg/runtimeapis/v1_9"
)

// ---

type PodSandbox_19 struct {
	inner *runtimeapi.PodSandbox
}

var _ PodSandbox = &PodSandbox_19{}

func (o *PodSandbox_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PodSandbox{}
	} else {
		o.inner = v.(*runtimeapi.PodSandbox)
	}
}
func (o *PodSandbox_19) Unwrap() interface{} { return o.inner }
func (o *PodSandbox_19) Copy() PodSandbox    { r := *o.inner; return &PodSandbox_19{&r} }
func (o *PodSandbox_19) Id() string          { return o.inner.Id }
func (o *PodSandbox_19) SetId(id string)     { o.inner.Id = id }

type Container_19 struct {
	inner *runtimeapi.Container
}

// ---

var _ Container = &Container_19{}

func (o *Container_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.Container{}
	} else {
		o.inner = v.(*runtimeapi.Container)
	}
}
func (o *Container_19) Unwrap() interface{}       { return o.inner }
func (o *Container_19) Copy() Container           { r := *o.inner; return &Container_19{&r} }
func (o *Container_19) Id() string                { return o.inner.Id }
func (o *Container_19) SetId(id string)           { o.inner.Id = id }
func (o *Container_19) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *Container_19) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }
func (o *Container_19) Image() string             { return o.inner.Image.GetImage() }
func (o *Container_19) SetImage(image string)     { o.inner.Image = &runtimeapi.ImageSpec{Image: image} }

// ---

type Image_19 struct {
	inner *runtimeapi.Image
}

var _ Image = &Image_19{}

func (o *Image_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.Image{}
	} else {
		o.inner = v.(*runtimeapi.Image)
	}
}
func (o *Image_19) Unwrap() interface{}           { return o.inner }
func (o *Image_19) Copy() Image                   { r := *o.inner; return &Image_19{&r} }
func (o *Image_19) Id() string                    { return o.inner.Id }
func (o *Image_19) SetId(id string)               { o.inner.Id = id }
func (o *Image_19) RepoTags() []string            { return o.inner.RepoTags }
func (o *Image_19) SetRepoTags(repoTags []string) { o.inner.RepoTags = repoTags }

// ---

type PodSandboxStatus_19 struct {
	inner *runtimeapi.PodSandboxStatus
}

var _ PodSandboxStatus = &PodSandboxStatus_19{}

func (o *PodSandboxStatus_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PodSandboxStatus{}
	} else {
		o.inner = v.(*runtimeapi.PodSandboxStatus)
	}
}
func (o *PodSandboxStatus_19) Unwrap() interface{}    { return o.inner }
func (o *PodSandboxStatus_19) Copy() PodSandboxStatus { r := *o.inner; return &PodSandboxStatus_19{&r} }
func (o *PodSandboxStatus_19) Id() string             { return o.inner.Id }
func (o *PodSandboxStatus_19) SetId(id string)        { o.inner.Id = id }

// ---

type ContainerStatus_19 struct {
	inner *runtimeapi.ContainerStatus
}

var _ ContainerStatus = &ContainerStatus_19{}

func (o *ContainerStatus_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStatus{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStatus)
	}
}
func (o *ContainerStatus_19) Unwrap() interface{}   { return o.inner }
func (o *ContainerStatus_19) Copy() ContainerStatus { r := *o.inner; return &ContainerStatus_19{&r} }
func (o *ContainerStatus_19) Id() string            { return o.inner.Id }
func (o *ContainerStatus_19) SetId(id string)       { o.inner.Id = id }
func (o *ContainerStatus_19) Image() string         { return o.inner.Image.GetImage() }
func (o *ContainerStatus_19) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type ContainerStats_19 struct {
	inner *runtimeapi.ContainerStats
}

var _ ContainerStats = &ContainerStats_19{}

func (o *ContainerStats_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStats{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStats)
	}
}
func (o *ContainerStats_19) Unwrap() interface{}  { return o.inner }
func (o *ContainerStats_19) Copy() ContainerStats { r := *o.inner; return &ContainerStats_19{&r} }
func (o *ContainerStats_19) Id() string           { return o.inner.Attributes.GetId() }
func (o *ContainerStats_19) SetId(id string) {
	if o.inner.Attributes == nil {
		o.inner.Attributes = &runtimeapi.ContainerAttributes{Id: id}
	} else {
		o.inner.Attributes.Id = id
	}
}

// ---

type FilesystemUsage_19 struct {
	inner *runtimeapi.FilesystemUsage
}

func (o *FilesystemUsage_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.FilesystemUsage{}
	} else {
		o.inner = v.(*runtimeapi.FilesystemUsage)
	}
}
func (o *FilesystemUsage_19) Unwrap() interface{} { return o.inner }

// ---

type VersionRequest_19 struct {
	inner *runtimeapi.VersionRequest
}

var _ VersionRequest = &VersionRequest_19{}

func (o *VersionRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.VersionRequest{}
	} else {
		o.inner = v.(*runtimeapi.VersionRequest)
	}
}
func (o *VersionRequest_19) Unwrap() interface{} { return o.inner }

// ---

type VersionResponse_19 struct {
	inner *runtimeapi.VersionResponse
}

var _ VersionResponse = &VersionResponse_19{}

func (o *VersionResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.VersionResponse{}
	} else {
		o.inner = v.(*runtimeapi.VersionResponse)
	}
}
func (o *VersionResponse_19) Unwrap() interface{} { return o.inner }

// ---

type StatusRequest_19 struct {
	inner *runtimeapi.StatusRequest
}

var _ StatusRequest = &StatusRequest_19{}

func (o *StatusRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StatusRequest{}
	} else {
		o.inner = v.(*runtimeapi.StatusRequest)
	}
}
func (o *StatusRequest_19) Unwrap() interface{} { return o.inner }

// ---

type StatusResponse_19 struct {
	inner *runtimeapi.StatusResponse
}

var _ StatusResponse = &StatusResponse_19{}

func (o *StatusResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StatusResponse{}
	} else {
		o.inner = v.(*runtimeapi.StatusResponse)
	}
}
func (o *StatusResponse_19) Unwrap() interface{} { return o.inner }

// ---

type UpdateRuntimeConfigRequest_19 struct {
	inner *runtimeapi.UpdateRuntimeConfigRequest
}

var _ UpdateRuntimeConfigRequest = &UpdateRuntimeConfigRequest_19{}

func (o *UpdateRuntimeConfigRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.UpdateRuntimeConfigRequest{}
	} else {
		o.inner = v.(*runtimeapi.UpdateRuntimeConfigRequest)
	}
}
func (o *UpdateRuntimeConfigRequest_19) Unwrap() interface{} { return o.inner }

// ---

type UpdateRuntimeConfigResponse_19 struct {
	inner *runtimeapi.UpdateRuntimeConfigResponse
}

var _ UpdateRuntimeConfigResponse = &UpdateRuntimeConfigResponse_19{}

func (o *UpdateRuntimeConfigResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.UpdateRuntimeConfigResponse{}
	} else {
		o.inner = v.(*runtimeapi.UpdateRuntimeConfigResponse)
	}
}
func (o *UpdateRuntimeConfigResponse_19) Unwrap() interface{} { return o.inner }

// ---

type RunPodSandboxRequest_19 struct {
	inner *runtimeapi.RunPodSandboxRequest
}

var _ RunPodSandboxRequest = &RunPodSandboxRequest_19{}

func (o *RunPodSandboxRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RunPodSandboxRequest{}
	} else {
		o.inner = v.(*runtimeapi.RunPodSandboxRequest)
	}
}
func (o *RunPodSandboxRequest_19) Unwrap() interface{} { return o.inner }
func (o *RunPodSandboxRequest_19) GetAnnotations() map[string]string {
	return o.inner.Config.GetAnnotations()
}

// ---

type RunPodSandboxResponse_19 struct {
	inner *runtimeapi.RunPodSandboxResponse
}

var _ RunPodSandboxResponse = &RunPodSandboxResponse_19{}

func (o *RunPodSandboxResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RunPodSandboxResponse{}
	} else {
		o.inner = v.(*runtimeapi.RunPodSandboxResponse)
	}
}
func (o *RunPodSandboxResponse_19) Unwrap() interface{}       { return o.inner }
func (o *RunPodSandboxResponse_19) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *RunPodSandboxResponse_19) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type ListPodSandboxRequest_19 struct {
	inner *runtimeapi.ListPodSandboxRequest
}

var _ ListPodSandboxRequest = &ListPodSandboxRequest_19{}

func (o *ListPodSandboxRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListPodSandboxRequest{}
	} else {
		o.inner = v.(*runtimeapi.ListPodSandboxRequest)
	}
}
func (o *ListPodSandboxRequest_19) Unwrap() interface{} { return o.inner }
func (o *ListPodSandboxRequest_19) IdFilter() string {
	return o.inner.Filter.GetId()
}

func (o *ListPodSandboxRequest_19) SetIdFilter(id string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.PodSandboxFilter{Id: id}
	} else {
		o.inner.Filter.Id = id
	}
}

// ---

type ListPodSandboxResponse_19 struct {
	inner *runtimeapi.ListPodSandboxResponse
}

var _ ListPodSandboxResponse = &ListPodSandboxResponse_19{}

func (o *ListPodSandboxResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListPodSandboxResponse{}
	} else {
		o.inner = v.(*runtimeapi.ListPodSandboxResponse)
	}
}
func (o *ListPodSandboxResponse_19) Unwrap() interface{} { return o.inner }
func (o *ListPodSandboxResponse_19) Items() []CRIObject {
	var r []CRIObject
	for _, sandbox := range o.inner.Items {
		r = append(r, &PodSandbox_19{sandbox})
	}
	return r
}
func (o *ListPodSandboxResponse_19) SetItems(items []CRIObject) {
	o.inner.Items = nil
	for _, wrapped := range items {
		o.inner.Items = append(o.inner.Items, wrapped.Unwrap().(*runtimeapi.PodSandbox))
	}
}

// ---

type StopPodSandboxRequest_19 struct {
	inner *runtimeapi.StopPodSandboxRequest
}

var _ StopPodSandboxRequest = &StopPodSandboxRequest_19{}

func (o *StopPodSandboxRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StopPodSandboxRequest{}
	} else {
		o.inner = v.(*runtimeapi.StopPodSandboxRequest)
	}
}
func (o *StopPodSandboxRequest_19) Unwrap() interface{}       { return o.inner }
func (o *StopPodSandboxRequest_19) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *StopPodSandboxRequest_19) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type StopPodSandboxResponse_19 struct {
	inner *runtimeapi.StopPodSandboxResponse
}

var _ StopPodSandboxResponse = &StopPodSandboxResponse_19{}

func (o *StopPodSandboxResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StopPodSandboxResponse{}
	} else {
		o.inner = v.(*runtimeapi.StopPodSandboxResponse)
	}
}
func (o *StopPodSandboxResponse_19) Unwrap() interface{} { return o.inner }

// ---

type RemovePodSandboxRequest_19 struct {
	inner *runtimeapi.RemovePodSandboxRequest
}

var _ RemovePodSandboxRequest = &RemovePodSandboxRequest_19{}

func (o *RemovePodSandboxRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemovePodSandboxRequest{}
	} else {
		o.inner = v.(*runtimeapi.RemovePodSandboxRequest)
	}
}
func (o *RemovePodSandboxRequest_19) Unwrap() interface{}       { return o.inner }
func (o *RemovePodSandboxRequest_19) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *RemovePodSandboxRequest_19) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type RemovePodSandboxResponse_19 struct {
	inner *runtimeapi.RemovePodSandboxResponse
}

var _ RemovePodSandboxResponse = &RemovePodSandboxResponse_19{}

func (o *RemovePodSandboxResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemovePodSandboxResponse{}
	} else {
		o.inner = v.(*runtimeapi.RemovePodSandboxResponse)
	}
}
func (o *RemovePodSandboxResponse_19) Unwrap() interface{} { return o.inner }

// ---

type PodSandboxStatusRequest_19 struct {
	inner *runtimeapi.PodSandboxStatusRequest
}

var _ PodSandboxStatusRequest = &PodSandboxStatusRequest_19{}

func (o *PodSandboxStatusRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PodSandboxStatusRequest{}
	} else {
		o.inner = v.(*runtimeapi.PodSandboxStatusRequest)
	}
}
func (o *PodSandboxStatusRequest_19) Unwrap() interface{}       { return o.inner }
func (o *PodSandboxStatusRequest_19) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *PodSandboxStatusRequest_19) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type PodSandboxStatusResponse_19 struct {
	inner *runtimeapi.PodSandboxStatusResponse
}

var _ PodSandboxStatusResponse = &PodSandboxStatusResponse_19{}

func (o *PodSandboxStatusResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PodSandboxStatusResponse{}
	} else {
		o.inner = v.(*runtimeapi.PodSandboxStatusResponse)
	}
}
func (o *PodSandboxStatusResponse_19) Unwrap() interface{} { return o.inner }
func (o *PodSandboxStatusResponse_19) Status() PodSandboxStatus {
	if o.inner.Status == nil {
		return nil
	}
	return &PodSandboxStatus_19{o.inner.Status}
}

// ---

type CreateContainerRequest_19 struct {
	inner *runtimeapi.CreateContainerRequest
}

var _ CreateContainerRequest = &CreateContainerRequest_19{}

func (o *CreateContainerRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.CreateContainerRequest{}
	} else {
		o.inner = v.(*runtimeapi.CreateContainerRequest)
	}
}
func (o *CreateContainerRequest_19) Unwrap() interface{}       { return o.inner }
func (o *CreateContainerRequest_19) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *CreateContainerRequest_19) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }
func (o *CreateContainerRequest_19) Image() string {
	if o.inner.Config == nil {
		return ""
	}
	return o.inner.Config.Image.GetImage()
}

func (o *CreateContainerRequest_19) SetImage(image string) {
	if o.inner.Config != nil {
		o.inner.Config.Image = &runtimeapi.ImageSpec{Image: image}
	}
}

// ---

type CreateContainerResponse_19 struct {
	inner *runtimeapi.CreateContainerResponse
}

var _ CreateContainerResponse = &CreateContainerResponse_19{}

func (o *CreateContainerResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.CreateContainerResponse{}
	} else {
		o.inner = v.(*runtimeapi.CreateContainerResponse)
	}
}
func (o *CreateContainerResponse_19) Unwrap() interface{}      { return o.inner }
func (o *CreateContainerResponse_19) ContainerId() string      { return o.inner.ContainerId }
func (o *CreateContainerResponse_19) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ListContainersRequest_19 struct {
	inner *runtimeapi.ListContainersRequest
}

var _ ListContainersRequest = &ListContainersRequest_19{}

func (o *ListContainersRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListContainersRequest{}
	} else {
		o.inner = v.(*runtimeapi.ListContainersRequest)
	}
}
func (o *ListContainersRequest_19) Unwrap() interface{} { return o.inner }
func (o *ListContainersRequest_19) IdFilter() string {
	return o.inner.Filter.GetId()
}

func (o *ListContainersRequest_19) SetIdFilter(id string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerFilter{Id: id}
	} else {
		o.inner.Filter.Id = id
	}
}

func (o *ListContainersRequest_19) PodSandboxIdFilter() string {
	return o.inner.Filter.GetPodSandboxId()
}

func (o *ListContainersRequest_19) SetPodSandboxIdFilter(podSandboxId string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerFilter{Id: podSandboxId}
	} else {
		o.inner.Filter.PodSandboxId = podSandboxId
	}
}

// ---

type ListContainersResponse_19 struct {
	inner *runtimeapi.ListContainersResponse
}

var _ ListContainersResponse = &ListContainersResponse_19{}

func (o *ListContainersResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListContainersResponse{}
	} else {
		o.inner = v.(*runtimeapi.ListContainersResponse)
	}
}
func (o *ListContainersResponse_19) Unwrap() interface{} { return o.inner }
func (o *ListContainersResponse_19) Items() []CRIObject {
	var r []CRIObject
	for _, container := range o.inner.Containers {
		r = append(r, &Container_19{container})
	}
	return r
}
func (o *ListContainersResponse_19) SetItems(items []CRIObject) {
	o.inner.Containers = nil
	for _, wrapped := range items {
		o.inner.Containers = append(o.inner.Containers, wrapped.Unwrap().(*runtimeapi.Container))
	}
}

// ---

type ListContainerStatsRequest_19 struct {
	inner *runtimeapi.ListContainerStatsRequest
}

var _ ListContainerStatsRequest = &ListContainerStatsRequest_19{}

func (o *ListContainerStatsRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListContainerStatsRequest{}
	} else {
		o.inner = v.(*runtimeapi.ListContainerStatsRequest)
	}
}
func (o *ListContainerStatsRequest_19) Unwrap() interface{} { return o.inner }
func (o *ListContainerStatsRequest_19) IdFilter() string {
	return o.inner.Filter.GetId()
}

func (o *ListContainerStatsRequest_19) SetIdFilter(id string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerStatsFilter{Id: id}
	} else {
		o.inner.Filter.Id = id
	}
}

func (o *ListContainerStatsRequest_19) PodSandboxIdFilter() string {
	return o.inner.Filter.GetPodSandboxId()
}

func (o *ListContainerStatsRequest_19) SetPodSandboxIdFilter(podSandboxId string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerStatsFilter{Id: podSandboxId}
	} else {
		o.inner.Filter.PodSandboxId = podSandboxId
	}
}

// ---

type ListContainerStatsResponse_19 struct {
	inner *runtimeapi.ListContainerStatsResponse
}

var _ ListContainerStatsResponse = &ListContainerStatsResponse_19{}

func (o *ListContainerStatsResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListContainerStatsResponse{}
	} else {
		o.inner = v.(*runtimeapi.ListContainerStatsResponse)
	}
}
func (o *ListContainerStatsResponse_19) Unwrap() interface{} { return o.inner }
func (o *ListContainerStatsResponse_19) Items() []CRIObject {
	var r []CRIObject
	for _, stats := range o.inner.Stats {
		r = append(r, &ContainerStats_19{stats})
	}
	return r
}
func (o *ListContainerStatsResponse_19) SetItems(items []CRIObject) {
	o.inner.Stats = nil
	for _, wrapped := range items {
		o.inner.Stats = append(o.inner.Stats, wrapped.Unwrap().(*runtimeapi.ContainerStats))
	}
}

// ---

type StartContainerRequest_19 struct {
	inner *runtimeapi.StartContainerRequest
}

var _ StartContainerRequest = &StartContainerRequest_19{}

func (o *StartContainerRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StartContainerRequest{}
	} else {
		o.inner = v.(*runtimeapi.StartContainerRequest)
	}
}
func (o *StartContainerRequest_19) Unwrap() interface{}      { return o.inner }
func (o *StartContainerRequest_19) ContainerId() string      { return o.inner.ContainerId }
func (o *StartContainerRequest_19) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type StartContainerResponse_19 struct {
	inner *runtimeapi.StartContainerResponse
}

var _ StartContainerResponse = &StartContainerResponse_19{}

func (o *StartContainerResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StartContainerResponse{}
	} else {
		o.inner = v.(*runtimeapi.StartContainerResponse)
	}
}
func (o *StartContainerResponse_19) Unwrap() interface{} { return o.inner }

// ---

type StopContainerRequest_19 struct {
	inner *runtimeapi.StopContainerRequest
}

var _ StopContainerRequest = &StopContainerRequest_19{}

func (o *StopContainerRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StopContainerRequest{}
	} else {
		o.inner = v.(*runtimeapi.StopContainerRequest)
	}
}
func (o *StopContainerRequest_19) Unwrap() interface{}      { return o.inner }
func (o *StopContainerRequest_19) ContainerId() string      { return o.inner.ContainerId }
func (o *StopContainerRequest_19) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type StopContainerResponse_19 struct {
	inner *runtimeapi.StopContainerResponse
}

var _ StopContainerResponse = &StopContainerResponse_19{}

func (o *StopContainerResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StopContainerResponse{}
	} else {
		o.inner = v.(*runtimeapi.StopContainerResponse)
	}
}
func (o *StopContainerResponse_19) Unwrap() interface{} { return o.inner }

// ---

type RemoveContainerRequest_19 struct {
	inner *runtimeapi.RemoveContainerRequest
}

var _ RemoveContainerRequest = &RemoveContainerRequest_19{}

func (o *RemoveContainerRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemoveContainerRequest{}
	} else {
		o.inner = v.(*runtimeapi.RemoveContainerRequest)
	}
}
func (o *RemoveContainerRequest_19) Unwrap() interface{}      { return o.inner }
func (o *RemoveContainerRequest_19) ContainerId() string      { return o.inner.ContainerId }
func (o *RemoveContainerRequest_19) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type RemoveContainerResponse_19 struct {
	inner *runtimeapi.RemoveContainerResponse
}

var _ RemoveContainerResponse = &RemoveContainerResponse_19{}

func (o *RemoveContainerResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemoveContainerResponse{}
	} else {
		o.inner = v.(*runtimeapi.RemoveContainerResponse)
	}
}
func (o *RemoveContainerResponse_19) Unwrap() interface{} { return o.inner }

// ---

type ContainerStatusRequest_19 struct {
	inner *runtimeapi.ContainerStatusRequest
}

var _ ContainerStatusRequest = &ContainerStatusRequest_19{}

func (o *ContainerStatusRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStatusRequest{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStatusRequest)
	}
}
func (o *ContainerStatusRequest_19) Unwrap() interface{}      { return o.inner }
func (o *ContainerStatusRequest_19) ContainerId() string      { return o.inner.ContainerId }
func (o *ContainerStatusRequest_19) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ContainerStatusResponse_19 struct {
	inner *runtimeapi.ContainerStatusResponse
}

var _ ContainerStatusResponse = &ContainerStatusResponse_19{}

func (o *ContainerStatusResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStatusResponse{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStatusResponse)
	}
}
func (o *ContainerStatusResponse_19) Unwrap() interface{} { return o.inner }
func (o *ContainerStatusResponse_19) Status() ContainerStatus {
	if o.inner.Status == nil {
		return nil
	}
	return &ContainerStatus_19{o.inner.Status}
}

// ---

type ContainerStatsRequest_19 struct {
	inner *runtimeapi.ContainerStatsRequest
}

var _ ContainerStatsRequest = &ContainerStatsRequest_19{}

func (o *ContainerStatsRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStatsRequest{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStatsRequest)
	}
}
func (o *ContainerStatsRequest_19) Unwrap() interface{}      { return o.inner }
func (o *ContainerStatsRequest_19) ContainerId() string      { return o.inner.ContainerId }
func (o *ContainerStatsRequest_19) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ContainerStatsResponse_19 struct {
	inner *runtimeapi.ContainerStatsResponse
}

var _ ContainerStatsResponse = &ContainerStatsResponse_19{}

func (o *ContainerStatsResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStatsResponse{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStatsResponse)
	}
}
func (o *ContainerStatsResponse_19) Unwrap() interface{} { return o.inner }
func (o *ContainerStatsResponse_19) Stats() ContainerStats {
	if o.inner.Stats == nil {
		return nil
	}
	return &ContainerStats_19{o.inner.Stats}
}

// ---

type ExecSyncRequest_19 struct {
	inner *runtimeapi.ExecSyncRequest
}

var _ ExecSyncRequest = &ExecSyncRequest_19{}

func (o *ExecSyncRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ExecSyncRequest{}
	} else {
		o.inner = v.(*runtimeapi.ExecSyncRequest)
	}
}
func (o *ExecSyncRequest_19) Unwrap() interface{}      { return o.inner }
func (o *ExecSyncRequest_19) ContainerId() string      { return o.inner.ContainerId }
func (o *ExecSyncRequest_19) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ExecSyncResponse_19 struct {
	inner *runtimeapi.ExecSyncResponse
}

var _ ExecSyncResponse = &ExecSyncResponse_19{}

func (o *ExecSyncResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ExecSyncResponse{}
	} else {
		o.inner = v.(*runtimeapi.ExecSyncResponse)
	}
}
func (o *ExecSyncResponse_19) Unwrap() interface{} { return o.inner }

// ---

type ExecRequest_19 struct {
	inner *runtimeapi.ExecRequest
}

var _ ExecRequest = &ExecRequest_19{}

func (o *ExecRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ExecRequest{}
	} else {
		o.inner = v.(*runtimeapi.ExecRequest)
		// FIXME: this is needed to support k8s 1.8
		if !o.inner.Stdout && !o.inner.Stderr {
			o.inner.Stdout = true
			o.inner.Stderr = true
		}
	}
}
func (o *ExecRequest_19) Unwrap() interface{}      { return o.inner }
func (o *ExecRequest_19) ContainerId() string      { return o.inner.ContainerId }
func (o *ExecRequest_19) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ExecResponse_19 struct {
	inner *runtimeapi.ExecResponse
}

var _ ExecResponse = &ExecResponse_19{}

func (o *ExecResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ExecResponse{}
	} else {
		o.inner = v.(*runtimeapi.ExecResponse)
	}
}
func (o *ExecResponse_19) Unwrap() interface{} { return o.inner }
func (o *ExecResponse_19) Url() string         { return o.inner.Url }
func (o *ExecResponse_19) SetUrl(url string)   { o.inner.Url = url }

// ---

type AttachRequest_19 struct {
	inner *runtimeapi.AttachRequest
}

var _ AttachRequest = &AttachRequest_19{}

func (o *AttachRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.AttachRequest{}
	} else {
		o.inner = v.(*runtimeapi.AttachRequest)
		// FIXME: this is needed to support k8s 1.8
		if !o.inner.Stdout && !o.inner.Stderr {
			o.inner.Stdout = true
			o.inner.Stderr = true
		}
	}
}
func (o *AttachRequest_19) Unwrap() interface{}      { return o.inner }
func (o *AttachRequest_19) ContainerId() string      { return o.inner.ContainerId }
func (o *AttachRequest_19) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type AttachResponse_19 struct {
	inner *runtimeapi.AttachResponse
}

var _ AttachResponse = &AttachResponse_19{}

func (o *AttachResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.AttachResponse{}
	} else {
		o.inner = v.(*runtimeapi.AttachResponse)
	}
}
func (o *AttachResponse_19) Unwrap() interface{} { return o.inner }
func (o *AttachResponse_19) Url() string         { return o.inner.Url }
func (o *AttachResponse_19) SetUrl(url string)   { o.inner.Url = url }

// ---

type PortForwardRequest_19 struct {
	inner *runtimeapi.PortForwardRequest
}

var _ PortForwardRequest = &PortForwardRequest_19{}

func (o *PortForwardRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PortForwardRequest{}
	} else {
		o.inner = v.(*runtimeapi.PortForwardRequest)
	}
}
func (o *PortForwardRequest_19) Unwrap() interface{}  { return o.inner }
func (o *PortForwardRequest_19) PodSandboxId() string { return o.inner.PodSandboxId }
func (o *PortForwardRequest_19) SetPodSandboxId(podSandboxId string) {
	o.inner.PodSandboxId = podSandboxId
}

// ---

type PortForwardResponse_19 struct {
	inner *runtimeapi.PortForwardResponse
}

var _ PortForwardResponse = &PortForwardResponse_19{}

func (o *PortForwardResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PortForwardResponse{}
	} else {
		o.inner = v.(*runtimeapi.PortForwardResponse)
	}
}
func (o *PortForwardResponse_19) Unwrap() interface{} { return o.inner }
func (o *PortForwardResponse_19) Url() string         { return o.inner.Url }
func (o *PortForwardResponse_19) SetUrl(url string)   { o.inner.Url = url }

// ---

type ListImagesRequest_19 struct {
	inner *runtimeapi.ListImagesRequest
}

var _ ListImagesRequest = &ListImagesRequest_19{}

func (o *ListImagesRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListImagesRequest{}
	} else {
		o.inner = v.(*runtimeapi.ListImagesRequest)
	}
}
func (o *ListImagesRequest_19) Unwrap() interface{} { return o.inner }
func (o *ListImagesRequest_19) ImageFilter() string { return o.inner.Filter.GetImage().GetImage() }
func (o *ListImagesRequest_19) SetImageFilter(image string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ImageFilter{
			Image: &runtimeapi.ImageSpec{Image: image},
		}
	} else {
		o.inner.Filter.Image = &runtimeapi.ImageSpec{Image: image}
	}
}

// ---

type ListImagesResponse_19 struct {
	inner *runtimeapi.ListImagesResponse
}

var _ ListImagesResponse = &ListImagesResponse_19{}

func (o *ListImagesResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListImagesResponse{}
	} else {
		o.inner = v.(*runtimeapi.ListImagesResponse)
	}
}
func (o *ListImagesResponse_19) Unwrap() interface{} { return o.inner }
func (o *ListImagesResponse_19) Items() []CRIObject {
	var r []CRIObject
	for _, image := range o.inner.Images {
		r = append(r, &Image_19{image})
	}
	return r
}
func (o *ListImagesResponse_19) SetItems(items []CRIObject) {
	o.inner.Images = nil
	for _, wrapped := range items {
		o.inner.Images = append(o.inner.Images, wrapped.Unwrap().(*runtimeapi.Image))
	}
}

// ---

type ImageStatusRequest_19 struct {
	inner *runtimeapi.ImageStatusRequest
}

var _ ImageStatusRequest = &ImageStatusRequest_19{}

func (o *ImageStatusRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ImageStatusRequest{}
	} else {
		o.inner = v.(*runtimeapi.ImageStatusRequest)
	}
}
func (o *ImageStatusRequest_19) Unwrap() interface{} { return o.inner }
func (o *ImageStatusRequest_19) Image() string       { return o.inner.Image.GetImage() }
func (o *ImageStatusRequest_19) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type ImageStatusResponse_19 struct {
	inner *runtimeapi.ImageStatusResponse
}

var _ ImageStatusResponse = &ImageStatusResponse_19{}

func (o *ImageStatusResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ImageStatusResponse{}
	} else {
		o.inner = v.(*runtimeapi.ImageStatusResponse)
	}
}
func (o *ImageStatusResponse_19) Unwrap() interface{} { return o.inner }
func (o *ImageStatusResponse_19) Image() Image {
	if o.inner.Image == nil {
		return nil
	}
	return &Image_19{o.inner.Image}
}
func (o *ImageStatusResponse_19) SetImage(image Image) {
	o.inner.Image = image.Unwrap().(*runtimeapi.Image)
}

// ---

type PullImageRequest_19 struct {
	inner *runtimeapi.PullImageRequest
}

var _ PullImageRequest = &PullImageRequest_19{}

func (o *PullImageRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PullImageRequest{}
	} else {
		o.inner = v.(*runtimeapi.PullImageRequest)
	}
}
func (o *PullImageRequest_19) Unwrap() interface{} { return o.inner }
func (o *PullImageRequest_19) Image() string       { return o.inner.Image.GetImage() }
func (o *PullImageRequest_19) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type PullImageResponse_19 struct {
	inner *runtimeapi.PullImageResponse
}

var _ PullImageResponse = &PullImageResponse_19{}

func (o *PullImageResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PullImageResponse{}
	} else {
		o.inner = v.(*runtimeapi.PullImageResponse)
	}
}
func (o *PullImageResponse_19) Unwrap() interface{}   { return o.inner }
func (o *PullImageResponse_19) Image() string         { return o.inner.ImageRef }
func (o *PullImageResponse_19) SetImage(image string) { o.inner.ImageRef = image }

// ---

type RemoveImageRequest_19 struct {
	inner *runtimeapi.RemoveImageRequest
}

var _ RemoveImageRequest = &RemoveImageRequest_19{}

func (o *RemoveImageRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemoveImageRequest{}
	} else {
		o.inner = v.(*runtimeapi.RemoveImageRequest)
	}
}
func (o *RemoveImageRequest_19) Unwrap() interface{} { return o.inner }
func (o *RemoveImageRequest_19) Image() string       { return o.inner.Image.GetImage() }
func (o *RemoveImageRequest_19) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type RemoveImageResponse_19 struct {
	inner *runtimeapi.RemoveImageResponse
}

var _ RemoveImageResponse = &RemoveImageResponse_19{}

func (o *RemoveImageResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemoveImageResponse{}
	} else {
		o.inner = v.(*runtimeapi.RemoveImageResponse)
	}
}
func (o *RemoveImageResponse_19) Unwrap() interface{} { return o.inner }

// ---

type ImageFsInfoRequest_19 struct {
	inner *runtimeapi.ImageFsInfoRequest
}

var _ ImageFsInfoRequest = &ImageFsInfoRequest_19{}

func (o *ImageFsInfoRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ImageFsInfoRequest{}
	} else {
		o.inner = v.(*runtimeapi.ImageFsInfoRequest)
	}
}
func (o *ImageFsInfoRequest_19) Unwrap() interface{} { return o.inner }

// ---

type ImageFsInfoResponse_19 struct {
	inner *runtimeapi.ImageFsInfoResponse
}

var _ ImageFsInfoResponse = &ImageFsInfoResponse_19{}

func (o *ImageFsInfoResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ImageFsInfoResponse{}
	} else {
		o.inner = v.(*runtimeapi.ImageFsInfoResponse)
	}
}
func (o *ImageFsInfoResponse_19) Unwrap() interface{} { return o.inner }
func (o *ImageFsInfoResponse_19) Items() []CRIObject {
	var r []CRIObject
	for _, fs := range o.inner.ImageFilesystems {
		r = append(r, &FilesystemUsage_19{fs})
	}
	return r
}
func (o *ImageFsInfoResponse_19) SetItems(items []CRIObject) {
	o.inner.ImageFilesystems = nil
	for _, wrapped := range items {
		o.inner.ImageFilesystems = append(o.inner.ImageFilesystems, wrapped.Unwrap().(*runtimeapi.FilesystemUsage))
	}
}

// --- 1.8+ only ---

type UpdateContainerResourcesRequest_19 struct {
	inner *runtimeapi.UpdateContainerResourcesRequest
}

var _ UpdateContainerResourcesRequest = &UpdateContainerResourcesRequest_19{}

func (o *UpdateContainerResourcesRequest_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.UpdateContainerResourcesRequest{}
	} else {
		o.inner = v.(*runtimeapi.UpdateContainerResourcesRequest)
	}
}
func (o *UpdateContainerResourcesRequest_19) Unwrap() interface{}      { return o.inner }
func (o *UpdateContainerResourcesRequest_19) ContainerId() string      { return o.inner.ContainerId }
func (o *UpdateContainerResourcesRequest_19) SetContainerId(id string) { o.inner.ContainerId = id }

// --- 1.8+ only ---

type UpdateContainerResourcesResponse_19 struct {
	inner *runtimeapi.UpdateContainerResourcesResponse
}

var _ UpdateContainerResourcesResponse = &UpdateContainerResourcesResponse_19{}

func (o *UpdateContainerResourcesResponse_19) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.UpdateContainerResourcesResponse{}
	} else {
		o.inner = v.(*runtimeapi.UpdateContainerResourcesResponse)
	}
}
func (o *UpdateContainerResourcesResponse_19) Unwrap() interface{} { return o.inner }

// ---

var cri19typeMatcher = newTypeMatcher()

func init() {
	cri19typeMatcher.registerTypes(
		&PodSandbox_19{},
		&Container_19{},
		&Image_19{},
		&PodSandboxStatus_19{},
		&ContainerStatus_19{},
		&ContainerStats_19{},
		&FilesystemUsage_19{},
		&VersionRequest_19{},
		&VersionResponse_19{},
		&StatusRequest_19{},
		&StatusResponse_19{},
		&UpdateRuntimeConfigRequest_19{},
		&UpdateRuntimeConfigResponse_19{},
		&RunPodSandboxRequest_19{},
		&RunPodSandboxResponse_19{},
		&ListPodSandboxRequest_19{},
		&ListPodSandboxResponse_19{},
		&StopPodSandboxRequest_19{},
		&StopPodSandboxResponse_19{},
		&RemovePodSandboxRequest_19{},
		&RemovePodSandboxResponse_19{},
		&PodSandboxStatusRequest_19{},
		&PodSandboxStatusResponse_19{},
		&CreateContainerRequest_19{},
		&CreateContainerResponse_19{},
		&ListContainersRequest_19{},
		&ListContainersResponse_19{},
		&ListContainerStatsRequest_19{},
		&ListContainerStatsResponse_19{},
		&StartContainerRequest_19{},
		&StartContainerResponse_19{},
		&StopContainerRequest_19{},
		&StopContainerResponse_19{},
		&RemoveContainerRequest_19{},
		&RemoveContainerResponse_19{},
		&ContainerStatusRequest_19{},
		&ContainerStatusResponse_19{},
		&ContainerStatsRequest_19{},
		&ContainerStatsResponse_19{},
		&ExecSyncRequest_19{},
		&ExecSyncResponse_19{},
		&ExecRequest_19{},
		&ExecResponse_19{},
		&AttachRequest_19{},
		&AttachResponse_19{},
		&PortForwardRequest_19{},
		&PortForwardResponse_19{},
		&ListImagesRequest_19{},
		&ListImagesResponse_19{},
		&ImageStatusRequest_19{},
		&ImageStatusResponse_19{},
		&PullImageRequest_19{},
		&PullImageResponse_19{},
		&RemoveImageRequest_19{},
		&RemoveImageResponse_19{},
		&ImageFsInfoRequest_19{},
		&ImageFsInfoResponse_19{},
		&UpdateContainerResourcesRequest_19{},
		&UpdateContainerResourcesResponse_19{},
	)
}

// CRI19 denotes CRI version 1.9 that's compatible with k8s 1.7, 1.8 and 1.9.
type CRI19 struct{}

var _ CRIVersion = &CRI19{}

func (c *CRI19) Register(server *grpc.Server) {
	runtimeapi.RegisterDummyRuntimeServiceServer(server)
	runtimeapi.RegisterDummyImageServiceServer(server)
}

func (c *CRI19) ProbeRequest() (interface{}, interface{}) {
	return &runtimeapi.VersionRequest{}, &runtimeapi.VersionResponse{}
}

func (c *CRI19) WrapObject(o interface{}) (CRIObject, CRIObject, error) {
	return wrapUsingMatcher(cri19typeMatcher, o)
}

func (c *CRI19) ProtoPackage() string { return "runtime" }

func (c *CRI19) UpgradesTo() CRIVersion {
	return &CRI110{}
}
