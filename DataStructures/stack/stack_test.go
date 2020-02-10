package stack

import "testing"

func TestStack(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if len(s.s) != 3 {
		t.Error("Length stack must be 3")
	}
	if v, _ := s.Pop(); v != 3 {
		t.Error("First value must be 3")
	}

	t.Log(s.Pop())
}
