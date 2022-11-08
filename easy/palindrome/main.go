package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(12210))
}

func isPalindrome(x int) bool {
	inputNum := x
	res := 0
	for x > 0 {
		remainder := x % 10
		res = (res * 10) + remainder
		x = x / 10
	}
	return res == inputNum
}
