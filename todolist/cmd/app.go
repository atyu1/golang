package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

const FILEPATH string = "./data.csv"

type todoItems struct {
	ID          uint32
	Description string
	DateCreated time.Time
	IsCompleted bool
}

// csv manager

func initCSV(filePath string) {
	defaultHeaders := []string{"ID", "Text", "DateCreated", "Done"}
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		fmt.Println("CSV File does not exists, creating one with default headers")

		f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal("Problem with openning the file for writing: ", filePath, err)
		}
		defer f.Close()

		csvWriter := csv.NewWriter(f)
		err = csvWriter.Write(defaultHeaders)
		if err != nil {
			log.Fatal("Problem with writing CSV to file: ", filePath, err)
		}
		defer csvWriter.Flush()
	}
}

func loadCSV(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Problem with openning the file: ", filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Problem with CSV parsing for file for reading: ", filePath, err)
	}

	return records
}

func appendCSV(filePath string, record todoItems) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Problem with openning the file for writing: ", filePath, err)
	}
	defer f.Close()

	csvWriter := csv.NewWriter(f)

	recordStringed := []string{fmt.Sprint(record.ID), record.Description, fmt.Sprint(record.DateCreated), fmt.Sprint(record.IsCompleted)}
	err = csvWriter.Write(recordStringed)
	if err != nil {
		log.Fatal("Problem with writing CSV to file: ", filePath, err)
	}
	defer csvWriter.Flush()
}

func writeToConsole(records [][]string) {
	w := new(tabwriter.Writer)
	defer w.Flush()

	w.Init(os.Stdout, 0, 18, 4, ' ', 0)
	for _, record := range records {
		fmt.Fprintln(w, strings.Join(record, "\t"))
	}
}

func getLastID(filePath string) uint32 {
	records := loadCSV(filePath)

	lastRecord := records[len(records)-1]
	lastId, err := strconv.Atoi(lastRecord[0])
	if err != nil {
		fmt.Println("Last ID is unkown, starting the numbering from 0")
		return uint32(0)
	}

	return uint32(lastId)
}

func findID(filePath string, id uint32) int {
	records := loadCSV(filePath)
	for counter, record := range records {
		if x, err := strconv.Atoi(record[0]); uint32(x) == id {
			if err != nil {
				log.Fatal("Cannot find the ID for: ", id, " Error: ", err)
			}
			return counter
		}
	}
	return -1
}

func deleteLine(records [][]string, line int) [][]string {
	var newRecords [][]string

	newRecords = make([][]string, len(records)-1)
	counter := 0
	for i, record := range records {
		if i == line {
			continue
		}
		newRecords[counter] = make([]string, len(record))
		newRecords[counter] = record
		counter += 1
	}

	return newRecords
}

func markDone(records [][]string, line int) [][]string {
	var newRecords [][]string

	newRecords = make([][]string, len(records))
	for i, record := range records {
		if i == line {
			record[len(record)-1] = "true"
		}
		newRecords[i] = make([]string, len(record))
		newRecords[i] = record
	}

	return newRecords
}

func replaceCSV(filePath string, records [][]string) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Problem with openning the file for writing: ", filePath, err)
	}
	defer f.Close()

	csvWriter := csv.NewWriter(f)

	err = csvWriter.WriteAll(records)
	if err != nil {
		log.Fatal("Problem with writing CSV to file: ", filePath, err)
	}
}
