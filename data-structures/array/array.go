package array

import (
	"fmt"
)

// DynamicArray .
type DynamicArray struct {
	capacity int
	size     int
	array    []int
}

func array(capacity int) *DynamicArray {
	return &DynamicArray{
		array:    make([]int, capacity),
		capacity: capacity,
	}
}

func (a *DynamicArray) len() int {
	return a.size
}

func (a *DynamicArray) cap() int {
	return a.capacity
}

func (a *DynamicArray) show() {
	for i := 0; i < a.size; i++ {
		print(a.array[i], " ")
	}
	println()
}

func (a *DynamicArray) get(index int) int {
	if index >= a.size {
		s := fmt.Sprintf("invalid array index %v (out of bounds for %v-element array)\n", index, a.size)
		panic(s)
	}
	return a.array[index]
}

func (a *DynamicArray) index(e int) int {
	for i, v := range a.array {
		if v == e {
			return i
		}
	}
	return -1
}

func (a *DynamicArray) insert(e int) {
	if a.size == a.capacity {
		a.capacity *= 2
		var temp = make([]int, a.capacity)
		for i := range a.array {
			temp[i] = a.array[i]
		}
		a.array = temp
	}
	a.array[a.size] = e
	a.size++
}

// func (a *DynamicArray) insertAt(index, e int) {
// 	if a.size == a.capacity {

// 	}
// 	a.size++
// }

// insert at beginning
// delete at beginning
// insert at end
// delete at end

// func (a *DynamicArray) remove(e int) {
// 	a.size--
// }

// func (a *DynamicArray) pop(index int) {
// 	a.size--
// }
