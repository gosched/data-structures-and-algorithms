package tree

import (
	// "math/rand"
	"fmt"
	"testing"
)

/*
                        		99
				54							104
	17	    					83  	103		107
						61
					60		75
*/

// 99 54 17 83 61 60 75 104 103 107
// 17 54 60 61 75 83 99 103 104 107
// 17 60 75 61 83 54 103 107 104 99

var (
	data = []int{99, 54, 104, 17, 83, 103, 107, 61, 60, 75}
	// data = rand.Perm(1000001)
	tree = &BinarySearchTree{}
)

func init() {
	for _, v := range data {
		tree.Insert(&Node{Key: v, Value: v})
	}
}

func TestTraverse(t *testing.T) {
	tree.Traverse("preorder", false)
	tree.Traverse("inorder", false)
	tree.Traverse("postorder", false)
	tree.Traverse("levelOrder", false)
}

func TestLevelOrderTraversal(t *testing.T) {
	tree.Traverse("levelOrder", false)

	result := LevelOrderRecursive(tree.Root)
	for _, v := range result {
		for _, n := range v {
			fmt.Print(n.Key, " ")
		}
		fmt.Println()
	}
}

func TestGetHeight(t *testing.T) {
	height := tree.Height()
	println(height)
}

func BenchmarkPreorderTraversalRecursive(t *testing.B) {
	tree.Traverse("preorder", true)
}

func BenchmarkPreorderTraversalIterative(t *testing.B) {
	tree.Traverse("preorder", false)
}

func BenchmarkInorderTraversalRecursive(t *testing.B) {
	tree.Traverse("inorder", true)
}

func BenchmarkInorderTraversalIterative(t *testing.B) {
	tree.Traverse("inorder", false)
}
