package single_number

func solution(nums []int) int {
	maps := map[int]int{}
	for _, num := range nums {
		if _, ok := maps[num]; !ok {
			maps[num] = 1
		} else {
			maps[num]++
		}
	}

	for k := range maps {
		if maps[k] == 1 {
			return k
		}
	}
	return 0
}
