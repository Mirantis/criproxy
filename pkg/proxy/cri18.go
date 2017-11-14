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
	"fmt"

	"google.golang.org/grpc"

	runtimeapi "github.com/Mirantis/criproxy/pkg/runtimeapi/v1_8"
)

// ---

type PodSandbox_18 struct {
	inner *runtimeapi.PodSandbox
}

var _ PodSandbox = &PodSandbox_18{}

func (o *PodSandbox_18) Unwrap() interface{} { return o.inner }
func (o *PodSandbox_18) Copy() PodSandbox    { r := *o.inner; return &PodSandbox_18{&r} }
func (o *PodSandbox_18) Id() string          { return o.inner.Id }
func (o *PodSandbox_18) SetId(id string)     { o.inner.Id = id }

type Container_18 struct {
	inner *runtimeapi.Container
}

// ---

var _ Container = &Container_18{}

func (o *Container_18) Unwrap() interface{}       { return o.inner }
func (o *Container_18) Copy() Container           { r := *o.inner; return &Container_18{&r} }
func (o *Container_18) Id() string                { return o.inner.Id }
func (o *Container_18) SetId(id string)           { o.inner.Id = id }
func (o *Container_18) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *Container_18) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }
func (o *Container_18) Image() string             { return o.inner.Image.GetImage() }
func (o *Container_18) SetImage(image string)     { o.inner.Image = &runtimeapi.ImageSpec{Image: image} }

// ---

type Image_18 struct {
	inner *runtimeapi.Image
}

var _ Image = &Image_18{}

func (o *Image_18) Unwrap() interface{}           { return o.inner }
func (o *Image_18) Copy() Image                   { r := *o.inner; return &Image_18{&r} }
func (o *Image_18) Id() string                    { return o.inner.Id }
func (o *Image_18) SetId(id string)               { o.inner.Id = id }
func (o *Image_18) RepoTags() []string            { return o.inner.RepoTags }
func (o *Image_18) SetRepoTags(repoTags []string) { o.inner.RepoTags = repoTags }

// ---

type PodSandboxStatus_18 struct {
	inner *runtimeapi.PodSandboxStatus
}

var _ PodSandboxStatus = &PodSandboxStatus_18{}

func (o *PodSandboxStatus_18) Unwrap() interface{}    { return o.inner }
func (o *PodSandboxStatus_18) Copy() PodSandboxStatus { r := *o.inner; return &PodSandboxStatus_18{&r} }
func (o *PodSandboxStatus_18) Id() string             { return o.inner.Id }
func (o *PodSandboxStatus_18) SetId(id string)        { o.inner.Id = id }

// ---

type ContainerStatus_18 struct {
	inner *runtimeapi.ContainerStatus
}

var _ ContainerStatus = &ContainerStatus_18{}

func (o *ContainerStatus_18) Unwrap() interface{}   { return o.inner }
func (o *ContainerStatus_18) Copy() ContainerStatus { r := *o.inner; return &ContainerStatus_18{&r} }
func (o *ContainerStatus_18) Id() string            { return o.inner.Id }
func (o *ContainerStatus_18) SetId(id string)       { o.inner.Id = id }
func (o *ContainerStatus_18) Image() string         { return o.inner.Image.GetImage() }
func (o *ContainerStatus_18) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type ContainerStats_18 struct {
	inner *runtimeapi.ContainerStats
}

var _ ContainerStats = &ContainerStats_18{}

func (o *ContainerStats_18) Unwrap() interface{}  { return o.inner }
func (o *ContainerStats_18) Copy() ContainerStats { r := *o.inner; return &ContainerStats_18{&r} }
func (o *ContainerStats_18) Id() string           { return o.inner.Attributes.GetId() }
func (o *ContainerStats_18) SetId(id string) {
	if o.inner.Attributes == nil {
		o.inner.Attributes = &runtimeapi.ContainerAttributes{Id: id}
	} else {
		o.inner.Attributes.Id = id
	}
}

// ---

type FilesystemUsage_18 struct {
	inner *runtimeapi.FilesystemUsage
}

