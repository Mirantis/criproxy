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

	runtimeapi "github.com/Mirantis/criproxy/pkg/runtimeapis/v1_10"
)

// ---

type PodSandbox_110 struct {
	inner *runtimeapi.PodSandbox
}

var _ PodSandbox = &PodSandbox_110{}

func (o *PodSandbox_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PodSandbox{}
	} else {
		o.inner = v.(*runtimeapi.PodSandbox)
	}
}
func (o *PodSandbox_110) Unwrap() interface{} { return o.inner }
func (o *PodSandbox_110) Copy() PodSandbox    { r := *o.inner; return &PodSandbox_110{&r} }
func (o *PodSandbox_110) Id() string          { return o.inner.Id }
func (o *PodSandbox_110) SetId(id string)     { o.inner.Id = id }

type Container_110 struct {
	inner *runtimeapi.Container
}

// ---

var _ Container = &Container_110{}

func (o *Container_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.Container{}
	} else {
		o.inner = v.(*runtimeapi.Container)
	}
}
func (o *Container_110) Unwrap() interface{}       { return o.inner }
func (o *Container_110) Copy() Container           { r := *o.inner; return &Container_110{&r} }
func (o *Container_110) Id() string                { return o.inner.Id }
func (o *Container_110) SetId(id string)           { o.inner.Id = id }
func (o *Container_110) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *Container_110) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }
func (o *Container_110) Image() string             { return o.inner.Image.GetImage() }
func (o *Container_110) SetImage(image string)     { o.inner.Image = &runtimeapi.ImageSpec{Image: image} }

// ---

type Image_110 struct {
	inner *runtimeapi.Image
}

var _ Image = &Image_110{}

func (o *Image_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.Image{}
	} else {
		o.inner = v.(*runtimeapi.Image)
	}
}
func (o *Image_110) Unwrap() interface{}           { return o.inner }
func (o *Image_110) Copy() Image                   { r := *o.inner; return &Image_110{&r} }
func (o *Image_110) Id() string                    { return o.inner.Id }
func (o *Image_110) SetId(id string)               { o.inner.Id = id }
func (o *Image_110) RepoTags() []string            { return o.inner.RepoTags }
func (o *Image_110) SetRepoTags(repoTags []string) { o.inner.RepoTags = repoTags }

// ---

type PodSandboxStatus_110 struct {
	inner *runtimeapi.PodSandboxStatus
}

var _ PodSandboxStatus = &PodSandboxStatus_110{}

func (o *PodSandboxStatus_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PodSandboxStatus{}
	} else {
		o.inner = v.(*runtimeapi.PodSandboxStatus)
	}
}
func (o *PodSandboxStatus_110) Unwrap() interface{} { return o.inner }
func (o *PodSandboxStatus_110) Copy() PodSandboxStatus {
	r := *o.inner
	return &PodSandboxStatus_110{&r}
}
func (o *PodSandboxStatus_110) Id() string      { return o.inner.Id }
func (o *PodSandboxStatus_110) SetId(id string) { o.inner.Id = id }

// ---

type ContainerStatus_110 struct {
	inner *runtimeapi.ContainerStatus
}

var _ ContainerStatus = &ContainerStatus_110{}

func (o *ContainerStatus_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStatus{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStatus)
	}
}
func (o *ContainerStatus_110) Unwrap() interface{}   { return o.inner }
func (o *ContainerStatus_110) Copy() ContainerStatus { r := *o.inner; return &ContainerStatus_110{&r} }
func (o *ContainerStatus_110) Id() string            { return o.inner.Id }
func (o *ContainerStatus_110) SetId(id string)       { o.inner.Id = id }
func (o *ContainerStatus_110) Image() string         { return o.inner.Image.GetImage() }
func (o *ContainerStatus_110) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type ContainerStats_110 struct {
	inner *runtimeapi.ContainerStats
}

var _ ContainerStats = &ContainerStats_110{}

func (o *ContainerStats_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStats{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStats)
	}
}
func (o *ContainerStats_110) Unwrap() interface{}  { return o.inner }
func (o *ContainerStats_110) Copy() ContainerStats { r := *o.inner; return &ContainerStats_110{&r} }
func (o *ContainerStats_110) Id() string           { return o.inner.Attributes.GetId() }
func (o *ContainerStats_110) SetId(id string) {
	if o.inner.Attributes == nil {
		o.inner.Attributes = &runtimeapi.ContainerAttributes{Id: id}
	} else {
		o.inner.Attributes.Id = id
	}
}

// ---

type FilesystemUsage_110 struct {
	inner *runtimeapi.FilesystemUsage
}

func (o *FilesystemUsage_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.FilesystemUsage{}
	} else {
		o.inner = v.(*runtimeapi.FilesystemUsage)
	}
}
func (o *FilesystemUsage_110) Unwrap() interface{} { return o.inner }

// ---

type VersionRequest_110 struct {
	inner *runtimeapi.VersionRequest
}

var _ VersionRequest = &VersionRequest_110{}

