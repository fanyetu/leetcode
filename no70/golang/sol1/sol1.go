package main

import "fmt"

func main() {
	fmt.Print(climbStairs(4))
}

/*
思路1：动态规划算法，将大问题分解为小问题
这里i级台阶的走法等于i-1级台阶的走法+i-2级台阶的走法
*/
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	dep1 := 2
	dep2 := 1
	for i := 2; i < n; i++ {
		cur := dep1 + dep2
		dep2 = dep1
		dep1 = cur
	}
	return dep1
}
