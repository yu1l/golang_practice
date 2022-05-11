// 実行速度遅いし多分間違ってるのであとで見直し
// 0 ~ 255

// $ go run ch2/ext_2_4.go 0
// => bit: [0, 0, 0, 0, 0, 0, 0, 0]
// => count: 0

// $ go run ch2/ext_2_4.go 255
// => bit: [1, 1, 1, 1, 1, 1, 1, 1]
// => count: 8
package main

import (
	"fmt"
	"os"
	"strconv"
)

var pc [8]byte

func PopCount() int {
	count := 0
	for i := 0; i < len(pc); i++ {
		num := pc[i]
		if num == 1 {
			count++
		}
	}
	return count
}

func main() {
	arg := os.Args[1]
	argInt, err := strconv.Atoi(arg)
	if err != nil {
		os.Exit(1)
	}
	index := 0
	for n := argInt; n > 0; n = int(byte(n >> 1)) {
		pc[index] = byte(n&1)
		index++
	}
	fmt.Println("bit: ", pc)
	count := PopCount()
	fmt.Println("count: ", count)
}
