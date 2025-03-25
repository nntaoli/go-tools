package linkedlist

import (
	"container/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	l := New[int]()
	l.PushBack(1)

	e2 := l.PushBack(2)
	l.PushFront(3)
	l.PushBack(4)

	t.Log(l.Front().V)
	t.Log(l.Front().Next().V)
	t.Log(l.Front().Next().Next().V)
	t.Log(l.Front().Next().Next().Next().V)

	l.Remove(e2)

	t.Log(l.Front().V)
	t.Log(l.Front().Next().V)

	assert.Equal(t, l.Front().Next().Next().V, 4)

	t.Log(l.Front().V, l.Back().V)
	t.Log(l.Back().Next())

	assert.Nil(t, l.Back().Next())
}

//BenchmarkNew-6   	13625920	        77.33 ns/op
func BenchmarkNew(b *testing.B) {
	l := New[int]()
	for i := 0; i < b.N; i++ {
		l.PushBack(i + 1)
	}
}

//BenchmarkList-6   	 8141215	       133.6 ns/op
func BenchmarkList(b *testing.B) {
	l := list.New()
	for i := 0; i < b.N; i++ {
		l.PushBack(i + 1)
	}

}
