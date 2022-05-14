// $ go run ch3/ext_3_12.go ABCB CBBA
package main

import (
	"bytes"
	"os"
	"fmt"
)

func isAnagram(byteA []byte, byteB []byte) bool {
	for {
		if len(byteA) == 0 || len(byteB) == 0 {
			break
		}
		emptyByte := []byte{}
		splitByte := []byte{byteA[0]}

		splitA := bytes.Split(byteA, splitByte)
		splitB := bytes.Split(byteB, splitByte)

		joinA := bytes.Join(splitA, emptyByte)
		joinB := bytes.Join(splitB, emptyByte)

		byteA = joinA
		byteB = joinB
	}

	if len(byteA) == 0 && len(byteB) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	a := os.Args[1]
	b := os.Args[2]

	// different character size
	if len(a) != len(b) {
		fmt.Println("not anagram")
		os.Exit(0)
	}

	byteA := []byte(a)
	byteB := []byte(b)

	comp := bytes.Compare(byteA, byteB)
	// exactly the same
	if comp == 0 {
		fmt.Println("not anagram")
		os.Exit(0)
	}

	if isAnagram(byteA, byteB) {
		fmt.Println("anagram")
	} else {
		fmt.Println("not anagram")
	}
}
