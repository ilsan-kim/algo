package reverse_words_in_a_string_3

import "strings"

func solution(s string) string {
	reverse := func(s string) string {
		start := 0
		end := len(s) - 1

		sb := []byte(s)
		for start < end {
			sb[start], sb[end] = sb[end], sb[start]
			start++
			end--
		}

		return string(sb)
	}

	stringSlice := strings.Split(s, " ")
	for i := range stringSlice {
		stringSlice[i] = reverse(stringSlice[i])
	}

	res := strings.Join(stringSlice, " ")
	return res
}
