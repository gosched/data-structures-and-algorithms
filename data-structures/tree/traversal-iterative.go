package tree

import (
	"container/list"
)

// PreorderIterative .
func PreorderIterative(root *Node) {
	if root == nil {
		return
	}

	stack, temp := list.New(), root

	stack.PushBack(temp)

	for stack.Len() > 0 {
		element := stack.Back()
		stack.Remove(element)
		temp = element.Value.(*Node)
		VisitNode(temp)
		if temp.Right != nil {
			stack.PushBack(temp.Right)
		}
		if temp.Left != nil {
			stack.PushBack(temp.Left)
		}
	}
}

// InorderIterative .
func InorderIterative(root *Node) {
	if root == nil {
		return
	}

	stack, mark := list.New(), root
	for stack.Len() > 0 || mark != nil {
		for mark != nil {
			stack.PushBack(mark)
			mark = mark.Left
		}
		element := stack.Back()
		stack.Remove(element)

		mark = element.Value.(*Node)

		VisitNode(mark)

		mark = mark.Right
	}
}

// PostorderIterative .
func PostorderIterative(root *Node) {
	if root == nil {
		return
	}

	stack, temp := list.New(), root

	var lastViewed *Node

	for stack.Len() > 0 || temp != nil {
		if temp != nil {
			stack.PushBack(temp)
			temp = temp.Left
		} else {
			element := stack.Back()
			stack.Remove(element)
			node := element.Value.(*Node)

			// right node not exist, or right node exist but visted
			if node.Right == nil || node.Right == lastViewed {
				VisitNode(node)
				lastViewed = node
			} else {
				stack.PushBack(node)
				temp = node.Right
			}
		}
	}
}

// LevelOrderIterative .
func LevelOrderIterative(root *Node) {
	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		count := queue.Len()
		for count > 0 {
			temp := queue.Front()
			queue.Remove(temp)
	
			node, ok := temp.Value.(*Node)
			if ok {
				VisitNode(node)
	
				if node.Left != nil {
					queue.PushBack(node.Left)
				}
		
				if node.Right != nil {
					queue.PushBack(node.Right)
				}	
			}

			count--
		}
		println()
	}
}

// specify level return number of nodes
