package tree

import (
	"fmt"
)

// Traversal

// Depth-first search
// pre, in, post

// Breadth-first search
// level

// PreorderRecursive .
func PreorderRecursive(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%v ", node.Value)
	PreorderRecursive(node.Left)
	PreorderRecursive(node.Right)
}

// InorderRecursive .
func InorderRecursive(node *Node) {
	if node == nil {
		return
	}
	InorderRecursive(node.Left)
	fmt.Printf("%v ", node.Value)
	InorderRecursive(node.Right)
}

// PostorderRecursive .
func PostorderRecursive(node *Node) {
	if node == nil {
		return
	}
	PostorderRecursive(node.Left)
	PostorderRecursive(node.Right)
	fmt.Printf("%v ", node.Value)
}

// LevelOrderRecursive .
func LevelOrderRecursive(root *Node) [][]*Node {
	var result = [][]*Node{}

	var dfs func(level int, node *Node)
	dfs = func (level int, node *Node) {
		if node == nil {
			return
		}
	
		if level >= len(result) {
			result = append(result, []*Node{})
		}
	
		result[level] = append(result[level], node)
		
		dfs(level+1, node.Left)
		dfs(level+1, node.Right)
	}

	dfs(0, root)
	return result
}

