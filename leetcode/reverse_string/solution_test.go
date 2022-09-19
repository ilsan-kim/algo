package reverse_string

import "testing"

func TestSolution(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		input := []byte{1, 2, 3, 3, 6}
		solution(input)
	})

	t.Run("case 2", func(t *testing.T) {
		input := []byte{5, 4, 3, 6, 1, 2}
		solution(input)
	})
}