func (o *FilesystemUsage_18) Unwrap() interface{} { return o.inner }

// ---

type VersionRequest_18 struct {
	inner *runtimeapi.VersionRequest
}

var _ VersionRequest = &VersionRequest_18{}

func (o *VersionRequest_18) Unwrap() interface{} { return o.inner }

// ---

type VersionResponse_18 struct {
	inner *runtimeapi.VersionResponse
}

var _ VersionResponse = &VersionResponse_18{}

func (o *VersionResponse_18) Unwrap() interface{} { return o.inner }

// ---

type StatusRequest_18 struct {
	inner *runtimeapi.StatusRequest
}

var _ StatusRequest = &StatusRequest_18{}

func (o *StatusRequest_18) Unwrap() interface{} { return o.inner }

// ---

type StatusResponse_18 struct {
	inner *runtimeapi.StatusResponse
}

var _ StatusResponse = &StatusResponse_18{}

func (o *StatusResponse_18) Unwrap() interface{} { return o.inner }

// ---

type UpdateRuntimeConfigRequest_18 struct {
	inner *runtimeapi.UpdateRuntimeConfigRequest
}

var _ UpdateRuntimeConfigRequest = &UpdateRuntimeConfigRequest_18{}

func (o *UpdateRuntimeConfigRequest_18) Unwrap() interface{} { return o.inner }

// ---

type UpdateRuntimeConfigResponse_18 struct {
	inner *runtimeapi.UpdateRuntimeConfigResponse
}

var _ UpdateRuntimeConfigResponse = &UpdateRuntimeConfigResponse_18{}

func (o *UpdateRuntimeConfigResponse_18) Unwrap() interface{} { return o.inner }

// ---

type RunPodSandboxRequest_18 struct {
	inner *runtimeapi.RunPodSandboxRequest
}

var _ RunPodSandboxRequest = &RunPodSandboxRequest_18{}

func (o *RunPodSandboxRequest_18) Unwrap() interface{} { return o.inner }
func (o *RunPodSandboxRequest_18) GetAnnotations() map[string]string {
	return o.inner.Config.GetAnnotations()
}

// ---

type RunPodSandboxResponse_18 struct {
	inner *runtimeapi.RunPodSandboxResponse
}

var _ RunPodSandboxResponse = &RunPodSandboxResponse_18{}

func (o *RunPodSandboxResponse_18) Unwrap() interface{}       { return o.inner }
func (o *RunPodSandboxResponse_18) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *RunPodSandboxResponse_18) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type ListPodSandboxRequest_18 struct {
	inner *runtimeapi.ListPodSandboxRequest
}

var _ ListPodSandboxRequest = &ListPodSandboxRequest_18{}

func (o *ListPodSandboxRequest_18) Unwrap() interface{} { return o.inner }
func (o *ListPodSandboxRequest_18) IdFilter() string {
	return o.inner.Filter.GetId()
}

func (o *ListPodSandboxRequest_18) SetIdFilter(id string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.PodSandboxFilter{Id: id}
	} else {
		o.inner.Filter.Id = id
	}
}

// ---

type ListPodSandboxResponse_18 struct {
	inner *runtimeapi.ListPodSandboxResponse
}

var _ ListPodSandboxResponse = &ListPodSandboxResponse_18{}

func (o *ListPodSandboxResponse_18) Unwrap() interface{} { return o.inner }
func (o *ListPodSandboxResponse_18) Items() []CRIObject {
	var r []CRIObject
	for _, sandbox := range o.inner.Items {
		r = append(r, &PodSandbox_18{sandbox})
	}
	return r
}
func (o *ListPodSandboxResponse_18) SetItems(items []CRIObject) {
	o.inner.Items = nil
	for _, wrapped := range items {
		o.inner.Items = append(o.inner.Items, wrapped.Unwrap().(*runtimeapi.PodSandbox))
	}
}

// ---

type StopPodSandboxRequest_18 struct {
	inner *runtimeapi.StopPodSandboxRequest
}

