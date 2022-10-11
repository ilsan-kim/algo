package fibonacci_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		actual := fib(4)
		expected := 3
		assert.Equal(t, expected, actual)
	})
}
