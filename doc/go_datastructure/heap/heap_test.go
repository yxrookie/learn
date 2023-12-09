package heap

import (
	"testing"
)

func TestHeap(t *testing.T) {
	aheap := &Heap{}
	aheap.initHeap(5)
	aheap.insrt(1)
	aheap.printHeap()
	aheap.insrt(2)
	aheap.printHeap()
	aheap.insrt(34)
	aheap.printHeap()
	aheap.insrt(16)
	aheap.printHeap()
	aheap.insrt(17)
	aheap.printHeap()
	aheap.insrt(5)
	aheap.printHeap() 

	aheap.removeMax()
	aheap.printHeap()
	aheap.removeMax()
	aheap.printHeap()
	aheap.removeMax()
	aheap.printHeap()
}