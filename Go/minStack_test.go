package main

import "testing"

func TestMinStack(t *testing.T) {
	stack := NewMinStack()

	stack.Push(5)
	if stack.Top() != 5 {
		t.Errorf("Top() = %v, want %v", stack.Top(), 5)
	}
	if stack.GetMin() != 5 {
		t.Errorf("GetMin() = %v, want %v", stack.GetMin(), 5)
	}

	stack.Push(3)
	if stack.Top() != 3 {
		t.Errorf("Top() = %v, want %v", stack.Top(), 3)
	}
	if stack.GetMin() != 3 {
		t.Errorf("GetMin() = %v, want %v", stack.GetMin(), 3)
	}

	stack.Push(7)
	if stack.Top() != 7 {
		t.Errorf("Top() = %v, want %v", stack.Top(), 7)
	}
	if stack.GetMin() != 3 {
		t.Errorf("GetMin() = %v, want %v", stack.GetMin(), 3)
	}

	stack.Pop()
	if stack.Top() != 3 {
		t.Errorf("Top() = %v, want %v", stack.Top(), 3)
	}
	if stack.GetMin() != 3 {
		t.Errorf("GetMin() = %v, want %v", stack.GetMin(), 3)
	}

	stack.Pop()
	if stack.Top() != 5 {
		t.Errorf("Top() = %v, want %v", stack.Top(), 5)
	}
	if stack.GetMin() != 5 {
		t.Errorf("GetMin() = %v, want %v", stack.GetMin(), 5)
	}
}