func (o *VersionRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.VersionRequest{}
	} else {
		o.inner = v.(*runtimeapi.VersionRequest)
	}
}
func (o *VersionRequest_110) Unwrap() interface{} { return o.inner }

// ---

type VersionResponse_110 struct {
	inner *runtimeapi.VersionResponse
}

var _ VersionResponse = &VersionResponse_110{}

func (o *VersionResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.VersionResponse{}
	} else {
		o.inner = v.(*runtimeapi.VersionResponse)
	}
}
func (o *VersionResponse_110) Unwrap() interface{} { return o.inner }

// ---

type StatusRequest_110 struct {
	inner *runtimeapi.StatusRequest
}

var _ StatusRequest = &StatusRequest_110{}

func (o *StatusRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StatusRequest{}
	} else {
		o.inner = v.(*runtimeapi.StatusRequest)
	}
}
func (o *StatusRequest_110) Unwrap() interface{} { return o.inner }

// ---

type StatusResponse_110 struct {
	inner *runtimeapi.StatusResponse
}

var _ StatusResponse = &StatusResponse_110{}

func (o *StatusResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StatusResponse{}
	} else {
		o.inner = v.(*runtimeapi.StatusResponse)
	}
}
func (o *StatusResponse_110) Unwrap() interface{} { return o.inner }

// ---

type UpdateRuntimeConfigRequest_110 struct {
	inner *runtimeapi.UpdateRuntimeConfigRequest
}

var _ UpdateRuntimeConfigRequest = &UpdateRuntimeConfigRequest_110{}

func (o *UpdateRuntimeConfigRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.UpdateRuntimeConfigRequest{}
	} else {
		o.inner = v.(*runtimeapi.UpdateRuntimeConfigRequest)
	}
}
func (o *UpdateRuntimeConfigRequest_110) Unwrap() interface{} { return o.inner }

// ---

type UpdateRuntimeConfigResponse_110 struct {
	inner *runtimeapi.UpdateRuntimeConfigResponse
}

var _ UpdateRuntimeConfigResponse = &UpdateRuntimeConfigResponse_110{}

func (o *UpdateRuntimeConfigResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.UpdateRuntimeConfigResponse{}
	} else {
		o.inner = v.(*runtimeapi.UpdateRuntimeConfigResponse)
	}
}
func (o *UpdateRuntimeConfigResponse_110) Unwrap() interface{} { return o.inner }

// ---

type RunPodSandboxRequest_110 struct {
	inner *runtimeapi.RunPodSandboxRequest
}

var _ RunPodSandboxRequest = &RunPodSandboxRequest_110{}

func (o *RunPodSandboxRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RunPodSandboxRequest{}
	} else {
		o.inner = v.(*runtimeapi.RunPodSandboxRequest)
	}
}
func (o *RunPodSandboxRequest_110) Unwrap() interface{} { return o.inner }
func (o *RunPodSandboxRequest_110) GetAnnotations() map[string]string {
	return o.inner.Config.GetAnnotations()
}

// ---

type RunPodSandboxResponse_110 struct {
	inner *runtimeapi.RunPodSandboxResponse
}

var _ RunPodSandboxResponse = &RunPodSandboxResponse_110{}

func (o *RunPodSandboxResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RunPodSandboxResponse{}
	} else {
		o.inner = v.(*runtimeapi.RunPodSandboxResponse)
	}
}
func (o *RunPodSandboxResponse_110) Unwrap() interface{}       { return o.inner }
func (o *RunPodSandboxResponse_110) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *RunPodSandboxResponse_110) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type ListPodSandboxRequest_110 struct {
	inner *runtimeapi.ListPodSandboxRequest
}

var _ ListPodSandboxRequest = &ListPodSandboxRequest_110{}

func (o *ListPodSandboxRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListPodSandboxRequest{}
	} else {
		o.inner = v.(*runtimeapi.ListPodSandboxRequest)
	}
}
func (o *ListPodSandboxRequest_110) Unwrap() interface{} { return o.inner }
func (o *ListPodSandboxRequest_110) IdFilter() string {
	return o.inner.Filter.GetId()
}

func (o *ListPodSandboxRequest_110) SetIdFilter(id string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.PodSandboxFilter{Id: id}
	} else {
		o.inner.Filter.Id = id
	}
}

// ---

type ListPodSandboxResponse_110 struct {
	inner *runtimeapi.ListPodSandboxResponse
}

var _ ListPodSandboxResponse = &ListPodSandboxResponse_110{}

func (o *ListPodSandboxResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListPodSandboxResponse{}
	} else {
		o.inner = v.(*runtimeapi.ListPodSandboxResponse)
	}
}
func (o *ListPodSandboxResponse_110) Unwrap() interface{} { return o.inner }
func (o *ListPodSandboxResponse_110) Items() []CRIObject {
	var r []CRIObject
	for _, sandbox := range o.inner.Items {
		r = append(r, &PodSandbox_110{sandbox})
	}
	return r
}
func (o *ListPodSandboxResponse_110) SetItems(items []CRIObject) {
	o.inner.Items = nil
	for _, wrapped := range items {
		o.inner.Items = append(o.inner.Items, wrapped.Unwrap().(*runtimeapi.PodSandbox))
	}
}

