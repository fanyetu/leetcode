package main

import "fmt"

func main() {
	result := isPalindrome(121)
	//result := isPalindrome(-121)
	//result := isPalindrome(10)
	fmt.Print(result)
}

// 考虑只用比较数字的前半部分和后半部分的反转数是否相同即可。如果数字长度为奇数，那么就删掉中间数比较即可
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var revertedNumber = 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x = x / 10
	}

	return x == revertedNumber || x == revertedNumber/10
}
