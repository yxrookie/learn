package stack

import "testing"

func TestStackFun(t *testing.T) {
	var s Stack
	s.initSatck(4)
	s.push(1)
	s.printStack()
	s.push(3)
	s.printStack()
	s.push(4)
	s.printStack()
	s.pop()
	s.printStack()
	s.pop()
	s.printStack()
}