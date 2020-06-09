package main

import "fmt"

func main() {
	//s := "babad"
	s := "cbbd"
	fmt.Print(longestPalindrome(s))
}

/*
中心扩展算法
*/
func longestPalindrome(s string) string {
	var (
		maxLen int = 1
		result string
	)
	if len(s) < 2 {
		return s
	}

	result = s[0:1]

	for i := 0; i < len(s)-1; i++ {
		oddStr := centerStr(s, i, i)
		evenStr := centerStr(s, i, i+1)
		maxLenStr := oddStr
		if len(evenStr) > len(oddStr) {
			maxLenStr = evenStr
		}

		if len(maxLenStr) > maxLen {
			maxLen = len(maxLenStr)
			result = maxLenStr
		}
	}

	return result
}

func centerStr(s string, left int, right int) string {
	// left==right，则中心是字符，回文长度是奇数
	// left+1==right，则中心是空字符，回文长度是偶数
	var (
		i int = left
		j int = right
	)

	for i >= 0 && j < len(s) {
		if s[i:i+1] == s[j:j+1] {
			i--
			j++
		} else {
			break
		}
	}

	return s[i+1 : j]
}
