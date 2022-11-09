package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(GenerateRandomChars(30))
	fmt.Println(GenerateRandomChars(30))
	fmt.Println(GenerateRandomChars(30))
	fmt.Println(GenerateRandomChars(30))
	fmt.Println(GenerateRandomChars(30))
	fmt.Println(GenerateRandomChars(30))
	fmt.Println(GenerateRandomChars(30))
	fmt.Println(GenerateRandomChars(30))
	fmt.Println(GenerateRandomChars(30))
}

// GenerateRandomChars Generates a random string with only charts
func GenerateRandomChars(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return strings.ToLower(string(b))
}
