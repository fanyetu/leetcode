package main

import "fmt"

func main() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Print(rob(nums))
}

/*
思路1: 动态规划
分解子问题，偷第k个房子时的最大值有两种方案，他们的最大值就是dp[k]的最大值
dp[k] = max(dp[k-1], dp[k-2] + nums[k])
*/
func rob(nums []int) int {
	// 边界情况0, 1
	if nums == nil || len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp1 := nums[0]
	dp2 := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp1, dp2 = dp2, max(dp1+nums[i], dp2)
	}
	return dp2
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
