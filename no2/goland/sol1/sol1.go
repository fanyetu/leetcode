package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func main() {
	list1 := genListNodes([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	list2 := genListNodes([]int{4, 6, 5})

	resp := addTwoNumbers(list1, list2)

	fmt.Print(resp)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		n1  *ListNode = l1
		one string

		n2  *ListNode = l2
		two string

		resp *ListNode
	)
	for {
		if n1 == nil {
			break
		}
		one = strconv.Itoa(n1.Val) + one
		n1 = n1.Next
	}

	for {
		if n2 == nil {
			break
		}
		two = strconv.Itoa(n2.Val) + two
		n2 = n2.Next
	}

	// 即使写出来也没用，这不是算法应该采用的做法
	fmt.Println(one, two)
	var onei *big.Int = big.NewInt(0)
	var twoi *big.Int = big.NewInt(0)
	onei.SetString(one, 0)
	twoi.SetString(two, 0)
	var result *big.Int = big.NewInt(0).Add(onei, twoi)

	length := len(result.String())

	for k := length - 1; k >= 0; k-- {
		m := int64(math.Pow10(k))
		o := big.NewInt(0).Div(result, big.NewInt(m))
		result = big.NewInt(0).Sub(result, big.NewInt(0).Mul(o, big.NewInt(m)))

		resp = &ListNode{Val: int(o.Int64()), Next: resp}
	}

	return resp
}

func genListNodes(arr []int) *ListNode {
	var result *ListNode
	for _, item := range arr {
		result = &ListNode{
			Val:  item,
			Next: result,
		}
	}
	return result
}

type ListNode struct {
	Val  int
	Next *ListNode
}
