// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package ex61

/**
 * Definition for a binary tree node. */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	res      = make([]int, 0)
	maxLevel = 0
)

func rightSideView(root *TreeNode) []int {
	res = make([]int, 0)
	maxLevel = 0
	dfs(1, root)
	return res
}

func dfs(level int, root *TreeNode) {
	if root == nil {
		return
	}
	// 主要逻辑
	if level > maxLevel {
		res = append(res, root.Val)
		maxLevel = level
	}

	if root.Right != nil {
		dfs(level+1, root.Right)
	}

	if root.Left != nil {
		dfs(level+1, root.Left)
	}

}
