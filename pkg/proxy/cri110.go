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

	runtimeapi "github.com/Mirantis/criproxy/pkg/runtimeapis/v1_12"
)

// ---

type PodSandbox_112 struct {
	inner *runtimeapi.PodSandbox
}

var _ PodSandbox = &PodSandbox_112{}

func (o *PodSandbox_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PodSandbox{}
	} else {
		o.inner = v.(*runtimeapi.PodSandbox)
	}
}
func (o *PodSandbox_112) Unwrap() interface{} { return o.inner }
func (o *PodSandbox_112) Copy() PodSandbox    { r := *o.inner; return &PodSandbox_112{&r} }
func (o *PodSandbox_112) Id() string          { return o.inner.Id }
func (o *PodSandbox_112) SetId(id string)     { o.inner.Id = id }

type Container_112 struct {
	inner *runtimeapi.Container
}

// ---

var _ Container = &Container_112{}

func (o *Container_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.Container{}
	} else {
		o.inner = v.(*runtimeapi.Container)
	}
}
func (o *Container_112) Unwrap() interface{}       { return o.inner }
func (o *Container_112) Copy() Container           { r := *o.inner; return &Container_112{&r} }
func (o *Container_112) Id() string                { return o.inner.Id }
func (o *Container_112) SetId(id string)           { o.inner.Id = id }
func (o *Container_112) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *Container_112) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }
func (o *Container_112) Image() string             { return o.inner.Image.GetImage() }
func (o *Container_112) SetImage(image string)     { o.inner.Image = &runtimeapi.ImageSpec{Image: image} }

// ---

type Image_112 struct {
	inner *runtimeapi.Image
}

var _ Image = &Image_112{}

func (o *Image_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.Image{}
	} else {
		o.inner = v.(*runtimeapi.Image)
	}
}
func (o *Image_112) Unwrap() interface{}                 { return o.inner }
func (o *Image_112) Copy() Image                         { r := *o.inner; return &Image_112{&r} }
func (o *Image_112) Id() string                          { return o.inner.Id }
func (o *Image_112) SetId(id string)                     { o.inner.Id = id }
func (o *Image_112) RepoTags() []string                  { return o.inner.RepoTags }
func (o *Image_112) SetRepoTags(repoTags []string)       { o.inner.RepoTags = repoTags }
func (o *Image_112) RepoDigests() []string               { return o.inner.RepoDigests }
func (o *Image_112) SetRepoDigests(repoDigests []string) { o.inner.RepoDigests = repoDigests }

// ---

type PodSandboxStatus_112 struct {
	inner *runtimeapi.PodSandboxStatus
}

var _ PodSandboxStatus = &PodSandboxStatus_112{}

func (o *PodSandboxStatus_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PodSandboxStatus{}
	} else {
		o.inner = v.(*runtimeapi.PodSandboxStatus)
	}
}
func (o *PodSandboxStatus_112) Unwrap() interface{} { return o.inner }
func (o *PodSandboxStatus_112) Copy() PodSandboxStatus {
	r := *o.inner
	return &PodSandboxStatus_112{&r}
}
func (o *PodSandboxStatus_112) Id() string      { return o.inner.Id }
func (o *PodSandboxStatus_112) SetId(id string) { o.inner.Id = id }

// ---

type ContainerStatus_112 struct {
	inner *runtimeapi.ContainerStatus
}

var _ ContainerStatus = &ContainerStatus_112{}

func (o *ContainerStatus_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStatus{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStatus)
	}
}
func (o *ContainerStatus_112) Unwrap() interface{}   { return o.inner }
func (o *ContainerStatus_112) Copy() ContainerStatus { r := *o.inner; return &ContainerStatus_112{&r} }
func (o *ContainerStatus_112) Id() string            { return o.inner.Id }
func (o *ContainerStatus_112) SetId(id string)       { o.inner.Id = id }
func (o *ContainerStatus_112) Image() string         { return o.inner.Image.GetImage() }
func (o *ContainerStatus_112) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type ContainerStats_112 struct {
	inner *runtimeapi.ContainerStats
}

var _ ContainerStats = &ContainerStats_112{}

func (o *ContainerStats_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStats{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStats)
	}
}
func (o *ContainerStats_112) Unwrap() interface{}  { return o.inner }
func (o *ContainerStats_112) Copy() ContainerStats { r := *o.inner; return &ContainerStats_112{&r} }
func (o *ContainerStats_112) Id() string           { return o.inner.Attributes.GetId() }
func (o *ContainerStats_112) SetId(id string) {
	if o.inner.Attributes == nil {
		o.inner.Attributes = &runtimeapi.ContainerAttributes{Id: id}
	} else {
		o.inner.Attributes.Id = id
	}
}

// ---

type FilesystemUsage_112 struct {
	inner *runtimeapi.FilesystemUsage
}

func (o *FilesystemUsage_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.FilesystemUsage{}
	} else {
		o.inner = v.(*runtimeapi.FilesystemUsage)
	}
}
func (o *FilesystemUsage_112) Unwrap() interface{} { return o.inner }

// ---

