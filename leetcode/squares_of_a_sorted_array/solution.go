package main

import (
	"log"
	"sort"
)

// https://leetcode.com/problems/squares-of-a-sorted-array/
func solution(nums []int) []int {
	res := make([]int, len(nums))
	for _, num := range nums {
		res = append(res, num*num)
	}
	sort.Ints(res)
	return res
}

func main() {
	a := []int{-4, -1, 0, 3, 10}
	b := solution(a)
	log.Println(b)
}
