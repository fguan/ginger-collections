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

// Package binary_tree implements a binary search tree
package binary_tree

import (
	"fmt"
)

type BinaryTree struct {
	Val int
	Left, Right *BinaryTree
}

func New(val int) *BinaryTree {
	return &BinaryTree{val, nil, nil}
}

// Return true if value is present in the true; false otherwise
func Search(tree *BinaryTree, val int) bool {
	if tree == nil {
		return false
	} else if val < tree.Val {
		Search(tree.Left, val)
	} else if val > tree.Val {
		Search(tree.Right, val)
	} 
	return true

}

// Insert value into the tree, duplicates allowed
func Insert(tree *BinaryTree, val int) {
	if tree == nil {
		tree = &BinaryTree{val, nil, nil}
	} else if val <= tree.Val {
		if tree.Left == nil {
			tree.Left = &BinaryTree{val, nil, nil}
		} else {
			Insert(tree.Left, val)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BinaryTree{val, nil, nil}
		} else {
			Insert(tree.Right, val)
		}
	}
}

// Return the node containing the minValue
func (tree *BinaryTree) minValue() *BinaryTree {
	if tree == nil {
		return nil
	} else {
		var currNode = tree
		for tree.Left != nil {
			currNode = tree.Left
		}
		return currNode
	}
}

// Print the tree in-order
func PrintInOrder(tree *BinaryTree) {
	if tree == nil {
		return
	} else {
		PrintInOrder(tree.Left)
		fmt.Printf("%d ", tree.Val)
		PrintInOrder(tree.Right)
	}
}

// Print the tree pre-order
func PrintPreOrder(tree *BinaryTree) {
	if tree == nil {
		return
	} else {
		fmt.Printf("%d ", tree.Val)
		PrintPreOrder(tree.Left)
		PrintPreOrder(tree.Right)
	}
}

// Print the tree post-order
func PrintPostOrder(tree *BinaryTree) {
	if tree == nil {
		return
	} else {
		PrintPostOrder(tree.Left)
		PrintPostOrder(tree.Right)
		fmt.Printf("%d ", tree.Val)
	}
}
