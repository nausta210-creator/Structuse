package main

import "fmt"

type FNode struct {
	Data string
	Next *FNode
}

type ForwardList struct {
	Head *FNode
}

func initFList() *ForwardList {
	return &ForwardList{Head: nil}
}

func (l *ForwardList) destroy() {
	l.Head = nil
}

// ---- ДОБАВЛЕНИЕ ----

func (l *ForwardList) pushHead(val string) {
	newNode := &FNode{Data: val, Next: l.Head}
	l.Head = newNode
}

func (l *ForwardList) pushTail(val string) {
	newNode := &FNode{Data: val, Next: nil}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
}

func (l *ForwardList) pushAfter(target, val string) {
	curr := l.Head
	for curr != nil && curr.Data != target {
		curr = curr.Next
	}
	if curr != nil {
		newNode := &FNode{Data: val, Next: curr.Next}
		curr.Next = newNode
	} else {
		fmt.Printf("Элемент '%s' не найден для вставки после\n", target)
	}
}

func (l *ForwardList) pushBefore(target, val string) {
	if l.Head == nil {
		return
	}
	if l.Head.Data == target {
		l.pushHead(val)
		return
	}
	curr := l.Head
	for curr.Next != nil && curr.Next.Data != target {
		curr = curr.Next
	}
	if curr.Next != nil {
		newNode := &FNode{Data: val, Next: curr.Next}
		curr.Next = newNode
	} else {
		fmt.Printf("Элемент '%s' не найден для вставки до\n", target)
	}
}

// ---- УДАЛЕНИЕ ----

func (l *ForwardList) popHead() {
	if l.Head == nil {
		return
	}
	l.Head = l.Head.Next
}

func (l *ForwardList) popTail() {
	if l.Head == nil {
		return
	}
	if l.Head.Next == nil {
		l.Head = nil
		return
	}
	curr := l.Head
	for curr.Next.Next != nil {
		curr = curr.Next
	}
	curr.Next = nil
}

func (l *ForwardList) removeValue(val string) {
	if l.Head == nil {
		return
	}
	if l.Head.Data == val {
		l.popHead()
		return
	}
	curr := l.Head
	for curr.Next != nil && curr.Next.Data != val {
		curr = curr.Next
	}
	if curr.Next != nil {
		curr.Next = curr.Next.Next
	}
}

func (l *ForwardList) removeBefore(target string) {
	if l.Head == nil || l.Head.Data == target {
		return
	}
	if l.Head.Next != nil && l.Head.Next.Data == target {
		l.popHead()
		return
	}
	curr := l.Head
	for curr.Next != nil && curr.Next.Next != nil && curr.Next.Next.Data != target {
		curr = curr.Next
	}
	if curr.Next != nil && curr.Next.Next != nil {
		curr.Next = curr.Next.Next
	}
}

func (l *ForwardList) removeAfter(target string) {
	curr := l.Head
	for curr != nil && curr.Data != target {
		curr = curr.Next
	}
	if curr != nil && curr.Next != nil {
		curr.Next = curr.Next.Next
	}
}

// ---- ПОИСК И ЧТЕНИЕ ----

func (l *ForwardList) findValue(val string) bool {
	curr := l.Head
	for curr != nil {
		if curr.Data == val {
			return true
		}
		curr = curr.Next
	}
	return false
}

func (l *ForwardList) getAt(index int) string {
	curr := l.Head
	curIdx := 0
	for curr != nil {
		if curIdx == index {
			return curr.Data
		}
		curr = curr.Next
		curIdx++
	}
	fmt.Println("Ошибка: Индекс выходит за границы списка")
	return ""
}

func (l *ForwardList) print() {
	if l.Head == nil {
		fmt.Println("[Пустой список]")
		return
	}
	fmt.Print("[ ")
	curr := l.Head
	for curr != nil {
		fmt.Print(curr.Data)
		if curr.Next != nil {
			fmt.Print(" -> ")
		}
		curr = curr.Next
	}
	fmt.Println(" ]")
}

func printFListRecursiveHelper(node *FNode) {
	if node == nil {
		return
	}
	fmt.Print(node.Data, " ")
	printFListRecursiveHelper(node.Next)
}

func (l *ForwardList) printRecursive() {
	fmt.Print("[ ")
	printFListRecursiveHelper(l.Head)
	fmt.Println("]")
}