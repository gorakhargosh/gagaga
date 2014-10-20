package stack

import "errors"

type Stack []interface{}

func (s Stack) Len() int {
	return len(s)
}

func (s *Stack) Push(x interface{}) {
	*s = append(*s, x)
}

func (s Stack) Top() (interface{}, error) {
	if len(s) == 0 {
		return nil, errors.New("can't Top() an empty stack")
	}
	return s[len(s)-1], nil
}

func (s *Stack) Pop() (interface{}, error) {
	stack := *s
	if len(stack) == 0 {
		return nil, errors.New("can't Pop() an empty stack.")
	}

	x := stack[len(stack)-1]
	*s = stack[:len(stack)-1]
	return x, nil
}
