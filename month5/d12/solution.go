/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 /**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	nodes := []*TreeNode{}
	result := [][]int{}
	result = append(result, []int{root.Val})
	nodes = append(nodes, root)
	
	tmpNodes := []*TreeNode{}

	for len(nodes) != 0 {
		node := nodes[0]
		nodes = nodes[1:]
		
		if node.Left != nil {
			tmpNodes = append(tmpNodes, node.Left)
		}

		if node.Right != nil {
			tmpNodes = append(tmpNodes, node.Right)
		}

		if len(nodes) == 0 {
			// copy(nodes, tmpNodes)
			tmpValues := []int{}
			for i := range tmpNodes {
				n := tmpNodes[i]
				nodes = append(nodes, n)
				tmpValues = append(tmpValues, n.Val)
			}
			tmpNodes = []*TreeNode{}
			if len(tmpValues) != 0 {
				result = append(result, tmpValues)
			}
		}
	}
	return result
}