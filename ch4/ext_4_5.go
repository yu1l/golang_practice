// $ go run ch4/ext_4_5.go heeeelllooo
// => helo
package main

import (
	"fmt"
	"bytes"
	"os"
)

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func removeNextDuplicatedCharInStringSlice(s string) string {
	var buf bytes.Buffer
	strSlice := []string{}
	for _, r := range s {
		strSlice = append(strSlice, string(r))
	}

	step := 1
	for index := 0; index < len(strSlice); index += step {
		if index == 0 {
			continue
		}

		prevChar := strSlice[index-1]
		currentChar := strSlice[index]
		if prevChar == currentChar {
			strSlice = remove(strSlice, index)
			step = 0
		} else {
			step = 1
		}
	}
	for _, r := range strSlice {
		fmt.Fprintf(&buf, "%s", string(r))
	}
	return buf.String()
}

func main() {
	s := os.Args[1]
	res := removeNextDuplicatedCharInStringSlice(s)
	fmt.Println(res)
}
