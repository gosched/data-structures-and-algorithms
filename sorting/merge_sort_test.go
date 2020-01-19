package sorting

import (
	"fmt"
	"testing"
)

// TestMergeSort .
func TestMergeSort(t *testing.T) {
	s := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	result := MergeSort(s)
	fmt.Println(result)
}

