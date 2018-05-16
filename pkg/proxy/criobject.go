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
	"google.golang.org/grpc"
)

type CRIObject interface {
	Wrap(v interface{})
	Unwrap() interface{}
}

type Upgradable interface {
	Upgrade() (interface{}, error)
}

type Downgradable interface {
	Downgrade() (interface{}, error)
}

type IdObject interface {
	Id() string
	SetId(string)
}

type PodSandboxIdObject interface {
	PodSandboxId() string
	SetPodSandboxId(string)
}

type ContainerIdObject interface {
	ContainerId() string
	SetContainerId(string)
}

type ImageObject interface {
	Image() string
	SetImage(string)
}

type IdFilterObject interface {
	IdFilter() string
	SetIdFilter(string)
}

type PodSandboxIdFilterObject interface {
	PodSandboxIdFilter() string
	SetPodSandboxIdFilter(string)
}

type ImageFilterObject interface {
	ImageFilter() string
	SetImageFilter(string)
}

type UrlObject interface {
	Url() string
	SetUrl(string)
}

type ObjectList interface {
	Items() []CRIObject
	SetItems([]CRIObject)
}

type PodSandbox interface {
	CRIObject
	IdObject
	Copy() PodSandbox
}

type Container interface {
	CRIObject
	IdObject
	PodSandboxIdObject
	ImageObject
	Copy() Container
}

type ContainerStats interface {
	CRIObject
	IdObject
	Copy() ContainerStats
}

type Image interface {
	CRIObject
	IdObject
	Copy() Image
	RepoTags() []string
	SetRepoTags([]string)
}

type PodSandboxStatus interface {
	CRIObject
	IdObject
	Copy() PodSandboxStatus
}

type ContainerStatus interface {
	CRIObject
	IdObject
	ImageObject
	Copy() ContainerStatus
}

type FilesystemUsage interface {
	CRIObject
}

type VersionRequest interface {
	CRIObject
}

type VersionResponse interface {
	CRIObject
}

type StatusRequest interface {
	CRIObject
}

type StatusResponse interface {
	CRIObject
}

type UpdateRuntimeConfigRequest interface {
	CRIObject
}

type UpdateRuntimeConfigResponse interface {
	CRIObject
}

type RunPodSandboxRequest interface {
	CRIObject
	GetAnnotations() map[string]string
}

type RunPodSandboxResponse interface {
	CRIObject
	PodSandboxIdObject
}

type ListPodSandboxRequest interface {
	CRIObject
	IdFilterObject
}

type ListPodSandboxResponse interface {
	CRIObject
	ObjectList
}

type StopPodSandboxRequest interface {
	CRIObject
	PodSandboxIdObject
}

type StopPodSandboxResponse interface {
	CRIObject
}

type RemovePodSandboxRequest interface {
	CRIObject
	PodSandboxIdObject
}

type RemovePodSandboxResponse interface {
	CRIObject
}

type PodSandboxStatusRequest interface {
	CRIObject
	PodSandboxIdObject
}

type PodSandboxStatusResponse interface {
	CRIObject
	Status() PodSandboxStatus
}

type CreateContainerRequest interface {
	CRIObject
	PodSandboxIdObject
	ImageObject
}

type CreateContainerResponse interface {
	CRIObject
	ContainerIdObject
}

type ListContainersRequest interface {
	CRIObject
	IdFilterObject
	PodSandboxIdFilterObject
}

type ListContainersResponse interface {
	CRIObject
	ObjectList
}

type ListContainerStatsRequest interface {
	CRIObject
	IdFilterObject
	PodSandboxIdFilterObject
}

type ListContainerStatsResponse interface {
	CRIObject
	ObjectList
}

type StartContainerRequest interface {
	CRIObject
	ContainerIdObject
}

type StartContainerResponse interface {
	CRIObject
}

type StopContainerRequest interface {
	CRIObject
	ContainerIdObject
}

type StopContainerResponse interface {
	CRIObject
}

type RemoveContainerRequest interface {
	CRIObject
	ContainerIdObject
}

type RemoveContainerResponse interface {
	CRIObject
}

type UpdateContainerResourcesRequest interface {
	CRIObject
	ContainerIdObject
}

type UpdateContainerResourcesResponse interface {
	CRIObject
}

type ContainerStatusRequest interface {
	CRIObject
	ContainerIdObject
}

type ContainerStatusResponse interface {
	CRIObject
	Status() ContainerStatus
}

type ContainerStatsRequest interface {
	CRIObject
	ContainerIdObject
}

type ContainerStatsResponse interface {
	CRIObject
	Stats() ContainerStats
}

type ExecSyncRequest interface {
	CRIObject
	ContainerIdObject
}

type ExecSyncResponse interface {
	CRIObject
}

type ExecRequest interface {
	CRIObject
	ContainerIdObject
}

type ExecResponse interface {
	CRIObject
	UrlObject
}

type AttachRequest interface {
	CRIObject
	ContainerIdObject
}

type AttachResponse interface {
	CRIObject
	UrlObject
}

type PortForwardRequest interface {
	CRIObject
	PodSandboxIdObject
}

type PortForwardResponse interface {
	CRIObject
	UrlObject
}

type ReopenContainerLogRequest interface {
	CRIObject
	ContainerIdObject
}

type ReopenContainerLogResponse interface {
	CRIObject
}

type ListImagesRequest interface {
	CRIObject
	ImageFilterObject
}

type ListImagesResponse interface {
	CRIObject
	ObjectList
}

type ImageStatusRequest interface {
	CRIObject
	ImageObject
}

type ImageStatusResponse interface {
	CRIObject
	Image() Image
	SetImage(Image)
}

type PullImageRequest interface {
	CRIObject
	ImageObject
}

type PullImageResponse interface {
	CRIObject
	ImageObject
}

type RemoveImageRequest interface {
	CRIObject
	ImageObject
}

type RemoveImageResponse interface {
	CRIObject
}

type ImageFsInfoRequest interface {
	CRIObject
}

type ImageFsInfoResponse interface {
	CRIObject
	ObjectList
}

type CRIVersion interface {
	Register(*grpc.Server)
	ProbeRequest() (interface{}, interface{})
	WrapObject(interface{}) (CRIObject, CRIObject, error)
	ProtoPackage() string
}

type UpgradableCRIVersion interface {
	CRIVersion
	UpgradesTo() CRIVersion
}

func wrapUsingMatcher(tm *typeMatcher, o interface{}) (CRIObject, CRIObject, error) {
	if o == nil {
		return nil, nil, nil
	}
	out, resp, err := tm.matchType(o)
	if err != nil {
		return nil, nil, err
	}
	out.(CRIObject).Wrap(o)
	if resp != nil {
		resp.(CRIObject).Wrap(nil)
		return out.(CRIObject), resp.(CRIObject), nil
	}
	return out.(CRIObject), nil, nil
}
