package heap

import "fmt"

// 删除大顶堆的堆顶元素
func (heap *Heap) removeMax() {
	if heap.count == 0 {
		fmt.Println("堆为空，无法删除堆顶元素")
		return
	}
	heap.a[1], heap.a[heap.count] = heap.a[heap.count], heap.a[1]
	heap.count --
	i := 1
	for i < heap.count {
		temp := i
		if i*2 <= heap.count && heap.a[i*2] > heap.a[i] {
			temp = i*2
		}
		if i*2+1 <= heap.count && heap.a[i*2+1] > heap.a[temp] {
			temp = i*2+1
		}
        if temp == i {
			break
		}
		heap.a[temp], heap.a[i] = heap.a[i], heap.a[temp]
		i = temp
	}
}