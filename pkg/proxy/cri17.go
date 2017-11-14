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
	"fmt"

	"google.golang.org/grpc"

	runtimeapi "github.com/Mirantis/criproxy/pkg/runtimeapi/v1_7"
)

// ---

type PodSandbox_17 struct {
	inner *runtimeapi.PodSandbox
}

var _ PodSandbox = &PodSandbox_17{}

func (o *PodSandbox_17) Unwrap() interface{} { return o.inner }
func (o *PodSandbox_17) Copy() PodSandbox    { r := *o.inner; return &PodSandbox_17{&r} }
func (o *PodSandbox_17) Id() string          { return o.inner.Id }
func (o *PodSandbox_17) SetId(id string)     { o.inner.Id = id }

type Container_17 struct {
	inner *runtimeapi.Container
}

// ---

var _ Container = &Container_17{}

func (o *Container_17) Unwrap() interface{}       { return o.inner }
func (o *Container_17) Copy() Container           { r := *o.inner; return &Container_17{&r} }
func (o *Container_17) Id() string                { return o.inner.Id }
func (o *Container_17) SetId(id string)           { o.inner.Id = id }
func (o *Container_17) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *Container_17) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }
func (o *Container_17) Image() string             { return o.inner.Image.GetImage() }
func (o *Container_17) SetImage(image string)     { o.inner.Image = &runtimeapi.ImageSpec{Image: image} }

// ---

type Image_17 struct {
	inner *runtimeapi.Image
}

var _ Image = &Image_17{}

func (o *Image_17) Unwrap() interface{}           { return o.inner }
func (o *Image_17) Copy() Image                   { r := *o.inner; return &Image_17{&r} }
func (o *Image_17) Id() string                    { return o.inner.Id }
func (o *Image_17) SetId(id string)               { o.inner.Id = id }
func (o *Image_17) RepoTags() []string            { return o.inner.RepoTags }
func (o *Image_17) SetRepoTags(repoTags []string) { o.inner.RepoTags = repoTags }

// ---

type PodSandboxStatus_17 struct {
	inner *runtimeapi.PodSandboxStatus
}

var _ PodSandboxStatus = &PodSandboxStatus_17{}

func (o *PodSandboxStatus_17) Unwrap() interface{}    { return o.inner }
func (o *PodSandboxStatus_17) Copy() PodSandboxStatus { r := *o.inner; return &PodSandboxStatus_17{&r} }
func (o *PodSandboxStatus_17) Id() string             { return o.inner.Id }
func (o *PodSandboxStatus_17) SetId(id string)        { o.inner.Id = id }

// ---

type ContainerStatus_17 struct {
	inner *runtimeapi.ContainerStatus
}

var _ ContainerStatus = &ContainerStatus_17{}

func (o *ContainerStatus_17) Unwrap() interface{}   { return o.inner }
func (o *ContainerStatus_17) Copy() ContainerStatus { r := *o.inner; return &ContainerStatus_17{&r} }
func (o *ContainerStatus_17) Id() string            { return o.inner.Id }
func (o *ContainerStatus_17) SetId(id string)       { o.inner.Id = id }
func (o *ContainerStatus_17) Image() string         { return o.inner.Image.GetImage() }
func (o *ContainerStatus_17) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type ContainerStats_17 struct {
	inner *runtimeapi.ContainerStats
}

var _ ContainerStats = &ContainerStats_17{}

func (o *ContainerStats_17) Unwrap() interface{}  { return o.inner }
func (o *ContainerStats_17) Copy() ContainerStats { r := *o.inner; return &ContainerStats_17{&r} }
func (o *ContainerStats_17) Id() string           { return o.inner.Attributes.GetId() }
func (o *ContainerStats_17) SetId(id string) {
	if o.inner.Attributes == nil {
		o.inner.Attributes = &runtimeapi.ContainerAttributes{Id: id}
	} else {
		o.inner.Attributes.Id = id
	}
}

// ---

type FilesystemUsage_17 struct {
	inner *runtimeapi.FilesystemUsage
}

