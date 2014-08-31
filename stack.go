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

type Stack struct {
	size int
	items []interface{}
}

func New() *Stack {
	return &Stack{0, nil}
}

func (stack *Stack) Push(item interface{}) {
	stack.items = append(stack.items, item)
	stack.size++
}

func (stack *Stack) Pop() interface{} {
	var item interface{} = nil
	if stack.size > 0 {
		item = stack.items[stack.size-1]
		stack.items[stack.size-1] = nil
		stack.size--
	}
	return item
}

func (stack *Stack) Empty() bool {
	return stack.size == 0
}

func (stack *Stack) Size() int {
	return stack.size
}

func (stack *Stack) Peek() interface{} {
	var item interface{} = nil
	if stack.size > 0 {
		item = stack.items[stack.size-1]
	}
	return item
}
