package config

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func CreateTree(root *Node, val int) *Node {
	if root == nil {
		return nil
	}
	if val == root.Val {
		root.Left = CreateTree(root.Left, val)

	} else {
		root.Right = CreateTree(root.Right, val)
	}
	return root
}

// (left, root, right)
func InOrderTraversal(root *Node) {
	if root != nil {
		InOrderTraversal(root.Left)
		fmt.Print(root.Val, " ")
		InOrderTraversal(root.Right)
	}
}

// (root, left, right)
func PreOrderTraversal(root *Node) {
	if root != nil {
		fmt.Print(root.Val, " ")
		PreOrderTraversal(root.Left)
		PreOrderTraversal(root.Right)
	}
}

// (left, right, root)
func PostOrderTraversal(root *Node) {
	if root != nil {
		PostOrderTraversal(root.Left)
		PostOrderTraversal(root.Right)
		fmt.Print(root.Val, " ")
	}
}