func (o *FilesystemUsage_17) Unwrap() interface{} { return o.inner }

// ---

type VersionRequest_17 struct {
	inner *runtimeapi.VersionRequest
}

var _ VersionRequest = &VersionRequest_17{}

func (o *VersionRequest_17) Unwrap() interface{} { return o.inner }

// ---

type VersionResponse_17 struct {
	inner *runtimeapi.VersionResponse
}

var _ VersionResponse = &VersionResponse_17{}

func (o *VersionResponse_17) Unwrap() interface{} { return o.inner }

// ---

type StatusRequest_17 struct {
	inner *runtimeapi.StatusRequest
}

var _ StatusRequest = &StatusRequest_17{}

func (o *StatusRequest_17) Unwrap() interface{} { return o.inner }

// ---

type StatusResponse_17 struct {
	inner *runtimeapi.StatusResponse
}

var _ StatusResponse = &StatusResponse_17{}

func (o *StatusResponse_17) Unwrap() interface{} { return o.inner }

// ---

type UpdateRuntimeConfigRequest_17 struct {
	inner *runtimeapi.UpdateRuntimeConfigRequest
}

var _ UpdateRuntimeConfigRequest = &UpdateRuntimeConfigRequest_17{}

func (o *UpdateRuntimeConfigRequest_17) Unwrap() interface{} { return o.inner }

// ---

type UpdateRuntimeConfigResponse_17 struct {
	inner *runtimeapi.UpdateRuntimeConfigResponse
}

var _ UpdateRuntimeConfigResponse = &UpdateRuntimeConfigResponse_17{}

func (o *UpdateRuntimeConfigResponse_17) Unwrap() interface{} { return o.inner }

// ---

type RunPodSandboxRequest_17 struct {
	inner *runtimeapi.RunPodSandboxRequest
}

var _ RunPodSandboxRequest = &RunPodSandboxRequest_17{}

func (o *RunPodSandboxRequest_17) Unwrap() interface{} { return o.inner }
func (o *RunPodSandboxRequest_17) GetAnnotations() map[string]string {
	return o.inner.Config.GetAnnotations()
}

// ---

type RunPodSandboxResponse_17 struct {
	inner *runtimeapi.RunPodSandboxResponse
}

var _ RunPodSandboxResponse = &RunPodSandboxResponse_17{}

func (o *RunPodSandboxResponse_17) Unwrap() interface{}       { return o.inner }
func (o *RunPodSandboxResponse_17) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *RunPodSandboxResponse_17) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type ListPodSandboxRequest_17 struct {
	inner *runtimeapi.ListPodSandboxRequest
}

var _ ListPodSandboxRequest = &ListPodSandboxRequest_17{}

func (o *ListPodSandboxRequest_17) Unwrap() interface{} { return o.inner }
func (o *ListPodSandboxRequest_17) IdFilter() string {
	return o.inner.Filter.GetId()
}

func (o *ListPodSandboxRequest_17) SetIdFilter(id string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.PodSandboxFilter{Id: id}
	} else {
		o.inner.Filter.Id = id
	}
}

// ---

type ListPodSandboxResponse_17 struct {
	inner *runtimeapi.ListPodSandboxResponse
}

var _ ListPodSandboxResponse = &ListPodSandboxResponse_17{}

func (o *ListPodSandboxResponse_17) Unwrap() interface{} { return o.inner }
func (o *ListPodSandboxResponse_17) Items() []CRIObject {
	var r []CRIObject
	for _, sandbox := range o.inner.Items {
		r = append(r, &PodSandbox_17{sandbox})
	}
	return r
}
func (o *ListPodSandboxResponse_17) SetItems(items []CRIObject) {
	o.inner.Items = nil
	for _, wrapped := range items {
		o.inner.Items = append(o.inner.Items, wrapped.Unwrap().(*runtimeapi.PodSandbox))
	}
}

// ---

type StopPodSandboxRequest_17 struct {
	inner *runtimeapi.StopPodSandboxRequest
}

