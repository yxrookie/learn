package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	var queue Queue
	queue = queue.InitQueue(4)
	queue.PushQueue(1)
	queue.PrintQueue()
	queue.PushQueue(12)
	queue.PrintQueue()
	queue.PushQueue(31)
	queue.PrintQueue()
	queue.PushQueue(120)
	queue.PrintQueue()
	queue.PushQueue(-1)
	queue.PrintQueue()
	queue.PopQueue()
	queue.PrintQueue()
	queue.PopQueue()
	queue.PrintQueue()
	queue.PopQueue()
	queue.PrintQueue()
}