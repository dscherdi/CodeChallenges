package main

// https://leetcode.com/problems/min-stack/?envType=study-plan-v2&envId=top-interview-150
type MinStack struct {
	value    int
	previous *MinStack
	minimum  int
}

func NewMinStack() *MinStack {
	return &MinStack{minimum: int(^uint(0) >> 1)}
}

func (stack *MinStack) Push(val int) {
	newMin := stack.minimum

	if val < newMin {
		newMin = val
	}

	newHead := &MinStack{
		value:    stack.value,
		previous: stack.previous,
		minimum:  stack.minimum,
	}

	stack.value = val
	stack.minimum = newMin
	stack.previous = newHead
}

func (stack *MinStack) Pop() {
	stack.value = stack.previous.value
	stack.minimum = stack.previous.minimum

	previous := stack.previous.previous
	stack.previous.previous = nil
	stack.previous = previous

}

func (stack *MinStack) Top() int {
	return stack.value
}

func (stack *MinStack) GetMin() int {
	return stack.minimum
}
