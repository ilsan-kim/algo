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

func solution(head *ListNode) *ListNode {
	var rootNode *ListNode = nil

	for head != nil {
		current := head
		head = head.Next
		current.Next = rootNode
		rootNode = current
	}

	return rootNode
}

func main() {
	node_1_3 := ListNode{Val: 4, Next: nil}
	node_1_2 := ListNode{Val: 2, Next: &node_1_3}
	node_1_1 := ListNode{Val: 1, Next: &node_1_2}
	a := solution(&node_1_1)
	fmt.Println(a)
}
