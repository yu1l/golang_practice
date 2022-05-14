package main

import (
	"os"
	"fmt"
)

func isAnagram(first string, second string) bool {
	// 本来はbytes.Buffer等を使う。
	aChars := make(map[string]int)
	bChars := make(map[string]int)
	for _, r := range first {
		aChars[string(r)]++
	}
	for _, r := range second {
		bChars[string(r)]++
	}
	for key, value := range aChars {
		if bChars[key] != value {
			return false
		}
	}
	return true
}

func main() {
	a := os.Args[1]
	b := os.Args[2]
	if a == b {
		fmt.Println("not anagram")
		os.Exit(0)
	}
	if len(a) != len(b) {
		fmt.Println("not anagram")
		os.Exit(0)
	}

	if isAnagram(a, b) {
		fmt.Println("anagram")
	} else {
		fmt.Println("not anagram")
	}
}
