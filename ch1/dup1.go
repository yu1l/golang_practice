// display lines which appear more than twice from standard input.
// using
// - if
// - map data type
// - bufio package

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// initial value of int is 0
	counts := make(map[string]int)

	// creates input stream
	input := bufio.NewScanner(os.Stdin)

	// input.Scan() reads next line text, and remove new line character from end of line.
	// input.Scan() returns true if line exists, and
	// returns false if no new line exits.
	for input.Scan() {
		counts[input.Text()]++

		// equals to
		// line := input.Text()
		// counts[line] = counts[line] + 1
	}

	// NOTICE: ignoring possible error from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