var _ StopPodSandboxRequest = &StopPodSandboxRequest_17{}

func (o *StopPodSandboxRequest_17) Unwrap() interface{}       { return o.inner }
func (o *StopPodSandboxRequest_17) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *StopPodSandboxRequest_17) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type StopPodSandboxResponse_17 struct {
	inner *runtimeapi.StopPodSandboxResponse
}

var _ StopPodSandboxResponse = &StopPodSandboxResponse_17{}

func (o *StopPodSandboxResponse_17) Unwrap() interface{} { return o.inner }

// ---

type RemovePodSandboxRequest_17 struct {
	inner *runtimeapi.RemovePodSandboxRequest
}

var _ RemovePodSandboxRequest = &RemovePodSandboxRequest_17{}

func (o *RemovePodSandboxRequest_17) Unwrap() interface{}       { return o.inner }
func (o *RemovePodSandboxRequest_17) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *RemovePodSandboxRequest_17) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type RemovePodSandboxResponse_17 struct {
	inner *runtimeapi.RemovePodSandboxResponse
}

var _ RemovePodSandboxResponse = &RemovePodSandboxResponse_17{}

func (o *RemovePodSandboxResponse_17) Unwrap() interface{} { return o.inner }

// ---

type PodSandboxStatusRequest_17 struct {
	inner *runtimeapi.PodSandboxStatusRequest
}

var _ PodSandboxStatusRequest = &PodSandboxStatusRequest_17{}

func (o *PodSandboxStatusRequest_17) Unwrap() interface{}       { return o.inner }
func (o *PodSandboxStatusRequest_17) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *PodSandboxStatusRequest_17) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type PodSandboxStatusResponse_17 struct {
	inner *runtimeapi.PodSandboxStatusResponse
}

var _ PodSandboxStatusResponse = &PodSandboxStatusResponse_17{}

func (o *PodSandboxStatusResponse_17) Unwrap() interface{} { return o.inner }
func (o *PodSandboxStatusResponse_17) Status() PodSandboxStatus {
	if o.inner.Status == nil {
		return nil
	}
	return &PodSandboxStatus_17{o.inner.Status}
}

// ---

type CreateContainerRequest_17 struct {
	inner *runtimeapi.CreateContainerRequest
}

var _ CreateContainerRequest = &CreateContainerRequest_17{}

func (o *CreateContainerRequest_17) Unwrap() interface{}       { return o.inner }
func (o *CreateContainerRequest_17) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *CreateContainerRequest_17) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }
func (o *CreateContainerRequest_17) Image() string {
	if o.inner.Config == nil {
		return ""
	}
	return o.inner.Config.Image.GetImage()
}

func (o *CreateContainerRequest_17) SetImage(image string) {
	if o.inner.Config != nil {
		o.inner.Config.Image = &runtimeapi.ImageSpec{Image: image}
	}
}

// ---

type CreateContainerResponse_17 struct {
	inner *runtimeapi.CreateContainerResponse
}

var _ CreateContainerResponse = &CreateContainerResponse_17{}

func (o *CreateContainerResponse_17) Unwrap() interface{}      { return o.inner }
func (o *CreateContainerResponse_17) ContainerId() string      { return o.inner.ContainerId }
func (o *CreateContainerResponse_17) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ListContainersRequest_17 struct {
	inner *runtimeapi.ListContainersRequest
}

var _ ListContainersRequest = &ListContainersRequest_17{}

func (o *ListContainersRequest_17) Unwrap() interface{} { return o.inner }
func (o *ListContainersRequest_17) IdFilter() string {
	return o.inner.Filter.GetId()
}

func (o *ListContainersRequest_17) SetIdFilter(id string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerFilter{Id: id}
	} else {
		o.inner.Filter.Id = id
	}
}

func (o *ListContainersRequest_17) PodSandboxIdFilter() string {
	return o.inner.Filter.GetPodSandboxId()
}

func (o *ListContainersRequest_17) SetPodSandboxIdFilter(podSandboxId string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerFilter{Id: podSandboxId}
	} else {
		o.inner.Filter.PodSandboxId = podSandboxId
	}
}

