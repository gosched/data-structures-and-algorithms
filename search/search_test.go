package algorithm

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	items1 := rand.Perm(20)
	items2 := rand.Perm(25)

	sort.Ints(items1)
	sort.Ints(items2)

	fmt.Println(items1)
	fmt.Println(items2)

	target1 := 11
	target2 := 18

	index1 := BinarySearch(items1, target1)
	index2 := BinarySearch(items2, target2)

	if index1 != -1 {
		fmt.Printf("index1:%v value1:%v\n", index1, items1[index1])
	}
	if index2 != -1 {
		fmt.Printf("index2:%v value2:%v\n", index2, items1[index2])
	}
}

func TestBinarySearch1(t *testing.T) {
	items1 := rand.Perm(20)
	items2 := rand.Perm(25)

	sort.Ints(items1)
	sort.Ints(items2)

	fmt.Println(items1)
	fmt.Println(items2)

	target1 := 11
	target2 := 18

	index1 := BinarySearch1(items1, target1)
	index2 := BinarySearch1(items2, target2)

	if index1 != -1 {
		fmt.Printf("index1:%v value1:%v\n", index1, items1[index1])
	}
	if index2 != -1 {
		fmt.Printf("index2:%v value2:%v\n", index2, items1[index2])
	}
}
