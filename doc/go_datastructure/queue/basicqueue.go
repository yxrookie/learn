package queue

import (
	"errors"
	"fmt"
)

type Queue struct {
	items    []int
	head     int
	tail     int
	capacity int
}

func (q *Queue) InitQueue(capacity int) Queue {
	return Queue{
		items:    make([]int, capacity),
		capacity: capacity,
	}
}

func (q *Queue) PushQueue(num int) error {
	if q.tail == q.capacity {
		return errors.New("队列已满，无法添加新的元素")
	}
	// 判断当 tail 到达切片末尾时，能否将 head~tail 内容整体向前挪动
	if q.tail == q.capacity {
		temp := q.tail-q.head
		for i := 0; i < q.tail-q.head; i++ {
			q.items[i] = q.items[i+q.head]
		}
		q.head = 0
		q.tail = temp
	}
	q.items[q.tail] = num
	q.tail ++
	return nil
}

func (q *Queue) PopQueue() error {
	if q.head == q.tail {
		return errors.New("当前队列为空，无法删除元素")
	}
	q.head ++
	return nil
}

func (q Queue) PrintQueue() {
	for _, data := range q.items[q.head:q.tail] {
		fmt.Printf("%d ", data)
	}
	fmt.Println()
}