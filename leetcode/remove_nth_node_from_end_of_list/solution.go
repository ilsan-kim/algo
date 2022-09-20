package remove_nth_node_from_end_of_list

import (
	"strconv"
	"strings"
)

/* https://leetcode.com/problems/remove-nth-node-from-end-of-list/ */

type ListNode struct {
	Val  int
	Next *ListNode
}

func solution(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}

	// find length
	length := 1
	next := head.Next
	for next != nil {
		next = next.Next
		length++
	}

	// if delete first node
	if length-n == 0 {
		head = head.Next
		//fmt.Println(head)
		return head
	}

	// change [n-1]th nodes next
	node := head
	for i := 0; i < length-n-1; i++ {
		node = node.Next
	}
	node.Next = node.Next.Next

	//fmt.Println(head)
	return head
}

func (l *ListNode) String() string {
	// for check result
	array := []string{}
	node := l
	for node != nil {
		msg := strconv.Itoa(node.Val)
		array = append(array, msg)
		node = node.Next
	}
	str := strings.Join(array, ",")
	return str
}
