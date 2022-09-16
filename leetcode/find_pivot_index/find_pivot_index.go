package main

import "log"

func bruteForce(nums []int) int {
	sum := func(arr []int) int {
		sum := 0
		for _, v := range arr {
			sum += v
		}
		return sum
	}

	for i := range nums {
		if i == 0 {
			if sum(nums[i+1:]) == 0 {
				return i
			}
		} else {
			if sum(nums[0:i]) == sum(nums[i+1:]) {
				return i
			}
		}
	}

	return -1
}

func faster(nums []int) int {
	sum := func(arr []int) int {
		sum := 0
		for _, v := range arr {
			sum += v
		}
		return sum
	}

	totalSum := sum(nums)
	leftSum := 0

	for i := range nums {
		log.Printf("leftSum(%v) == totalSum(%v)-leftSum(%v)-me(%v)", leftSum, totalSum, leftSum, nums[i])
		if !(leftSum == totalSum-leftSum-nums[i]) {
			leftSum += nums[i]
			log.Println(i, leftSum)
		} else {
			return i
		}
	}

	return -1
}

func main() {
	nums := []int{1, 2, 3}
	a := bruteForce(nums)
	b := faster(nums)
	log.Println(a)
	log.Println(b)

}
