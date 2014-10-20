package stack

import "errors"

type Stack []interface{}

func (s Stack) Len() int {
	return len(s)
}

func (s *Stack) Push(x interface{}) {
	*stack = append(*stack, x)
}

func (s Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("can't Top() an empty stack")
	}
	return stack[len(stack)-1], nil
}