// ---

type StopPodSandboxRequest_110 struct {
	inner *runtimeapi.StopPodSandboxRequest
}

var _ StopPodSandboxRequest = &StopPodSandboxRequest_110{}

func (o *StopPodSandboxRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StopPodSandboxRequest{}
	} else {
		o.inner = v.(*runtimeapi.StopPodSandboxRequest)
	}
}
func (o *StopPodSandboxRequest_110) Unwrap() interface{}       { return o.inner }
func (o *StopPodSandboxRequest_110) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *StopPodSandboxRequest_110) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type StopPodSandboxResponse_110 struct {
	inner *runtimeapi.StopPodSandboxResponse
}

var _ StopPodSandboxResponse = &StopPodSandboxResponse_110{}

func (o *StopPodSandboxResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StopPodSandboxResponse{}
	} else {
		o.inner = v.(*runtimeapi.StopPodSandboxResponse)
	}
}
func (o *StopPodSandboxResponse_110) Unwrap() interface{} { return o.inner }

// ---

type RemovePodSandboxRequest_110 struct {
	inner *runtimeapi.RemovePodSandboxRequest
}

var _ RemovePodSandboxRequest = &RemovePodSandboxRequest_110{}

func (o *RemovePodSandboxRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemovePodSandboxRequest{}
	} else {
		o.inner = v.(*runtimeapi.RemovePodSandboxRequest)
	}
}
func (o *RemovePodSandboxRequest_110) Unwrap() interface{}       { return o.inner }
func (o *RemovePodSandboxRequest_110) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *RemovePodSandboxRequest_110) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type RemovePodSandboxResponse_110 struct {
	inner *runtimeapi.RemovePodSandboxResponse
}

var _ RemovePodSandboxResponse = &RemovePodSandboxResponse_110{}

func (o *RemovePodSandboxResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemovePodSandboxResponse{}
	} else {
		o.inner = v.(*runtimeapi.RemovePodSandboxResponse)
	}
}
func (o *RemovePodSandboxResponse_110) Unwrap() interface{} { return o.inner }

// ---

type PodSandboxStatusRequest_110 struct {
	inner *runtimeapi.PodSandboxStatusRequest
}

var _ PodSandboxStatusRequest = &PodSandboxStatusRequest_110{}

func (o *PodSandboxStatusRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PodSandboxStatusRequest{}
	} else {
		o.inner = v.(*runtimeapi.PodSandboxStatusRequest)
	}
}
func (o *PodSandboxStatusRequest_110) Unwrap() interface{}       { return o.inner }
func (o *PodSandboxStatusRequest_110) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *PodSandboxStatusRequest_110) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type PodSandboxStatusResponse_110 struct {
	inner *runtimeapi.PodSandboxStatusResponse
}

var _ PodSandboxStatusResponse = &PodSandboxStatusResponse_110{}

func (o *PodSandboxStatusResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PodSandboxStatusResponse{}
	} else {
		o.inner = v.(*runtimeapi.PodSandboxStatusResponse)
	}
}
func (o *PodSandboxStatusResponse_110) Unwrap() interface{} { return o.inner }
func (o *PodSandboxStatusResponse_110) Status() PodSandboxStatus {
	if o.inner.Status == nil {
		return nil
	}
	return &PodSandboxStatus_110{o.inner.Status}
}

// ---

type CreateContainerRequest_110 struct {
	inner *runtimeapi.CreateContainerRequest
}

var _ CreateContainerRequest = &CreateContainerRequest_110{}

func (o *CreateContainerRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.CreateContainerRequest{}
	} else {
		o.inner = v.(*runtimeapi.CreateContainerRequest)
	}
}
func (o *CreateContainerRequest_110) Unwrap() interface{}       { return o.inner }
func (o *CreateContainerRequest_110) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *CreateContainerRequest_110) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }
func (o *CreateContainerRequest_110) Image() string {
	if o.inner.Config == nil {
		return ""
	}
	return o.inner.Config.Image.GetImage()
}

func (o *CreateContainerRequest_110) SetImage(image string) {
	if o.inner.Config != nil {
		o.inner.Config.Image = &runtimeapi.ImageSpec{Image: image}
	}
}

// ---

type CreateContainerResponse_110 struct {
	inner *runtimeapi.CreateContainerResponse
}

var _ CreateContainerResponse = &CreateContainerResponse_110{}

func (o *CreateContainerResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.CreateContainerResponse{}
	} else {
		o.inner = v.(*runtimeapi.CreateContainerResponse)
	}
}
func (o *CreateContainerResponse_110) Unwrap() interface{}      { return o.inner }
func (o *CreateContainerResponse_110) ContainerId() string      { return o.inner.ContainerId }
func (o *CreateContainerResponse_110) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ListContainersRequest_110 struct {
	inner *runtimeapi.ListContainersRequest
}

var _ ListContainersRequest = &ListContainersRequest_110{}

