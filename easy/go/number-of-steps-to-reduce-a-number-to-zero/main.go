package main

import "fmt"

func main() {
	fmt.Println(numberOfSteps(14))
	fmt.Println(numberOfSteps(8))
	fmt.Println(numberOfSteps(123))
	fmt.Println(numberOfSteps(11))
}

func numberOfSteps(num int) int {
	steps := 0
	for num != 0 {
		if num%2 == 0 {
			num = num / 2
			steps = steps + 1
		} else {
			num = num - 1
			steps = steps + 1
		}
	}
	return steps
}
