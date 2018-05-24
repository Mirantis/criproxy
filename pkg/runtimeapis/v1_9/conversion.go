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

package runtime

import (
	"unsafe"

	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"

	v1_10 "github.com/Mirantis/criproxy/pkg/runtimeapis/v1_10"
)

var (
	Scheme = runtime.NewScheme()
	// localSchemeBuilder is only needed to keep the generated code happy
	localSchemeBuilder = runtime.NewSchemeBuilder()
)

func init() {
	RegisterConversions(Scheme)
}

func Convert_v1_10_ContainerConfig_To_v1_9_ContainerConfig(in *v1_10.ContainerConfig, out *ContainerConfig, s conversion.Scope) error {
	// the body of this function is a copy of
	// Convert_v1_9_ContainerConfig_To_v1_10_ContainerConfig,
	// so leaving unsafe.Pointer stuff in place.
	out.Metadata = (*ContainerMetadata)(unsafe.Pointer(in.Metadata))
	out.Image = (*ImageSpec)(unsafe.Pointer(in.Image))
	out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
	out.Args = *(*[]string)(unsafe.Pointer(&in.Args))
	out.WorkingDir = in.WorkingDir
	out.Envs = *(*[]*KeyValue)(unsafe.Pointer(&in.Envs))
	out.Mounts = *(*[]*Mount)(unsafe.Pointer(&in.Mounts))
	out.Devices = *(*[]*Device)(unsafe.Pointer(&in.Devices))
	out.Labels = *(*map[string]string)(unsafe.Pointer(&in.Labels))
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	out.LogPath = in.LogPath
	out.Stdin = in.Stdin
	out.StdinOnce = in.StdinOnce
	out.Tty = in.Tty
	if in.Linux != nil {
		in, out := &in.Linux, &out.Linux
		*out = new(LinuxContainerConfig)
		if err := Convert_v1_10_LinuxContainerConfig_To_v1_9_LinuxContainerConfig(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Linux = nil
	}
	// no out.Windows
	return nil
}

func Convert_v1_10_FilesystemUsage_To_v1_9_FilesystemUsage(in *v1_10.FilesystemUsage, out *FilesystemUsage, s conversion.Scope) error {
	out.Timestamp = in.Timestamp
	// XXX: can't get old StorageId (UUID) from the new FsId which contains just a mount point
	out.UsedBytes = (*UInt64Value)(in.UsedBytes)
	out.InodesUsed = (*UInt64Value)(in.InodesUsed)
	return nil
}

func Convert_v1_9_FilesystemUsage_To_v1_10_FilesystemUsage(in *FilesystemUsage, out *v1_10.FilesystemUsage, s conversion.Scope) error {
	out.Timestamp = in.Timestamp
	// XXX: can't get new FsId which contains just a mount point from the old StorageId (UUID)
	out.UsedBytes = (*v1_10.UInt64Value)(in.UsedBytes)
	out.InodesUsed = (*v1_10.UInt64Value)(in.InodesUsed)
	return nil
}

func Convert_v1_10_NamespaceOption_To_v1_9_NamespaceOption(in *v1_10.NamespaceOption, out *NamespaceOption, s conversion.Scope) error {
	out.HostNetwork = in.Network == v1_10.NamespaceMode_NODE
	out.HostPid = in.Pid == v1_10.NamespaceMode_NODE
	out.HostIpc = in.Ipc == v1_10.NamespaceMode_NODE
	return nil
}

func Convert_v1_9_NamespaceOption_To_v1_10_NamespaceOption(in *NamespaceOption, out *v1_10.NamespaceOption, s conversion.Scope) error {
	out.Network = v1_10.NamespaceMode_POD
	if in.HostNetwork {
		out.Network = v1_10.NamespaceMode_NODE
	}
	out.Pid = v1_10.NamespaceMode_POD
	if in.HostPid {
		out.Pid = v1_10.NamespaceMode_NODE
	}
	out.Ipc = v1_10.NamespaceMode_POD
	if in.HostIpc {
		out.Ipc = v1_10.NamespaceMode_NODE
	}
	return nil
}

func Convert_v1_10_LinuxContainerSecurityContext_To_v1_9_LinuxContainerSecurityContext(in *v1_10.LinuxContainerSecurityContext, out *LinuxContainerSecurityContext, s conversion.Scope) error {
	// the body of this function is a copy of
	// Convert_v1_9_LinuxContainerSecurityContext_To_v1_10_LinuxContainerSecurityContext
	// so leaving unsafe.Pointer stuff in place.
	out.Capabilities = (*Capability)(unsafe.Pointer(in.Capabilities))
	out.Privileged = in.Privileged
	if in.NamespaceOptions != nil {
		in, out := &in.NamespaceOptions, &out.NamespaceOptions
		*out = new(NamespaceOption)
		if err := Convert_v1_10_NamespaceOption_To_v1_9_NamespaceOption(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.NamespaceOptions = nil
	}
	out.SelinuxOptions = (*SELinuxOption)(unsafe.Pointer(in.SelinuxOptions))
	out.RunAsUser = (*Int64Value)(unsafe.Pointer(in.RunAsUser))
	// no out.RunAsGroup
	out.RunAsUsername = in.RunAsUsername
	out.ReadonlyRootfs = in.ReadonlyRootfs
	out.SupplementalGroups = *(*[]int64)(unsafe.Pointer(&in.SupplementalGroups))
	out.ApparmorProfile = in.ApparmorProfile
	out.SeccompProfilePath = in.SeccompProfilePath
	out.NoNewPrivs = in.NoNewPrivs
	return nil
}

func Convert_v1_10_LinuxSandboxSecurityContext_To_v1_9_LinuxSandboxSecurityContext(in *v1_10.LinuxSandboxSecurityContext, out *LinuxSandboxSecurityContext, s conversion.Scope) error {
	if in.NamespaceOptions != nil {
		in, out := &in.NamespaceOptions, &out.NamespaceOptions
		*out = new(NamespaceOption)
		if err := Convert_v1_10_NamespaceOption_To_v1_9_NamespaceOption(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.NamespaceOptions = nil
	}
	out.SelinuxOptions = (*SELinuxOption)(unsafe.Pointer(in.SelinuxOptions))
	out.RunAsUser = (*Int64Value)(unsafe.Pointer(in.RunAsUser))
	out.ReadonlyRootfs = in.ReadonlyRootfs
	out.SupplementalGroups = *(*[]int64)(unsafe.Pointer(&in.SupplementalGroups))
	out.Privileged = in.Privileged
	out.SeccompProfilePath = in.SeccompProfilePath
	// no out.RunAsGroup
	return nil
}
