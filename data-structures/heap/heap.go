package heap

// Heap Sort
// Best: O(n log(n))
// Average: O(n log(n))
// Worst: O(n log(n))
// Space: O(1)}

// parent  = (child - 1) / 2
// child 1 = parent * 2 + 1
// child 2 = parent * 2 + 2

// https://golang.org/src/container/heap/heap.go

// MinHeap .
type MinHeap struct {
	items []int
}

// Init .
func (h *MinHeap) Init() {
	h.items = []int{}
}

// InitializeFrom .
// 一次性新增多個元素
func (h *MinHeap) InitializeFrom(items []int) {
	for _, v := range items {
		h.items = append(h.items, v)
	}

	n := len(h.items)
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
}

// Push .
func (h *MinHeap) Push(item int) {
	h.items = append(h.items, item)
	h.up(h.Len() - 1)
}

// Top .
func (h *MinHeap) Top() int {
	return h.items[0]
}

// Clear .
func (h *MinHeap) Clear() {
	h.items = []int{}
}

// Pop .
// 小的元素索引較小
// 最小的元素索引為零
// 把最小的元素跟尾端的元素互換
// down
// 備份尾端的最小元素
// 移除處於尾端的最小元素
// 回傳最小元素
func (h *MinHeap) Pop() int {
	n := h.Len()

	index := n - 1
	h.Swap(0, index)
	h.down(0, index)

	item := h.items[n-1]
	h.items = h.items[:n-1]
	return item
}

// Len .
func (h *MinHeap) Len() int {
	return len(h.items)
}

// Less .
func (h *MinHeap) Less(i, j int) bool {
	return h.items[i] < h.items[j]
}

// Swap .
func (h *MinHeap) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

// Empty .
func (h *MinHeap) Empty() bool {
	return len(h.items) == 0
}

func (h *MinHeap) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.items[i], h.items[j] = h.items[j], h.items[i]
		j = i
	}
}

func (h *MinHeap) down(p, n int) bool {
	i := p // parent
	for {
		j1 := 2*i + 1          // left child
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j // next parent
	}
	return i > p
}
