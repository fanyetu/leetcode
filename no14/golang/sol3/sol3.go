package main

import (
	"fmt"
	"strings"
)

func main() {
	//strs := []string{"flower", "flow", "flight"}
	//strs := []string{"dog","racecar","car"}
	strs := []string{"aa", "a"}
	//strs := []string{"c", "acc", "ccc"}
	//strs := []string{"c", "c"}
	s := longestCommonPrefix(strs)
	fmt.Print(s)
}

/*
分治法
*/
func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) < 1 {
		return ""
	}
	return longestCommonPrefixIn(strs, 0, len(strs)-1)
}

func longestCommonPrefixIn(strs []string, start int, end int) string {
	if start >= end {
		return strs[start]
	}

	mid := (start + end) / 2
	left := longestCommonPrefixIn(strs, start, mid)
	right := longestCommonPrefixIn(strs, mid+1, end)
	return commonPrefix(left, right)
}

func commonPrefix(left string, right string) string {
	index := len(left)
	for ; index > 0; index-- {
		sub := left[0:index]
		if strings.Index(right, sub) == 0 {
			return sub
		}
	}
	return ""
}
