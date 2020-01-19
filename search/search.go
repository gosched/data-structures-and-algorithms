package algorithm

// BinarySearch .
// BinarySearch 僅適用於排序好的資料
func BinarySearch(items []int, target int) int {
	l, r := 0, len(items)-1
	for l <= r {
		mid := (l + r) / 2
		if items[mid] == target {
			return mid
		}
		if items[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

// BinarySearch1 .
func BinarySearch1(items []int, target int) int {
	l, r := 0, len(items)
	for l < r {
		mid := l + (r-l)/2
		if items[mid] == target {
			return mid
		}
		if items[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return -1
}
