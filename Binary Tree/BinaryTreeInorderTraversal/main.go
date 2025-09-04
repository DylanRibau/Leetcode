package main

import "fmt"

/*
Given the root of a binary tree, return the inorder traversal of its nodes' values.



Example 1:

Input: root = [1,null,2,3]

Output: [1,3,2]

Explanation:

Example 2:

Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]

Output: [4,2,6,5,7,1,3,9,8]

Explanation:

Example 3:

Input: root = []

Output: []

Example 4:

Input: root = [1]

Output: [1]



Constraints:

    The number of nodes in the tree is in the range [0, 100].
    -100 <= Node.val <= 100

*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func main() {
	oneNode3 := Node{Val: 3}
	oneNode2 := Node{Val: 2, Left: &oneNode3}
	oneNode1 := Node{Val: 1, Right: &oneNode2}

	fmt.Printf("Calling inorderTraversal: %v", inorderTraversal(&oneNode1))
}

func inorderTraversal(root *Node) []int {
	return search(root, make([]int, 0))
}

func search(node *Node, result []int) []int {
	if node == nil {
		return result
	}
	if node.Left != nil {
		result = search(node.Left, result)
	}
	result = append(result, node.Val)
	if node.Right != nil {
		result = search(node.Right, result)
	}
	return result
}
