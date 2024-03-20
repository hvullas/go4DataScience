package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	//open the file
	csvfile, err := os.Open("C:/Users/91973/Desktop/ullas/projects/readCSV/csvfiles/survey.csv")
	if err != nil {
		log.Fatal(err)
	}

	//implement workflow

	//parse the data

	data := csv.NewReader(csvfile)

	for {
		record, err := data.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}
}
