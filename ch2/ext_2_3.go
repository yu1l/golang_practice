package main

import "fmt"

var pc [256]byte

func PopCount(x uint64) int {
	count := 0
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x >> (i * 8))])
	}
	return count
}

func main() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	count := PopCount(255)
	fmt.Println("count: ", count)
}
