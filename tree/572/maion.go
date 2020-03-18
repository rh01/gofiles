/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
 /**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasSubtree(s *TreeNode, t *TreeNode) bool {

    if s == nil && t == nil {
        return true
    }
    
    return s!=nil && t != nil && s.Val == t.Val && hasSubtree(s.Left, t.Left) && hasSubtree(s.Right, t.Right)
		
	
	// return 
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
    if s == nil && t == nil {
        return true
    }
    
    if s == nil && t != nil {
        return false
    }
	return hasSubtree(s, t) || isSubtree(s.Left, t) ||  isSubtree(s.Right, t) 
}