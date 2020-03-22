package sorting

import "testing"

func Test(t *testing.T) {
	data := data(40)
	test(data, bubble, insertion, shell, selection, heap)
}
