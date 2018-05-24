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
	"reflect"
	"strings"
)

type typeMatcher struct {
	typeMap     map[string]reflect.Type
	responseMap map[string]reflect.Type
}

func newTypeMatcher() *typeMatcher {
	return &typeMatcher{
		typeMap:     make(map[string]reflect.Type),
		responseMap: make(map[string]reflect.Type),
	}
}

func (tm *typeMatcher) registerTypes(objects ...interface{}) {
	for _, o := range objects {
		t := reflect.TypeOf(o)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		typeName := t.Name()
		p := strings.LastIndex(typeName, "_")
		if p > 0 {
			typeName = typeName[:p]
		}
		tm.typeMap[typeName] = t
		if strings.HasSuffix(typeName, "Response") {
			requestTypeName := typeName[:len(typeName)-8] + "Request"
			tm.responseMap[requestTypeName] = t
		}
	}
}

func (tm *typeMatcher) matchType(o interface{}) (interface{}, interface{}, error) {
	t := reflect.TypeOf(o)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	targetType, found := tm.typeMap[t.Name()]
	if !found {
		return nil, nil, fmt.Errorf("matching type not found for %T", o)
	}
	var resp interface{}
	respType, respFound := tm.responseMap[t.Name()]
	if respFound {
		resp = reflect.New(respType).Interface()
	}
	return reflect.New(targetType).Interface(), resp, nil
}
