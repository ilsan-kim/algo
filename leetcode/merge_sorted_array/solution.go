package merge_sorted_array

import (
	"fmt"
	"sort"
)

func solution(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[0:m]
	nums2 = nums2[0:n]
	arr := append(nums1, nums2...)
	sort.Ints(arr)
	fmt.Println(arr)
}
