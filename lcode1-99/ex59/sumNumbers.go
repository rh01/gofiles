// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/**
 * Definition for a binary tree node.
 */

package ex59

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	//curPathSumNumber := 0
	totalSumNumber := 0
	walk(root, 0, &totalSumNumber)
	return totalSumNumber
}

func walk(root *TreeNode, curPathSumNumber int, totalSumNumber *int) {
	if root == nil {
		return
	}

	curPathSumNumber = curPathSumNumber*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*totalSumNumber += curPathSumNumber
		return
	}

	walk(root.Left, curPathSumNumber, totalSumNumber)
	walk(root.Right, curPathSumNumber, totalSumNumber)
	return
}
