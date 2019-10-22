// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex31

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	// 初始化存放结果的切片
	res := make([][]int, 0)
	hqueue := make([]*TreeNode, 0)
	hqueue = append(hqueue, root)
	// 将root节点胡值加入到结果队列中
	res = append(res, []int{root.Val})
	// 保存临时的
	localqueue := make([]*TreeNode, 0)
	// 利用队列来做
	for len(hqueue) != 0 {
		// 当前节点
		cur := hqueue[0]
		hqueue = hqueue[1:]
		// 如果当前节点的左节点不为空，则加入到队列中
		if cur.Left != nil {
			localqueue = append(localqueue, cur.Left)
		}
		// 同理，若当前节点的右节点不为空，则加入到队列中
		if cur.Right != nil {
			localqueue = append(localqueue, cur.Right)
		}
		// 检查hqueue是否为空，若为空，则将临时的队列里面的元素
		// 加入到结果队列里面
		if len(hqueue) == 0 {
			data := make([]int, 0)
			for i := 0; i < len(localqueue); i++ {
				curNode := localqueue[i]
				curVal := curNode.Val
				hqueue = append(hqueue, curNode)
				data = append(data, curVal)
			}
			// append data to res
			if len(data) != 0 {
				res = append(res, data)
			}
			// make localqueue to nil
			localqueue = make([]*TreeNode, 0)
		}
	}
	// reverse(res)
	return res
}

func reverse(s [][]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
