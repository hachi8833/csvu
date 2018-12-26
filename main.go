package main

import (
	"encoding/csv"
	"fmt"
	"github.com/logrusorgru/aurora"
	"io"
	"os"
)

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

func rotateColor(c int, col interface{}) {
	switch c % 7 {
	case 0:
		fmt.Print(aurora.Magenta(col), " ")
	case 1:
		fmt.Print(aurora.Blue(col), " ")
	case 2:
		fmt.Print(aurora.Brown(col), " ")
	case 3:
		fmt.Print(aurora.Green(col), " ")
	case 4:
		fmt.Print(aurora.Gray(col), " ")
	case 5:
		fmt.Print(aurora.Cyan(col), " ")
	case 6:
		fmt.Print(aurora.Red(col), " ")
	}
	return
}

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