type VersionRequest_112 struct {
	inner *runtimeapi.VersionRequest
}

var _ VersionRequest = &VersionRequest_112{}

func (o *VersionRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.VersionRequest{}
	} else {
		o.inner = v.(*runtimeapi.VersionRequest)
	}
}
func (o *VersionRequest_112) Unwrap() interface{} { return o.inner }

// ---

type VersionResponse_112 struct {
	inner *runtimeapi.VersionResponse
}

var _ VersionResponse = &VersionResponse_112{}

func (o *VersionResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.VersionResponse{}
	} else {
		o.inner = v.(*runtimeapi.VersionResponse)
	}
}
func (o *VersionResponse_112) Unwrap() interface{} { return o.inner }

// ---

type StatusRequest_112 struct {
	inner *runtimeapi.StatusRequest
}

var _ StatusRequest = &StatusRequest_112{}

func (o *StatusRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StatusRequest{}
	} else {
		o.inner = v.(*runtimeapi.StatusRequest)
	}
}
func (o *StatusRequest_112) Unwrap() interface{} { return o.inner }

// ---

type StatusResponse_112 struct {
	inner *runtimeapi.StatusResponse
}

var _ StatusResponse = &StatusResponse_112{}

func (o *StatusResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StatusResponse{}
	} else {
		o.inner = v.(*runtimeapi.StatusResponse)
	}
}
func (o *StatusResponse_112) Unwrap() interface{} { return o.inner }

// ---

type UpdateRuntimeConfigRequest_112 struct {
	inner *runtimeapi.UpdateRuntimeConfigRequest
}

var _ UpdateRuntimeConfigRequest = &UpdateRuntimeConfigRequest_112{}

func (o *UpdateRuntimeConfigRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.UpdateRuntimeConfigRequest{}
	} else {
		o.inner = v.(*runtimeapi.UpdateRuntimeConfigRequest)
	}
}
func (o *UpdateRuntimeConfigRequest_112) Unwrap() interface{} { return o.inner }

// ---

type UpdateRuntimeConfigResponse_112 struct {
	inner *runtimeapi.UpdateRuntimeConfigResponse
}

var _ UpdateRuntimeConfigResponse = &UpdateRuntimeConfigResponse_112{}

func (o *UpdateRuntimeConfigResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.UpdateRuntimeConfigResponse{}
	} else {
		o.inner = v.(*runtimeapi.UpdateRuntimeConfigResponse)
	}
}
func (o *UpdateRuntimeConfigResponse_112) Unwrap() interface{} { return o.inner }

// ---

type RunPodSandboxRequest_112 struct {
	inner *runtimeapi.RunPodSandboxRequest
}

var _ RunPodSandboxRequest = &RunPodSandboxRequest_112{}

func (o *RunPodSandboxRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RunPodSandboxRequest{}
	} else {
		o.inner = v.(*runtimeapi.RunPodSandboxRequest)
	}
}
func (o *RunPodSandboxRequest_112) Unwrap() interface{} { return o.inner }
func (o *RunPodSandboxRequest_112) GetAnnotations() map[string]string {
	return o.inner.Config.GetAnnotations()
}

// ---

type RunPodSandboxResponse_112 struct {
	inner *runtimeapi.RunPodSandboxResponse
}

var _ RunPodSandboxResponse = &RunPodSandboxResponse_112{}

func (o *RunPodSandboxResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RunPodSandboxResponse{}
	} else {
		o.inner = v.(*runtimeapi.RunPodSandboxResponse)
	}
}
func (o *RunPodSandboxResponse_112) Unwrap() interface{}       { return o.inner }
func (o *RunPodSandboxResponse_112) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *RunPodSandboxResponse_112) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type ListPodSandboxRequest_112 struct {
	inner *runtimeapi.ListPodSandboxRequest
}

var _ ListPodSandboxRequest = &ListPodSandboxRequest_112{}

func (o *ListPodSandboxRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListPodSandboxRequest{}
	} else {
		o.inner = v.(*runtimeapi.ListPodSandboxRequest)
	}
}
func (o *ListPodSandboxRequest_112) Unwrap() interface{} { return o.inner }
func (o *ListPodSandboxRequest_112) IdFilter() string {
	return o.inner.Filter.GetId()
}

func (o *ListPodSandboxRequest_112) SetIdFilter(id string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.PodSandboxFilter{Id: id}
	} else {
		o.inner.Filter.Id = id
	}
}

// ---

type ListPodSandboxResponse_112 struct {
	inner *runtimeapi.ListPodSandboxResponse
}

var _ ListPodSandboxResponse = &ListPodSandboxResponse_112{}

func (o *ListPodSandboxResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListPodSandboxResponse{}
	} else {
		o.inner = v.(*runtimeapi.ListPodSandboxResponse)
	}
}
func (o *ListPodSandboxResponse_112) Unwrap() interface{} { return o.inner }
func (o *ListPodSandboxResponse_112) Items() []CRIObject {
	var r []CRIObject
	for _, sandbox := range o.inner.Items {
		r = append(r, &PodSandbox_112{sandbox})
	}
	return r
}
func (o *ListPodSandboxResponse_112) SetItems(items []CRIObject) {
	o.inner.Items = nil
	for _, wrapped := range items {
		o.inner.Items = append(o.inner.Items, wrapped.Unwrap().(*runtimeapi.PodSandbox))
	}
}

