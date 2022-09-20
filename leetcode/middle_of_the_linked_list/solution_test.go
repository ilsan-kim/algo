package middle_of_the_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		l5 := &ListNode{Val: 5, Next: nil}
		l4 := &ListNode{Val: 4, Next: l5}
		l3 := &ListNode{Val: 3, Next: l4}
		l2 := &ListNode{Val: 2, Next: l3}
		l1 := &ListNode{Val: 1, Next: l2}

		actual := solution(l1)
		expected := l3
		assert.Equal(t, expected, actual)
	})

	t.Run("case 2", func(t *testing.T) {
		l6 := &ListNode{Val: 6, Next: nil}
		l5 := &ListNode{Val: 5, Next: l6}
		l4 := &ListNode{Val: 4, Next: l5}
		l3 := &ListNode{Val: 3, Next: l4}
		l2 := &ListNode{Val: 2, Next: l3}
		l1 := &ListNode{Val: 1, Next: l2}

		actual := solution(l1)
		expected := l4
		assert.Equal(t, expected, actual)
	})
}
