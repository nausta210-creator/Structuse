package main

import "fmt"

type SNode struct {
	Data string
	Next *SNode
}

type Stack struct {
	Top *SNode
}

func initStack() *Stack {
	return &Stack{Top: nil}
}

func (s *Stack) destroy() {
	s.Top = nil
}

func (s *Stack) push(val string) {
	newNode := &SNode{Data: val, Next: s.Top}
	s.Top = newNode
}

func (s *Stack) pop() string {
	if s.Top == nil {
		return ""
	}
	val := s.Top.Data
	s.Top = s.Top.Next
	return val
}

func (s *Stack) print() {
	if s.Top == nil {
		fmt.Println("[Пустой стек]")
		return
	}
	fmt.Print("[ ")
	curr := s.Top
	for curr != nil {
		fmt.Print(curr.Data)
		if curr.Next != nil {
			fmt.Print(" -> ")
		}
		curr = curr.Next
	}
	fmt.Println(" ]")
}