var _ StopPodSandboxRequest = &StopPodSandboxRequest_18{}

func (o *StopPodSandboxRequest_18) Unwrap() interface{}       { return o.inner }
func (o *StopPodSandboxRequest_18) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *StopPodSandboxRequest_18) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type StopPodSandboxResponse_18 struct {
	inner *runtimeapi.StopPodSandboxResponse
}

var _ StopPodSandboxResponse = &StopPodSandboxResponse_18{}

func (o *StopPodSandboxResponse_18) Unwrap() interface{} { return o.inner }

// ---

type RemovePodSandboxRequest_18 struct {
	inner *runtimeapi.RemovePodSandboxRequest
}

var _ RemovePodSandboxRequest = &RemovePodSandboxRequest_18{}

func (o *RemovePodSandboxRequest_18) Unwrap() interface{}       { return o.inner }
func (o *RemovePodSandboxRequest_18) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *RemovePodSandboxRequest_18) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type RemovePodSandboxResponse_18 struct {
	inner *runtimeapi.RemovePodSandboxResponse
}

var _ RemovePodSandboxResponse = &RemovePodSandboxResponse_18{}

func (o *RemovePodSandboxResponse_18) Unwrap() interface{} { return o.inner }

// ---

type PodSandboxStatusRequest_18 struct {
	inner *runtimeapi.PodSandboxStatusRequest
}

var _ PodSandboxStatusRequest = &PodSandboxStatusRequest_18{}

func (o *PodSandboxStatusRequest_18) Unwrap() interface{}       { return o.inner }
func (o *PodSandboxStatusRequest_18) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *PodSandboxStatusRequest_18) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type PodSandboxStatusResponse_18 struct {
	inner *runtimeapi.PodSandboxStatusResponse
}

var _ PodSandboxStatusResponse = &PodSandboxStatusResponse_18{}

func (o *PodSandboxStatusResponse_18) Unwrap() interface{} { return o.inner }
func (o *PodSandboxStatusResponse_18) Status() PodSandboxStatus {
	if o.inner.Status == nil {
		return nil
	}
	return &PodSandboxStatus_18{o.inner.Status}
}

// ---

type CreateContainerRequest_18 struct {
	inner *runtimeapi.CreateContainerRequest
}

var _ CreateContainerRequest = &CreateContainerRequest_18{}

func (o *CreateContainerRequest_18) Unwrap() interface{}       { return o.inner }
func (o *CreateContainerRequest_18) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *CreateContainerRequest_18) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }
func (o *CreateContainerRequest_18) Image() string {
	if o.inner.Config == nil {
		return ""
	}
	return o.inner.Config.Image.GetImage()
}

func (o *CreateContainerRequest_18) SetImage(image string) {
	if o.inner.Config != nil {
		o.inner.Config.Image = &runtimeapi.ImageSpec{Image: image}
	}
}

// ---

type CreateContainerResponse_18 struct {
	inner *runtimeapi.CreateContainerResponse
}

var _ CreateContainerResponse = &CreateContainerResponse_18{}

func (o *CreateContainerResponse_18) Unwrap() interface{}      { return o.inner }
func (o *CreateContainerResponse_18) ContainerId() string      { return o.inner.ContainerId }
func (o *CreateContainerResponse_18) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ListContainersRequest_18 struct {
	inner *runtimeapi.ListContainersRequest
}

var _ ListContainersRequest = &ListContainersRequest_18{}

func (o *ListContainersRequest_18) Unwrap() interface{} { return o.inner }
func (o *ListContainersRequest_18) IdFilter() string {
	return o.inner.Filter.GetId()
}

func (o *ListContainersRequest_18) SetIdFilter(id string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerFilter{Id: id}
	} else {
		o.inner.Filter.Id = id
	}
}

func (o *ListContainersRequest_18) PodSandboxIdFilter() string {
	return o.inner.Filter.GetPodSandboxId()
}

func (o *ListContainersRequest_18) SetPodSandboxIdFilter(podSandboxId string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerFilter{Id: podSandboxId}
	} else {
		o.inner.Filter.PodSandboxId = podSandboxId
	}
}

