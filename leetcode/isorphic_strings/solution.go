package main

import (
	"log"
)

// https://leetcode.com/problems/isomorphic-strings/
func solution(s string, t string) bool {
	matched := map[string]string{}

	if len(s) != len(t) {
		return false
	}

	for i, v := range s {
		if matched[string(v)] == "" {
			matched[string(v)] = string(t[i])
		} else {
			//처음 넣는 문자열이 아닌경우
			if matched[string(v)] == string(t[i]) {
				continue
			} else {
				return false
			}
		}
	}

	count := map[string]int{}
	for _, v := range matched {
		if count[v] == 0 {
			count[v] += 1
		} else {
			return false
		}
	}
	return true
}

func main() {
	s := "tile"
	t := "tilt"
	a := solution(s, t)
	log.Println(a)
}
