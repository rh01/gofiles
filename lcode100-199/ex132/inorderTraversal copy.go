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
	stack := make([]*TreeNode, 0)
	res = make([]int, 0)

	for len(stack) >= 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			if len(stack) > 0 {
				root = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				res = append(res, root.Val)
				root = root.Right
			}
		}
	}

	return res
}

// 中序遍历
// 左根右
// 右根左

// 给定一个二叉树，返回它的中序 遍历。

// 示例:

// 输入: [1,null,2,3]
//    1
//  /  \
// 0    2
//     / \
//    3   4

// 输出: [1,3,2]
// 4 2 3 1
// 2 1 0
// 4 2 3

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
