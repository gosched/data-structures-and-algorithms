package array

import "testing"

func Test(t *testing.T) {
	a := array(10)
	println(a.len(), a.cap())

	for i := 0; i < 30; i++ {
		a.insert(i)
		a.show()
	}
}
