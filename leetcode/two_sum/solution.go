package two_sum

/* https://leetcode.com/problems/two-sum/submissions/ */

func solution(nums []int, target int) []int {
	answer := []int{}
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				answer = append(answer, i)
				answer = append(answer, j)
				return answer
			}
		}
	}
	return answer
}
