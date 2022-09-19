package contains_duplicate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		nums := []int{1, 2, 3, 1}
		actual := solution(nums)
		assert.Equal(t, true, actual)
	})

	t.Run("case 2", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		actual := solution(nums)
		assert.Equal(t, false, actual)
	})

	t.Run("case 3", func(t *testing.T) {
		nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
		actual := solution(nums)
		assert.Equal(t, true, actual)
	})
}