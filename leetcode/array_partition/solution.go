package array_partition

import (
	"sort"
)

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i, _ := range nums {
		if i%2 == 0 {
			sum += nums[i]
		}
	}
	return sum
}
