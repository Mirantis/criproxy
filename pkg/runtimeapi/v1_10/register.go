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

package v1alpha2

import (
	"google.golang.org/grpc"
)

// CRI Proxy does all the work in its grpc interceptor,
// wo we don't need real handlers. Let's cheat grpc a bit

func RegisterDummyRuntimeServiceServer(s *grpc.Server) {
	desc := _RuntimeService_serviceDesc
	desc.HandlerType = (*interface{})(nil)
	s.RegisterService(&desc, struct{}{})
}

func RegisterDummyImageServiceServer(s *grpc.Server) {
	desc := _ImageService_serviceDesc
	desc.HandlerType = (*interface{})(nil)
	s.RegisterService(&desc, struct{}{})
}
