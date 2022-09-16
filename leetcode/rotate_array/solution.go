package main

import "fmt"

// https://leetcode.com/problems/rotate-array/
func solution(nums []int, k int) {
	t := k % len(nums)

	temp := make([]int, len(nums))
	for i, _ := range nums {
		temp[i] = nums[i]
	}
	a := temp[len(temp)-t:]
	b := temp[:len(temp)-t]

	for i := range nums {
		if i < t {
			fmt.Printf("nums[%v] = a[%v] // a[%v] = %v\n", i, i, i, a[i])
			nums[i] = a[i]
		} else {
			fmt.Printf("nums[%v] = b[%v] // b[%v] = %v\n", i, i-t, i-t, b[i-t])
			nums[i] = b[i-t]
		}
	}
}

func main() {
	//nums := []int{-1, -100, 3, 99}
	nums := []int{1, 2, 3, 4, 5, 6, 7}

	k := 3

	solution(nums, k)
}
