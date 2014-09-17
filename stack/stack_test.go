// Copyright 2014 Frank Guan
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package stack

import (
	"testing"
)

var expected = []interface{}{1, 2, 3, "hello"}

func TestStack(t *testing.T) {
	s := New()

	if !s.Empty() {
		t.Errorf("expected empty stack")
	}

	if s.Size() != 0 {
		t.Errorf("expected stack of size 0")
	}
	
	// push items from expected
	for _, v := range expected {
		s.Push(v)
	}
	
	if s.Size() != len(expected) {
		t.Errorf("expected %v, got %v", len(expected), s.Size())
	}
	
	
	item := s.Pop()
	for j := len(expected)-1; item != nil; j-- {
		t.Logf("item is: %v", item)
		if item != expected[j] {
			t.Errorf("expected %v, got %v", expected[j], item)
		}
		item = s.Pop()
	}
}