// ---

type StopPodSandboxRequest_112 struct {
	inner *runtimeapi.StopPodSandboxRequest
}

var _ StopPodSandboxRequest = &StopPodSandboxRequest_112{}

func (o *StopPodSandboxRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StopPodSandboxRequest{}
	} else {
		o.inner = v.(*runtimeapi.StopPodSandboxRequest)
	}
}
func (o *StopPodSandboxRequest_112) Unwrap() interface{}       { return o.inner }
func (o *StopPodSandboxRequest_112) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *StopPodSandboxRequest_112) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type StopPodSandboxResponse_112 struct {
	inner *runtimeapi.StopPodSandboxResponse
}

var _ StopPodSandboxResponse = &StopPodSandboxResponse_112{}

func (o *StopPodSandboxResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StopPodSandboxResponse{}
	} else {
		o.inner = v.(*runtimeapi.StopPodSandboxResponse)
	}
}
func (o *StopPodSandboxResponse_112) Unwrap() interface{} { return o.inner }

// ---

type RemovePodSandboxRequest_112 struct {
	inner *runtimeapi.RemovePodSandboxRequest
}

var _ RemovePodSandboxRequest = &RemovePodSandboxRequest_112{}

func (o *RemovePodSandboxRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemovePodSandboxRequest{}
	} else {
		o.inner = v.(*runtimeapi.RemovePodSandboxRequest)
	}
}
func (o *RemovePodSandboxRequest_112) Unwrap() interface{}       { return o.inner }
func (o *RemovePodSandboxRequest_112) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *RemovePodSandboxRequest_112) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type RemovePodSandboxResponse_112 struct {
	inner *runtimeapi.RemovePodSandboxResponse
}

var _ RemovePodSandboxResponse = &RemovePodSandboxResponse_112{}

func (o *RemovePodSandboxResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemovePodSandboxResponse{}
	} else {
		o.inner = v.(*runtimeapi.RemovePodSandboxResponse)
	}
}
func (o *RemovePodSandboxResponse_112) Unwrap() interface{} { return o.inner }

// ---

type PodSandboxStatusRequest_112 struct {
	inner *runtimeapi.PodSandboxStatusRequest
}

var _ PodSandboxStatusRequest = &PodSandboxStatusRequest_112{}

func (o *PodSandboxStatusRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PodSandboxStatusRequest{}
	} else {
		o.inner = v.(*runtimeapi.PodSandboxStatusRequest)
	}
}
func (o *PodSandboxStatusRequest_112) Unwrap() interface{}       { return o.inner }
func (o *PodSandboxStatusRequest_112) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *PodSandboxStatusRequest_112) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type PodSandboxStatusResponse_112 struct {
	inner *runtimeapi.PodSandboxStatusResponse
}

var _ PodSandboxStatusResponse = &PodSandboxStatusResponse_112{}

func (o *PodSandboxStatusResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PodSandboxStatusResponse{}
	} else {
		o.inner = v.(*runtimeapi.PodSandboxStatusResponse)
	}
}
func (o *PodSandboxStatusResponse_112) Unwrap() interface{} { return o.inner }
func (o *PodSandboxStatusResponse_112) Status() PodSandboxStatus {
	if o.inner.Status == nil {
		return nil
	}
	return &PodSandboxStatus_112{o.inner.Status}
}

// ---

type CreateContainerRequest_112 struct {
	inner *runtimeapi.CreateContainerRequest
}

var _ CreateContainerRequest = &CreateContainerRequest_112{}

func (o *CreateContainerRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.CreateContainerRequest{}
	} else {
		o.inner = v.(*runtimeapi.CreateContainerRequest)
	}
}
func (o *CreateContainerRequest_112) Unwrap() interface{}       { return o.inner }
func (o *CreateContainerRequest_112) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *CreateContainerRequest_112) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }
func (o *CreateContainerRequest_112) Image() string {
	if o.inner.Config == nil {
		return ""
	}
	return o.inner.Config.Image.GetImage()
}

func (o *CreateContainerRequest_112) SetImage(image string) {
	if o.inner.Config != nil {
		o.inner.Config.Image = &runtimeapi.ImageSpec{Image: image}
	}
}

// ---

type CreateContainerResponse_112 struct {
	inner *runtimeapi.CreateContainerResponse
}

var _ CreateContainerResponse = &CreateContainerResponse_112{}

func (o *CreateContainerResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.CreateContainerResponse{}
	} else {
		o.inner = v.(*runtimeapi.CreateContainerResponse)
	}
}
func (o *CreateContainerResponse_112) Unwrap() interface{}      { return o.inner }
func (o *CreateContainerResponse_112) ContainerId() string      { return o.inner.ContainerId }
func (o *CreateContainerResponse_112) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ListContainersRequest_112 struct {
	inner *runtimeapi.ListContainersRequest
}

