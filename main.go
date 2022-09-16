package main

import (
	"fmt"
	"log"
)

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2

		if nums[mid] < target {
			fmt.Printf("mid > %v , nums[mid] > %v :::::: nums[mid] < target\n", mid, nums[mid])
			low = mid + 1
			fmt.Printf("low changed > %v (high > %v)\n", low, high)
		} else if nums[mid] > target {
			fmt.Printf("mid > %v , nums[mid] > %v :::::: nums[mid] < target\n", mid, nums[mid])
			high = mid - 1
			fmt.Printf("high changed > %v (low > %v)\n", high, low)
		} else {
			return mid
		}
	}

	if low == len(nums) || nums[low] != target {
		if low == len(nums) {
			return low
		}
		if target > nums[low] {
			return low + 1
		} else {
			return low
		}
	} else {
		return low
	}
}

func main() {
	//nums := []int{-1, 0, 3, 5, 9, 12}
	//target := 10
	nums := []int{1, 3, 5, 6}
	target := 7
	a := searchInsert(nums, target)
	log.Println(a)
}