// ---

type ListContainersResponse_18 struct {
	inner *runtimeapi.ListContainersResponse
}

var _ ListContainersResponse = &ListContainersResponse_18{}

func (o *ListContainersResponse_18) Unwrap() interface{} { return o.inner }
func (o *ListContainersResponse_18) Items() []CRIObject {
	var r []CRIObject
	for _, container := range o.inner.Containers {
		r = append(r, &Container_18{container})
	}
	return r
}
func (o *ListContainersResponse_18) SetItems(items []CRIObject) {
	o.inner.Containers = nil
	for _, wrapped := range items {
		o.inner.Containers = append(o.inner.Containers, wrapped.Unwrap().(*runtimeapi.Container))
	}
}

// ---

type ListContainerStatsRequest_18 struct {
	inner *runtimeapi.ListContainerStatsRequest
}

var _ ListContainerStatsRequest = &ListContainerStatsRequest_18{}

func (o *ListContainerStatsRequest_18) Unwrap() interface{} { return o.inner }
func (o *ListContainerStatsRequest_18) IdFilter() string {
	return o.inner.Filter.GetId()
}

func (o *ListContainerStatsRequest_18) SetIdFilter(id string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerStatsFilter{Id: id}
	} else {
		o.inner.Filter.Id = id
	}
}

func (o *ListContainerStatsRequest_18) PodSandboxIdFilter() string {
	return o.inner.Filter.GetPodSandboxId()
}

func (o *ListContainerStatsRequest_18) SetPodSandboxIdFilter(podSandboxId string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerStatsFilter{Id: podSandboxId}
	} else {
		o.inner.Filter.PodSandboxId = podSandboxId
	}
}

// ---

type ListContainerStatsResponse_18 struct {
	inner *runtimeapi.ListContainerStatsResponse
}

var _ ListContainerStatsResponse = &ListContainerStatsResponse_18{}

func (o *ListContainerStatsResponse_18) Unwrap() interface{} { return o.inner }
func (o *ListContainerStatsResponse_18) Items() []CRIObject {
	var r []CRIObject
	for _, stats := range o.inner.Stats {
		r = append(r, &ContainerStats_18{stats})
	}
	return r
}
func (o *ListContainerStatsResponse_18) SetItems(items []CRIObject) {
	o.inner.Stats = nil
	for _, wrapped := range items {
		o.inner.Stats = append(o.inner.Stats, wrapped.Unwrap().(*runtimeapi.ContainerStats))
	}
}

// ---

type StartContainerRequest_18 struct {
	inner *runtimeapi.StartContainerRequest
}

var _ StartContainerRequest = &StartContainerRequest_18{}

func (o *StartContainerRequest_18) Unwrap() interface{}      { return o.inner }
func (o *StartContainerRequest_18) ContainerId() string      { return o.inner.ContainerId }
func (o *StartContainerRequest_18) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type StartContainerResponse_18 struct {
	inner *runtimeapi.StartContainerResponse
}

var _ StartContainerResponse = &StartContainerResponse_18{}

func (o *StartContainerResponse_18) Unwrap() interface{} { return o.inner }

// ---

type StopContainerRequest_18 struct {
	inner *runtimeapi.StopContainerRequest
}

var _ StopContainerRequest = &StopContainerRequest_18{}

func (o *StopContainerRequest_18) Unwrap() interface{}      { return o.inner }
func (o *StopContainerRequest_18) ContainerId() string      { return o.inner.ContainerId }
func (o *StopContainerRequest_18) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type StopContainerResponse_18 struct {
	inner *runtimeapi.StopContainerResponse
}

var _ StopContainerResponse = &StopContainerResponse_18{}

func (o *StopContainerResponse_18) Unwrap() interface{} { return o.inner }

// ---

type RemoveContainerRequest_18 struct {
	inner *runtimeapi.RemoveContainerRequest
}

var _ RemoveContainerRequest = &RemoveContainerRequest_18{}

