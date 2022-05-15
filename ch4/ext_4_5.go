// $ go run ch4/ext_4_5.go heeeelllooo
package main

import (
	"fmt"
	"bytes"
	"os"
)

func removeNextDuplicatedChars(s string) string {
	var buf bytes.Buffer
	stack := []string{}
	for _, r := range s {
		str := string(r)
		if len(stack) == 0 {
			stack = append(stack, str)
		} else {
			lastChar := stack[len(stack)-1]
			if lastChar != str {
				stack = append(stack, str)
			}
		}
	}
	for _, r := range stack {
		fmt.Fprintf(&buf, "%s", r)
	}
	return buf.String()
}

func main() {
	s := os.Args[1]
	res := removeNextDuplicatedChars(s)
	fmt.Println(res)
}
