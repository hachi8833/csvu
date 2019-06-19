package main

import (
	"encoding/csv"
	"fmt"
	"github.com/logrusorgru/aurora"
	"io"
	"os"
)

// DELIMITER holds a string to delimition.
const DELIMITER = " "

func main() {

	fp := readStdin()

	reader := csv.NewReader(fp)
	reader.Comma = ','
	reader.LazyQuotes = true

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		for i, col := range record {
			rotateColor(i, col)
		}
		fmt.Println()
	}
}

// rotateColor exports the colored columns with circuration.
func rotateColor(i int, col interface{}) {
	switch i % 7 {
	case 0:
		fmt.Print(aurora.Magenta(col), DELIMITER)
	case 1:
		fmt.Print(aurora.Blue(col), DELIMITER)
	case 2:
		fmt.Print(aurora.Brown(col), DELIMITER)
	case 3:
		fmt.Print(aurora.Green(col), DELIMITER)
	case 4:
		fmt.Print(aurora.White(col), DELIMITER)
	case 5:
		fmt.Print(aurora.Cyan(col), DELIMITER)
	case 6:
		fmt.Print(aurora.Red(col), DELIMITER)
	}
	return
}

// readStdin returns the CSV from stdin.
func readStdin() *os.File {
	var fp *os.File
	if len(os.Args) < 2 {
		fp = os.Stdin
	} else {
		var err error
		fp, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}
	return fp
}
