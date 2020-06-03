package main

import (
	"fmt"
	"math"
)

func main() {
	n1 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 4, Next: n1}
	n3 := &ListNode{Val: 2, Next: n2}

	l1 := &ListNode{Val: 4}
	l2 := &ListNode{Val: 6, Next: l1}
	l3 := &ListNode{Val: 5, Next: l2}

	addTwoNumbers(n3, l3)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		i   int       = 0
		n1  *ListNode = l1
		one float64

		j    int       = 0
		n2   *ListNode = l2
		two  float64
		plus float64 = 0
	)
	for {
		if n1 == nil {
			break
		}
		plus = math.Pow10(i)
		one += float64(n1.Val) * plus
		i++
		n1 = n1.Next
	}

	for {
		if n2 == nil {
			break
		}
		plus = math.Pow10(j)
		two += float64(n2.Val) * plus
		j++
		n2 = n2.Next
	}

	fmt.Println(one, two)

	result := int64(one + two)

	fmt.Print(result)
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}
