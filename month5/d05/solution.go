/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  二叉搜索树， 中序遍历为有序
func isValidBST(root *TreeNode) bool {

	orderList := []int{}
	var inOrderRetrave func(*TreeNode)
	inOrderRetrave = func(node *TreeNode) {
		if node == nil {
			return
		}

		inOrderRetrave(node.Left)
		orderList = append(orderList, root.Val)
		inOrderRetrave(node.Right)
	}
	inOrderRetrave(root)
	
	for i := 1; i < len(orderList); i++ {
		if orderList[i] < orderList[i-1] {
			return false
		}
	}
	return true
}

