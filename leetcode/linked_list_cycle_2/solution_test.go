package linked_list_cycle_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		node4 := &ListNode{Val: -4, Next: nil}
		node3 := &ListNode{Val: 0, Next: node4}
		node2 := &ListNode{Val: 2, Next: node3}
		node1 := &ListNode{Val: 3, Next: node2}
		node4.Next = node2

		expected := node2
		actual := solution(node1)

		assert.Equal(t, expected, actual)
	})

	t.Run("case 2", func(t *testing.T) {
		node2 := &ListNode{Val: 2, Next: nil}
		node1 := &ListNode{Val: 1, Next: node2}
		node2.Next = node1

		expected := node1
		actual := solution(node1)

		assert.Equal(t, expected, actual)
	})

	t.Run("case 3", func(t *testing.T) {
		node1 := &ListNode{Val: 1, Next: nil}

		var expected *ListNode = nil
		actual := solution(node1)

		assert.Equal(t, expected, actual)
	})
}
