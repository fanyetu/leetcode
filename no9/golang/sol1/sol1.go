package main

import (
	"fmt"
	"strconv"
)

func main() {
	//result := isPalindrome(121)
	result := isPalindrome(-121)
	fmt.Print(result)
}

// 字符串暴力解法
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	str = reverse(str)
	result, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return result == x
}

func reverse(s string) string {
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		result += s[i : i+1]
	}
	return result
}
