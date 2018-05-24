/*
Copyright 2016 Mirantis

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Based on pkg/kubelet/api/testing/utils.go from Kubernetes project.
Original copyright notice follows:

Copyright 2016 The Kubernetes Authors.

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

package testing

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

const (
	journalWaitInterval = 1000 * time.Millisecond
	journalWaitCount    = 1500
)

var (
	version = "0.1.0"

	FakeRuntimeName  = "fakeRuntime"
	FakePodSandboxIP = "192.168.192.168"
)

func filterInLabels(filter, labels map[string]string) bool {
	for k, v := range filter {
		if value, ok := labels[k]; ok {
			if value != v {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

// Journal records a series of events (items) represented as strings in a
// thread-safe way
type Journal interface {
	// Record saves the specified item in the journal
	Record(item string)
}

// SimpleJournal is an implementation of Journal that has methods
// for waiting for particular events and verifying journal contents
type SimpleJournal struct {
	sync.Mutex
	Items []string
	skip  map[string]bool
}

// NewSimpleJournal creates an instance of SimpleJournal
func NewSimpleJournal() *SimpleJournal {
	return &SimpleJournal{skip: make(map[string]bool)}
}

// Skip makes SimpleJournal skip the specified items when
// they occur
func (j *SimpleJournal) Skip(item string) {
	j.skip[item] = true
}

// Record implements Record method of Journal interface
func (j *SimpleJournal) Record(item string) {
	j.Lock()
	defer j.Unlock()

	if !j.skip[item] {
		j.Items = append(j.Items, item)
	}
}

// Verify verifies that the current contents of the journal is expectedItems,
// returns nil if so or an error otherwise
func (j *SimpleJournal) Verify(expectedItems []string) error {
	j.Lock()
	defer j.Unlock()

	actualItems := j.Items
	j.Items = nil
	if !reflect.DeepEqual(actualItems, expectedItems) {
		return fmt.Errorf("bad journal items. Expected %v, got %v", expectedItems, actualItems)
	}
	return nil
}

// PrefixJournal is an implementation of Journal interface that prefixes
// every item passed to it with the specified prefix before passing it on
// to the underlying Journal
type PrefixJournal struct {
	journal Journal
	prefix  string
}

// NewPrefixJournal creates an instance of PrefixJournal with the
// specified underlying journal and prefix
func NewPrefixJournal(journal Journal, prefix string) *PrefixJournal {
	return &PrefixJournal{journal, prefix}
}

// Record implements Record method of Journal interface
func (j *PrefixJournal) Record(item string) {
	j.journal.Record(j.prefix + item)
}
