package heap

import "fmt"

// 建堆
func (heap *Heap) buildHeap(n int, nums []int) {
	if n > heap.capacity {
		fmt.Println("给定数组大于堆的容量，建堆失败")
		return
	}
	for i := 1; i <= n; i++ {
		heap.a[i] = nums[i-1]
	}
	heap.count = n
	for i := n/2; i >= 1; i-- {
		heap.heapify(i)
	}
}

func (heap *Heap) heapify(i int) {
	index := i
	for index < heap.count {
		temp := index
		if index*2 <= heap.count && heap.a[index*2] > heap.a[index] {
			temp = index*2
		}
		if index*2+1 <= heap.count && heap.a[index*2+1] > heap.a[temp] {
			temp = index*2+1
		}
		if temp == index {
			break
		}
		heap.a[temp], heap.a[index] = heap.a[index], heap.a[temp]
		index = temp
	}
}

func(heap *Heap) heapSort(n int, nums []int) []int {
	heap.buildHeap(n, nums)

	res := make([]int, heap.count)
	i := heap.count
	for i >= 1 {
		res[i-1] = heap.a[1]
		heap.a[i], heap.a[1] = heap.a[1], heap.a[i]
		i --
		heap.count --
		heap.heapify(1)
	}
	return res
}