func (o *RemoveContainerRequest_18) Unwrap() interface{}      { return o.inner }
func (o *RemoveContainerRequest_18) ContainerId() string      { return o.inner.ContainerId }
func (o *RemoveContainerRequest_18) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type RemoveContainerResponse_18 struct {
	inner *runtimeapi.RemoveContainerResponse
}

var _ RemoveContainerResponse = &RemoveContainerResponse_18{}

func (o *RemoveContainerResponse_18) Unwrap() interface{} { return o.inner }

// ---

type ContainerStatusRequest_18 struct {
	inner *runtimeapi.ContainerStatusRequest
}

var _ ContainerStatusRequest = &ContainerStatusRequest_18{}

func (o *ContainerStatusRequest_18) Unwrap() interface{}      { return o.inner }
func (o *ContainerStatusRequest_18) ContainerId() string      { return o.inner.ContainerId }
func (o *ContainerStatusRequest_18) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ContainerStatusResponse_18 struct {
	inner *runtimeapi.ContainerStatusResponse
}

var _ ContainerStatusResponse = &ContainerStatusResponse_18{}

func (o *ContainerStatusResponse_18) Unwrap() interface{} { return o.inner }
func (o *ContainerStatusResponse_18) Status() ContainerStatus {
	if o.inner.Status == nil {
		return nil
	}
	return &ContainerStatus_18{o.inner.Status}
}

// ---

type ContainerStatsRequest_18 struct {
	inner *runtimeapi.ContainerStatsRequest
}

var _ ContainerStatsRequest = &ContainerStatsRequest_18{}

func (o *ContainerStatsRequest_18) Unwrap() interface{}      { return o.inner }
func (o *ContainerStatsRequest_18) ContainerId() string      { return o.inner.ContainerId }
func (o *ContainerStatsRequest_18) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ContainerStatsResponse_18 struct {
	inner *runtimeapi.ContainerStatsResponse
}

var _ ContainerStatsResponse = &ContainerStatsResponse_18{}

func (o *ContainerStatsResponse_18) Unwrap() interface{} { return o.inner }
func (o *ContainerStatsResponse_18) Stats() ContainerStats {
	if o.inner.Stats == nil {
		return nil
	}
	return &ContainerStats_18{o.inner.Stats}
}

// ---

type ExecSyncRequest_18 struct {
	inner *runtimeapi.ExecSyncRequest
}

var _ ExecSyncRequest = &ExecSyncRequest_18{}

func (o *ExecSyncRequest_18) Unwrap() interface{}      { return o.inner }
func (o *ExecSyncRequest_18) ContainerId() string      { return o.inner.ContainerId }
func (o *ExecSyncRequest_18) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ExecSyncResponse_18 struct {
	inner *runtimeapi.ExecSyncResponse
}

var _ ExecSyncResponse = &ExecSyncResponse_18{}

func (o *ExecSyncResponse_18) Unwrap() interface{} { return o.inner }

// ---

type ExecRequest_18 struct {
	inner *runtimeapi.ExecRequest
}

var _ ExecRequest = &ExecRequest_18{}

func (o *ExecRequest_18) Unwrap() interface{}      { return o.inner }
func (o *ExecRequest_18) ContainerId() string      { return o.inner.ContainerId }
func (o *ExecRequest_18) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ExecResponse_18 struct {
	inner *runtimeapi.ExecResponse
}

var _ ExecResponse = &ExecResponse_18{}

func (o *ExecResponse_18) Unwrap() interface{} { return o.inner }
func (o *ExecResponse_18) Url() string         { return o.inner.Url }
func (o *ExecResponse_18) SetUrl(url string)   { o.inner.Url = url }

// ---

type AttachRequest_18 struct {
	inner *runtimeapi.AttachRequest
}

var _ AttachRequest = &AttachRequest_18{}

func (o *AttachRequest_18) Unwrap() interface{}      { return o.inner }
func (o *AttachRequest_18) ContainerId() string      { return o.inner.ContainerId }
func (o *AttachRequest_18) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type AttachResponse_18 struct {
	inner *runtimeapi.AttachResponse
}

