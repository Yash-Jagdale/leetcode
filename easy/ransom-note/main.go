package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "aab"))
	fmt.Println(canConstruct("aab", "aab"))
	fmt.Println(canConstruct("yash", "hhysawant"))
}
func canConstruct(ransomNote string, magazine string) bool {

	for _, i := range ransomNote {
		if strings.Contains(magazine, string(i)) {
			magazine = strings.Replace(magazine, string(i), "", 1)
		} else {
			return false
		}
	}
	return true
}
