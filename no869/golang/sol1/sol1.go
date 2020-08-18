package main

import (
	"fmt"
	"reflect"
)

func main() {
	var flag = reorderedPowerOf2(46)
	fmt.Print(flag)
}

/*
思路1：
计算N的从0到9的个数，比如1122233就有2个1、3个2、2个3
然后将N的个数和2的i次幂的个数相比较，如果相同，则其可以变换为2的i次幂
*/
func reorderedPowerOf2(N int) bool {
	var (
		originArr []int
		arr       []int
	)
	originArr = count(N)
	for i := 0; i < 32; i++ {
		arr = count(1 << i)
		if reflect.DeepEqual(originArr, arr) {
			return true
		}
	}
	return false
}

// 计算这个数字中每一个数字的数量
func count(N int) []int {
	var arr []int = make([]int, 10)
	for N > 0 {
		arr[N%10]++
		N = N / 10
	}
	return arr
}
