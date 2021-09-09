package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/xuri/excelize/v2"
)

const RowSymbols string = "ABCDEFGHIJKLMN"

func main() {
	logger := log.Default()
	// get file paths
	var files []string
	err := filepath.Walk("csv", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if regexp.MustCompile(`.+.csv`).Match([]byte(path)) {
			files = append(files, info.Name())
		}
		return nil
	})
	if err != nil {
		logger.Fatal(err)
	}

	// create new excelfile
	excel := excelize.NewFile()
	for _, f := range files {
		logger.Printf("READ: %v", f)
		// create sheet
		_ = excel.NewSheet(f)
		// read csv
		openedFile, err := os.Open(fmt.Sprintf("csv/%s", f))
		defer openedFile.Close()
		if err != nil {
			logger.Fatal(err)
		}
		reader := csv.NewReader(openedFile)
		count := 1
		for {
			line, err := reader.Read()
			logger.Printf("WRITE: %v", line)
			if err != nil {
				logger.Printf(err.Error() + "\n")
				break
			}
			// add value
			for i, v := range line {
				excel.SetCellValue(f, fmt.Sprint(RowSymbols[i:i+1], count), v)
			}
			// tweak layout
			excel.SetColWidth(f, "A", "A", 40)
			excel.SetColWidth(f, "B", "B", 30)
			excel.SetColWidth(f, "C", "C", 100)
			count++
		}
	}
	// save excel
	excel.DeleteSheet("Sheet1")
	if err := excel.SaveAs(fmt.Sprintf("output/%s.xlsx", time.Now().Format("20060102_150405"))); err != nil {
		log.Fatal(err)
	}
	logger.Println("**** FINISH ****")
}
