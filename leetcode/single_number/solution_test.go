package single_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		nums := []int{2, 2, 1}
		actual := solution(nums)
		expected := 1
		assert.Equal(t, expected, actual)
	})

	t.Run("case 2", func(t *testing.T) {
		nums := []int{4, 1, 2, 1, 2}
		actual := solution(nums)
		expected := 4
		assert.Equal(t, expected, actual)
	})

}
