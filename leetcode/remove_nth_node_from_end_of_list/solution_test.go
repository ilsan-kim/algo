package remove_nth_node_from_end_of_list

import (
	"testing"
)

func TestName(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		l5 := &ListNode{Val: 5, Next: nil}
		l4 := &ListNode{Val: 4, Next: l5}
		l3 := &ListNode{Val: 3, Next: l4}
		l2 := &ListNode{Val: 2, Next: l3}
		l1 := &ListNode{Val: 1, Next: l2}

		solution(l1, 2)
	})

	t.Run("case 2", func(t *testing.T) {
		l1 := &ListNode{Val: 1, Next: nil}

		solution(l1, 1)
	})

	t.Run("case 3", func(t *testing.T) {
		l2 := &ListNode{Val: 2, Next: nil}
		l1 := &ListNode{Val: 1, Next: l2}

		solution(l1, 1)
	})

	t.Run("case 4", func(t *testing.T) {
		l2 := &ListNode{Val: 2, Next: nil}
		l1 := &ListNode{Val: 1, Next: l2}

		solution(l1, 2)
	})
}
