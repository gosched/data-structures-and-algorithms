package list

import "testing"

func Test(t *testing.T) {

	l1, l2 := &List{}, &List{}

	l1.New()
	l2.New()

	for i := 0; i < 10; i++ {
		l1.Append(i)
		l2.Append(i * i)
	}

	l1.Traversal()
	l2.Traversal()

	l1.Extend(l2.list)

	l1.Traversal()

}