var _ ListContainersRequest = &ListContainersRequest_112{}

func (o *ListContainersRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListContainersRequest{}
	} else {
		o.inner = v.(*runtimeapi.ListContainersRequest)
	}
}
func (o *ListContainersRequest_112) Unwrap() interface{} { return o.inner }
func (o *ListContainersRequest_112) IdFilter() string {
	return o.inner.Filter.GetId()
}

func (o *ListContainersRequest_112) SetIdFilter(id string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerFilter{Id: id}
	} else {
		o.inner.Filter.Id = id
	}
}

func (o *ListContainersRequest_112) PodSandboxIdFilter() string {
	return o.inner.Filter.GetPodSandboxId()
}

func (o *ListContainersRequest_112) SetPodSandboxIdFilter(podSandboxId string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerFilter{Id: podSandboxId}
	} else {
		o.inner.Filter.PodSandboxId = podSandboxId
	}
}

// ---

type ListContainersResponse_112 struct {
	inner *runtimeapi.ListContainersResponse
}

var _ ListContainersResponse = &ListContainersResponse_112{}

func (o *ListContainersResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListContainersResponse{}
	} else {
		o.inner = v.(*runtimeapi.ListContainersResponse)
	}
}
func (o *ListContainersResponse_112) Unwrap() interface{} { return o.inner }
func (o *ListContainersResponse_112) Items() []CRIObject {
	var r []CRIObject
	for _, container := range o.inner.Containers {
		r = append(r, &Container_112{container})
	}
	return r
}
func (o *ListContainersResponse_112) SetItems(items []CRIObject) {
	o.inner.Containers = nil
	for _, wrapped := range items {
		o.inner.Containers = append(o.inner.Containers, wrapped.Unwrap().(*runtimeapi.Container))
	}
}

// ---

type ListContainerStatsRequest_112 struct {
	inner *runtimeapi.ListContainerStatsRequest
}

var _ ListContainerStatsRequest = &ListContainerStatsRequest_112{}

func (o *ListContainerStatsRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListContainerStatsRequest{}
	} else {
		o.inner = v.(*runtimeapi.ListContainerStatsRequest)
	}
}
func (o *ListContainerStatsRequest_112) Unwrap() interface{} { return o.inner }
func (o *ListContainerStatsRequest_112) IdFilter() string {
	return o.inner.Filter.GetId()
}

func (o *ListContainerStatsRequest_112) SetIdFilter(id string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerStatsFilter{Id: id}
	} else {
		o.inner.Filter.Id = id
	}
}

func (o *ListContainerStatsRequest_112) PodSandboxIdFilter() string {
	return o.inner.Filter.GetPodSandboxId()
}

func (o *ListContainerStatsRequest_112) SetPodSandboxIdFilter(podSandboxId string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerStatsFilter{Id: podSandboxId}
	} else {
		o.inner.Filter.PodSandboxId = podSandboxId
	}
}

// ---

type ListContainerStatsResponse_112 struct {
	inner *runtimeapi.ListContainerStatsResponse
}

var _ ListContainerStatsResponse = &ListContainerStatsResponse_112{}

func (o *ListContainerStatsResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListContainerStatsResponse{}
	} else {
		o.inner = v.(*runtimeapi.ListContainerStatsResponse)
	}
}
func (o *ListContainerStatsResponse_112) Unwrap() interface{} { return o.inner }
func (o *ListContainerStatsResponse_112) Items() []CRIObject {
	var r []CRIObject
	for _, stats := range o.inner.Stats {
		r = append(r, &ContainerStats_112{stats})
	}
	return r
}
func (o *ListContainerStatsResponse_112) SetItems(items []CRIObject) {
	o.inner.Stats = nil
	for _, wrapped := range items {
		o.inner.Stats = append(o.inner.Stats, wrapped.Unwrap().(*runtimeapi.ContainerStats))
	}
}

// ---

type StartContainerRequest_112 struct {
	inner *runtimeapi.StartContainerRequest
}

var _ StartContainerRequest = &StartContainerRequest_112{}

func (o *StartContainerRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StartContainerRequest{}
	} else {
		o.inner = v.(*runtimeapi.StartContainerRequest)
	}
}
func (o *StartContainerRequest_112) Unwrap() interface{}      { return o.inner }
func (o *StartContainerRequest_112) ContainerId() string      { return o.inner.ContainerId }
func (o *StartContainerRequest_112) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type StartContainerResponse_112 struct {
	inner *runtimeapi.StartContainerResponse
}

var _ StartContainerResponse = &StartContainerResponse_112{}

func (o *StartContainerResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StartContainerResponse{}
	} else {
		o.inner = v.(*runtimeapi.StartContainerResponse)
	}
}
func (o *StartContainerResponse_112) Unwrap() interface{} { return o.inner }

// ---

type StopContainerRequest_112 struct {
	inner *runtimeapi.StopContainerRequest
}

var _ StopContainerRequest = &StopContainerRequest_112{}