var _ AttachResponse = &AttachResponse_18{}

func (o *AttachResponse_18) Unwrap() interface{} { return o.inner }
func (o *AttachResponse_18) Url() string         { return o.inner.Url }
func (o *AttachResponse_18) SetUrl(url string)   { o.inner.Url = url }

// ---

type PortForwardRequest_18 struct {
	inner *runtimeapi.PortForwardRequest
}

var _ PortForwardRequest = &PortForwardRequest_18{}

func (o *PortForwardRequest_18) Unwrap() interface{}  { return o.inner }
func (o *PortForwardRequest_18) PodSandboxId() string { return o.inner.PodSandboxId }
func (o *PortForwardRequest_18) SetPodSandboxId(podSandboxId string) {
	o.inner.PodSandboxId = podSandboxId
}

// ---

type PortForwardResponse_18 struct {
	inner *runtimeapi.PortForwardResponse
}

var _ PortForwardResponse = &PortForwardResponse_18{}

func (o *PortForwardResponse_18) Unwrap() interface{} { return o.inner }
func (o *PortForwardResponse_18) Url() string         { return o.inner.Url }
func (o *PortForwardResponse_18) SetUrl(url string)   { o.inner.Url = url }

// ---

type ListImagesRequest_18 struct {
	inner *runtimeapi.ListImagesRequest
}

var _ ListImagesRequest = &ListImagesRequest_18{}

func (o *ListImagesRequest_18) Unwrap() interface{} { return o.inner }
func (o *ListImagesRequest_18) ImageFilter() string { return o.inner.Filter.GetImage().GetImage() }
func (o *ListImagesRequest_18) SetImageFilter(image string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ImageFilter{
			Image: &runtimeapi.ImageSpec{Image: image},
		}
	} else {
		o.inner.Filter.Image = &runtimeapi.ImageSpec{Image: image}
	}
}

// ---

type ListImagesResponse_18 struct {
	inner *runtimeapi.ListImagesResponse
}

var _ ListImagesResponse = &ListImagesResponse_18{}

func (o *ListImagesResponse_18) Unwrap() interface{} { return o.inner }
func (o *ListImagesResponse_18) Items() []CRIObject {
	var r []CRIObject
	for _, image := range o.inner.Images {
		r = append(r, &Image_18{image})
	}
	return r
}
func (o *ListImagesResponse_18) SetItems(items []CRIObject) {
	o.inner.Images = nil
	for _, wrapped := range items {
		o.inner.Images = append(o.inner.Images, wrapped.Unwrap().(*runtimeapi.Image))
	}
}

// ---

type ImageStatusRequest_18 struct {
	inner *runtimeapi.ImageStatusRequest
}

var _ ImageStatusRequest = &ImageStatusRequest_18{}

func (o *ImageStatusRequest_18) Unwrap() interface{} { return o.inner }
func (o *ImageStatusRequest_18) Image() string       { return o.inner.Image.GetImage() }
func (o *ImageStatusRequest_18) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type ImageStatusResponse_18 struct {
	inner *runtimeapi.ImageStatusResponse
}

var _ ImageStatusResponse = &ImageStatusResponse_18{}

func (o *ImageStatusResponse_18) Unwrap() interface{} { return o.inner }
func (o *ImageStatusResponse_18) Image() Image        { return &Image_18{o.inner.Image} }
func (o *ImageStatusResponse_18) SetImage(image Image) {
	o.inner.Image = image.Unwrap().(*runtimeapi.Image)
}

// ---

type PullImageRequest_18 struct {
	inner *runtimeapi.PullImageRequest
}

var _ PullImageRequest = &PullImageRequest_18{}

func (o *PullImageRequest_18) Unwrap() interface{} { return o.inner }
func (o *PullImageRequest_18) Image() string       { return o.inner.Image.GetImage() }
func (o *PullImageRequest_18) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type PullImageResponse_18 struct {
	inner *runtimeapi.PullImageResponse
}

