package main

func bf(nums []int) []int {
	var res []int
	for i, v := range nums {
		if i == 0 {
			res = append(res, v)
		} else {
			sum := 0
			for _, k := range nums[0 : i+1] {
				sum += k
			}
			res = append(res, sum)
		}
	}
	return res
}

func faster(nums []int) []int {
	res := make([]int, len(nums))
	sum_before := 0

	for i, v := range nums {
		res[i] = v + sum_before
		sum_before = v + sum_before
	}
	return res
}

func main() {
	nums := []int{3, 1, 2, 10, 1}
	bf(nums)
	faster(nums)
}