func (o *StopContainerRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StopContainerRequest{}
	} else {
		o.inner = v.(*runtimeapi.StopContainerRequest)
	}
}
func (o *StopContainerRequest_112) Unwrap() interface{}      { return o.inner }
func (o *StopContainerRequest_112) ContainerId() string      { return o.inner.ContainerId }
func (o *StopContainerRequest_112) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type StopContainerResponse_112 struct {
	inner *runtimeapi.StopContainerResponse
}

var _ StopContainerResponse = &StopContainerResponse_112{}

func (o *StopContainerResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StopContainerResponse{}
	} else {
		o.inner = v.(*runtimeapi.StopContainerResponse)
	}
}
func (o *StopContainerResponse_112) Unwrap() interface{} { return o.inner }

// ---

type RemoveContainerRequest_112 struct {
	inner *runtimeapi.RemoveContainerRequest
}

var _ RemoveContainerRequest = &RemoveContainerRequest_112{}

func (o *RemoveContainerRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemoveContainerRequest{}
	} else {
		o.inner = v.(*runtimeapi.RemoveContainerRequest)
	}
}
func (o *RemoveContainerRequest_112) Unwrap() interface{}      { return o.inner }
func (o *RemoveContainerRequest_112) ContainerId() string      { return o.inner.ContainerId }
func (o *RemoveContainerRequest_112) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type RemoveContainerResponse_112 struct {
	inner *runtimeapi.RemoveContainerResponse
}

var _ RemoveContainerResponse = &RemoveContainerResponse_112{}

func (o *RemoveContainerResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemoveContainerResponse{}
	} else {
		o.inner = v.(*runtimeapi.RemoveContainerResponse)
	}
}
func (o *RemoveContainerResponse_112) Unwrap() interface{} { return o.inner }

// ---

type ReopenContainerLogRequest_112 struct {
	inner *runtimeapi.ReopenContainerLogRequest
}

var _ ReopenContainerLogRequest = &ReopenContainerLogRequest_112{}

func (o *ReopenContainerLogRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ReopenContainerLogRequest{}
	} else {
		o.inner = v.(*runtimeapi.ReopenContainerLogRequest)
	}
}
func (o *ReopenContainerLogRequest_112) Unwrap() interface{}      { return o.inner }
func (o *ReopenContainerLogRequest_112) ContainerId() string      { return o.inner.ContainerId }
func (o *ReopenContainerLogRequest_112) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ReopenContainerLogResponse_112 struct {
	inner *runtimeapi.ReopenContainerLogResponse
}

var _ ReopenContainerLogResponse = &ReopenContainerLogResponse_112{}

func (o *ReopenContainerLogResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ReopenContainerLogResponse{}
	} else {
		o.inner = v.(*runtimeapi.ReopenContainerLogResponse)
	}
}
func (o *ReopenContainerLogResponse_112) Unwrap() interface{} { return o.inner }

// ---

type ContainerStatusRequest_112 struct {
	inner *runtimeapi.ContainerStatusRequest
}

var _ ContainerStatusRequest = &ContainerStatusRequest_112{}

func (o *ContainerStatusRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStatusRequest{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStatusRequest)
	}
}
func (o *ContainerStatusRequest_112) Unwrap() interface{}      { return o.inner }
func (o *ContainerStatusRequest_112) ContainerId() string      { return o.inner.ContainerId }
func (o *ContainerStatusRequest_112) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ContainerStatusResponse_112 struct {
	inner *runtimeapi.ContainerStatusResponse
}

var _ ContainerStatusResponse = &ContainerStatusResponse_112{}

func (o *ContainerStatusResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStatusResponse{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStatusResponse)
	}
}
func (o *ContainerStatusResponse_112) Unwrap() interface{} { return o.inner }
func (o *ContainerStatusResponse_112) Status() ContainerStatus {
	if o.inner.Status == nil {
		return nil
	}
	return &ContainerStatus_112{o.inner.Status}
}

// ---

type ContainerStatsRequest_112 struct {
	inner *runtimeapi.ContainerStatsRequest
}

var _ ContainerStatsRequest = &ContainerStatsRequest_112{}

func (o *ContainerStatsRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStatsRequest{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStatsRequest)
	}
}
func (o *ContainerStatsRequest_112) Unwrap() interface{}      { return o.inner }
func (o *ContainerStatsRequest_112) ContainerId() string      { return o.inner.ContainerId }
func (o *ContainerStatsRequest_112) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ContainerStatsResponse_112 struct {
	inner *runtimeapi.ContainerStatsResponse
}

var _ ContainerStatsResponse = &ContainerStatsResponse_112{}

func (o *ContainerStatsResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStatsResponse{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStatsResponse)
	}
}
func (o *ContainerStatsResponse_112) Unwrap() interface{} { return o.inner }
func (o *ContainerStatsResponse_112) Stats() ContainerStats {
	if o.inner.Stats == nil {
		return nil
	}
	return &ContainerStats_112{o.inner.Stats}
}

// ---

type ExecSyncRequest_112 struct {
	inner *runtimeapi.ExecSyncRequest
}

