// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex02

// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res = make([][]int, 0)

	if root == nil {
		return res
	}

	var q = []*TreeNode{}
	q = append(q, root)
	res = append(res, []int{root.Val})
	local := []*TreeNode{}
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]

		if cur.Left != nil {
			local = append(local, cur.Left)
		}

		if cur.Right != nil {
			local = append(local, cur.Right)
		}

		if len(q) == 0 {
			localData := make([]int, 0)
			for _, value := range local {
				localData = append(localData, value.Val)
				q = append(q, value)
			}
			res = append(res, localData)
			local = []*TreeNode{}
		}
	}
	return res[:len(res)-1]
}

type NodeLevel struct {
	Node  *TreeNode
	Level int
}

func levelOrderv2(root *TreeNode) [][]int {
	result := [][]int{}
	level := 0
	// Use a queue as FIFO access
	queue := []NodeLevel{}

	if root == nil {
		return result
	}

	queue = append(queue, NodeLevel{Node: root, Level: level})
	result = append(result, []int{})
	for len(queue) > 0 {
		// Grab first element
		curr, rest := queue[0], queue[1:]
		queue = rest
		if len(result)-1 < curr.Level {
			result = append(result, []int{})
		}
		result[curr.Level] = append(result[curr.Level], curr.Node.Val)
		// Push left and right children into the queue
		if curr.Node.Left != nil {
			queue = append(queue, NodeLevel{Node: curr.Node.Left, Level: curr.Level + 1})
		}
		if curr.Node.Right != nil {
			queue = append(queue, NodeLevel{Node: curr.Node.Right, Level: curr.Level + 1})
		}
	}

	return result
}
