package main

import "fmt"

func main() {
	//s := "III"
	//s := "IV"
	s := "IX"
	result := romanToInt(s)
	fmt.Print(result)
}

func romanToInt(s string) int {
	var (
		smap map[string]int = map[string]int{
			"I": 1,
			"V": 5,
			"X": 10,
			"L": 50,
			"C": 100,
			"D": 500,
			"M": 1000,
		}
		result int = 0
	)

	for i := 0; i < len(s); i++ {
		x := smap[s[i:i+1]]
		if i < len(s)-1 {
			y := smap[s[i+1:i+2]]
			if y > x {
				x = -x
			}
		}

		result = result + x
	}
	return result
}
