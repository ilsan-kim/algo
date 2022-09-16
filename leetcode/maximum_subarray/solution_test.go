package maximum_subarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
		actual := solution(nums)
		expected := 6
		assert.Equal(t, expected, actual)
	})
}
