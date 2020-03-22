package sorting

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"time"
)

func data(length int) []int {
	return rand.Perm(length)
}

func shuffle(data []int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range data {
		randIndex := r.Intn(len(data))
		data[i], data[randIndex] = data[randIndex], data[i]
	}
}

func test(data []int, funcs ...func([]int)) {
	for _, f := range funcs {
		fmt.Println(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
		fmt.Println(data)
		f(data)
		fmt.Println(data)
		shuffle(data)
		fmt.Println(data)
		println()
	}
}

func bubble(data []int) {
	length := len(data)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func insertion(data []int) {
	length := len(data)
	for i := 1; i < length; i++ {
		for j := i; j > 0 && data[j-1] > data[j]; j-- {
			data[j-1], data[j] = data[j], data[j-1]
		}
	}
}

func insertion2(data []int) {
	length := len(data)
	for i := 1; i < length; i++ {
		temp := data[i]
		j := i - 1
		for j >= 0 && data[j] > temp {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

func shell(data []int) {

}

func selection(data []int) {
	length := len(data)
	for i := 0; i < length; i++ {
		minIndex, change := i, false
		for j := i + 1; j < length; j++ {
			if data[j] < data[minIndex] {
				minIndex, change = j, true
			}
		}
		if change {
			data[i], data[minIndex] = data[minIndex], data[i]
		}
	}
}

func heap(data []int) {

}

// ShowMapByKeys .
func ShowMapByKeys(m map[int]int) {
	m[0] = 123123
	m[1] = 666
	m[2] = 101

	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		print(m[k], " ")
	}

	println()
}

// ShowMapByValues .
func ShowMapByValues(m map[int]int) {

}
