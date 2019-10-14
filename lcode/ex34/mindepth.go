// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package ex34

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return walk(root, 1)
}

func walk(root *TreeNode, level int) int {
	if root.Left == nil && root.Right == nil {
		return level
	} else if root.Left == nil && root.Right != nil {
		return walk(root.Right, level+1)
	} else if root.Left != nil && root.Right == nil {
		return walk(root.Left, level+1)
	} else {
		left := walk(root.Left, level +1)
		right := walk(root.Right, level+1)
		return min(left, right)
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
