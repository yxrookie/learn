package stack

import (
	"container/list"
	"fmt"
)

type lStack struct {
	list *list.List
}

func NewStack() *lStack {
	return &lStack{
		list: list.New(),
	}
}

func (s *lStack) Push(value interface{}) {
	s.list.PushBack(value)
}

func (s *lStack) Pop() (interface{}, error) {
	if s.list.Len() == 0 {
		return nil, fmt.Errorf("栈为空")
	}
	front := s.list.Back()
	s.list.Remove(front)
	// .Value 允许我们在链表中存储和检索不同类型的数据。
	return front.Value, nil
}

func (s *lStack) Peek() (interface{}, error) {
	if s.list.Len() == 0 {
		return nil, fmt.Errorf("当前栈为空，无发取出栈顶元素")
	}
	return s.list.Back().Value, nil
}

func (s *lStack) Len() int {
	return s.list.Len()
}

func (s *lStack) isEmpty() bool {
	/* if s.list.Len() == 0 {
		return true
	}
	return false */
	return s.list.Len() == 0
}

func (s lStack) PrintStack() {
	if s.list.Len() == 0 {
		fmt.Println("当前栈空")
		return
	}
	for e := s.list.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
}