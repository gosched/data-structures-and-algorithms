package tree

import (
	"container/list"
	"fmt"
)

// VisitNode .
func VisitNode(node *Node) {
	if node != nil {
		fmt.Printf("%v ", node.Value)
	} else {
		fmt.Print("nil ")
	}
}

// Height .
// find height of a specific node
func Height(node *Node) int {
	if node == nil {
		return 0
	}
	
	height := 1
	leftHeight := Height(node.Left)
	rightHeight := Height(node.Right)
	
	if leftHeight > rightHeight {
		height += leftHeight
	} else {
		height += rightHeight
	}

	return height
}

// HeightIterative .
func HeightIterative(root *Node) int {
	if root == nil {
		return 0
	}

	level := 0

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		count := queue.Len()
		for count > 0 {
			temp := queue.Front()
			queue.Remove(temp)
	
			node, ok := temp.Value.(*Node)
			if ok {	
				if node.Left != nil {
					queue.PushBack(node.Left)
				}
		
				if node.Right != nil {
					queue.PushBack(node.Right)
				}	
			}

			count--
		}
		level++
	}
	return level
}
