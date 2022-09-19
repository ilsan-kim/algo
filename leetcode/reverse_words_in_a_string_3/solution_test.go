package reverse_words_in_a_string_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		input := "Let's take LeetCode contest"
		expected := "s'teL ekat edoCteeL tsetnoc"
		actual := solution(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("case 2", func(t *testing.T) {
		input := "God Ding"
		expected := "doG gniD"
		actual := solution(input)
		assert.Equal(t, expected, actual)
	})
}
