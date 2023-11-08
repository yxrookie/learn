package queue

import (
	"container/list"
	"fmt"
)

type Lqueue struct {
	list *list.List
}

func NewLqueue() *Lqueue {
	return &Lqueue{
		list: list.New(),
	}
}

func (q *Lqueue) Push(value interface{}) {
	q.list.PushBack(value)
}

func (q *Lqueue) Pop() (interface{}, error) {
	if q.list.Len() == 0 {
		return nil, fmt.Errorf("队列为空，无法删除队头元素")
	}
	value := q.list.Front()
	q.list.Remove(value)
	return value.Value, nil
}

func (q Lqueue) PrintQueue() {
	if q.list.Len() == 0 {
		fmt.Println("队列为空")
		return
	}
	for e := q.list.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
}