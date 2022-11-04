package reverse_vowels_of_a_string

func solution(s string) string {
	// TODO two pointer로 풀어야하네.........
	// 양끝에서 시작해서 만날때마다 자리 바꿔주면 더 빠를듯
	res := ""
	vowelMap := map[string]bool{"a": true, "e": true, "i": true, "o": true, "u": true, "A": true, "E": true, "I": true, "O": true, "U": true}
	numVowel := 0
	vowels := []string{}

	for i := range s {
		if ok := vowelMap[string(s[i])]; ok {
			vowels = append(vowels, string(s[i]))
			numVowel++
		}
	}

	// reverse
	for i, j := 0, len(vowels)-1; i < j; i, j = i+1, j-1 {
		vowels[i], vowels[j] = vowels[j], vowels[i]
	}

	// make new string
	c := 0
	for i := range s {
		if ok := vowelMap[string(s[i])]; ok {
			res += vowels[c]
			c++
		} else {
			res += string(s[i])
		}
	}

	return res
}
