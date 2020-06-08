package main

import (
	"fmt"
	"math"
)

func main() {
	//result := isPalindrome(121)
	//result := isPalindrome(-121)
	result := isPalindrome(10)
	fmt.Print(result)
}

// 将数字整体反转，但是这样可能会出现数字越界的问题，即反转后的数据超过了int的最大值
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	var y int = x
	var i int = 1
	var list []int = make([]int, 0)
	for y >= 10 {
		item := y % 10
		pow := int(math.Pow10(i))
		y = x / pow
		list = append(list, item)
		i++
	}
	if y != 0 {
		list = append(list, y)
	}

	var result int = 0
	for j := 0; j < len(list); j++ {
		pow := int(math.Pow10(len(list) - j - 1))
		item := list[j]
		result += item * pow
	}

	return result == x
}
