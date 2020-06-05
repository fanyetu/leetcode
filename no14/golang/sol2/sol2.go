package main

import (
	"fmt"
	"strings"
)

func main() {
	//strs := []string{"flower", "flow", "flight"}
	//strs := []string{"dog","racecar","car"}
	strs := []string{"aa", "a"}
	s := longestCommonPrefix(strs)
	fmt.Print(s)
}

/*
https://leetcode-cn.com/problems/longest-common-prefix/solution/zui-chang-gong-gong-qian-zhui-by-leetcode/
*/
func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) < 1 {
		return ""
	}

	var prefix string = strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) < 0 {
			prefix = prefix[0 : len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}
	}
	return prefix
}
