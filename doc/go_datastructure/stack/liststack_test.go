package stack

import (
	"fmt"
	"testing"
)

func TestListstack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.PrintStack()
	stack.Push(2)
	stack.PrintStack()
	stack.Push(3)
	stack.PrintStack()
	if !stack.isEmpty() {
		fmt.Println("当前栈不为空")
	}
	stack.Pop()
	stack.PrintStack()
	stack.Pop()
	stack.PrintStack()
	stack.Pop()
	stack.PrintStack()
}