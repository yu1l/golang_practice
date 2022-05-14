// $ go run ch3/ext_3_10.go 1000000
// => 1,000,000
package main

import (
	"bytes"
	"os"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	if len(s) <= 3 {
		fmt.Print(s)
		os.Exit(0)
	}
	prefix := len(s) % 3
	buf.WriteString(s[:prefix])
	commaS := s[prefix:]
	for i := 0; i < len(commaS) ; i += 3 {
		fmt.Fprintf(&buf, ",")
		buf.WriteString(commaS[i:i+3])
	}
	return buf.String()
}

func main() {
	s := os.Args[1]
	fmt.Println(comma(s))
}
