package main

import "log"

func trap(height []int) int {
	maxFunc := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	if len(height) == 0 {
		return 0
	}

	volume := 0
	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]

	for left < right {
		leftMax, rightMax = maxFunc(height[left], leftMax), maxFunc(height[right], rightMax)
		if leftMax <= rightMax {
			volume += leftMax - height[left]
			left += 1
		} else {
			volume += rightMax - height[right]
			right -= 1
		}
	}
	return volume
}

func main() {
	a := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	log.Println(trap(a))
}