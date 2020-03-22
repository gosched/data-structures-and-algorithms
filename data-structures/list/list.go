package list

import (
	"container/list"
	"fmt"
)

// List .
type List struct {
	list *list.List
}

// l.list.Back()
// l.list.Front()
// l.list.InsertAfter()
// l.list.InsertBefore()
// l.list.MoveAfter()
// l.list.MoveAfter()
// l.list.MoveToBack()
// l.list.MoveToFront()
// l.list.PushFront()

// New .
func (l *List) New() {
	l.list = list.New()
}

// Init .
func (l *List) Init() {
	l.list.Init()
}

// Len .
func (l *List) Len(other *list.List) int {
	return l.list.Len()
}

// Append .
func (l *List) Append(value interface{}) {
	l.list.PushBack(value)
	// e := l.list.PushBack(value)
}

// Insert .
func (l *List) Insert(index int, value interface{}) {

}

// Extend .
func (l *List) Extend(other *list.List) {
	l.list.PushBackList(other)
}

// FrontExtend .
func (l *List) FrontExtend(other *list.List) {
	l.list.PushFrontList(other)
}

// Remove .
func (l *List) Remove(e *list.Element) {
	l.list.Remove(e)
	// value := l.list.Remove(e)
	// return value
}

// Pop .
func (l *List) Pop() *list.Element {
	e := l.list.Back()
	return e
}

// Traversal .
func (l *List) Traversal() {
	for e := l.list.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	println()
}

// Iterate .
func (l *List) Iterate(f func() error) {
	for e := l.list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

// var stack []string

// stack = append(stack, "world!")
// stack = append(stack, "Hello ")

// for len(stack) > 0 {
//     n := len(stack) - 1
//     fmt.Print(stack[n])

//     stack = stack[:n]
// }
