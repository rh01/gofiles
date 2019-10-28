// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex132

/**
 * Definition for a binary tree node.
 *
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	res = make([]int, 0)
)

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res = make([]int, 0)
	walk(root, &res)
	return res
}

func walk(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	walk(root.Left, res)
	*res = append(*res, root.Val)
	walk(root.Right, res)
	return
}