// ---

type ListContainersResponse_17 struct {
	inner *runtimeapi.ListContainersResponse
}

var _ ListContainersResponse = &ListContainersResponse_17{}

func (o *ListContainersResponse_17) Unwrap() interface{} { return o.inner }
func (o *ListContainersResponse_17) Items() []CRIObject {
	var r []CRIObject
	for _, container := range o.inner.Containers {
		r = append(r, &Container_17{container})
	}
	return r
}
func (o *ListContainersResponse_17) SetItems(items []CRIObject) {
	o.inner.Containers = nil
	for _, wrapped := range items {
		o.inner.Containers = append(o.inner.Containers, wrapped.Unwrap().(*runtimeapi.Container))
	}
}

// ---

type ListContainerStatsRequest_17 struct {
	inner *runtimeapi.ListContainerStatsRequest
}

var _ ListContainerStatsRequest = &ListContainerStatsRequest_17{}

func (o *ListContainerStatsRequest_17) Unwrap() interface{} { return o.inner }
func (o *ListContainerStatsRequest_17) IdFilter() string {
	return o.inner.Filter.GetId()
}

func (o *ListContainerStatsRequest_17) SetIdFilter(id string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerStatsFilter{Id: id}
	} else {
		o.inner.Filter.Id = id
	}
}

func (o *ListContainerStatsRequest_17) PodSandboxIdFilter() string {
	return o.inner.Filter.GetPodSandboxId()
}

func (o *ListContainerStatsRequest_17) SetPodSandboxIdFilter(podSandboxId string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerStatsFilter{Id: podSandboxId}
	} else {
		o.inner.Filter.PodSandboxId = podSandboxId
	}
}

// ---

type ListContainerStatsResponse_17 struct {
	inner *runtimeapi.ListContainerStatsResponse
}

var _ ListContainerStatsResponse = &ListContainerStatsResponse_17{}

func (o *ListContainerStatsResponse_17) Unwrap() interface{} { return o.inner }
func (o *ListContainerStatsResponse_17) Items() []CRIObject {
	var r []CRIObject
	for _, stats := range o.inner.Stats {
		r = append(r, &ContainerStats_17{stats})
	}
	return r
}
func (o *ListContainerStatsResponse_17) SetItems(items []CRIObject) {
	o.inner.Stats = nil
	for _, wrapped := range items {
		o.inner.Stats = append(o.inner.Stats, wrapped.Unwrap().(*runtimeapi.ContainerStats))
	}
}

// ---

type StartContainerRequest_17 struct {
	inner *runtimeapi.StartContainerRequest
}

var _ StartContainerRequest = &StartContainerRequest_17{}

func (o *StartContainerRequest_17) Unwrap() interface{}      { return o.inner }
func (o *StartContainerRequest_17) ContainerId() string      { return o.inner.ContainerId }
func (o *StartContainerRequest_17) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type StartContainerResponse_17 struct {
	inner *runtimeapi.StartContainerResponse
}

var _ StartContainerResponse = &StartContainerResponse_17{}

func (o *StartContainerResponse_17) Unwrap() interface{} { return o.inner }

// ---

type StopContainerRequest_17 struct {
	inner *runtimeapi.StopContainerRequest
}

var _ StopContainerRequest = &StopContainerRequest_17{}

func (o *StopContainerRequest_17) Unwrap() interface{}      { return o.inner }
func (o *StopContainerRequest_17) ContainerId() string      { return o.inner.ContainerId }
func (o *StopContainerRequest_17) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type StopContainerResponse_17 struct {
	inner *runtimeapi.StopContainerResponse
}

var _ StopContainerResponse = &StopContainerResponse_17{}

func (o *StopContainerResponse_17) Unwrap() interface{} { return o.inner }

// ---

type RemoveContainerRequest_17 struct {
	inner *runtimeapi.RemoveContainerRequest
}

var _ RemoveContainerRequest = &RemoveContainerRequest_17{}

