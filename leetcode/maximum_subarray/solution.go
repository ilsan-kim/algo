package maximum_subarray

import "math"

func solution(nums []int) int {
	maxFunc := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	bestSum := math.MinInt
	localSum := 0

	for _, v := range nums {
		localSum = maxFunc(localSum+v, v)
		bestSum = maxFunc(localSum, bestSum)
	}
	return bestSum
}
