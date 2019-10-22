// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex30

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root==nil{
		return right==nil
	}

	return isChildSymmetric(root.Left, root.Right)
}

func isChildSymmetric(left, right *TreeNode) bool {
	if left == nil {
		return true
	}

	if right == nil {
		return false
	}

	return left.Val == right.Val && isChildSymmetric(left.Left, right.Right) && isChildSymmetric(left.Right, right.Left) 
}
