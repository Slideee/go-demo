package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	set := make(map[int]struct{})
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if _, ok := set[k-node.Val]; ok {
			return true
		}
		set[node.Val] = struct{}{}
		return dfs(node.Left) || dfs(node.Right)
	}
	return dfs(root)
}
