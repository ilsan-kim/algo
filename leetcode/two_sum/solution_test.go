package two_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		nums := []int{2, 7, 11, 15}
		target := 9

		actual := solution(nums, target)
		expected := []int{0, 1}

		assert.Equal(t, expected, actual)
	})

	t.Run("case 2", func(t *testing.T) {
		nums := []int{3, 2, 4}
		target := 6

		actual := solution(nums, target)
		expected := []int{1, 2}

		assert.Equal(t, expected, actual)
	})

	t.Run("case 3", func(t *testing.T) {
		nums := []int{3, 3}
		target := 6

		actual := solution(nums, target)
		expected := []int{0, 1}

		assert.Equal(t, expected, actual)
	})
}
