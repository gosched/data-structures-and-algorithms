package tree

// Node .
type Node struct {
	Left  *Node
	Right *Node
	Key   int
	Value interface{}
}

// BinarySearchTree .
type BinarySearchTree struct {
	Root *Node
}

// Compare .
func (node *Node) Compare(node2 *Node) int {
	if node.Key > node2.Key {
		return 1
	}
	if node.Key < node2.Key {
		return -1
	}
	return 0
}

// Insert .
func (bst *BinarySearchTree) Insert(node *Node) {
	if node == nil {
		return
	}

	if bst.Root == nil {
		bst.Root = node
		return
	}
	
	temp := bst.Root
	for temp != nil {
		if temp.Compare(node) == 1 {
			if temp.Left != nil {
				temp = temp.Left
			} else {
				temp.Left = node
				break
			}
		} else {
			if temp.Right != nil {
				temp = temp.Right
			} else {
				temp.Right = node
				break
			}
		}
	}
}

// Traverse .
func (bst *BinarySearchTree) Traverse(order string, recursive bool) {
	if recursive {
		switch order {
		case "preorder":
			PreorderRecursive(bst.Root)
		case "inorder":
			InorderRecursive(bst.Root)
		case "postorder":
			PostorderRecursive(bst.Root)
		}
	} else {
		switch order {
		case "preorder":
			PreorderIterative(bst.Root)
		case "inorder":
			InorderIterative(bst.Root)
		case "postorder":
			PostorderIterative(bst.Root)
		case "levelOrder":
			LevelOrderIterative(bst.Root)
		}
	}
	println()
}

// Height .
func (bst *BinarySearchTree) Height() int {
	return Height(bst.Root)
}

// is balance

// balance