package stack

import (
	"container/list"
	"fmt"
)

type Stack struct {
	root *list.List
}

func New() *Stack {
	s := new(Stack)
	s.root = list.New()
	return s
}

func (s *Stack) Push(v interface{}) {
	s.root.PushBack(v)
}

func (s *Stack) Pop() interface{} {
	return s.root.Remove(s.root.Back())
}

func (s *Stack) Print() {
	for e := s.root.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
