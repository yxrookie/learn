package stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	items  []int
	count  int
	maxlen int
}

func (s *Stack) initSatck(capacity int) *Stack {
	s.items = make([]int, capacity)
	s.maxlen = capacity 
	return s
}

func (s *Stack) push(num int) error {
	if s.count >= s.maxlen {
		return errors.New("栈已满，无法添加新的元素")
	}
	s.items[s.count] = num
	s.count ++
	return nil
}

func  (s *Stack) pop() (int, error) {
	if s.count <= 0 {
		return 0, errors.New("栈当前容量为空，无法删除元素")
	}
	delitem := s.items[s.count]
	s.count--
	return delitem, nil
}

func(s Stack) printStack() {
	for i := 0; i < s.count; i++ {
		fmt.Printf("%d ", s.items[i])
	}
	fmt.Println()
}
