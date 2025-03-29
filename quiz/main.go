package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func readCSV(f string) [][]string {
	csv_file, err := os.Open(f)
	if err != nil {
		log.Fatal("Cannot open the file")
		os.Exit(1)
	}
	content := csv.NewReader(csv_file)
	questions, err := content.ReadAll()
	if err != nil {
		log.Fatal("Cannot read the csv file content")
		os.Exit(1)
	}

	return questions
}

func main() {
	csvfilename := flag.String("csv", "questions.csv", "CSV File with questions")
	flag.Parse()

	qs := readCSV(*csvfilename)
	var correct int = 0
	var incorrect int = 0

	for _, q := range qs {
		fmt.Print(q[0], "=")
		var a string
		_, _ = fmt.Scanln(&a)
		if a == "end" {
			break
		}

		//fmt.Println("Debug: ", a)
		if a == q[1] {
			correct += 1
		} else {
			incorrect += 1
		}
	}
	fmt.Println("----- Results -----")
	fmt.Println("# Correct: ", correct)
	fmt.Println("# Incorrect: ", incorrect)
}
