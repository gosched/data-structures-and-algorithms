package unique

import (
	"sort"
)

// RemoveDuplicates .
// Filter for unique values and remove duplicates
func RemoveDuplicates(data []int) []int {
	var result = []int{}

	for _, v := range data {
		var ok = true
		for _, e := range result {
			if v == e {
				ok = false
			}
		}
		if ok {
			result = append(result, v)
		}
	}

	return result
}

// RemoveDuplicates2 .
func RemoveDuplicates2(elements []int) []int {
	encountered := map[int]bool{}
	result := []int{}

	for _, v := range elements {
		if encountered[v] == true {
			continue
		}
		encountered[v] = true
		result = append(result, v)
	}

	return result
}

// RemoveDuplicates3 .
func RemoveDuplicates3(data []int) []int {
	if len(data) < 2 {
		return data
	}

	sort.Ints(data)

	var result = []int{data[0]}

	for i := 1; i < len(data); i++ {
		if data[i] != data[i-1] {
			result = append(result, data[i])
		}
	}

	return result
}
