package main

import "fmt"

func solution(survey []string, choices []int) string {
	points := map[string]int{
		"R": 0,
		"T": 0,
		"C": 0,
		"F": 0,
		"J": 0,
		"M": 0,
		"A": 0,
		"N": 0,
	}
	res := ""

	for i, v := range survey {
		if choices[i] < 4 {
			points[string(v[0])] += 4 - choices[i]
		} else if choices[i] > 4 {
			points[string(v[1])] += choices[i] - 4
		}
	}

	if points["T"] > points["R"] {
		res += "T"
	} else {
		res += "R"
	}

	if points["F"] > points["C"] {
		res += "F"
	} else {
		res += "C"
	}

	if points["M"] > points["J"] {
		res += "M"
	} else {
		res += "J"
	}

	if points["N"] > points["A"] {
		res += "N"
	} else {
		res += "A"
	}

	return res
}

func main() {
	s := []string{"AN", "CF", "MJ", "RT", "NA"}
	c := []int{5, 3, 2, 7, 5}

	r := solution(s, c)
	fmt.Println(r)
}