func (o *RemoveContainerRequest_17) Unwrap() interface{}      { return o.inner }
func (o *RemoveContainerRequest_17) ContainerId() string      { return o.inner.ContainerId }
func (o *RemoveContainerRequest_17) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type RemoveContainerResponse_17 struct {
	inner *runtimeapi.RemoveContainerResponse
}

var _ RemoveContainerResponse = &RemoveContainerResponse_17{}

func (o *RemoveContainerResponse_17) Unwrap() interface{} { return o.inner }

// ---

type ContainerStatusRequest_17 struct {
	inner *runtimeapi.ContainerStatusRequest
}

var _ ContainerStatusRequest = &ContainerStatusRequest_17{}

func (o *ContainerStatusRequest_17) Unwrap() interface{}      { return o.inner }
func (o *ContainerStatusRequest_17) ContainerId() string      { return o.inner.ContainerId }
func (o *ContainerStatusRequest_17) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ContainerStatusResponse_17 struct {
	inner *runtimeapi.ContainerStatusResponse
}

var _ ContainerStatusResponse = &ContainerStatusResponse_17{}

func (o *ContainerStatusResponse_17) Unwrap() interface{} { return o.inner }
func (o *ContainerStatusResponse_17) Status() ContainerStatus {
	if o.inner.Status == nil {
		return nil
	}
	return &ContainerStatus_17{o.inner.Status}
}

// ---

type ContainerStatsRequest_17 struct {
	inner *runtimeapi.ContainerStatsRequest
}

var _ ContainerStatsRequest = &ContainerStatsRequest_17{}

func (o *ContainerStatsRequest_17) Unwrap() interface{}      { return o.inner }
func (o *ContainerStatsRequest_17) ContainerId() string      { return o.inner.ContainerId }
func (o *ContainerStatsRequest_17) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ContainerStatsResponse_17 struct {
	inner *runtimeapi.ContainerStatsResponse
}

var _ ContainerStatsResponse = &ContainerStatsResponse_17{}

func (o *ContainerStatsResponse_17) Unwrap() interface{} { return o.inner }
func (o *ContainerStatsResponse_17) Stats() ContainerStats {
	if o.inner.Stats == nil {
		return nil
	}
	return &ContainerStats_17{o.inner.Stats}
}

// ---

type ExecSyncRequest_17 struct {
	inner *runtimeapi.ExecSyncRequest
}

var _ ExecSyncRequest = &ExecSyncRequest_17{}

func (o *ExecSyncRequest_17) Unwrap() interface{}      { return o.inner }
func (o *ExecSyncRequest_17) ContainerId() string      { return o.inner.ContainerId }
func (o *ExecSyncRequest_17) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ExecSyncResponse_17 struct {
	inner *runtimeapi.ExecSyncResponse
}

var _ ExecSyncResponse = &ExecSyncResponse_17{}

func (o *ExecSyncResponse_17) Unwrap() interface{} { return o.inner }

// ---

type ExecRequest_17 struct {
	inner *runtimeapi.ExecRequest
}

var _ ExecRequest = &ExecRequest_17{}

func (o *ExecRequest_17) Unwrap() interface{}      { return o.inner }
func (o *ExecRequest_17) ContainerId() string      { return o.inner.ContainerId }
func (o *ExecRequest_17) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ExecResponse_17 struct {
	inner *runtimeapi.ExecResponse
}

var _ ExecResponse = &ExecResponse_17{}

func (o *ExecResponse_17) Unwrap() interface{} { return o.inner }
func (o *ExecResponse_17) Url() string         { return o.inner.Url }
func (o *ExecResponse_17) SetUrl(url string)   { o.inner.Url = url }

// ---

type AttachRequest_17 struct {
	inner *runtimeapi.AttachRequest
}

var _ AttachRequest = &AttachRequest_17{}

func (o *AttachRequest_17) Unwrap() interface{}      { return o.inner }
func (o *AttachRequest_17) ContainerId() string      { return o.inner.ContainerId }
func (o *AttachRequest_17) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type AttachResponse_17 struct {
	inner *runtimeapi.AttachResponse
}

