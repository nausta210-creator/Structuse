package main

import "fmt"

type DNode struct {
	Data string
	Prev *DNode
	Next *DNode
}

type DoublyLinkedList struct {
	Head *DNode
	Tail *DNode
}

func initDList() *DoublyLinkedList {
	return &DoublyLinkedList{Head: nil, Tail: nil}
}

func (l *DoublyLinkedList) destroy() {
	l.Head = nil
	l.Tail = nil
}

// ---- ДОБАВЛЕНИЕ ----

func (l *DoublyLinkedList) pushHead(val string) {
	newNode := &DNode{Data: val, Prev: nil, Next: l.Head}
	if l.Head != nil {
		l.Head.Prev = newNode
	} else {
		l.Tail = newNode
	}
	l.Head = newNode
}

func (l *DoublyLinkedList) pushTail(val string) {
	newNode := &DNode{Data: val, Prev: l.Tail, Next: nil}
	if l.Tail != nil {
		l.Tail.Next = newNode
	} else {
		l.Head = newNode
	}
	l.Tail = newNode
}

func (l *DoublyLinkedList) pushAfter(target, val string) {
	curr := l.Head
	for curr != nil && curr.Data != target {
		curr = curr.Next
	}
	if curr != nil {
		newNode := &DNode{Data: val, Prev: curr, Next: curr.Next}
		if curr.Next != nil {
			curr.Next.Prev = newNode
		} else {
			l.Tail = newNode
		}
		curr.Next = newNode
	} else {
		fmt.Printf("Элемент '%s' не найден\n", target)
	}
}

func (l *DoublyLinkedList) pushBefore(target, val string) {
	curr := l.Head
	for curr != nil && curr.Data != target {
		curr = curr.Next
	}
	if curr != nil {
		newNode := &DNode{Data: val, Prev: curr.Prev, Next: curr}
		if curr.Prev != nil {
			curr.Prev.Next = newNode
		} else {
			l.Head = newNode
		}
		curr.Prev = newNode
	} else {
		fmt.Printf("Элемент '%s' не найден\n", target)
	}
}

// ---- УДАЛЕНИЕ ----

func (l *DoublyLinkedList) popHead() {
	if l.Head == nil {
		return
	}
	l.Head = l.Head.Next
	if l.Head != nil {
		l.Head.Prev = nil
	} else {
		l.Tail = nil
	}
}

func (l *DoublyLinkedList) popTail() {
	if l.Tail == nil {
		return
	}
	l.Tail = l.Tail.Prev
	if l.Tail != nil {
		l.Tail.Next = nil
	} else {
		l.Head = nil
	}
}

func (l *DoublyLinkedList) removeValue(val string) {
	curr := l.Head
	for curr != nil && curr.Data != val {
		curr = curr.Next
	}
	if curr == nil {
		return
	}
	if curr == l.Head {
		l.popHead()
	} else if curr == l.Tail {
		l.popTail()
	} else {
		curr.Prev.Next = curr.Next
		curr.Next.Prev = curr.Prev
	}
}

func (l *DoublyLinkedList) removeBefore(target string) {
	curr := l.Head
	for curr != nil && curr.Data != target {
		curr = curr.Next
	}
	if curr != nil && curr.Prev != nil {
		toDelete := curr.Prev
		if toDelete == l.Head {
			l.popHead()
		} else {
			toDelete.Prev.Next = curr
			curr.Prev = toDelete.Prev
		}
	}
}

func (l *DoublyLinkedList) removeAfter(target string) {
	curr := l.Head
	for curr != nil && curr.Data != target {
		curr = curr.Next
	}
	if curr != nil && curr.Next != nil {
		toDelete := curr.Next
		if toDelete == l.Tail {
			l.popTail()
		} else {
			curr.Next = toDelete.Next
			toDelete.Next.Prev = curr
		}
	}
}

// ---- ПОИСК И ЧТЕНИЕ ----

func (l *DoublyLinkedList) findValue(val string) bool {
	curr := l.Head
	for curr != nil {
		if curr.Data == val {
			return true
		}
		curr = curr.Next
	}
	return false
}

func (l *DoublyLinkedList) getAt(index int) string {
	curr := l.Head
	curIdx := 0
	for curr != nil {
		if curIdx == index {
			return curr.Data
		}
		curr = curr.Next
		curIdx++
	}
	fmt.Println("Ошибка: Индекс выходит за границы")
	return ""
}

func (l *DoublyLinkedList) print() {
	if l.Head == nil {
		fmt.Println("[Пустой список]")
		return
	}
	fmt.Print("[ ")
	curr := l.Head
	for curr != nil {
		fmt.Print(curr.Data)
		if curr.Next != nil {
			fmt.Print(" <=> ")
		}
		curr = curr.Next
	}
	fmt.Println(" ]")
}

func (l *DoublyLinkedList) printReverse() {
	if l.Tail == nil {
		fmt.Println("[Пустой список]")
		return
	}
	fmt.Print("[ ")
	curr := l.Tail
	for curr != nil {
		fmt.Print(curr.Data)
		if curr.Prev != nil {
			fmt.Print(" <=> ")
		}
		curr = curr.Prev
	}
	fmt.Println(" ]")
}