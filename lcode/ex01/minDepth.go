// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex01

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// var (
// 	level = 1
// )

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return depth(root, 1)
}

func depth(cRoot *TreeNode, level int) int {
	if cRoot.Left == nil && cRoot.Right == nil {
		return level
	} else if cRoot.Left == nil && cRoot.Right != nil {
		return depth(cRoot.Right, level+1)
	} else if cRoot.Left != nil && cRoot.Right == nil {
		return depth(cRoot.Right, level+1)
	} else {
		return min(depth(cRoot.Left, level+1), depth(cRoot.Right, level+1))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
