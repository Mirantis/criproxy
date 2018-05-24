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
	"net"
	"os"
	"syscall"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Interceptor specifies an interceptor to be used by gRPC server.
type Interceptor interface {
	// Register registers CRI services for proxy's CRI version within the Server.
	Register(s *grpc.Server)
	// Match checks whether fullMethod can be handled by this proxy instance.
	Match(fullMethod string) bool
	// Intercept handles a CRI request. It's invoked from a gRPC interceptor.
	Intercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
	// Stop disconnects from all the CRI servers
	Stop()
}

// Server denotes a gRPC server.
type Server struct {
	server       *grpc.Server
	interceptors []Interceptor
}

// NewServer makes a new gRPC server.
func NewServer(interceptors []Interceptor, hook func()) *Server {
	s := &Server{interceptors: interceptors}
	s.server = grpc.NewServer(grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if hook != nil {
			hook()
		}
		return s.intercept(ctx, req, info, handler)
	}))
	for _, intc := range s.interceptors {
		intc.Register(s.server)
	}
	return s
}

func (s *Server) intercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	for _, intc := range s.interceptors {
		if intc.Match(info.FullMethod) {
			return intc.Intercept(ctx, req, info, handler)
		}
	}
	return nil, fmt.Errorf("no interceptor for method %q", info.FullMethod)
}

// Serve makes the server listen on the specified addr. If readyCh is
// not nil, it'll be closed when the server is ready to accept
// connections.
func (s *Server) Serve(addr string, readyCh chan struct{}) error {
	if err := syscall.Unlink(addr); err != nil && !os.IsNotExist(err) {
		return err
	}
	ln, err := net.Listen("unix", addr)
	if err != nil {
		return err
	}
	defer ln.Close()
	if readyCh != nil {
		close(readyCh)
	}
	return s.server.Serve(ln)
}

// Stop stops the server.
func (s *Server) Stop() {
	for _, intc := range s.interceptors {
		intc.Stop()
	}
	s.server.GracefulStop()
}
