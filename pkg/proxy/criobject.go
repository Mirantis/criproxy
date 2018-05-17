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

// CRIObject denotes a wrapped CRI object.
type CRIObject interface {
	// Wrap sets the inner object to specified value.  If v is
	// nil, a new instance of the corresponding inner type is used
	// to initialize the inner object. Otherwise v must be a pointer
	// to a raw CRI object.
	Wrap(v interface{})
	// Unwrap returns the underlying raw CRI object.
	Unwrap() interface{}
}

// Upgradable denotes a CRI object that can be upgraded to a newer CRI version.
type Upgradable interface {
	// Upgrade returns the underlying object converted to a newer
	// CRI version as a raw CRI object.
	Upgrade() (interface{}, error)
}

// Downgradable denotes a CRI object that can be downgraded to an older CRI version.
type Downgradable interface {
	// Downgraded returns the underlying object converted to an
	// older CRI version as a raw CRI object.
	Downgrade() (interface{}, error)
}

// IdObject is a wrapped CRI object that has an id.
type IdObject interface {
	// Id returns the id of the object.
	Id() string
	// SetId sets the id of the object.
	SetId(string)
}

// PodSandboxIdObject is a wrapped CRI object that contains a pod sandbox id.
type PodSandboxIdObject interface {
	// PodSandboxId returns the pod sandbox id of the object.
	PodSandboxId() string
	// PodSandboxId set the pod sandbox id of the object.
	SetPodSandboxId(string)
}

// ContainerIdObject is a wrapped CRI object that contains a container id.
type ContainerIdObject interface {
	// ContainerId returns the container id of the object.
	ContainerId() string
	// SetContainerId returns the container id of the object.
	SetContainerId(string)
}

// ImageObject is a wrapped CRI object that contains an image reference.
type ImageObject interface {
	// Image returns the image reference of the object.
	Image() string
	// SEtImage sets the image reference of the object.
	SetImage(string)
}

// IdFilterObject is a wrapped CRI object that denotes a filter that uses an id.
type IdFilterObject interface {
	// IdFilter returns the id used by the filter.
	IdFilter() string
	// SetIdFilter sets the id used by the filter.
	SetIdFilter(string)
}

// IdFilterObject is a wrapped CRI object that denotes a filter that uses a pod sandbox id.
type PodSandboxIdFilterObject interface {
	// PodSandboxIdFilter returns the pod sadnbox id used by the filter.
	PodSandboxIdFilter() string
	// SetPodSandboxIdFilter sets the pod sadnbox id used by the filter.
	SetPodSandboxIdFilter(string)
}

// ImageFilterObject is a wrapped CRI object that denotes a filter that uses an image reference.
type ImageFilterObject interface {
	// ImageFilter returns the image reference that's used by the filter.
	ImageFilter() string
	// ImageFilter sets the image reference that's used by the filter.
	SetImageFilter(string)
}

// UrlObject is a wrapped CRI object that contains an URL.
type UrlObject interface {
	// Url returns the url contained in the object.
	Url() string
	// Url sets the url contained in the object.
	SetUrl(string)
}

// ObjectList denotes a wrapped CRI object that denotes a list of other CRI objects.
type ObjectList interface {
	// Items returns a slice of CRI objects that are contained in the list.
	Items() []CRIObject
	// SetItems sets the list of objects to the specified slice.
	SetItems([]CRIObject)
}

// PodSandbox wraps a CRI PodSandbox object
type PodSandbox interface {
	CRIObject
	IdObject
	Copy() PodSandbox
}

// Container wraps a CRI Container object
type Container interface {
	CRIObject
	IdObject
	PodSandboxIdObject
	ImageObject
	Copy() Container
}

// ContainerStats wraps a CRI ContainerStats object
type ContainerStats interface {
	CRIObject
	IdObject
	Copy() ContainerStats
}

// Image wraps a CRI Image object
type Image interface {
	CRIObject
	IdObject
	Copy() Image
	RepoTags() []string
	SetRepoTags([]string)
}

// PodSandboxStatus wraps a CRI PodSandboxStatus object
type PodSandboxStatus interface {
	CRIObject
	IdObject
	Copy() PodSandboxStatus
}

// ContainerStatus wraps a CRI ContainerStatus object
type ContainerStatus interface {
	CRIObject
	IdObject
	ImageObject
	Copy() ContainerStatus
}

// FilesystemUsage wraps a CRI FilesystemUsage object
type FilesystemUsage interface {
	CRIObject
}

