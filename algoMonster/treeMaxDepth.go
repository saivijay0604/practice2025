package algomonster

import (
	"leetcode/practice2025/config"
)

//Max depth of a binary tree is the longest root-to-leaf path.
//Given a binary tree, find its max depth. Here, we define the length of the path to be the number of edges on that path, not the number of nodes.

func TreeMaxDepth(root *config.Node) int {
	if root == nil {
		return 0
	}
	leftDepth := TreeMaxDepth(root.Left)
	rightDepth := TreeMaxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}

	return rightDepth + 1
}