var _ AttachResponse = &AttachResponse_17{}

func (o *AttachResponse_17) Unwrap() interface{} { return o.inner }
func (o *AttachResponse_17) Url() string         { return o.inner.Url }
func (o *AttachResponse_17) SetUrl(url string)   { o.inner.Url = url }

// ---

type PortForwardRequest_17 struct {
	inner *runtimeapi.PortForwardRequest
}

var _ PortForwardRequest = &PortForwardRequest_17{}

func (o *PortForwardRequest_17) Unwrap() interface{}  { return o.inner }
func (o *PortForwardRequest_17) PodSandboxId() string { return o.inner.PodSandboxId }
func (o *PortForwardRequest_17) SetPodSandboxId(podSandboxId string) {
	o.inner.PodSandboxId = podSandboxId
}

// ---

type PortForwardResponse_17 struct {
	inner *runtimeapi.PortForwardResponse
}

var _ PortForwardResponse = &PortForwardResponse_17{}

func (o *PortForwardResponse_17) Unwrap() interface{} { return o.inner }
func (o *PortForwardResponse_17) Url() string         { return o.inner.Url }
func (o *PortForwardResponse_17) SetUrl(url string)   { o.inner.Url = url }

// ---

type ListImagesRequest_17 struct {
	inner *runtimeapi.ListImagesRequest
}

var _ ListImagesRequest = &ListImagesRequest_17{}

func (o *ListImagesRequest_17) Unwrap() interface{} { return o.inner }
func (o *ListImagesRequest_17) ImageFilter() string { return o.inner.Filter.GetImage().GetImage() }
func (o *ListImagesRequest_17) SetImageFilter(image string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ImageFilter{
			Image: &runtimeapi.ImageSpec{Image: image},
		}
	} else {
		o.inner.Filter.Image = &runtimeapi.ImageSpec{Image: image}
	}
}

// ---

type ListImagesResponse_17 struct {
	inner *runtimeapi.ListImagesResponse
}

var _ ListImagesResponse = &ListImagesResponse_17{}

func (o *ListImagesResponse_17) Unwrap() interface{} { return o.inner }
func (o *ListImagesResponse_17) Items() []CRIObject {
	var r []CRIObject
	for _, image := range o.inner.Images {
		r = append(r, &Image_17{image})
	}
	return r
}
func (o *ListImagesResponse_17) SetItems(items []CRIObject) {
	o.inner.Images = nil
	for _, wrapped := range items {
		o.inner.Images = append(o.inner.Images, wrapped.Unwrap().(*runtimeapi.Image))
	}
}

// ---

type ImageStatusRequest_17 struct {
	inner *runtimeapi.ImageStatusRequest
}

var _ ImageStatusRequest = &ImageStatusRequest_17{}

func (o *ImageStatusRequest_17) Unwrap() interface{} { return o.inner }
func (o *ImageStatusRequest_17) Image() string       { return o.inner.Image.GetImage() }
func (o *ImageStatusRequest_17) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type ImageStatusResponse_17 struct {
	inner *runtimeapi.ImageStatusResponse
}

var _ ImageStatusResponse = &ImageStatusResponse_17{}

func (o *ImageStatusResponse_17) Unwrap() interface{} { return o.inner }
func (o *ImageStatusResponse_17) Image() Image        { return &Image_17{o.inner.Image} }
func (o *ImageStatusResponse_17) SetImage(image Image) {
	o.inner.Image = image.Unwrap().(*runtimeapi.Image)
}

// ---

type PullImageRequest_17 struct {
	inner *runtimeapi.PullImageRequest
}

var _ PullImageRequest = &PullImageRequest_17{}

func (o *PullImageRequest_17) Unwrap() interface{} { return o.inner }
func (o *PullImageRequest_17) Image() string       { return o.inner.Image.GetImage() }
func (o *PullImageRequest_17) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type PullImageResponse_17 struct {
	inner *runtimeapi.PullImageResponse
}

