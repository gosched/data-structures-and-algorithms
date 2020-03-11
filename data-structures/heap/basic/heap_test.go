package heap

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func Test(t *testing.T) {
	items := rand.Perm(15)

	fmt.Println(items)

	var h = &MinHeap{}

	h.Init()

	for _, v := range items {
		h.Push(v)
	}

	fmt.Println(h.items)

	for h.Len() > 0 {
		fmt.Printf("%v ", h.Top())
		h.Pop()
	}
	fmt.Println()

	h.Clear()
	fmt.Println(h.items)

	h.InitializeFrom(items)
	fmt.Println(h.items)

	sort.Ints(h.items)
	fmt.Println(h.items)
}

// [0 5 1 10 7 2 3 13 11 12 8 4 9 14 6]
// 0
// 5          1
// 10    7    2   3
// 13 11 12 8 4 9 14 6

// [0 4 1 7 5 2 3 11 13 10 8 12 9 14 6]
// 0
// 4          1
// 7     5    2    3
// 11 13 10 8 12 9 14 6
