// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

var (
	res = []int{}
)

// 返回第k个最小的节点
func KthNode(pRoot *TreeNode, k int) *TreeNode {
	inOrderRetraval(pRoot)
	if k < 1 && k > len(res) {
		return nil
	}

	return &TreeNode{res[k-1], nil, nil}
}

func inOrderRetraval(pRoot *TreeNode) {
	if pRoot == nil {
		return
	}

	inOrderRetraval(pRoot.left)
	res = append(res, pRoot.val)
	inOrderRetraval(pRoot.right)
}

func main() {
	pRoot := &TreeNode{
		5,
		&TreeNode{
			3,
			&TreeNode{
				2,
				nil,
				nil,
			},
			&TreeNode{
				4,
				nil,
				nil,
			},
		},
		&TreeNode{
			7,
			nil,
			nil,
		},
	}

	fmt.Println(KthNode(pRoot, 2))
}
