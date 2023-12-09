package heap

import "fmt"
type Heap struct {
	a []int 
	capacity int
	count int
}

func (heap *Heap) initHeap(n int){
	heap.a = make([]int, n+1)
	heap.capacity = n
	heap.count = 0
}

func (heap *Heap) insrt(data int) {
	if heap.count >= heap.capacity {
		fmt.Println("堆容量已满")
		return
	}
	heap.count++
	heap.a[heap.count] = data
	i := heap.count
	// 在切片最后一个位置插入元素后，自下而上进行堆化
	for i/2 > 0 && heap.a[i/2] < heap.a[i]{
		heap.a[i/2], heap.a[i] = heap.a[i], heap.a[i/2]
		i = i/2
	}
}

func (heap *Heap) printHeap() {
	for i := 1; i <= heap.count; i++ {
		fmt.Print(heap.a[i], " ")
	}
	fmt.Println()
}