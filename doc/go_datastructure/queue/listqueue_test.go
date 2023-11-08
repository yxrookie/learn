package queue

import "testing"

func TestListqueue(t *testing.T) {
	lq := NewLqueue()
	lq.Push(1)
	lq.PrintQueue()
	lq.Push(12)
	lq.PrintQueue()
	lq.Push(8)
	lq.PrintQueue()
	lq.Pop()
	lq.PrintQueue()
	lq.Pop()
	lq.PrintQueue()
	lq.Pop()
	lq.PrintQueue()
}