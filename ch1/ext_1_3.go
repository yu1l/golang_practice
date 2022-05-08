package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func with_strings() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

// with_strings() is faster than echo1()
func main() {
	start_echo1 := time.Now()

	fmt.Println("start - echo1()")
	echo1()

	finish_echo1 := time.Now()
	elapsed_echo1 := finish_echo1.Sub(start_echo1)
	fmt.Println(elapsed_echo1)


	start_with_strings := time.Now()

	fmt.Println("start - with_strings()")
	with_strings()

	finish_with_strings := time.Now()
	elapsed_with_strings := finish_with_strings.Sub(start_with_strings)
	fmt.Println(elapsed_with_strings)
}
