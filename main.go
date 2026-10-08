package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// splitQuery — разбивает строку на слова
func splitQuery(query string) []string {
	return strings.Fields(query)
}

// isNumber — проверяет, что строка состоит только из цифр
func isNumber(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func main() {
	filename := ""
	query := ""

	// Разбор аргументов
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if arg == "--file" && i+1 < len(os.Args) {
			i++
			filename = os.Args[i]
		} else if arg == "--query" && i+1 < len(os.Args) {
			i++
			query = os.Args[i]
		}
	}

	if filename == "" || query == "" {
		fmt.Fprintln(os.Stderr, "Использование: ./dbms --file <file.data> --query '<COMMAND> <name> [args]'")
		os.Exit(1)
	}

	records := loadFromFile(filename)
	args := splitQuery(query)
	if len(args) == 0 {
		return
	}

	command := args[0]

	structName := ""
	if len(args) > 1 {
		structName = args[1]
	}

	structType := ""
	if command[0] == 'M' {
		structType = "M"
	} else if command[0] == 'F' {
		structType = "F"
	} else if command[0] == 'L' {
		structType = "L"
	} else if command[0] == 'S' {
		structType = "S"
	} else if command[0] == 'Q' {
		structType = "Q"
	} else if command[0] == 'T' {
		structType = "T"
	} else if command == "PRINT" {
		for _, rec := range records {
			if rec.Name == structName {
				structType = rec.Type
				break
			}
		}
	}

	if structType == "" {
		fmt.Fprintln(os.Stderr, "Неизвестная команда или структура: "+command)
		os.Exit(1)
	}

	modified := false

	// ---------------- МАССИВ ----------------
	if structType == "M" {
		recIdx := -1
		for i := range records {
			if records[i].Name == structName && records[i].Type == "M" {
				recIdx = i
				break
			}
		}
		arr := initArray(4)
		if recIdx != -1 {
			for _, val := range records[recIdx].Items {
				arr.pushBack(val)
			}
		}

		if command == "MPUSH" {
			if len(args) == 3 {
				arr.pushBack(args[2])
				modified = true
			} else if len(args) == 4 {
				idx, _ := strconv.Atoi(args[2])
				arr.insertAt(idx, args[3])
				modified = true
			}
		} else if command == "MDEL" && len(args) == 3 {
			idx, _ := strconv.Atoi(args[2])
			arr.removeAt(idx)
			modified = true
		} else if command == "MGET" && len(args) == 3 {
			idx, _ := strconv.Atoi(args[2])
			fmt.Println(arr.getAt(idx))
		} else if command == "MSET" && len(args) == 4 {
			idx, _ := strconv.Atoi(args[2])
			arr.setAt(idx, args[3])
			modified = true
		} else if command == "MLEN" {
			fmt.Println(arr.getSize())
		} else if command == "PRINT" {
			arr.print()
		}

		if modified {
			newItems := []string{}
			for i := 0; i < arr.Length; i++ {
				newItems = append(newItems, arr.Data[i])
			}
			if recIdx != -1 {
				records[recIdx].Items = newItems
			} else {
				records = append(records, DataRecord{Type: "M", Name: structName, Items: newItems})
			}
			saveToFile(filename, records)
		}
		arr.destroy()
	} else if structType == "F" {
		// ---------------- ОДНОСВЯЗНЫЙ СПИСОК ----------------
		recIdx := -1
		for i := range records {
			if records[i].Name == structName && records[i].Type == "F" {
				recIdx = i
				break
			}
		}
		list := initFList()
		if recIdx != -1 {
			for _, val := range records[recIdx].Items {
				list.pushTail(val)
			}
		}

		if command == "FPUSH" && len(args) >= 4 {
			mode := args[2]
			if mode == "HEAD" {
				list.pushHead(args[3])
			} else if mode == "TAIL" {
				list.pushTail(args[3])
			} else if mode == "AFTER" && len(args) == 5 {
				list.pushAfter(args[3], args[4])
			} else if mode == "BEFORE" && len(args) == 5 {
				list.pushBefore(args[3], args[4])
			}
			modified = true
		} else if command == "FDEL" && len(args) >= 3 {
			mode := args[2]
			if mode == "HEAD" {
				list.popHead()
			} else if mode == "TAIL" {
				list.popTail()
			} else if mode == "VAL" && len(args) == 4 {
				list.removeValue(args[3])
			} else if mode == "BEFORE" && len(args) == 4 {
				list.removeBefore(args[3])
			} else if mode == "AFTER" && len(args) == 4 {
				list.removeAfter(args[3])
			}
			modified = true
		} else if command == "FGET" && len(args) == 3 {
			if isNumber(args[2]) {
				idx, _ := strconv.Atoi(args[2])
				val := list.getAt(idx)
				if val != "" {
					fmt.Println(val)
				}
			} else {
				found := list.findValue(args[2])
				if found {
					fmt.Println("TRUE")
				} else {
					fmt.Println("FALSE")
				}
			}
		} else if command == "PRINT" {
			list.print()
		} else if command == "FPRINT_REC" {
			list.printRecursive()
		}

		if modified {
			newItems := []string{}
			curr := list.Head
			for curr != nil {
				newItems = append(newItems, curr.Data)
				curr = curr.Next
			}
			if recIdx != -1 {
				records[recIdx].Items = newItems
			} else {
				records = append(records, DataRecord{Type: "F", Name: structName, Items: newItems})
			}
			saveToFile(filename, records)
		}
		list.destroy()
	} else if structType == "L" {
		// ---------------- ДВУСВЯЗНЫЙ СПИСОК ----------------
		recIdx := -1
		for i := range records {
			if records[i].Name == structName && records[i].Type == "L" {
				recIdx = i
				break
			}
		}
		list := initDList()
		if recIdx != -1 {
			for _, val