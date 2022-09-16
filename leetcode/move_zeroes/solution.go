package main

import "fmt"

// https://leetcode.com/problems/move-zeroes/
func solution(nums []int) {
	leftIdx := 0
	rightIdx := len(nums) - 1

	//for leftIdx < rightIdx && (nums[leftIdx] != 0 || nums[rightIdx] != 0) {
	for leftIdx < rightIdx {
		if nums[leftIdx] == 0 {
			nums = append(nums[:leftIdx], nums[leftIdx+1:]...)
			nums = append(nums, 0)
		} else {
			leftIdx++
		}

		if nums[rightIdx] == 0 {
			rightIdx--
		}
		fmt.Println(leftIdx, rightIdx, nums)
	}
}

func main() {
	nums := []int{0, 0, 1}
	//nums := []int{1, 0}
	//nums := []int{0}
	//nums := []int{0, 1, 0}
	//nums := []int{0, 1, 0, 2, 7, 212, 4, 2, 0, 12, 445, 0, 2}
	solution(nums)
}
