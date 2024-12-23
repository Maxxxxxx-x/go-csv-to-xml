package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"

	"github.com/Maxxxxxx-x/go-csv-to-xml/model"
	"github.com/gocarina/gocsv"
)

func parseCSVFile(filePath string) ([]model.Book, error) {
	var results []model.Book

	file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return results, err
	}

	defer file.Close()
	if err := gocsv.Unmarshal(file, &results); err != nil {
		return results, err
	}

	return results, nil
}

func main() {
	inputFileNamePtr := flag.String("i", "", "CSV File to convert to XML")
	outputFileNamePtr := flag.String("o", "", "XML File output name")
	flag.Parse()

	if *inputFileNamePtr == "" {
		flag.PrintDefaults()
		return
	}
	if *outputFileNamePtr == "" {
		*outputFileNamePtr = "output.xml"
	}

	result, err := parseCSVFile(*inputFileNamePtr)
	if err != nil {
		fmt.Printf("Failed to parse CSV File: %s\n", err.Error())
        return
	}

    books := model.BookList{
        Books: result,
    }

    output, _ := xml.MarshalIndent(books, "", " ")
    if err := os.WriteFile(*outputFileNamePtr, output, 0644); err != nil {
        fmt.Printf("Failed to create XML File: %s\n", err.Error())
        return
    }
    fmt.Printf("Converted %s to %s\n", *inputFileNamePtr, *outputFileNamePtr)
}
