package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {

	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {

	item := s.items[len(s.items)-1]

	s.items = s.items[:len(s.items)-1]
	return item

}

func (s *Stack) Peek() (int, error) {
	if len(s.items) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func stackTest() {

	stack := Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("Stack after pushing 1, 2, and 3:", stack.items)

	item := stack.Pop()

	fmt.Println("Popped item from stack:", item)

}