func (o *ListContainersRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListContainersRequest{}
	} else {
		o.inner = v.(*runtimeapi.ListContainersRequest)
	}
}
func (o *ListContainersRequest_110) Unwrap() interface{} { return o.inner }
func (o *ListContainersRequest_110) IdFilter() string {
	return o.inner.Filter.GetId()
}

func (o *ListContainersRequest_110) SetIdFilter(id string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerFilter{Id: id}
	} else {
		o.inner.Filter.Id = id
	}
}

func (o *ListContainersRequest_110) PodSandboxIdFilter() string {
	return o.inner.Filter.GetPodSandboxId()
}

func (o *ListContainersRequest_110) SetPodSandboxIdFilter(podSandboxId string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerFilter{Id: podSandboxId}
	} else {
		o.inner.Filter.PodSandboxId = podSandboxId
	}
}

// ---

type ListContainersResponse_110 struct {
	inner *runtimeapi.ListContainersResponse
}

var _ ListContainersResponse = &ListContainersResponse_110{}

func (o *ListContainersResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListContainersResponse{}
	} else {
		o.inner = v.(*runtimeapi.ListContainersResponse)
	}
}
func (o *ListContainersResponse_110) Unwrap() interface{} { return o.inner }
func (o *ListContainersResponse_110) Items() []CRIObject {
	var r []CRIObject
	for _, container := range o.inner.Containers {
		r = append(r, &Container_110{container})
	}
	return r
}
func (o *ListContainersResponse_110) SetItems(items []CRIObject) {
	o.inner.Containers = nil
	for _, wrapped := range items {
		o.inner.Containers = append(o.inner.Containers, wrapped.Unwrap().(*runtimeapi.Container))
	}
}

// ---

type ListContainerStatsRequest_110 struct {
	inner *runtimeapi.ListContainerStatsRequest
}

var _ ListContainerStatsRequest = &ListContainerStatsRequest_110{}

func (o *ListContainerStatsRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListContainerStatsRequest{}
	} else {
		o.inner = v.(*runtimeapi.ListContainerStatsRequest)
	}
}
func (o *ListContainerStatsRequest_110) Unwrap() interface{} { return o.inner }
func (o *ListContainerStatsRequest_110) IdFilter() string {
	return o.inner.Filter.GetId()
}

func (o *ListContainerStatsRequest_110) SetIdFilter(id string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerStatsFilter{Id: id}
	} else {
		o.inner.Filter.Id = id
	}
}

func (o *ListContainerStatsRequest_110) PodSandboxIdFilter() string {
	return o.inner.Filter.GetPodSandboxId()
}

func (o *ListContainerStatsRequest_110) SetPodSandboxIdFilter(podSandboxId string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerStatsFilter{Id: podSandboxId}
	} else {
		o.inner.Filter.PodSandboxId = podSandboxId
	}
}

// ---

type ListContainerStatsResponse_110 struct {
	inner *runtimeapi.ListContainerStatsResponse
}

var _ ListContainerStatsResponse = &ListContainerStatsResponse_110{}

func (o *ListContainerStatsResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListContainerStatsResponse{}
	} else {
		o.inner = v.(*runtimeapi.ListContainerStatsResponse)
	}
}
func (o *ListContainerStatsResponse_110) Unwrap() interface{} { return o.inner }
func (o *ListContainerStatsResponse_110) Items() []CRIObject {
	var r []CRIObject
	for _, stats := range o.inner.Stats {
		r = append(r, &ContainerStats_110{stats})
	}
	return r
}
func (o *ListContainerStatsResponse_110) SetItems(items []CRIObject) {
	o.inner.Stats = nil
	for _, wrapped := range items {
		o.inner.Stats = append(o.inner.Stats, wrapped.Unwrap().(*runtimeapi.ContainerStats))
	}
}

// ---

type StartContainerRequest_110 struct {
	inner *runtimeapi.StartContainerRequest
}

var _ StartContainerRequest = &StartContainerRequest_110{}

func (o *StartContainerRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StartContainerRequest{}
	} else {
		o.inner = v.(*runtimeapi.StartContainerRequest)
	}
}
func (o *StartContainerRequest_110) Unwrap() interface{}      { return o.inner }
func (o *StartContainerRequest_110) ContainerId() string      { return o.inner.ContainerId }
func (o *StartContainerRequest_110) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type StartContainerResponse_110 struct {
	inner *runtimeapi.StartContainerResponse
}

var _ StartContainerResponse = &StartContainerResponse_110{}

func (o *StartContainerResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StartContainerResponse{}
	} else {
		o.inner = v.(*runtimeapi.StartContainerResponse)
	}
}
func (o *StartContainerResponse_110) Unwrap() interface{} { return o.inner }

// ---

type StopContainerRequest_110 struct {
	inner *runtimeapi.StopContainerRequest
}

var _ StopContainerRequest = &StopContainerRequest_110{}

