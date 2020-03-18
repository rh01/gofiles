package main

import "strings"

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

 func serializeTree(root *TreeNode) string {
	if root == nil {
		return ""
	}
	res := []string{}
    var cal func(*TreeNode) 
	cal = func(node *TreeNode) {
		if node == nil {
			res = append(res, "#")
			return 
		} 
		res =  append(res, string(node.Val))
		cal(node.Left)
		cal(node.Right)
		return 
	}
    
    cal(root)

	ret := strings.Join(res, "_")
	return ret
}

func movePrefixTable(prefix []int, n int) {
	for l := n -1; l > 0; l--{
		prefix[l] = prefix[l-1]
	}
	prefix[0] = -1
}


func prefixTable(pattern string, n int) []int {
	prefix := make([]int, n)
	prefix[0] = 0
	l := 0
	i := 1

	for i < n {
		if pattern[i] == pattern[l] {
			l++
			prefix[i] = l
			i++
		} else {
			if l > 0 {
				l = prefix[l-1]
			}else {
				prefix[i] = l 
				i++
			}
		}
	}

	movePrefixTable(prefix, n)
	return prefix
}

func kmpSearch(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)
	i, j := 0, 0
	
	prefix := prefixTable(s2, n)
	for i < m {
		if j == n - 1 && s1[i] == s2[j] {
			return true
		}
		if s1[i] == s2[j] {
			i++
			j++
		} else {
			j = prefix[j]
			if j == -1 {
				i++
				j++
			}
		}
	}

	return false
}

func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	s1 := serializeTree(t1)
	s2 := serializeTree(t2)

	return kmpSearch(s1, s2)
}

func main() {
	
}