var _ ExecSyncRequest = &ExecSyncRequest_112{}

func (o *ExecSyncRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ExecSyncRequest{}
	} else {
		o.inner = v.(*runtimeapi.ExecSyncRequest)
	}
}
func (o *ExecSyncRequest_112) Unwrap() interface{}      { return o.inner }
func (o *ExecSyncRequest_112) ContainerId() string      { return o.inner.ContainerId }
func (o *ExecSyncRequest_112) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ExecSyncResponse_112 struct {
	inner *runtimeapi.ExecSyncResponse
}

var _ ExecSyncResponse = &ExecSyncResponse_112{}

func (o *ExecSyncResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ExecSyncResponse{}
	} else {
		o.inner = v.(*runtimeapi.ExecSyncResponse)
	}
}
func (o *ExecSyncResponse_112) Unwrap() interface{} { return o.inner }

// ---

type ExecRequest_112 struct {
	inner *runtimeapi.ExecRequest
}

var _ ExecRequest = &ExecRequest_112{}

func (o *ExecRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ExecRequest{}
	} else {
		o.inner = v.(*runtimeapi.ExecRequest)
	}
}
func (o *ExecRequest_112) Unwrap() interface{}      { return o.inner }
func (o *ExecRequest_112) ContainerId() string      { return o.inner.ContainerId }
func (o *ExecRequest_112) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ExecResponse_112 struct {
	inner *runtimeapi.ExecResponse
}

var _ ExecResponse = &ExecResponse_112{}

func (o *ExecResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ExecResponse{}
	} else {
		o.inner = v.(*runtimeapi.ExecResponse)
	}
}
func (o *ExecResponse_112) Unwrap() interface{} { return o.inner }
func (o *ExecResponse_112) Url() string         { return o.inner.Url }
func (o *ExecResponse_112) SetUrl(url string)   { o.inner.Url = url }

// ---

type AttachRequest_112 struct {
	inner *runtimeapi.AttachRequest
}

var _ AttachRequest = &AttachRequest_112{}

func (o *AttachRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.AttachRequest{}
	} else {
		o.inner = v.(*runtimeapi.AttachRequest)
	}
}
func (o *AttachRequest_112) Unwrap() interface{}      { return o.inner }
func (o *AttachRequest_112) ContainerId() string      { return o.inner.ContainerId }
func (o *AttachRequest_112) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type AttachResponse_112 struct {
	inner *runtimeapi.AttachResponse
}

var _ AttachResponse = &AttachResponse_112{}

func (o *AttachResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.AttachResponse{}
	} else {
		o.inner = v.(*runtimeapi.AttachResponse)
	}
}
func (o *AttachResponse_112) Unwrap() interface{} { return o.inner }
func (o *AttachResponse_112) Url() string         { return o.inner.Url }
func (o *AttachResponse_112) SetUrl(url string)   { o.inner.Url = url }

// ---

type PortForwardRequest_112 struct {
	inner *runtimeapi.PortForwardRequest
}

var _ PortForwardRequest = &PortForwardRequest_112{}

func (o *PortForwardRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PortForwardRequest{}
	} else {
		o.inner = v.(*runtimeapi.PortForwardRequest)
	}
}
func (o *PortForwardRequest_112) Unwrap() interface{}  { return o.inner }
func (o *PortForwardRequest_112) PodSandboxId() string { return o.inner.PodSandboxId }
func (o *PortForwardRequest_112) SetPodSandboxId(podSandboxId string) {
	o.inner.PodSandboxId = podSandboxId
}

// ---

type PortForwardResponse_112 struct {
	inner *runtimeapi.PortForwardResponse
}

var _ PortForwardResponse = &PortForwardResponse_112{}

func (o *PortForwardResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PortForwardResponse{}
	} else {
		o.inner = v.(*runtimeapi.PortForwardResponse)
	}
}
func (o *PortForwardResponse_112) Unwrap() interface{} { return o.inner }
func (o *PortForwardResponse_112) Url() string         { return o.inner.Url }
func (o *PortForwardResponse_112) SetUrl(url string)   { o.inner.Url = url }

// ---

type ListImagesRequest_112 struct {
	inner *runtimeapi.ListImagesRequest
}

var _ ListImagesRequest = &ListImagesRequest_112{}

func (o *ListImagesRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListImagesRequest{}
	} else {
		o.inner = v.(*runtimeapi.ListImagesRequest)
	}
}
func (o *ListImagesRequest_112) Unwrap() interface{} { return o.inner }
func (o *ListImagesRequest_112) ImageFilter() string { return o.inner.Filter.GetImage().GetImage() }
func (o *ListImagesRequest_112) SetImageFilter(image string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ImageFilter{
			Image: &runtimeapi.ImageSpec{Image: image},
		}
	} else {
		o.inner.Filter.Image = &runtimeapi.ImageSpec{Image: image}
	}
}

// ---

type ListImagesResponse_112 struct {
	inner *runtimeapi.ListImagesResponse
}

var _ ListImagesResponse = &ListImagesResponse_112{}

