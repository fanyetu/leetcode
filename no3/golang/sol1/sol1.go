package main

import "fmt"

func main() {
	//var s string = "abcabcbb"
	var s string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\\\"#$%&'()*+,-./:;<=>?@[\\\\]^_`{|}~"
	//var s string = "bbbbb"
	//var s string = "pwwkew"
	var result = lengthOfLongestSubstring(s)
	fmt.Print(result)
}

// 循环过多，时间超长了
func lengthOfLongestSubstring(s string) int {
	var (
		maxWinSize int = len(s)
		winSize    int = maxWinSize
		result     int = 0
	)

	for ; winSize >= 1; winSize-- {
		f := true
		for i := 0; i <= len(s)-winSize; i++ {
			subStr := s[i : i+winSize]
			flag := true
			smap := make(map[uint8]int)
			for j := 0; j < len(subStr); j++ {
				item := subStr[j]
				if _, ok := smap[item]; ok {
					flag = false
					break
				}
				smap[item] = 0
			}

			smap = make(map[uint8]int)
			if flag {
				f = true
				break
			} else {
				f = false
			}
		}

		if f {
			result = winSize
			break
		}
	}

	return result
}
