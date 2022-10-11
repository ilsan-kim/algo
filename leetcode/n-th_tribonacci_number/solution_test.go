package n_th_tribonacci_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		expected := 1389537
		actual := tribonacci(25)
		assert.Equal(t, expected, actual)
	})
}
