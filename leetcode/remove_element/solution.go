package remove_element

func removeElement(nums []int, val int) int {
	cnt := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(nums) - cnt
}
