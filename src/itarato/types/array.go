package types

import (
	"errors"
	"log"
)

type StackItem struct {
	item interface{}
	prev *StackItem
}

type Stack struct {
	top   *StackItem
	count uint8
}

func NewStack() (s *Stack) {
	return &Stack{}
}

func (s *Stack) Push(item interface{}) {
	s.count++
	s.top = &StackItem{item, s.top}
	log.Println("Stack", s, "push:", item)
}

func (s *Stack) Pop() (item interface{}, err error) {
	if s.count <= 0 {
		return nil, errors.New("Empty stack")
	}

	s.count--
	item = s.top
	s.top = s.top.prev
	return item, nil
}

func (s *Stack) Top() (item interface{}, err error) {
	if s.count <= 0 {
		return nil, errors.New("Empty stack")
	}

	item = s.top
	return item, nil
}

func (s *Stack) Length() (count uint8) {
	return s.count
}