var _ PullImageResponse = &PullImageResponse_17{}

func (o *PullImageResponse_17) Unwrap() interface{}   { return o.inner }
func (o *PullImageResponse_17) Image() string         { return o.inner.ImageRef }
func (o *PullImageResponse_17) SetImage(image string) { o.inner.ImageRef = image }

// ---

type RemoveImageRequest_17 struct {
	inner *runtimeapi.RemoveImageRequest
}

var _ RemoveImageRequest = &RemoveImageRequest_17{}

func (o *RemoveImageRequest_17) Unwrap() interface{} { return o.inner }
func (o *RemoveImageRequest_17) Image() string       { return o.inner.Image.GetImage() }
func (o *RemoveImageRequest_17) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type RemoveImageResponse_17 struct {
	inner *runtimeapi.RemoveImageResponse
}

var _ RemoveImageResponse = &RemoveImageResponse_17{}

func (o *RemoveImageResponse_17) Unwrap() interface{} { return o.inner }

// ---

type ImageFsInfoRequest_17 struct {
	inner *runtimeapi.ImageFsInfoRequest
}

var _ ImageFsInfoRequest = &ImageFsInfoRequest_17{}

func (o *ImageFsInfoRequest_17) Unwrap() interface{} { return o.inner }

// ---

type ImageFsInfoResponse_17 struct {
	inner *runtimeapi.ImageFsInfoResponse
}

var _ ImageFsInfoResponse = &ImageFsInfoResponse_17{}

func (o *ImageFsInfoResponse_17) Unwrap() interface{} { return o.inner }
func (o *ImageFsInfoResponse_17) Items() []CRIObject {
	var r []CRIObject
	for _, fs := range o.inner.ImageFilesystems {
		r = append(r, &FilesystemUsage_17{fs})
	}
	return r
}
func (o *ImageFsInfoResponse_17) SetItems(items []CRIObject) {
	o.inner.ImageFilesystems = nil
	for _, wrapped := range items {
		o.inner.ImageFilesystems = append(o.inner.ImageFilesystems, wrapped.Unwrap().(*runtimeapi.FilesystemUsage))
	}
}

// ---

type CRI17 struct{}

var _ CRIVersion = &CRI17{}

func (c *CRI17) Register(server *grpc.Server) {
	runtimeapi.RegisterDummyRuntimeServiceServer(server)
	runtimeapi.RegisterDummyImageServiceServer(server)
}

func (c *CRI17) ProbeRequest() (interface{}, interface{}) {
	return &runtimeapi.VersionRequest{}, &runtimeapi.VersionResponse{}
}

