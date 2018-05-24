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
	"reflect"
	"testing"
)

type SomeValue struct{}
type SomeRequest struct{}
type SomeResponse struct{}

type SomeValue_123 struct{}
type SomeRequest_123 struct{}
type SomeResponse_123 struct{}

func TestTypeMatcher(t *testing.T) {
	tm := newTypeMatcher()
	tm.registerTypes(SomeValue_123{}, &SomeRequest_123{}, &SomeResponse_123{})
	for _, tc := range []struct {
		name                string
		in                  interface{}
		expectedWrapper     interface{}
		expectedRespWrapper interface{}
	}{
		{
			name:            "SomeValue",
			in:              &SomeValue{},
			expectedWrapper: &SomeValue_123{},
		},
		{
			name:                "SomeRequest",
			in:                  &SomeRequest{},
			expectedWrapper:     &SomeRequest_123{},
			expectedRespWrapper: &SomeResponse_123{},
		},
		{
			name:            "SomeResponse",
			in:              &SomeResponse{},
			expectedWrapper: &SomeResponse_123{},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out, resp, err := tm.matchType(tc.in)
			if err != nil {
				t.Fatalf("matchType(): %v", err)
			}
			if reflect.TypeOf(out) != reflect.TypeOf(tc.expectedWrapper) {
				t.Errorf("bad wrapper: %T instead of %T", out, tc.expectedWrapper)
			}
			switch {
			case resp == nil && tc.expectedRespWrapper == nil:
				// ok
			case tc.expectedRespWrapper == nil:
				t.Errorf("got unexpected response wrapper")
			case resp == nil:
				t.Errorf("didn't get expected response wrapper")
			case reflect.TypeOf(resp) != reflect.TypeOf(tc.expectedRespWrapper):
				t.Errorf("bad response wrapper: %T instead of %T", resp, tc.expectedRespWrapper)
			}
		})
	}
}
