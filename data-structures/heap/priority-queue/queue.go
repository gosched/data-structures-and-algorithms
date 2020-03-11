package pq

import "container/heap"

// PriorityQueue

// Node .
type Node struct {
	Val int
}

// Item .
type Item struct {
	value    *Node
	priority int
	index    int
}

// PriorityQueue .
type PriorityQueue []*Item

func (h PriorityQueue) Len() int           { return len(h) }
func (h PriorityQueue) Less(i, j int) bool { return h[i].priority < h[j].priority }
func (h PriorityQueue) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

// Push .
func (h *PriorityQueue) Push(x interface{}) {
	n := len(*h)
	item := &Item{
		value: x.(*Node),
		index: n,
	}
	item.priority = item.value.Val
	*h = append(*h, item)
}

// Pop .
func (h *PriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)

	item := old[n-1]
	item.index = -1

	*h = old[0 : n-1]

	return item
}

// Update .
func (h *PriorityQueue) Update(item *Item, value *Node, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(h, item.index)
}
