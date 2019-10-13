// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ex28

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	return createBST(1, n)
}

func createBST(start int, end int) []*TreeNode {
	var res []*TreeNode
	if start > end {
		res = append(res, nil)
		return res
	}

	for i := start; i < end+1; i++ {
		left := createBST(start, i-1) // left child bST
		right := createBST(i+1, end)  // right child bst
		lsize := len(left)
		rsize := len(right)
		for j := 0; j < lsize; j++ {
			for k := 0; k < rsize; k++ {
				root := &TreeNode{i, nil, nil}
				root.Left = left[j]
				root.Right = right[k]
				res = append(res, root)
			}
		}
	}
	return res
}

// 给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
//
// 示例:

/*
输入: 3
输出:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
解释:
以上的输出对应以下 5 种不同结构的二叉搜索树：

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-binary-search-trees-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
