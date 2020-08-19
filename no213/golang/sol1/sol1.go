package main

import "fmt"

func main() {
	//nums := []int{2, 7, 9, 3, 1}
	//nums := []int{2,3,2}
	nums := []int{1, 2, 3, 1}
	fmt.Print(rob(nums))
}

/*
思路：使用动态规划算法，将大问题分解为子问题
这个题目实际上就是198的升级版，只需要计算0——n-1和1——n的最大值即可
*/
func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	return max(robArr(nums[0:len(nums)-1]), robArr(nums[1:len(nums)]))
}

func robArr(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	pd1 := max(nums[0], nums[1])
	pd2 := nums[0]
	for i := 2; i < len(nums); i++ {
		pd1, pd2 = max(pd2+nums[i], pd1), pd1
	}
	return pd1
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
