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
	"os"
	"strings"
	"time"

	proxy "github.com/Mirantis/criproxy/pkg/proxy"
	"github.com/golang/glog"
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
	apiServerHost = flag.String("apiserver", "", "apiserver URL")
)

// runCriProxy starts CRI proxy
func runCriProxy(connect, listen string) error {
	addrs := strings.Split(connect, ",")
	proxy, err := proxy.NewRuntimeProxy(addrs, connectionTimeout, nil)
	if err != nil {
		return fmt.Errorf("Error starting CRI proxy: %v", err)
	}
	glog.V(1).Infof("Starting CRI proxy on socket %s", listen)
	if err := proxy.Serve(listen, nil); err != nil {
		return fmt.Errorf("Serving failed: %v", err)
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
