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

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return []int{}
	}
	res = append(res, root.Val)
	hqueue := make([]*TreeNode, 0)
	hqueue = append(hqueue, root)
	localqueue := make([]*TreeNode, 0)
	for len(hqueue) != 0 {
		cur := hqueue[len(hqueue)-1]
		hqueue = hqueue[:len(hqueue)-1]
		if cur.Left != nil {
			localqueue = append(localqueue, cur.Left)
		}

		if cur.Right != nil {
			localqueue = append(localqueue, cur.Right)
		}

		if len(hqueue) == 0 {
			if len(localqueue) != 0 {
				res = append(res, localqueue[len(localqueue)-1].Val)
			}
			for i := 0; i < len(localqueue); i++ {
				hqueue = append(hqueue, localqueue[i])
			}
			localqueue = make([]*TreeNode, 0)
		}
	}

	return res
}

// func dfs(root *TreeNode, res *[]int) {
// 	if root == nil {
// 		return
// 	}

// 	*res = append(*res, root.Val)
// 	if root.Left == nil && root.Right == nil {
// 		return
// 	}

// 	if root.Right != nil {
// 		dfs(root.Right, res)
// 	} else {
// 		dfs(root.Left, res)
// 	}
// }