var _ PullImageResponse = &PullImageResponse_18{}

func (o *PullImageResponse_18) Unwrap() interface{}   { return o.inner }
func (o *PullImageResponse_18) Image() string         { return o.inner.ImageRef }
func (o *PullImageResponse_18) SetImage(image string) { o.inner.ImageRef = image }

// ---

type RemoveImageRequest_18 struct {
	inner *runtimeapi.RemoveImageRequest
}

var _ RemoveImageRequest = &RemoveImageRequest_18{}

func (o *RemoveImageRequest_18) Unwrap() interface{} { return o.inner }
func (o *RemoveImageRequest_18) Image() string       { return o.inner.Image.GetImage() }
func (o *RemoveImageRequest_18) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type RemoveImageResponse_18 struct {
	inner *runtimeapi.RemoveImageResponse
}

var _ RemoveImageResponse = &RemoveImageResponse_18{}

func (o *RemoveImageResponse_18) Unwrap() interface{} { return o.inner }

// ---

type ImageFsInfoRequest_18 struct {
	inner *runtimeapi.ImageFsInfoRequest
}

var _ ImageFsInfoRequest = &ImageFsInfoRequest_18{}

func (o *ImageFsInfoRequest_18) Unwrap() interface{} { return o.inner }

// ---

type ImageFsInfoResponse_18 struct {
	inner *runtimeapi.ImageFsInfoResponse
}

var _ ImageFsInfoResponse = &ImageFsInfoResponse_18{}

func (o *ImageFsInfoResponse_18) Unwrap() interface{} { return o.inner }
func (o *ImageFsInfoResponse_18) Items() []CRIObject {
	var r []CRIObject
	for _, fs := range o.inner.ImageFilesystems {
		r = append(r, &FilesystemUsage_18{fs})
	}
	return r
}
func (o *ImageFsInfoResponse_18) SetItems(items []CRIObject) {
	o.inner.ImageFilesystems = nil
	for _, wrapped := range items {
		o.inner.ImageFilesystems = append(o.inner.ImageFilesystems, wrapped.Unwrap().(*runtimeapi.FilesystemUsage))
	}
}

// ---

type CRI18 struct{}

var _ CRIVersion = &CRI18{}

func (c *CRI18) Register(server *grpc.Server) {
	runtimeapi.RegisterDummyRuntimeServiceServer(server)
	runtimeapi.RegisterDummyImageServiceServer(server)
}

func (c *CRI18) ProbeRequest() (interface{}, interface{}) {
	return &runtimeapi.VersionRequest{}, &runtimeapi.VersionResponse{}
}

