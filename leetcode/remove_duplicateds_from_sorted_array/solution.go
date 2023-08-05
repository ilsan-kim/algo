package remove_duplicateds_from_sorted_array

func removeDuplicates(nums []int) int {
	prev := nums[len(nums)-1]
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			continue
		}
		if prev == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
		}
		prev = nums[i]
	}
	return len(nums)
}