// VersionRequest wraps a CRI VersionRequest object
type VersionRequest interface {
	CRIObject
}

// VersionResponse wraps a CRI VersionResponse object
type VersionResponse interface {
	CRIObject
}

// StatusRequest wraps a CRI StatusRequest object
type StatusRequest interface {
	CRIObject
}

// StatusResponse wraps a CRI StatusResponse object
type StatusResponse interface {
	CRIObject
}

// UpdateRuntimeConfigRequest wraps a CRI UpdateRuntimeConfigRequest object
type UpdateRuntimeConfigRequest interface {
	CRIObject
}

// UpdateRuntimeConfigResponse wraps a CRI UpdateRuntimeConfigResponse object
type UpdateRuntimeConfigResponse interface {
	CRIObject
}

// RunPodSandboxRequest wraps a CRI RunPodSandboxRequest object
type RunPodSandboxRequest interface {
	CRIObject
	GetAnnotations() map[string]string
}

// RunPodSandboxResponse wraps a CRI RunPodSandboxResponse object
type RunPodSandboxResponse interface {
	CRIObject
	PodSandboxIdObject
}

// ListPodSandboxRequest wraps a CRI ListPodSandboxRequest object
type ListPodSandboxRequest interface {
	CRIObject
	IdFilterObject
}

// ListPodSandboxResponse wraps a CRI ListPodSandboxResponse object
type ListPodSandboxResponse interface {
	CRIObject
	ObjectList
}

// StopPodSandboxRequest wraps a CRI StopPodSandboxRequest object
type StopPodSandboxRequest interface {
	CRIObject
	PodSandboxIdObject
}

// StopPodSandboxResponse wraps a CRI StopPodSandboxResponse object
type StopPodSandboxResponse interface {
	CRIObject
}

// RemovePodSandboxRequest wraps a CRI RemovePodSandboxRequest object
type RemovePodSandboxRequest interface {
	CRIObject
	PodSandboxIdObject
}

// RemovePodSandboxResponse wraps a CRI RemovePodSandboxResponse object
type RemovePodSandboxResponse interface {
	CRIObject
}

// PodSandboxStatusRequest wraps a CRI PodSandboxStatusRequest object
type PodSandboxStatusRequest interface {
	CRIObject
	PodSandboxIdObject
}

// PodSandboxStatusResponse wraps a CRI PodSandboxStatusResponse object
type PodSandboxStatusResponse interface {
	CRIObject
	Status() PodSandboxStatus
}

// CreateContainerRequest wraps a CRI CreateContainerRequest object
type CreateContainerRequest interface {
	CRIObject
	PodSandboxIdObject
	ImageObject
}

// CreateContainerResponse wraps a CRI CreateContainerResponse object
type CreateContainerResponse interface {
	CRIObject
	ContainerIdObject
}

// ListContainersRequest wraps a CRI ListContainersRequest object
type ListContainersRequest interface {
	CRIObject
	IdFilterObject
	PodSandboxIdFilterObject
}

// ListContainersResponse wraps a CRI ListContainersResponse object
type ListContainersResponse interface {
	CRIObject
	ObjectList
}

// ListContainerStatsRequest wraps a CRI ListContainerStatsRequest object
type ListContainerStatsRequest interface {
	CRIObject
	IdFilterObject
	PodSandboxIdFilterObject
}

// ListContainerStatsResponse wraps a CRI ListContainerStatsResponse object
type ListContainerStatsResponse interface {
	CRIObject
	ObjectList
}

// StartContainerRequest wraps a CRI StartContainerRequest object
type StartContainerRequest interface {
	CRIObject
	ContainerIdObject
}

// StartContainerResponse wraps a CRI StartContainerResponse object
type StartContainerResponse interface {
	CRIObject
}

// StopContainerRequest wraps a CRI StopContainerRequest object
type StopContainerRequest interface {
	CRIObject
	ContainerIdObject
}

// StopContainerResponse wraps a CRI StopContainerResponse object
type StopContainerResponse interface {
	CRIObject
}

// RemoveContainerRequest wraps a CRI RemoveContainerRequest object
type RemoveContainerRequest interface {
	CRIObject
	ContainerIdObject
}

// RemoveContainerResponse wraps a CRI RemoveContainerResponse object
type RemoveContainerResponse interface {
	CRIObject
}

// UpdateContainerResourcesRequest wraps a CRI UpdateContainerResourcesRequest object
type UpdateContainerResourcesRequest interface {
	CRIObject
	ContainerIdObject
}

