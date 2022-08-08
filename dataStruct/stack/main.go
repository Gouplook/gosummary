package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int
	Top    int
	Arr    []int
}

func (s *Stack) Push(val int) (err error) {
	if s.Top == s.MaxTop-1 {
		return errors.New("stack full ..")
	}
	s.Top++
	s.Arr[s.Top] = val
	return
}

func (s *Stack) Pull() (val int, err error) {
	if s.Top == -1 {
		return 0, errors.New("stack empty")
	}
	val = s.Arr[s.Top]
	s.Top--
	return val, nil
}
func (s *Stack) List() {
	if s.Top == -1 {
		fmt.Println("stack empty ")
	}
	for i := s.Top; i >= 0; i-- {
		fmt.Printf("s.Arr[%d]=%d\n", i, s.Arr[i])
	}
}

func main() {
	stack := &Stack{
		MaxTop: 5,
		Top:    -1,
		Arr:    make([]int, 5),
	}
	_ = stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	stack.List()

}
