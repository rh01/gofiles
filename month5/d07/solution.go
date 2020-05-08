type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 第一种方法  传统的方法
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t != nil {
		return false
	} else if s == nil && t == nil {
		return true
	} else if s != nil && t == nil {
		return false
	} 

	return hasSubtree(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func hasSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t != nil {
		return false
	} else if s == nil && t == nil {
		return true
	} else if s != nil && t == nil {
		return false
	} 
		
	
	return s.Val == t.Val && hasSubtree(s.Left, t.Left) && hasSubtree(s.Right, t.Right)
}