// UpdateContainerResourcesResponse wraps a CRI UpdateContainerResourcesResponse object
type UpdateContainerResourcesResponse interface {
	CRIObject
}

// ContainerStatusRequest wraps a CRI ContainerStatusRequest object
type ContainerStatusRequest interface {
	CRIObject
	ContainerIdObject
}

// ContainerStatusResponse wraps a CRI ContainerStatusResponse object
type ContainerStatusResponse interface {
	CRIObject
	Status() ContainerStatus
}

// ContainerStatsRequest wraps a CRI ContainerStatsRequest object
type ContainerStatsRequest interface {
	CRIObject
	ContainerIdObject
}

// ContainerStatsResponse wraps a CRI ContainerStatsResponse object
type ContainerStatsResponse interface {
	CRIObject
	Stats() ContainerStats
}

// ExecSyncRequest wraps a CRI ExecSyncRequest object
type ExecSyncRequest interface {
	CRIObject
	ContainerIdObject
}

// ExecSyncResponse wraps a CRI ExecSyncResponse object
type ExecSyncResponse interface {
	CRIObject
}

// ExecRequest wraps a CRI ExecRequest object
type ExecRequest interface {
	CRIObject
	ContainerIdObject
}

// ExecResponse wraps a CRI ExecResponse object
type ExecResponse interface {
	CRIObject
	UrlObject
}

// AttachRequest wraps a CRI AttachRequest object
type AttachRequest interface {
	CRIObject
	ContainerIdObject
}

// AttachResponse wraps a CRI AttachResponse object
type AttachResponse interface {
	CRIObject
	UrlObject
}

// PortForwardRequest wraps a CRI PortForwardRequest object
type PortForwardRequest interface {
	CRIObject
	PodSandboxIdObject
}

// PortForwardResponse wraps a CRI PortForwardResponse object
type PortForwardResponse interface {
	CRIObject
	UrlObject
}

// ReopenContainerLogRequest wraps a CRI ReopenContainerLogRequest object
type ReopenContainerLogRequest interface {
	CRIObject
	ContainerIdObject
}

// ReopenContainerLogResponse wraps a CRI ReopenContainerLogResponse object
type ReopenContainerLogResponse interface {
	CRIObject
}

// ListImagesRequest wraps a CRI ListImagesRequest object
type ListImagesRequest interface {
	CRIObject
	ImageFilterObject
}

// ListImagesResponse wraps a CRI ListImagesResponse object
type ListImagesResponse interface {
	CRIObject
	ObjectList
}

// ImageStatusRequest wraps a CRI ImageStatusRequest object
type ImageStatusRequest interface {
	CRIObject
	ImageObject
}

// ImageStatusResponse wraps a CRI ImageStatusResponse object
type ImageStatusResponse interface {
	CRIObject
	Image() Image
	SetImage(Image)
}

// PullImageRequest wraps a CRI PullImageRequest object
type PullImageRequest interface {
	CRIObject
	ImageObject
}

// PullImageResponse wraps a CRI PullImageResponse object
type PullImageResponse interface {
	CRIObject
	ImageObject
}

// RemoveImageRequest wraps a CRI RemoveImageRequest object
type RemoveImageRequest interface {
	CRIObject
	ImageObject
}

// RemoveImageResponse wraps a CRI RemoveImageResponse object
type RemoveImageResponse interface {
	CRIObject
}

// ImageFsInfoRequest wraps a CRI ImageFsInfoRequest object
type ImageFsInfoRequest interface {
	CRIObject
}

// ImageFsInfoResponse wraps a CRI ImageFsInfoResponse object
type ImageFsInfoResponse interface {
	CRIObject
	ObjectList
}

// CRI version denotes a version of CRI.
type CRIVersion interface {
	// Register registers the CRI version with a gRPC Server.
	Register(*grpc.Server)
	// ProbeRequest returns raw CRI request and response objects
	// that can be used to check the server availability and
	// compatibility with this CRI version.
	ProbeRequest() (interface{}, interface{})
	// WrapObject wraps a raw CRI object and returns the wrapped
	// source object, and, in case if the object is a Request,
	// also an empty Response object that matches it
	WrapObject(interface{}) (CRIObject, CRIObject, error)
	// ProtoPackage returns proto package used by the CRI version.
	ProtoPackage() string
}

// UpgradableCRIVersion is a CRI version that supports upgrading of
// the objects for a newer CRI version.
type UpgradableCRIVersion interface {
	CRIVersion
	// UpgradesTo returns a CRI version this one upgrades to.
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
