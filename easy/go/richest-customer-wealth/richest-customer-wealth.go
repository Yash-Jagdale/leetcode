package main

import "fmt"

func main() {
	fmt.Println(maximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))
	fmt.Println(maximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}}))
	fmt.Println(maximumWealth([][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}))
}

func maximumWealth(accounts [][]int) int {
	max := 0
	for _, account := range accounts {
		sum := 0
		for _, i := range account {
			sum = sum + i
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