func (o *StopContainerRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StopContainerRequest{}
	} else {
		o.inner = v.(*runtimeapi.StopContainerRequest)
	}
}
func (o *StopContainerRequest_110) Unwrap() interface{}      { return o.inner }
func (o *StopContainerRequest_110) ContainerId() string      { return o.inner.ContainerId }
func (o *StopContainerRequest_110) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type StopContainerResponse_110 struct {
	inner *runtimeapi.StopContainerResponse
}

var _ StopContainerResponse = &StopContainerResponse_110{}

func (o *StopContainerResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StopContainerResponse{}
	} else {
		o.inner = v.(*runtimeapi.StopContainerResponse)
	}
}
func (o *StopContainerResponse_110) Unwrap() interface{} { return o.inner }

// ---

type RemoveContainerRequest_110 struct {
	inner *runtimeapi.RemoveContainerRequest
}

var _ RemoveContainerRequest = &RemoveContainerRequest_110{}

func (o *RemoveContainerRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemoveContainerRequest{}
	} else {
		o.inner = v.(*runtimeapi.RemoveContainerRequest)
	}
}
func (o *RemoveContainerRequest_110) Unwrap() interface{}      { return o.inner }
func (o *RemoveContainerRequest_110) ContainerId() string      { return o.inner.ContainerId }
func (o *RemoveContainerRequest_110) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type RemoveContainerResponse_110 struct {
	inner *runtimeapi.RemoveContainerResponse
}

var _ RemoveContainerResponse = &RemoveContainerResponse_110{}

func (o *RemoveContainerResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemoveContainerResponse{}
	} else {
		o.inner = v.(*runtimeapi.RemoveContainerResponse)
	}
}
func (o *RemoveContainerResponse_110) Unwrap() interface{} { return o.inner }

// ---

type ReopenContainerLogRequest_110 struct {
	inner *runtimeapi.ReopenContainerLogRequest
}

var _ ReopenContainerLogRequest = &ReopenContainerLogRequest_110{}

func (o *ReopenContainerLogRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ReopenContainerLogRequest{}
	} else {
		o.inner = v.(*runtimeapi.ReopenContainerLogRequest)
	}
}
func (o *ReopenContainerLogRequest_110) Unwrap() interface{}      { return o.inner }
func (o *ReopenContainerLogRequest_110) ContainerId() string      { return o.inner.ContainerId }
func (o *ReopenContainerLogRequest_110) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ReopenContainerLogResponse_110 struct {
	inner *runtimeapi.ReopenContainerLogResponse
}

var _ ReopenContainerLogResponse = &ReopenContainerLogResponse_110{}

func (o *ReopenContainerLogResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ReopenContainerLogResponse{}
	} else {
		o.inner = v.(*runtimeapi.ReopenContainerLogResponse)
	}
}
func (o *ReopenContainerLogResponse_110) Unwrap() interface{} { return o.inner }

// ---

type ContainerStatusRequest_110 struct {
	inner *runtimeapi.ContainerStatusRequest
}

var _ ContainerStatusRequest = &ContainerStatusRequest_110{}

func (o *ContainerStatusRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStatusRequest{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStatusRequest)
	}
}
func (o *ContainerStatusRequest_110) Unwrap() interface{}      { return o.inner }
func (o *ContainerStatusRequest_110) ContainerId() string      { return o.inner.ContainerId }
func (o *ContainerStatusRequest_110) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ContainerStatusResponse_110 struct {
	inner *runtimeapi.ContainerStatusResponse
}

var _ ContainerStatusResponse = &ContainerStatusResponse_110{}

func (o *ContainerStatusResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStatusResponse{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStatusResponse)
	}
}
func (o *ContainerStatusResponse_110) Unwrap() interface{} { return o.inner }
func (o *ContainerStatusResponse_110) Status() ContainerStatus {
	if o.inner.Status == nil {
		return nil
	}
	return &ContainerStatus_110{o.inner.Status}
}

// ---

type ContainerStatsRequest_110 struct {
	inner *runtimeapi.ContainerStatsRequest
}

var _ ContainerStatsRequest = &ContainerStatsRequest_110{}

func (o *ContainerStatsRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStatsRequest{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStatsRequest)
	}
}
func (o *ContainerStatsRequest_110) Unwrap() interface{}      { return o.inner }
func (o *ContainerStatsRequest_110) ContainerId() string      { return o.inner.ContainerId }
func (o *ContainerStatsRequest_110) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ContainerStatsResponse_110 struct {
	inner *runtimeapi.ContainerStatsResponse
}

var _ ContainerStatsResponse = &ContainerStatsResponse_110{}

func (o *ContainerStatsResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStatsResponse{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStatsResponse)
	}
}
func (o *ContainerStatsResponse_110) Unwrap() interface{} { return o.inner }
func (o *ContainerStatsResponse_110) Stats() ContainerStats {
	if o.inner.Stats == nil {
		return nil
	}
	return &ContainerStats_110{o.inner.Stats}
}

// ---

type ExecSyncRequest_110 struct {
	inner *runtimeapi.ExecSyncRequest
}

var _ ExecSyncRequest = &ExecSyncRequest_110{}

