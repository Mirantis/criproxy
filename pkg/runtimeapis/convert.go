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
	"fmt"
	"reflect"

	v1_9 "github.com/Mirantis/criproxy/pkg/runtimeapis/v1_9"
	"github.com/gogo/protobuf/proto"
)

func convertTo(in interface{}, targetProtoPackage string) (interface{}, error) {
	targetTypeName := fmt.Sprintf("%s.%s", targetProtoPackage, reflect.TypeOf(in).Elem().Name())
	mtype := proto.MessageType(targetTypeName)
	if mtype == nil {
		return nil, fmt.Errorf("target type for %T not found in proto package %q", in, targetProtoPackage)
	}
	if reflect.TypeOf(in) == mtype {
		return in, nil
	}
	out := reflect.New(mtype.Elem()).Interface()
	return out, v1_9.Scheme.Convert(in, out, nil)
}

// Upgrade converts CRI 1.9 object to CRI 1.10 one. It just returns
// the object if it's already CRI 1.10.
func Upgrade(in interface{}) (interface{}, error) {
	return convertTo(in, "runtime.v1alpha2")
}

// Downgrade converts CRI 1.10 object to CRI 1.9 one. It just returns
// the object if it's already CRI 1.9.
func Downgrade(in interface{}) (interface{}, error) {
	return convertTo(in, "runtime")
}
