package main

import (
	"bufio"
	"os"
	"strings"
)

type DataRecord struct {
	Type  string
	Name  string
	Items []string
}

// loadFromFile — читает все структуры из файла
func loadFromFile(filename string) []DataRecord {
	records := []DataRecord{}

	file, err := os.Open(filename)
	if err != nil {
		return records // файл не существует — пустой список
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		record := DataRecord{
			Type:  parts[0],
			Name:  parts[1],
			Items: parts[2:],
		}
		records = append(records, record)
	}
	return records
}

// saveToFile — записывает все структуры в файл
func saveToFile(filename string, records []DataRecord) {
	file, err := os.Create(filename)
	if err != nil {
		os.Stderr.WriteString("Ошибка открытия файла для записи: " + filename + "\n")
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, r := range records {
		writer.WriteString(r.Type + " " + r.Name)
		for _, item := range r.Items {
			writer.WriteString(" " + item)
		}
		writer.WriteString("\n")
	}
	writer.Flush()
}