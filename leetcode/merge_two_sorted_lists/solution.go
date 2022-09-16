package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Expr() string {
	if l == nil {
		return "[]"
	}

	lst := []string{}
	next := l
	for {
		strVal := strconv.Itoa(next.Val)
		lst = append(lst, strVal)
		if next.Next == nil {
			break
		}
		next = next.Next
	}
	res := strings.Join(lst, ", ")
	return "[" + res + "]"
}

func solution(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	var startNode = &ListNode{}
	var tempNode = startNode
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tempNode.Next = list1
			list1 = list1.Next
		} else {
			tempNode.Next = list2
			list2 = list2.Next
		}
		tempNode = tempNode.Next
	}

	if list1 != nil {
		tempNode.Next = list1
	} else {
		tempNode.Next = list2
	}
	return startNode.Next
}

func main() {
	node_1_3 := ListNode{Val: 4, Next: nil}
	node_1_2 := ListNode{Val: 2, Next: &node_1_3}
	node_1_1 := ListNode{Val: 1, Next: &node_1_2}
	node_2_3 := ListNode{Val: 4, Next: nil}
	node_2_2 := ListNode{Val: 3, Next: &node_2_3}
	node_2_1 := ListNode{Val: 1, Next: &node_2_2}
	a := solution(&node_1_1, &node_2_1)
	fmt.Println(a.Expr())
}
