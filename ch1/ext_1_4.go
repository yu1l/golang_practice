// display all file names
// which include every duplicated lines
// $ go run ch1/ext_1_4.go ch1/sample.txt ch1/sample1.txt
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			for _, arg := range files {
				f, err := os.Open(arg)
				if err != nil {
					fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
					continue
				}
				if isLineExistInFile(f, line) {
					fmt.Printf("at\t%s\n", arg)
				}
				f.Close()
			}
		}
	}
}

func isLineExistInFile(f *os.File, line string) bool {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == line {
			return true
		}
	}
	return false
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTICE: ignoring errors from input.Err()
}
