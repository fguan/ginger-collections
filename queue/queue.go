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

// Package queue implements a FIFO queue
package queue

type Queue struct {
	size int
	head, tail *Node
}

type Node struct {
	val interface{}
	next *Node
}

func New() *Queue {
	return &Queue{0, nil, nil}
}

func (queue *Queue) Enqueue(item interface{}) {
	newNode := &Node{item, nil}
	if queue.size == 0 {
		queue.head = newNode
		queue.tail = newNode
	} else {
		queue.tail.next = newNode
		queue.tail = newNode
	}
	queue.size++
}

func (queue *Queue) Dequeue() interface{} {
	var item interface{} = nil
	if queue.size > 0 {
		item = queue.head.val
		old := queue.head
		queue.head = queue.head.next
		old.next = nil
		queue.size--
	}
	return item
}

func (queue *Queue) Empty() bool {
	return queue.size == 0
}

func (queue *Queue) Size() int {
	return queue.size
}

func (queue *Queue) Peek() interface{} {
	var item interface{} = nil
	if queue.size > 0 {
		item = queue.head.val
	}
	return item
}
