// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package ex58

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := 1<<32 - 1
	walk(root, &maxSum)
	return maxSum
}

func walk(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	leftSum := walk(root.Left, maxSum)
	rightSum := walk(root.Right, maxSum)

	value1 := root.Val
	value2 := leftSum + root.Val
	value3 := rightSum + root.Val
	value4 := leftSum + rightSum + root.Val

	*maxSum = max(value1, value2, value3, value4)

	return max(value1, value2, value3)
}

func max(a1 int, as ...int) int {
	for _, val := range as {
		if val > a1 {
			a1 = val
		}
	}
	return a1
}