func (c *CRI18) WrapObject(o interface{}) (CRIObject, CRIObject, error) {
	switch v := o.(type) {
	case *runtimeapi.PodSandbox:
		return &PodSandbox_18{v}, nil, nil
	case *runtimeapi.Container:
		return &Container_18{v}, nil, nil
	case *runtimeapi.Image:
		return &Image_18{v}, nil, nil
	case *runtimeapi.PodSandboxStatus:
		return &PodSandboxStatus_18{v}, nil, nil
	case *runtimeapi.VersionRequest:
		return &VersionRequest_18{v}, &VersionResponse_18{&runtimeapi.VersionResponse{}}, nil
	case *runtimeapi.StatusRequest:
		return &StatusRequest_18{v}, &StatusResponse_18{&runtimeapi.StatusResponse{}}, nil
	case *runtimeapi.UpdateRuntimeConfigRequest:
		return &UpdateRuntimeConfigRequest_18{v}, &UpdateRuntimeConfigResponse_18{&runtimeapi.UpdateRuntimeConfigResponse{}}, nil
	case *runtimeapi.RunPodSandboxRequest:
		return &RunPodSandboxRequest_18{v}, &RunPodSandboxResponse_18{&runtimeapi.RunPodSandboxResponse{}}, nil
	case *runtimeapi.ListPodSandboxRequest:
		return &ListPodSandboxRequest_18{v}, &ListPodSandboxResponse_18{&runtimeapi.ListPodSandboxResponse{}}, nil
	case *runtimeapi.StopPodSandboxRequest:
		return &StopPodSandboxRequest_18{v}, &StopPodSandboxResponse_18{&runtimeapi.StopPodSandboxResponse{}}, nil
	case *runtimeapi.RemovePodSandboxRequest:
		return &RemovePodSandboxRequest_18{v}, &RemovePodSandboxResponse_18{&runtimeapi.RemovePodSandboxResponse{}}, nil
	case *runtimeapi.PodSandboxStatusRequest:
		return &PodSandboxStatusRequest_18{v}, &PodSandboxStatusResponse_18{&runtimeapi.PodSandboxStatusResponse{}}, nil
	case *runtimeapi.CreateContainerRequest:
		return &CreateContainerRequest_18{v}, &CreateContainerResponse_18{&runtimeapi.CreateContainerResponse{}}, nil
	case *runtimeapi.ListContainersRequest:
		return &ListContainersRequest_18{v}, &ListContainersResponse_18{&runtimeapi.ListContainersResponse{}}, nil
	case *runtimeapi.ListContainerStatsRequest:
		return &ListContainerStatsRequest_18{v}, &ListContainerStatsResponse_18{&runtimeapi.ListContainerStatsResponse{}}, nil
	case *runtimeapi.StartContainerRequest:
		return &StartContainerRequest_18{v}, &StartContainerResponse_18{&runtimeapi.StartContainerResponse{}}, nil
	case *runtimeapi.StopContainerRequest:
		return &StopContainerRequest_18{v}, &StopContainerResponse_18{&runtimeapi.StopContainerResponse{}}, nil
	case *runtimeapi.RemoveContainerRequest:
		return &RemoveContainerRequest_18{v}, &RemoveContainerResponse_18{&runtimeapi.RemoveContainerResponse{}}, nil
	case *runtimeapi.ContainerStatusRequest:
		return &ContainerStatusRequest_18{v}, &ContainerStatusResponse_18{&runtimeapi.ContainerStatusResponse{}}, nil
	case *runtimeapi.ContainerStatsRequest:
		return &ContainerStatsRequest_18{v}, &ContainerStatsResponse_18{&runtimeapi.ContainerStatsResponse{}}, nil
	case *runtimeapi.ExecSyncRequest:
		return &ExecSyncRequest_18{v}, &ExecSyncResponse_18{&runtimeapi.ExecSyncResponse{}}, nil
	case *runtimeapi.ExecRequest:
		return &ExecRequest_18{v}, &ExecResponse_18{&runtimeapi.ExecResponse{}}, nil
	case *runtimeapi.AttachRequest:
		return &AttachRequest_18{v}, &AttachResponse_18{&runtimeapi.AttachResponse{}}, nil
	case *runtimeapi.PortForwardRequest:
		return &PortForwardRequest_18{v}, &PortForwardResponse_18{&runtimeapi.PortForwardResponse{}}, nil
	case *runtimeapi.ListImagesRequest:
		return &ListImagesRequest_18{v}, &ListImagesResponse_18{&runtimeapi.ListImagesResponse{}}, nil
	case *runtimeapi.ImageStatusRequest:
		return &ImageStatusRequest_18{v}, &ImageStatusResponse_18{&runtimeapi.ImageStatusResponse{}}, nil
	case *runtimeapi.PullImageRequest:
		return &PullImageRequest_18{v}, &PullImageResponse_18{&runtimeapi.PullImageResponse{}}, nil
	case *runtimeapi.RemoveImageRequest:
		return &RemoveImageRequest_18{v}, &RemoveImageResponse_18{&runtimeapi.RemoveImageResponse{}}, nil
	case *runtimeapi.ImageFsInfoRequest:
		return &ImageFsInfoRequest_18{v}, &ImageFsInfoResponse_18{&runtimeapi.ImageFsInfoResponse{}}, nil
	default:
		return nil, nil, fmt.Errorf("can't wrap %T", o)
	}
}
