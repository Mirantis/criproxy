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

package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/golang/glog"

	"github.com/Mirantis/criproxy/pkg/proxy"
	"github.com/Mirantis/criproxy/pkg/utils"
)

const (
	// XXX: don't hardcode
	connectionTimeout = 30 * time.Second
)

var (
	listen = flag.String("listen", "/run/criproxy.sock",
		"The unix socket to listen on, e.g. /run/virtlet.sock")
	connect = flag.String("connect", "/var/run/dockershim.sock",
		"CRI runtime ids and unix socket(s) to connect to, e.g. /var/run/dockershim.sock,alt:/var/run/another.sock")
	streamPort    = flag.Int("streamPort", 11250, "streaming port of the default runtime")
	streamUrl     = flag.String("streamUrl", "", "streaming url of the default runtime (-streamPort is ignored if this value is set)")
	apiServerHost = flag.String("apiserver", "", "apiserver URL")
	criVersions   = []proxy.CRIVersion{&proxy.CRI19{}, &proxy.CRI110{}}
)

// runCriProxy starts CRI proxy
func runCriProxy(connect, listen string) error {
	addrs := strings.Split(connect, ",")
	var err error
	var realStreamUrl *url.URL
	if *streamUrl == "" {
		if realStreamUrl, err = utils.GetStreamUrl(*streamPort); err != nil {
			return fmt.Errorf("can't get stream url: %v", err)
		}
	} else {
		if realStreamUrl, err = url.Parse(*streamUrl); err != nil {
			return fmt.Errorf("invalid stream url %q: %v", *streamUrl, err)
		}
	}
	var interceptors []proxy.Interceptor
	for _, criVersion := range criVersions {
		proxy, err := proxy.NewRuntimeProxy(criVersion, addrs, connectionTimeout, realStreamUrl)
		if err != nil {
			return fmt.Errorf("error initializing CRI proxy: %v", err)
		}
		interceptors = append(interceptors, proxy)
	}
	glog.V(1).Infof("Starting CRI proxy on socket %s", listen)
	server := proxy.NewServer(interceptors, nil)
	if err := server.Serve(listen, nil); err != nil {
		return fmt.Errorf("serving failed: %v", err)
	}
	return nil
}

func main() {
	flag.Parse()
	if err := runCriProxy(*connect, *listen); err != nil {
		glog.Error(err)
		os.Exit(1)
	}
}