func (o *ExecSyncRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ExecSyncRequest{}
	} else {
		o.inner = v.(*runtimeapi.ExecSyncRequest)
	}
}
func (o *ExecSyncRequest_110) Unwrap() interface{}      { return o.inner }
func (o *ExecSyncRequest_110) ContainerId() string      { return o.inner.ContainerId }
func (o *ExecSyncRequest_110) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ExecSyncResponse_110 struct {
	inner *runtimeapi.ExecSyncResponse
}

var _ ExecSyncResponse = &ExecSyncResponse_110{}

func (o *ExecSyncResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ExecSyncResponse{}
	} else {
		o.inner = v.(*runtimeapi.ExecSyncResponse)
	}
}
func (o *ExecSyncResponse_110) Unwrap() interface{} { return o.inner }

// ---

type ExecRequest_110 struct {
	inner *runtimeapi.ExecRequest
}

var _ ExecRequest = &ExecRequest_110{}

func (o *ExecRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ExecRequest{}
	} else {
		o.inner = v.(*runtimeapi.ExecRequest)
	}
}
func (o *ExecRequest_110) Unwrap() interface{}      { return o.inner }
func (o *ExecRequest_110) ContainerId() string      { return o.inner.ContainerId }
func (o *ExecRequest_110) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ExecResponse_110 struct {
	inner *runtimeapi.ExecResponse
}

var _ ExecResponse = &ExecResponse_110{}

func (o *ExecResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ExecResponse{}
	} else {
		o.inner = v.(*runtimeapi.ExecResponse)
	}
}
func (o *ExecResponse_110) Unwrap() interface{} { return o.inner }
func (o *ExecResponse_110) Url() string         { return o.inner.Url }
func (o *ExecResponse_110) SetUrl(url string)   { o.inner.Url = url }

// ---

type AttachRequest_110 struct {
	inner *runtimeapi.AttachRequest
}

var _ AttachRequest = &AttachRequest_110{}

func (o *AttachRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.AttachRequest{}
	} else {
		o.inner = v.(*runtimeapi.AttachRequest)
	}
}
func (o *AttachRequest_110) Unwrap() interface{}      { return o.inner }
func (o *AttachRequest_110) ContainerId() string      { return o.inner.ContainerId }
func (o *AttachRequest_110) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type AttachResponse_110 struct {
	inner *runtimeapi.AttachResponse
}

var _ AttachResponse = &AttachResponse_110{}

func (o *AttachResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.AttachResponse{}
	} else {
		o.inner = v.(*runtimeapi.AttachResponse)
	}
}
func (o *AttachResponse_110) Unwrap() interface{} { return o.inner }
func (o *AttachResponse_110) Url() string         { return o.inner.Url }
func (o *AttachResponse_110) SetUrl(url string)   { o.inner.Url = url }

// ---

type PortForwardRequest_110 struct {
	inner *runtimeapi.PortForwardRequest
}

var _ PortForwardRequest = &PortForwardRequest_110{}

func (o *PortForwardRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PortForwardRequest{}
	} else {
		o.inner = v.(*runtimeapi.PortForwardRequest)
	}
}
func (o *PortForwardRequest_110) Unwrap() interface{}  { return o.inner }
func (o *PortForwardRequest_110) PodSandboxId() string { return o.inner.PodSandboxId }
func (o *PortForwardRequest_110) SetPodSandboxId(podSandboxId string) {
	o.inner.PodSandboxId = podSandboxId
}

// ---

type PortForwardResponse_110 struct {
	inner *runtimeapi.PortForwardResponse
}

var _ PortForwardResponse = &PortForwardResponse_110{}

func (o *PortForwardResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PortForwardResponse{}
	} else {
		o.inner = v.(*runtimeapi.PortForwardResponse)
	}
}
func (o *PortForwardResponse_110) Unwrap() interface{} { return o.inner }
func (o *PortForwardResponse_110) Url() string         { return o.inner.Url }
func (o *PortForwardResponse_110) SetUrl(url string)   { o.inner.Url = url }

// ---

type ListImagesRequest_110 struct {
	inner *runtimeapi.ListImagesRequest
}

var _ ListImagesRequest = &ListImagesRequest_110{}

func (o *ListImagesRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListImagesRequest{}
	} else {
		o.inner = v.(*runtimeapi.ListImagesRequest)
	}
}
func (o *ListImagesRequest_110) Unwrap() interface{} { return o.inner }
func (o *ListImagesRequest_110) ImageFilter() string { return o.inner.Filter.GetImage().GetImage() }
func (o *ListImagesRequest_110) SetImageFilter(image string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ImageFilter{
			Image: &runtimeapi.ImageSpec{Image: image},
		}
	} else {
		o.inner.Filter.Image = &runtimeapi.ImageSpec{Image: image}
	}
}

// ---

type ListImagesResponse_110 struct {
	inner *runtimeapi.ListImagesResponse
}

var _ ListImagesResponse = &ListImagesResponse_110{}