func (o *ListImagesResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListImagesResponse{}
	} else {
		o.inner = v.(*runtimeapi.ListImagesResponse)
	}
}
func (o *ListImagesResponse_112) Unwrap() interface{} { return o.inner }
func (o *ListImagesResponse_112) Items() []CRIObject {
	var r []CRIObject
	for _, image := range o.inner.Images {
		r = append(r, &Image_112{image})
	}
	return r
}
func (o *ListImagesResponse_112) SetItems(items []CRIObject) {
	o.inner.Images = nil
	for _, wrapped := range items {
		o.inner.Images = append(o.inner.Images, wrapped.Unwrap().(*runtimeapi.Image))
	}
}

// ---

type ImageStatusRequest_112 struct {
	inner *runtimeapi.ImageStatusRequest
}

var _ ImageStatusRequest = &ImageStatusRequest_112{}

func (o *ImageStatusRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ImageStatusRequest{}
	} else {
		o.inner = v.(*runtimeapi.ImageStatusRequest)
	}
}
func (o *ImageStatusRequest_112) Unwrap() interface{} { return o.inner }
func (o *ImageStatusRequest_112) Image() string       { return o.inner.Image.GetImage() }
func (o *ImageStatusRequest_112) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type ImageStatusResponse_112 struct {
	inner *runtimeapi.ImageStatusResponse
}

var _ ImageStatusResponse = &ImageStatusResponse_112{}

func (o *ImageStatusResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ImageStatusResponse{}
	} else {
		o.inner = v.(*runtimeapi.ImageStatusResponse)
	}
}
func (o *ImageStatusResponse_112) Unwrap() interface{} { return o.inner }
func (o *ImageStatusResponse_112) Image() Image {
	if o.inner.Image == nil {
		return nil
	}
	return &Image_112{o.inner.Image}
}
func (o *ImageStatusResponse_112) SetImage(image Image) {
	o.inner.Image = image.Unwrap().(*runtimeapi.Image)
}

// ---

type PullImageRequest_112 struct {
	inner *runtimeapi.PullImageRequest
}

var _ PullImageRequest = &PullImageRequest_112{}

func (o *PullImageRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PullImageRequest{}
	} else {
		o.inner = v.(*runtimeapi.PullImageRequest)
	}
}
func (o *PullImageRequest_112) Unwrap() interface{} { return o.inner }
func (o *PullImageRequest_112) Image() string       { return o.inner.Image.GetImage() }
func (o *PullImageRequest_112) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type PullImageResponse_112 struct {
	inner *runtimeapi.PullImageResponse
}

var _ PullImageResponse = &PullImageResponse_112{}

func (o *PullImageResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PullImageResponse{}
	} else {
		o.inner = v.(*runtimeapi.PullImageResponse)
	}
}
func (o *PullImageResponse_112) Unwrap() interface{}   { return o.inner }
func (o *PullImageResponse_112) Image() string         { return o.inner.ImageRef }
func (o *PullImageResponse_112) SetImage(image string) { o.inner.ImageRef = image }

// ---

type RemoveImageRequest_112 struct {
	inner *runtimeapi.RemoveImageRequest
}

var _ RemoveImageRequest = &RemoveImageRequest_112{}

func (o *RemoveImageRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemoveImageRequest{}
	} else {
		o.inner = v.(*runtimeapi.RemoveImageRequest)
	}
}
func (o *RemoveImageRequest_112) Unwrap() interface{} { return o.inner }
func (o *RemoveImageRequest_112) Image() string       { return o.inner.Image.GetImage() }
func (o *RemoveImageRequest_112) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type RemoveImageResponse_112 struct {
	inner *runtimeapi.RemoveImageResponse
}

var _ RemoveImageResponse = &RemoveImageResponse_112{}

func (o *RemoveImageResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemoveImageResponse{}
	} else {
		o.inner = v.(*runtimeapi.RemoveImageResponse)
	}
}
func (o *RemoveImageResponse_112) Unwrap() interface{} { return o.inner }

// ---

type ImageFsInfoRequest_112 struct {
	inner *runtimeapi.ImageFsInfoRequest
}

var _ ImageFsInfoRequest = &ImageFsInfoRequest_112{}

func (o *ImageFsInfoRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ImageFsInfoRequest{}
	} else {
		o.inner = v.(*runtimeapi.ImageFsInfoRequest)
	}
}
func (o *ImageFsInfoRequest_112) Unwrap() interface{} { return o.inner }

// ---

type ImageFsInfoResponse_112 struct {
	inner *runtimeapi.ImageFsInfoResponse
}

var _ ImageFsInfoResponse = &ImageFsInfoResponse_112{}

func (o *ImageFsInfoResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ImageFsInfoResponse{}
	} else {
		o.inner = v.(*runtimeapi.ImageFsInfoResponse)
	}
}
func (o *ImageFsInfoResponse_112) Unwrap() interface{} { return o.inner }
func (o *ImageFsInfoResponse_112) Items() []CRIObject {
	var r []CRIObject
	for _, fs := range o.inner.ImageFilesystems {
		r = append(r, &FilesystemUsage_112{fs})
	}
	return r
}
func (o *ImageFsInfoResponse_112) SetItems(items []CRIObject) {
	o.inner.ImageFilesystems = nil
	for _, wrapped := range items {
		o.inner.ImageFilesystems = append(o.inner.ImageFilesystems, wrapped.Unwrap().(*runtimeapi.FilesystemUsage))
	}
}

