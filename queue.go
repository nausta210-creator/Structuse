package main

import "fmt"

type QNode struct {
	Data string
	Next *QNode
}

type Queue struct {
	Head *QNode
	Tail *QNode
}

func initQueue() *Queue {
	return &Queue{Head: nil, Tail: nil}
}

func (q *Queue) destroy() {
	q.Head = nil
	q.Tail = nil
}

func (q *Queue) push(val string) {
	newNode := &QNode{Data: val, Next: nil}
	if q.Tail != nil {
		q.Tail.Next = newNode
	} else {
		q.Head = newNode
	}
	q.Tail = newNode
}

func (q *Queue) pop() string {
	if q.Head == nil {
		return ""
	}
	val := q.Head.Data
	q.Head = q.Head.Next
	if q.Head == nil {
		q.Tail = nil
	}
	return val
}

func (q *Queue) print() {
	if q.Head == nil {
		fmt.Println("[Пустая очередь]")
		return
	}
	fmt.Print("[ ")
	curr := q.Head
	for curr != nil {
		fmt.Print(curr.Data)
		if curr.Next != nil {
			fmt.Print(" <- ")
		}
		curr = curr.Next
	}
	fmt.Println(" ]")
}