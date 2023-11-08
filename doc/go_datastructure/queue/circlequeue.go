package queue

import (
	"errors"
	"fmt"
)

type Circlequeue struct {
	items    []int
	head     int
	tail     int
	capacity int
}

func (c *Circlequeue) initQueue(capacity int) Circlequeue {
	return Circlequeue{
		items:    make([]int, capacity),
		capacity: capacity,
	}
}

func (c *Circlequeue) Push(num int) error {
	if (c.tail+1)%c.capacity == c.head {
		return errors.New("队列已满，无法添加新的元素")
	}
	c.items[c.tail] = num
	c.tail = (c.tail+1) % c.capacity
	return nil
}

func (c *Circlequeue) Pop() error {
	if c.tail == c.head {
		return errors.New("队列为空，无法删除对头元素")
	}
	c.head = (c.head+1)%c.capacity
	return nil
}

func (c *Circlequeue) PrintQueue() {
	// 注意队列的数据范围
	for _, data := range c.items[c.head:c.tail] {
		fmt.Printf("%d ", data)
	}
	fmt.Println()
}