func (c *CRI17) WrapObject(o interface{}) (CRIObject, CRIObject, error) {
	switch v := o.(type) {
	case *runtimeapi.PodSandbox:
		return &PodSandbox_17{v}, nil, nil
	case *runtimeapi.Container:
		return &Container_17{v}, nil, nil
	case *runtimeapi.Image:
		return &Image_17{v}, nil, nil
	case *runtimeapi.PodSandboxStatus:
		return &PodSandboxStatus_17{v}, nil, nil
	case *runtimeapi.VersionRequest:
		return &VersionRequest_17{v}, &VersionResponse_17{&runtimeapi.VersionResponse{}}, nil
	case *runtimeapi.StatusRequest:
		return &StatusRequest_17{v}, &StatusResponse_17{&runtimeapi.StatusResponse{}}, nil
	case *runtimeapi.UpdateRuntimeConfigRequest:
		return &UpdateRuntimeConfigRequest_17{v}, &UpdateRuntimeConfigResponse_17{&runtimeapi.UpdateRuntimeConfigResponse{}}, nil
	case *runtimeapi.RunPodSandboxRequest:
		return &RunPodSandboxRequest_17{v}, &RunPodSandboxResponse_17{&runtimeapi.RunPodSandboxResponse{}}, nil
	case *runtimeapi.ListPodSandboxRequest:
		return &ListPodSandboxRequest_17{v}, &ListPodSandboxResponse_17{&runtimeapi.ListPodSandboxResponse{}}, nil
	case *runtimeapi.StopPodSandboxRequest:
		return &StopPodSandboxRequest_17{v}, &StopPodSandboxResponse_17{&runtimeapi.StopPodSandboxResponse{}}, nil
	case *runtimeapi.RemovePodSandboxRequest:
		return &RemovePodSandboxRequest_17{v}, &RemovePodSandboxResponse_17{&runtimeapi.RemovePodSandboxResponse{}}, nil
	case *runtimeapi.PodSandboxStatusRequest:
		return &PodSandboxStatusRequest_17{v}, &PodSandboxStatusResponse_17{&runtimeapi.PodSandboxStatusResponse{}}, nil
	case *runtimeapi.CreateContainerRequest:
		return &CreateContainerRequest_17{v}, &CreateContainerResponse_17{&runtimeapi.CreateContainerResponse{}}, nil
	case *runtimeapi.ListContainersRequest:
		return &ListContainersRequest_17{v}, &ListContainersResponse_17{&runtimeapi.ListContainersResponse{}}, nil
	case *runtimeapi.ListContainerStatsRequest:
		return &ListContainerStatsRequest_17{v}, &ListContainerStatsResponse_17{&runtimeapi.ListContainerStatsResponse{}}, nil
	case *runtimeapi.StartContainerRequest:
		return &StartContainerRequest_17{v}, &StartContainerResponse_17{&runtimeapi.StartContainerResponse{}}, nil
	case *runtimeapi.StopContainerRequest:
		return &StopContainerRequest_17{v}, &StopContainerResponse_17{&runtimeapi.StopContainerResponse{}}, nil
	case *runtimeapi.RemoveContainerRequest:
		return &RemoveContainerRequest_17{v}, &RemoveContainerResponse_17{&runtimeapi.RemoveContainerResponse{}}, nil
	case *runtimeapi.ContainerStatusRequest:
		return &ContainerStatusRequest_17{v}, &ContainerStatusResponse_17{&runtimeapi.ContainerStatusResponse{}}, nil
	case *runtimeapi.ContainerStatsRequest:
		return &ContainerStatsRequest_17{v}, &ContainerStatsResponse_17{&runtimeapi.ContainerStatsResponse{}}, nil
	case *runtimeapi.ExecSyncRequest:
		return &ExecSyncRequest_17{v}, &ExecSyncResponse_17{&runtimeapi.ExecSyncResponse{}}, nil
	case *runtimeapi.ExecRequest:
		return &ExecRequest_17{v}, &ExecResponse_17{&runtimeapi.ExecResponse{}}, nil
	case *runtimeapi.AttachRequest:
		return &AttachRequest_17{v}, &AttachResponse_17{&runtimeapi.AttachResponse{}}, nil
	case *runtimeapi.PortForwardRequest:
		return &PortForwardRequest_17{v}, &PortForwardResponse_17{&runtimeapi.PortForwardResponse{}}, nil
	case *runtimeapi.ListImagesRequest:
		return &ListImagesRequest_17{v}, &ListImagesResponse_17{&runtimeapi.ListImagesResponse{}}, nil
	case *runtimeapi.ImageStatusRequest:
		return &ImageStatusRequest_17{v}, &ImageStatusResponse_17{&runtimeapi.ImageStatusResponse{}}, nil
	case *runtimeapi.PullImageRequest:
		return &PullImageRequest_17{v}, &PullImageResponse_17{&runtimeapi.PullImageResponse{}}, nil
	case *runtimeapi.RemoveImageRequest:
		return &RemoveImageRequest_17{v}, &RemoveImageResponse_17{&runtimeapi.RemoveImageResponse{}}, nil
	case *runtimeapi.ImageFsInfoRequest:
		return &ImageFsInfoRequest_17{v}, &ImageFsInfoResponse_17{&runtimeapi.ImageFsInfoResponse{}}, nil
	default:
		return nil, nil, fmt.Errorf("can't wrap %T", o)
	}
}
