package heap

import (
	"fmt"
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


	bheap := &Heap{}
	bheap.initHeap(10)
	nums := []int{1, 34, 25, 7, 18}
	res := bheap.heapSort(5, nums)
	fmt.Println(res)
	
	cheap := &Heap{}
	cheap.initHeap(10)
	nums2 := []int{101, 99, 87, 65, 18}
	res2 := cheap.heapSort(5, nums2)
	fmt.Println(res2)
}