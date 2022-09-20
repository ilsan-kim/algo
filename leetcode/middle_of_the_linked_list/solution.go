package middle_of_the_linked_list

import "fmt"

/* https://leetcode.com/problems/middle-of-the-linked-list */

type ListNode struct {
	Val  int
	Next *ListNode
}

func solution(head *ListNode) *ListNode {
	length := 1
	next := head.Next
	for next != nil {
		next = next.Next
		length++
	}

	for i := 0; i < length/2; i++ {
		head = head.Next
	}

	fmt.Println(head.Val)
	return head
}
