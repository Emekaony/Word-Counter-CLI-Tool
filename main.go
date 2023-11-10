package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// Defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	// parse the flags provided by the user
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines))
}

// a function that counts command line args by words
func count(r io.Reader, countLines bool) int {
	// an io Reader is a go type (interface) from which u can read data
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// if the count lines flag is not set, we want to count words so we define
	// the scanner split type to words (default behavior is to split by lines)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	// Defining a counter
	wc := 0

	// for every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	// return the total
	return wc
}