func (o *ListImagesResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListImagesResponse{}
	} else {
		o.inner = v.(*runtimeapi.ListImagesResponse)
	}
}
func (o *ListImagesResponse_110) Unwrap() interface{} { return o.inner }
func (o *ListImagesResponse_110) Items() []CRIObject {
	var r []CRIObject
	for _, image := range o.inner.Images {
		r = append(r, &Image_110{image})
	}
	return r
}
func (o *ListImagesResponse_110) SetItems(items []CRIObject) {
	o.inner.Images = nil
	for _, wrapped := range items {
		o.inner.Images = append(o.inner.Images, wrapped.Unwrap().(*runtimeapi.Image))
	}
}

// ---

type ImageStatusRequest_110 struct {
	inner *runtimeapi.ImageStatusRequest
}

var _ ImageStatusRequest = &ImageStatusRequest_110{}

func (o *ImageStatusRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ImageStatusRequest{}
	} else {
		o.inner = v.(*runtimeapi.ImageStatusRequest)
	}
}
func (o *ImageStatusRequest_110) Unwrap() interface{} { return o.inner }
func (o *ImageStatusRequest_110) Image() string       { return o.inner.Image.GetImage() }
func (o *ImageStatusRequest_110) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type ImageStatusResponse_110 struct {
	inner *runtimeapi.ImageStatusResponse
}

var _ ImageStatusResponse = &ImageStatusResponse_110{}

func (o *ImageStatusResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ImageStatusResponse{}
	} else {
		o.inner = v.(*runtimeapi.ImageStatusResponse)
	}
}
func (o *ImageStatusResponse_110) Unwrap() interface{} { return o.inner }
func (o *ImageStatusResponse_110) Image() Image {
	if o.inner.Image == nil {
		return nil
	}
	return &Image_110{o.inner.Image}
}
func (o *ImageStatusResponse_110) SetImage(image Image) {
	o.inner.Image = image.Unwrap().(*runtimeapi.Image)
}

// ---

type PullImageRequest_110 struct {
	inner *runtimeapi.PullImageRequest
}

var _ PullImageRequest = &PullImageRequest_110{}

func (o *PullImageRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PullImageRequest{}
	} else {
		o.inner = v.(*runtimeapi.PullImageRequest)
	}
}
func (o *PullImageRequest_110) Unwrap() interface{} { return o.inner }
func (o *PullImageRequest_110) Image() string       { return o.inner.Image.GetImage() }
func (o *PullImageRequest_110) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type PullImageResponse_110 struct {
	inner *runtimeapi.PullImageResponse
}

var _ PullImageResponse = &PullImageResponse_110{}

func (o *PullImageResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PullImageResponse{}
	} else {
		o.inner = v.(*runtimeapi.PullImageResponse)
	}
}
func (o *PullImageResponse_110) Unwrap() interface{}   { return o.inner }
func (o *PullImageResponse_110) Image() string         { return o.inner.ImageRef }
func (o *PullImageResponse_110) SetImage(image string) { o.inner.ImageRef = image }

// ---

type RemoveImageRequest_110 struct {
	inner *runtimeapi.RemoveImageRequest
}

var _ RemoveImageRequest = &RemoveImageRequest_110{}

func (o *RemoveImageRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemoveImageRequest{}
	} else {
		o.inner = v.(*runtimeapi.RemoveImageRequest)
	}
}
func (o *RemoveImageRequest_110) Unwrap() interface{} { return o.inner }
func (o *RemoveImageRequest_110) Image() string       { return o.inner.Image.GetImage() }
func (o *RemoveImageRequest_110) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type RemoveImageResponse_110 struct {
	inner *runtimeapi.RemoveImageResponse
}

var _ RemoveImageResponse = &RemoveImageResponse_110{}

func (o *RemoveImageResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemoveImageResponse{}
	} else {
		o.inner = v.(*runtimeapi.RemoveImageResponse)
	}
}
func (o *RemoveImageResponse_110) Unwrap() interface{} { return o.inner }

// ---

type ImageFsInfoRequest_110 struct {
	inner *runtimeapi.ImageFsInfoRequest
}

var _ ImageFsInfoRequest = &ImageFsInfoRequest_110{}

func (o *ImageFsInfoRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ImageFsInfoRequest{}
	} else {
		o.inner = v.(*runtimeapi.ImageFsInfoRequest)
	}
}
func (o *ImageFsInfoRequest_110) Unwrap() interface{} { return o.inner }

// ---

type ImageFsInfoResponse_110 struct {
	inner *runtimeapi.ImageFsInfoResponse
}

var _ ImageFsInfoResponse = &ImageFsInfoResponse_110{}

func (o *ImageFsInfoResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ImageFsInfoResponse{}
	} else {
		o.inner = v.(*runtimeapi.ImageFsInfoResponse)
	}
}
func (o *ImageFsInfoResponse_110) Unwrap() interface{} { return o.inner }
func (o *ImageFsInfoResponse_110) Items() []CRIObject {
	var r []CRIObject
	for _, fs := range o.inner.ImageFilesystems {
		r = append(r, &FilesystemUsage_110{fs})
	}
	return r
}
func (o *ImageFsInfoResponse_110) SetItems(items []CRIObject) {
	o.inner.ImageFilesystems = nil
	for _, wrapped := range items {
		o.inner.ImageFilesystems = append(o.inner.ImageFilesystems, wrapped.Unwrap().(*runtimeapi.FilesystemUsage))
	}
}