// --- 1.8+ only ---

type UpdateContainerResourcesRequest_112 struct {
	inner *runtimeapi.UpdateContainerResourcesRequest
}

var _ UpdateContainerResourcesRequest = &UpdateContainerResourcesRequest_112{}

func (o *UpdateContainerResourcesRequest_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.UpdateContainerResourcesRequest{}
	} else {
		o.inner = v.(*runtimeapi.UpdateContainerResourcesRequest)
	}
}
func (o *UpdateContainerResourcesRequest_112) Unwrap() interface{}      { return o.inner }
func (o *UpdateContainerResourcesRequest_112) ContainerId() string      { return o.inner.ContainerId }
func (o *UpdateContainerResourcesRequest_112) SetContainerId(id string) { o.inner.ContainerId = id }

// --- 1.8+ only ---

type UpdateContainerResourcesResponse_112 struct {
	inner *runtimeapi.UpdateContainerResourcesResponse
}

var _ UpdateContainerResourcesResponse = &UpdateContainerResourcesResponse_112{}

func (o *UpdateContainerResourcesResponse_112) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.UpdateContainerResourcesResponse{}
	} else {
		o.inner = v.(*runtimeapi.UpdateContainerResourcesResponse)
	}
}
func (o *UpdateContainerResourcesResponse_112) Unwrap() interface{} { return o.inner }

// ---

var cri112typeMatcher = newTypeMatcher()

func init() {
	cri112typeMatcher.registerTypes(
		&PodSandbox_112{},
		&Container_112{},
		&Image_112{},
		&PodSandboxStatus_112{},
		&ContainerStatus_112{},
		&ContainerStats_112{},
		&FilesystemUsage_112{},
		&VersionRequest_112{},
		&VersionResponse_112{},
		&StatusRequest_112{},
		&StatusResponse_112{},
		&UpdateRuntimeConfigRequest_112{},
		&UpdateRuntimeConfigResponse_112{},
		&RunPodSandboxRequest_112{},
		&RunPodSandboxResponse_112{},
		&ListPodSandboxRequest_112{},
		&ListPodSandboxResponse_112{},
		&StopPodSandboxRequest_112{},
		&StopPodSandboxResponse_112{},
		&RemovePodSandboxRequest_112{},
		&RemovePodSandboxResponse_112{},
		&PodSandboxStatusRequest_112{},
		&PodSandboxStatusResponse_112{},
		&CreateContainerRequest_112{},
		&CreateContainerResponse_112{},
		&ListContainersRequest_112{},
		&ListContainersResponse_112{},
		&ListContainerStatsRequest_112{},
		&ListContainerStatsResponse_112{},
		&StartContainerRequest_112{},
		&StartContainerResponse_112{},
		&StopContainerRequest_112{},
		&StopContainerResponse_112{},
		&RemoveContainerRequest_112{},
		&RemoveContainerResponse_112{},
		&ReopenContainerLogRequest_112{},
		&ReopenContainerLogResponse_112{},
		&ContainerStatusRequest_112{},
		&ContainerStatusResponse_112{},
		&ContainerStatsRequest_112{},
		&ContainerStatsResponse_112{},
		&ExecSyncRequest_112{},
		&ExecSyncResponse_112{},
		&ExecRequest_112{},
		&ExecResponse_112{},
		&AttachRequest_112{},
		&AttachResponse_112{},
		&PortForwardRequest_112{},
		&PortForwardResponse_112{},
		&ListImagesRequest_112{},
		&ListImagesResponse_112{},
		&ImageStatusRequest_112{},
		&ImageStatusResponse_112{},
		&PullImageRequest_112{},
		&PullImageResponse_112{},
		&RemoveImageRequest_112{},
		&RemoveImageResponse_112{},
		&ImageFsInfoRequest_112{},
		&ImageFsInfoResponse_112{},
		&UpdateContainerResourcesRequest_112{},
		&UpdateContainerResourcesResponse_112{},
	)
}

// CRI112 denotes the CRI version 1.10
type CRI112 struct{}

var _ CRIVersion = &CRI112{}

func (c *CRI112) Register(server *grpc.Server) {
	runtimeapi.RegisterDummyRuntimeServiceServer(server)
	runtimeapi.RegisterDummyImageServiceServer(server)
}

func (c *CRI112) ProbeRequest() (interface{}, interface{}) {
	return &runtimeapi.VersionRequest{}, &runtimeapi.VersionResponse{}
}

func (c *CRI112) WrapObject(o interface{}) (CRIObject, CRIObject, error) {
	return wrapUsingMatcher(cri112typeMatcher, o)
}

func (c *CRI112) ProtoPackage() string { return "runtime.v1alpha2" }
