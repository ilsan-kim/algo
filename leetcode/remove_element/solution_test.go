package remove_element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
		res := removeElement(nums, 2)
		assert.Equal(t, 5, res)
	})

	t.Run("case 2", func(t *testing.T) {
		nums := []int{3, 2, 2, 3}
		res := removeElement(nums, 3)
		assert.Equal(t, 2, res)
	})
}
