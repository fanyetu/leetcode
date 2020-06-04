package main

import "fmt"

func main() {
	s := "III"
	//s := "IV"
	//s := "IX"
	result := romanToInt(s)
	fmt.Print(result)
}

func romanToInt(s string) int {
	var (
		sum    int = 0
		preNum int = getValue(s[0:1])
	)
	for i := 1; i < len(s); i++ {
		num := getValue(s[i : i+1])
		if preNum >= num {
			sum += preNum
		} else {
			sum -= preNum
		}
		preNum = num
	}

	sum += preNum

	return sum
}

func getValue(s string) int {
	switch s {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	}
	return 0
}
