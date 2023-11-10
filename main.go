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

	// Defininf a boolean flag -b to count bytes
	b := flag.Bool("b", false, "Count Bytes")
	// parse the flags provided by the user
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *b))
}

// a function that counts command line args by words
func count(r io.Reader, countLines bool, countBytes bool) int {
	var answer int = 0

	// an io Reader is a go type (interface) from which u can read data
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// if the count lines flag is not set, we want to count words so we define
	// the scanner split type to words (default behavior is to split by lines)
	if countBytes {
		scanner.Split(bufio.ScanBytes)
		byte_count := 0

		for scanner.Scan() {
			token := scanner.Text()
			fmt.Println("Token is " + token)
			byte_count++
		}
		answer = byte_count
	} else if !countLines == true {
		scanner.Split(bufio.ScanWords)
		wc := 0

		for scanner.Scan() {
			wc++
		}

		answer = wc
	}

	// return the total
	return answer
}
