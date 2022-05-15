// 合ってるか分からないからあとで見直し
// $ go run ch4/ext_4_3.go
package main

import (
	"fmt"
)

func reverse(s *[4]int) [4]int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return *s
}

func main () {
	s := [...]int{0, 1, 2, 3}
	fmt.Println(reverse(&s))
}
