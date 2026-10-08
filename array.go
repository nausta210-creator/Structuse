package main

import "fmt"

type DynamicArray struct {
	Data     []string
	Capacity int
	Length   int
}

func initArray(initialCapacity int) *DynamicArray {
	if initialCapacity <= 0 {
		initialCapacity = 4
	}
	return &DynamicArray{
		Data:     make([]string, initialCapacity),
		Capacity: initialCapacity,
		Length:   0,
	}
}

func (a *DynamicArray) resize(newCapacity int) {
	newData := make([]string, newCapacity)
	copy(newData, a.Data[:a.Length])
	a.Data = newData
	a.Capacity = newCapacity
}

func (a *DynamicArray) pushBack(value string) {
	if a.Length == a.Capacity {
		a.resize(a.Capacity * 2)
	}
	a.Data[a.Length] = value
	a.Length++
}

func (a *DynamicArray) insertAt(index int, value string) {
	if index < 0 || index > a.Length {
		fmt.Println("Ошибка: Индекс выходит за границы")
		return
	}
	if a.Length == a.Capacity {
		a.resize(a.Capacity * 2)
	}
	for i := a.Length; i > index; i-- {
		a.Data[i] = a.Data[i-1]
	}
	a.Data[index] = value
	a.Length++
}

func (a *DynamicArray) getAt(index int) string {
	if index < 0 || index >= a.Length {
		fmt.Println("Ошибка: Индекс выходит за границы")
		return ""
	}
	return a.Data[index]
}

func (a *DynamicArray) removeAt(index int) {
	if index < 0 || index >= a.Length {
		fmt.Println("Ошибка: Индекс выходит за границы")
		return
	}
	for i := index; i < a.Length-1; i++ {
		a.Data[i] = a.Data[i+1]
	}
	a.Length--
}

func (a *DynamicArray) setAt(index int, value string) {
	if index < 0 || index >= a.Length {
		fmt.Println("Ошибка: Индекс выходит за границы")
		return
	}
	a.Data[index] = value
}

func (a *DynamicArray) getSize() int {
	return a.Length
}

func (a *DynamicArray) print() {
	if a.Length == 0 {
		fmt.Println("[Пустой массив]")
		return
	}
	fmt.Print("[ ")
	for i := 0; i < a.Length; i++ {
		fmt.Print(a.Data[i])
		if i+1 < a.Length {
			fmt.Print(", ")
		}
	}
	fmt.Println(" ]")
}