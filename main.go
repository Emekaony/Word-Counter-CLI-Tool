package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// calling the count function to print the number of words
	// received form the standard input and printing it out
	fmt.Println(count(os.Stdin))
}

// an io Reader is a go type (interface) from which u can read data
func count(r io.Reader) int {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// Define the scanner split type to words (default is split by lines)
	scanner.Split(bufio.ScanWords)

	// Defining a counter
	wc := 0

	// for every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	// return the total
	return wc
}
