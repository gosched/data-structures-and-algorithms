package sorting

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestQuickSort .
func TestQuickSort(t *testing.T) {
	source := rand.Perm(20)
	target := rand.Perm(20)

	QuickSort(source)
	sort.Ints(target)

	// fmt.Println(source)
	// fmt.Println(target)

	assert.Equal(t, source, target, "they should be equal")

	source2 := rand.Perm(25)
	target2 := rand.Perm(25)

	QuickSort(source2)
	sort.Ints(target2)

	// fmt.Println(source2)
	// fmt.Println(target2)

	assert.Equal(t, source2, target2, "they should be equal")
}

func TestQuickSort2(t *testing.T) {
	source := []int{4, 5, 5, 2, 2, 3, 3, 3, 5, 5, 5, 1, 1, 4}
	target := []int{1, 1, 2, 2, 3, 3, 3, 4, 4, 5, 5, 5, 5, 5}

	QuickSort2(source)
	sort.Ints(target)

	// fmt.Println(source)
	// fmt.Println(target)

	assert.Equal(t, source, target, "they should be equal")

	source2 := []int{4, 5, 5, 2, 2, 3, 3, 3, 5, 5, 5, 1, 4}
	target2 := []int{1, 2, 2, 3, 3, 3, 4, 4, 5, 5, 5, 5, 5}

	QuickSort2(source2)
	sort.Ints(target2)

	// fmt.Println(source2)
	// fmt.Println(target2)

	assert.Equal(t, source2, target2, "they should be equal")
}

func TestQuickSort3(t *testing.T) {
	source := rand.Perm(20)
	target := rand.Perm(20)

	QuickSort3(source)
	sort.Ints(target)

	// fmt.Println(source)
	// fmt.Println(target)

	assert.Equal(t, source, target, "they should be equal")

	source2 := rand.Perm(25)
	target2 := rand.Perm(25)

	QuickSort3(source2)
	sort.Ints(target2)

	// fmt.Println(source2)
	// fmt.Println(target2)

	assert.Equal(t, source2, target2, "they should be equal")
}
