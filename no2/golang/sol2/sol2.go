package main

import "fmt"

func main() {
	list1 := genListNodes([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	list2 := genListNodes([]int{4, 6, 5})

	resp := addTwoNumbers2(list1, list2)

	printListNode(resp)
}
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		result  *ListNode = &ListNode{Val: 0}
		cur     *ListNode = result
		com     int
		n1      *ListNode = l1
		n2      *ListNode = l2
		v1      int       = 0
		v2      int       = 0
		sumVal  int       = 0
		nextVal int       = 0
	)

	for n1 != nil || n2 != nil {
		if n1 != nil {
			v1 = n1.Val
		} else {
			v1 = 0
		}

		if n2 != nil {
			v2 = n2.Val
		} else {
			v2 = 0
		}

		sumVal = v1 + v2 + com

		nextVal = sumVal % 10
		if sumVal > 9 {
			com = 1
		} else {
			com = 0
		}

		cur.Next = &ListNode{Val: nextVal}
		cur = cur.Next

		if n1 != nil {
			n1 = n1.Next
		}
		if n2 != nil {
			n2 = n2.Next
		}
	}
	if com > 0 {
		cur.Next = &ListNode{Val: com}
	}

	return result.Next
}

func printListNode(l *ListNode) {
	var (
		n *ListNode = l
	)
	fmt.Print("[")
	for n != nil {
		fmt.Print(n.Val)
		n = n.Next
		if n != nil {
			fmt.Print(",")
		}
	}
	fmt.Print("]")
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
