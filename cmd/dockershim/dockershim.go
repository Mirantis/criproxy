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

Based on kubelet code from Kubernetes project.
Original copyright notice follows:

Copyright 2015 The Kubernetes Authors.

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

package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/golang/glog"
	"github.com/spf13/pflag"

	"k8s.io/apiserver/pkg/util/flag"
	"k8s.io/apiserver/pkg/util/logs"
	"k8s.io/kubernetes/cmd/kubelet/app/options"
	"k8s.io/kubernetes/pkg/apis/componentconfig"
	"k8s.io/kubernetes/pkg/kubelet"
	"k8s.io/kubernetes/pkg/kubelet/dockershim"
	"k8s.io/kubernetes/pkg/kubelet/dockershim/libdocker"
	dockerremote "k8s.io/kubernetes/pkg/kubelet/dockershim/remote"
	"k8s.io/kubernetes/pkg/kubelet/server/streaming"
	"k8s.io/kubernetes/pkg/version/verflag"
)

func getAddressForStreaming() (string, error) {
	// FIXME: see kubelet.setNodeAddress
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	addrs, err := net.LookupIP(hostname)
	for _, addr := range addrs {
		if !addr.IsLoopback() && addr.To4() != nil {
			return addr.String(), nil
		}
	}
	return "", errors.New("unable to get IP address")
}

// RunDockershim only starts the dockershim in current process. This is only used for cri validate testing purpose
// TODO(random-liu): Move this to a separate binary.
func RunDockershim(c *componentconfig.KubeletConfiguration, r *options.ContainerRuntimeOptions) error {
	streamingAddr, err := getAddressForStreaming()
	if err != nil {
		return err
	}

	// Create docker client.
	dockerClient := libdocker.ConnectToDockerOrDie(r.DockerEndpoint, c.RuntimeRequestTimeout.Duration,
		r.ImagePullProgressDeadline.Duration)

	// Initialize network plugin settings.
	binDir := r.CNIBinDir
	if binDir == "" {
		binDir = r.NetworkPluginDir
	}
	nh := &kubelet.NoOpLegacyHost{}
	pluginSettings := dockershim.NetworkPluginSettings{
		HairpinMode:       componentconfig.HairpinMode(c.HairpinMode),
		NonMasqueradeCIDR: c.NonMasqueradeCIDR,
		PluginName:        r.NetworkPluginName,
		PluginConfDir:     r.CNIConfDir,
		PluginBinDir:      binDir,
		MTU:               int(r.NetworkPluginMTU),
		LegacyRuntimeHost: nh,
	}

	// Initialize streaming configuration. (Not using TLS now)
	streamingConfig := &streaming.Config{
		Addr:                            streamingAddr + ":12345", // FIXME
		StreamIdleTimeout:               c.StreamingConnectionIdleTimeout.Duration,
		StreamCreationTimeout:           streaming.DefaultConfig.StreamCreationTimeout,
		SupportedRemoteCommandProtocols: streaming.DefaultConfig.SupportedRemoteCommandProtocols,
		SupportedPortForwardProtocols:   streaming.DefaultConfig.SupportedPortForwardProtocols,
	}

	ds, err := dockershim.NewDockerService(dockerClient, c.SeccompProfileRoot, r.PodSandboxImage,
		streamingConfig, &pluginSettings, c.RuntimeCgroups, c.CgroupDriver, r.DockerExecHandlerName, r.DockershimRootDirectory,
		r.DockerDisableSharedPID)
	if err != nil {
		return err
	}
	if err := ds.Start(); err != nil {
		return err
	}

	glog.V(2).Infof("Starting the GRPC server for the docker CRI shim.")
	server := dockerremote.NewDockerServer(c.RemoteRuntimeEndpoint, ds)
	if err := server.Start(); err != nil {
		return err
	}

	// Start the streaming server
	return http.ListenAndServe(streamingConfig.Addr, ds)
}

func main() {
	s := options.NewKubeletServer()
	s.AddFlags(pflag.CommandLine)

	flag.InitFlags()
	logs.InitLogs()
	defer logs.FlushLogs()

	verflag.PrintAndExitIfRequested()

	if err := RunDockershim(&s.KubeletConfiguration, &s.ContainerRuntimeOptions); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}