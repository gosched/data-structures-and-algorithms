package sorting

// MergeSort .
func MergeSort(s []int) []int {
	result := mergeSort(s)
	return result
}

// worst case
// n log n
func mergeSort(s []int) []int {
	l := len(s)
	if l < 2 {
		return s
	}
	mid := l / 2
	return merge(mergeSort(s[:mid]), mergeSort(s[mid:]))
}

func merge(left, right []int) []int {
	lLen, rLen, i, j := len(left), len(right), 0, 0

	result := make([]int, lLen+rLen)

	for k := 0; k < lLen+rLen; k++ {
		if i < lLen && j >= rLen {
			result[k] = left[i]
			i++
		} else if j < rLen && i >= lLen {
			result[k] = right[j]
			j++
		} else if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
	}

	return result
}

func merge2(left, right []int) []int {
	lLen, rLen, i, j, k := len(left), len(right), 0, 0, 0

	result := make([]int, lLen+rLen)

	for i < lLen && j < rLen {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	for i < lLen {
		result[k] = left[i]
		i++
		k++
	}

	for j < rLen {
		result[k] = right[j]
		j++
		k++
	}

	return result
}