// --- 1.8+ only ---

type UpdateContainerResourcesRequest_110 struct {
	inner *runtimeapi.UpdateContainerResourcesRequest
}

var _ UpdateContainerResourcesRequest = &UpdateContainerResourcesRequest_110{}

func (o *UpdateContainerResourcesRequest_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.UpdateContainerResourcesRequest{}
	} else {
		o.inner = v.(*runtimeapi.UpdateContainerResourcesRequest)
	}
}
func (o *UpdateContainerResourcesRequest_110) Unwrap() interface{}      { return o.inner }
func (o *UpdateContainerResourcesRequest_110) ContainerId() string      { return o.inner.ContainerId }
func (o *UpdateContainerResourcesRequest_110) SetContainerId(id string) { o.inner.ContainerId = id }

// --- 1.8+ only ---

type UpdateContainerResourcesResponse_110 struct {
	inner *runtimeapi.UpdateContainerResourcesResponse
}

var _ UpdateContainerResourcesResponse = &UpdateContainerResourcesResponse_110{}

func (o *UpdateContainerResourcesResponse_110) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.UpdateContainerResourcesResponse{}
	} else {
		o.inner = v.(*runtimeapi.UpdateContainerResourcesResponse)
	}
}
func (o *UpdateContainerResourcesResponse_110) Unwrap() interface{} { return o.inner }

// ---

var cri110typeMatcher = newTypeMatcher()

func init() {
	cri110typeMatcher.registerTypes(
		&PodSandbox_110{},
		&Container_110{},
		&Image_110{},
		&PodSandboxStatus_110{},
		&ContainerStatus_110{},
		&ContainerStats_110{},
		&FilesystemUsage_110{},
		&VersionRequest_110{},
		&VersionResponse_110{},
		&StatusRequest_110{},
		&StatusResponse_110{},
		&UpdateRuntimeConfigRequest_110{},
		&UpdateRuntimeConfigResponse_110{},
		&RunPodSandboxRequest_110{},
		&RunPodSandboxResponse_110{},
		&ListPodSandboxRequest_110{},
		&ListPodSandboxResponse_110{},
		&StopPodSandboxRequest_110{},
		&StopPodSandboxResponse_110{},
		&RemovePodSandboxRequest_110{},
		&RemovePodSandboxResponse_110{},
		&PodSandboxStatusRequest_110{},
		&PodSandboxStatusResponse_110{},
		&CreateContainerRequest_110{},
		&CreateContainerResponse_110{},
		&ListContainersRequest_110{},
		&ListContainersResponse_110{},
		&ListContainerStatsRequest_110{},
		&ListContainerStatsResponse_110{},
		&StartContainerRequest_110{},
		&StartContainerResponse_110{},
		&StopContainerRequest_110{},
		&StopContainerResponse_110{},
		&RemoveContainerRequest_110{},
		&RemoveContainerResponse_110{},
		&ReopenContainerLogRequest_110{},
		&ReopenContainerLogResponse_110{},
		&ContainerStatusRequest_110{},
		&ContainerStatusResponse_110{},
		&ContainerStatsRequest_110{},
		&ContainerStatsResponse_110{},
		&ExecSyncRequest_110{},
		&ExecSyncResponse_110{},
		&ExecRequest_110{},
		&ExecResponse_110{},
		&AttachRequest_110{},
		&AttachResponse_110{},
		&PortForwardRequest_110{},
		&PortForwardResponse_110{},
		&ListImagesRequest_110{},
		&ListImagesResponse_110{},
		&ImageStatusRequest_110{},
		&ImageStatusResponse_110{},
		&PullImageRequest_110{},
		&PullImageResponse_110{},
		&RemoveImageRequest_110{},
		&RemoveImageResponse_110{},
		&ImageFsInfoRequest_110{},
		&ImageFsInfoResponse_110{},
		&UpdateContainerResourcesRequest_110{},
		&UpdateContainerResourcesResponse_110{},
	)
}

// CRI19 denotes the CRI version 1.10
type CRI110 struct{}

var _ CRIVersion = &CRI110{}

func (c *CRI110) Register(server *grpc.Server) {
	runtimeapi.RegisterDummyRuntimeServiceServer(server)
	runtimeapi.RegisterDummyImageServiceServer(server)
}

func (c *CRI110) ProbeRequest() (interface{}, interface{}) {
	return &runtimeapi.VersionRequest{}, &runtimeapi.VersionResponse{}
}

func (c *CRI110) WrapObject(o interface{}) (CRIObject, CRIObject, error) {
	return wrapUsingMatcher(cri110typeMatcher, o)
}

func (c *CRI110) ProtoPackage() string { return "runtime.v1alpha2" }
