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

func PopCount(pointers *[8]byte) int {
	count := 0
	for i := 0; i < len(pointers); i++ {
		num := pointers[i]
		if byte(num&1) == 1 {
			count++
		}
	}
	return count
}

func CreateByteArray(number int) [8]byte {
	var pc [8]byte
	index := 0
	for n := number; n > 0; n = int(byte(n >> 1)) {
		pc[index] = byte(n&1)
		index++
	}
	return pc
}

func main() {
	arg := os.Args[1]
	argInt, err := strconv.Atoi(arg)
	if err != nil {
		os.Exit(1)
	}
	pc := CreateByteArray(argInt)
	fmt.Println("bit: ", pc)
	count := PopCount(&pc)
	fmt.Println("count: ", count)
}
