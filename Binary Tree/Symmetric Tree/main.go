package main

import "fmt"

func main() {
	// node3l := TreeNode{Val: 3}
	// node4l := TreeNode{Val: 4}
	// node4r := TreeNode{Val: 4}
	// node3r := TreeNode{Val: 3}
	// node2l := TreeNode{Val: 2, Left: &node3l, Right: &node4l}
	// node2r := TreeNode{Val: 2, Left: &node4r, Right: &node3r}
	// root1 := TreeNode{Val: 1, Left: &node2l, Right: &node2r}

	node3l := TreeNode{Val: 3}
	node3r := TreeNode{Val: 3}
	node2l := TreeNode{Val: 2, Right: &node3l}
	node2r := TreeNode{Val: 2, Right: &node3r}
	root1 := TreeNode{Val: 1, Left: &node2l, Right: &node2r}

	fmt.Printf("isSymmetric? %v", isSymmetric(&root1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return search(root.Left, root.Right)
}

func search(lnode *TreeNode, rnode *TreeNode) bool {
	if lnode == nil && rnode == nil {
		return true
	}
	if !compare(lnode, rnode) {
		return false
	}
	result := search(lnode.Left, rnode.Right)
	if !result {
		return result
	}
	result = search(lnode.Right, rnode.Left)
	return result
}

func compare(node1 *TreeNode, node2 *TreeNode) bool {
	if (node1 == nil && node2 != nil) || (node2 == nil && node1 != nil) {
		return false
	}
	return node1.Val == node2.Val
}
