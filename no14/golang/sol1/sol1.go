package main

import "fmt"

func main() {
	strs := []string{"flower", "flow", "flight"}
	//strs := []string{"dog","racecar","car"}
	//strs := []string{"aa", "a"}
	s := longestCommonPrefix(strs)
	fmt.Print(s)
}

/*
想象数组的末尾有一个非常短的字符串，使用上述方法依旧会进行 S​​ 次比较。优化这类情况的一种方法就是水平扫描。我们从前往后枚举字符串的每一列，先比较每个字符串相同列上的字符（即不同字符串相同下标的字符）然后再进行对下一列的比较。

复杂度分析

时间复杂度：O(S)，S 是所有字符串中字符数量的总和。

最坏情况下，输入数据为 nn 个长度为 mm 的相同字符串，算法会进行 S = m*n 次比较。可以看到最坏情况下，本算法的效率与算法一相同，但是最好的情况下，算法只需要进行 n*minLen 次比较，其中 minLen 是数组中最短字符串的长度。

空间复杂度：O(1)，我们只需要使用常数级别的额外空间。

作者：LeetCode
链接：https://leetcode-cn.com/problems/longest-common-prefix/solution/zui-chang-gong-gong-qian-zhui-by-leetcode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) < 1 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i : i+1]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || c != strs[j][i:i+1] {
				return strs[0][0:i]
			}
		}
	}
	return strs[0]
}
