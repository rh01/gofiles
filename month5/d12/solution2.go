/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func levelOrder(root *TreeNode) [][]int {
	// nil situation
	if root == nil {
		return nil
	}

	// next initialize some contianer
	res := [][]int{}
	hqueue := make([]*TreeNode, 0)
	res = append(res, []int{root.Val})
	hqueue = append(hqueue, root)

	// abort situation
	for len(hqueue) != 0 {
		// initialize tmp nodes and vals
		tmpNodes := []*TreeNode{}
		tmpValues := []int{}
		for len(hqueue) != 0 {
			node := hqueue[0]
			// pop
			hqueue = hqueue[1:]

			if node.Left != nil {
				tmpNodes = append(tmpNodes, node.Left)
				tmpValues = append(tmpValues, node.Left.Val)
			}

			if node.Right != nil {
				tmpNodes = append(tmpNodes, node.Right)
				tmpValues = append(tmpValues, node.Right.Val)
			}
		}
		hqueue = append(hqueue, tmpNodes...)
		if len(tmpValues) != 0 {
			res = append(res, tmpValues)
		}
	}
	return res
}