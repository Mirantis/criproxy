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

package runtimeapis

import (
	"log"
	"reflect"
	"testing"

	"github.com/ghodss/yaml"

	v1_10 "github.com/Mirantis/criproxy/pkg/runtimeapis/v1_10"
	v1_9 "github.com/Mirantis/criproxy/pkg/runtimeapis/v1_9"
)

const (
	podUid1       = "4bde9008-4663-4342-84ed-310cea787f95"
	podSandboxId1 = "pod-1-1_default_" + podUid1 + "_0"
)

func mustYaml(v interface{}) string {
	out, err := yaml.Marshal(v)
	if err != nil {
		log.Panicf("Error marshalling yaml: %v", err)
	}
	return string(out)
}

func podSandboxConfig9(nsOption *v1_9.NamespaceOption) *v1_9.PodSandboxConfig {
	r := &v1_9.PodSandboxConfig{
		Metadata: &v1_9.PodSandboxMetadata{
			Name:      "pod-1-1",
			Uid:       podUid1,
			Namespace: "default",
			Attempt:   0,
		},
		Labels: map[string]string{"name": "pod-1-1"},
	}
	if nsOption != nil {
		r.Linux = &v1_9.LinuxPodSandboxConfig{
			SecurityContext: &v1_9.LinuxSandboxSecurityContext{
				NamespaceOptions: nsOption,
			},
		}
	}
	return r
}

func podSandboxConfig10(nsOption *v1_10.NamespaceOption) *v1_10.PodSandboxConfig {
	r := &v1_10.PodSandboxConfig{
		Metadata: &v1_10.PodSandboxMetadata{
			Name:      "pod-1-1",
			Uid:       podUid1,
			Namespace: "default",
			Attempt:   0,
		},
		Labels: map[string]string{"name": "pod-1-1"},
	}
	if nsOption != nil {
		r.Linux = &v1_10.LinuxPodSandboxConfig{
			SecurityContext: &v1_10.LinuxSandboxSecurityContext{
				NamespaceOptions: nsOption,
			},
		}
	}
	return r
}

func TestConversion(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   interface{}
		out  interface{}
	}{
		{
			name: "RunPodSandboxRequest",
			in: &v1_9.RunPodSandboxRequest{
				Config: podSandboxConfig9(nil),
			},
			out: &v1_10.RunPodSandboxRequest{
				Config: podSandboxConfig10(nil),
			},
		},
		{
			name: "RunPodSandboxRequest (host network)",
			in: &v1_9.RunPodSandboxRequest{
				Config: podSandboxConfig9(&v1_9.NamespaceOption{HostNetwork: true}),
			},
			out: &v1_10.RunPodSandboxRequest{
				Config: podSandboxConfig10(&v1_10.NamespaceOption{
					Network: v1_10.NamespaceMode_NODE,
				}),
			},
		},
		{
			name: "RunPodSandboxRequest (host pid)",
			in: &v1_9.RunPodSandboxRequest{
				Config: podSandboxConfig9(&v1_9.NamespaceOption{HostPid: true}),
			},
			out: &v1_10.RunPodSandboxRequest{
				Config: podSandboxConfig10(&v1_10.NamespaceOption{
					Pid: v1_10.NamespaceMode_NODE,
				}),
			},
		},
		{
			name: "RunPodSandboxRequest (host ipc)",
			in: &v1_9.RunPodSandboxRequest{
				Config: podSandboxConfig9(&v1_9.NamespaceOption{HostIpc: true}),
			},
			out: &v1_10.RunPodSandboxRequest{
				Config: podSandboxConfig10(&v1_10.NamespaceOption{
					Ipc: v1_10.NamespaceMode_NODE,
				}),
			},
		},
		{
			name: "CreateContainerRequest",
			in: &v1_9.CreateContainerRequest{
				PodSandboxId: podSandboxId1,
				Config: &v1_9.ContainerConfig{
					Metadata: &v1_9.ContainerMetadata{
						Name:    "container1",
						Attempt: 0,
					},
					Image: &v1_9.ImageSpec{
						Image: "image1-1",
					},
				},
			},
			out: &v1_10.CreateContainerRequest{
				PodSandboxId: podSandboxId1,
				Config: &v1_10.ContainerConfig{
					Metadata: &v1_10.ContainerMetadata{
						Name:    "container1",
						Attempt: 0,
					},
					Image: &v1_10.ImageSpec{
						Image: "image1-1",
					},
				},
			},
		},
	} {
		out, err := Upgrade(tc.in)
		switch {
		case err != nil:
			t.Fatalf("Upgrade: %v", err)
		case !reflect.DeepEqual(out, tc.out):
			t.Errorf("bad conversion: expected:\n%s\nactual:\n%s", mustYaml(tc.out), mustYaml(out))
		}

		switch out1, err := Upgrade(out); {
		case err != nil:
			t.Errorf("Upgrade (repeated): %v", err)
		case out1 != out:
			t.Errorf("Upgrade is not idempotent")
		}

		back, err := Downgrade(out)
		switch {
		case err != nil:
			t.Errorf("Downgrade: %v", err)
		case !reflect.DeepEqual(back, tc.in):
			t.Errorf("bad reverse conversion: expected:\n%s\nactual:\n%s", mustYaml(tc.in), mustYaml(back))
		}

		switch back1, err := Downgrade(back); {
		case err != nil:
			t.Errorf("Downgrade (repeated): %v", err)
		case back1 != back:
			t.Errorf("Downgrade is not idempotent")
		}
	}
}
