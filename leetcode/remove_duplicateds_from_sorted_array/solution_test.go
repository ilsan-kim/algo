package remove_duplicateds_from_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		res := removeDuplicates(nums)
		assert.Equal(t, 5, res)
	})

	t.Run("case 2", func(t *testing.T) {
		nums := []int{1, 1, 2}
		res := removeDuplicates(nums)
		assert.Equal(t, 2, res)
	})
}
