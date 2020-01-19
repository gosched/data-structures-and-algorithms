package sorting

// QuickSort .
// Lomuto partition scheme
func QuickSort(items []int) {
	if len(items) < 2 {
		return
	}

	i := partition(items)
	QuickSort(items[:i])
	QuickSort(items[i+1:])
}

func partition(items []int) int {
	high := len(items) - 1
	pivot := items[high]
	i := 0
	for j := 0; j < high; j++ {
		if items[j] < pivot {
			items[i], items[j] = items[j], items[i]
			i++
		}
	}
	items[i], items[high] = items[high], items[i]
	return i
}

// QuickSort2 .
// Repeated elements
// Dutch national flag problem
func QuickSort2(items []int) {
	quickSort2(items, 0, len(items)-1)
}

func quickSort2(items []int, low, high int) {
	if low < high {
		left, right := partition2(items, low, high)
		quickSort2(items, low, left-1)
		quickSort2(items, right+1, high)
	}
}

func partition2(items []int, low, high int) (int, int) {
	pivot := items[high]

	i := low
	j := low
	k := high

	for j <= k {
		if items[j] < pivot {
			items[j], items[i] = items[i], items[j]
			i++
			j++
		} else if items[j] > pivot {
			items[j], items[k] = items[k], items[j]
			k--
		} else {
			j++
		}
	}

	return i, k
}

// QuickSort3 .
// Hoare partition scheme
func QuickSort3(items []int) {
	quickSort3(items, 0, len(items)-1)
}

func quickSort3(items []int, low, high int) {
	if low < high {
		p := partition3(items, low, high)
		quickSort3(items, low, p)
		quickSort3(items, p+1, high)
	}
}

func partition3(items []int, low, high int) int {
	pivot := items[low+(high-low)/2]

	i, j := low, high

	for {
		for items[i] < pivot {
			i++
		}

		for items[j] > pivot {
			j--
		}

		if i >= j {
			return j
		}

		items[i], items[j] = items[j], items[i]
	}
}
