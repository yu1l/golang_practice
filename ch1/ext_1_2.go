package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		var index string = fmt.Sprintf("%d", i)
		var arg string = os.Args[i]
		var output string = index + " : " + arg
		fmt.Println(output)
	}
}
