package homework

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		plans := [][]string{{"korean", "11:40", "30"}, {"english", "12:10", "20"}, {"math", "12:30", "40"}}
		expected := []string{"korean", "english", "math"}
		actual := solution(plans)
		assert.Equal(t, expected, actual)
	})

	t.Run("case 2", func(t *testing.T) {
		plans := [][]string{{"science", "12:40", "50"}, {"music", "12:20", "40"}, {"history", "14:00", "30"}, {"computer", "12:30", "100"}}
		expected := []string{"science", "history", "computer", "music"}
		actual := solution(plans)
		assert.Equal(t, expected, actual)
	})

	t.Run("case 3", func(t *testing.T) {
		plans := [][]string{{"aaa", "12:00", "20"}, {"bbb", "12:10", "30"}, {"ccc", "12:40", "10"}}
		expected := []string{"bbb", "ccc", "aaa"}
		actual := solution(plans)
		assert.Equal(t, expected, actual)
	})

	t.Run("case 4", func(t *testing.T) {
		plans := [][]string{{"aaa", "12:00", "20"}, {"bbb", "12:10", "30"}, {"ccc", "12:40", "10"}, {"ddd", "13:55", "60"}}
		expected := []string{"bbb", "ccc", "aaa", "ddd"}
		actual := solution(plans)
		assert.Equal(t, expected, actual)
	})
}
