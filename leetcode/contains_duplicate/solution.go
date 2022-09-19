package contains_duplicate

func solution(nums []int) bool {
	counter := map[int]bool{}
	for _, v := range nums {
		if _, ok := counter[v]; !ok {
			counter[v] = true
		} else {
			return true
		}
	}
	return false
}
