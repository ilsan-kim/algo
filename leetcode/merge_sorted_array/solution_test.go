package merge_sorted_array

import "testing"

func TestSolution(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		nums1 := []int{1, 2, 3, 0, 0, 0}
		m := 3
		nums2 := []int{2, 5, 6}
		n := 3
		solution(nums1, m, nums2, n)
	})

	t.Run("case 2", func(t *testing.T) {
		nums1 := []int{1}
		m := 1
		nums2 := []int{}
		n := 0
		solution(nums1, m, nums2, n)
	})

	t.Run("case 3", func(t *testing.T) {
		nums1 := []int{0}
		m := 0
		nums2 := []int{1}
		n := 1
		solution(nums1, m, nums2, n)
	})
}
