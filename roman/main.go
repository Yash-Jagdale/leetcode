package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MMMCMXLV"))

}

func romanToInt(s string) int {
	var num = 0
	roman := getRomanToInt()

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if val, ok := withPrefix()[s[i:i+2]]; ok {
				num = num + val
				i++
				continue
			}
			num = num + roman[string(s[i])]
			continue
		}
		num = num + roman[string(s[i])]
	}
	return num
}

func getRomanToInt() map[string]int {
	return map[string]int{"I": 1, "V": 5, "X": 10, "L": 50,
		"C": 100,
		"D": 500,
		"M": 1000}
}

func withPrefix() map[string]int {
	return map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
}
