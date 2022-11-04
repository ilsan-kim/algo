package reverse_vowels_of_a_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		s := "hello"
		expected := "holle"
		actual := solution(s)
		assert.Equal(t, expected, actual)
	})

	t.Run("t2", func(t *testing.T) {
		s := "leetcode"
		expected := "leotcede"
		actual := solution(s)
		assert.Equal(t, expected, actual)
	})
}
