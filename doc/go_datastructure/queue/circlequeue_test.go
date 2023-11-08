package queue

import "testing"

func TestCircleQueue(t *testing.T) {
	var cq Circlequeue
	cq = cq.initQueue(4)
	cq.Push(1)
	cq.PrintQueue()
	cq.Push(10)
	cq.PrintQueue()
	cq.Push(2)
	cq.PrintQueue()
	cq.Push(13)
	cq.PrintQueue()
	cq.Push(-1)
	cq.PrintQueue()
	cq.Pop()
	cq.PrintQueue()
	cq.Pop()
	cq.PrintQueue()
	cq.Pop()
	cq.PrintQueue()
	cq.Pop()
	cq.PrintQueue()
}