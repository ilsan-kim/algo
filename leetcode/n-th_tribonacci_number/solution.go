package n_th_tribonacci_number

/* https://leetcode.com/problems/n-th-tribonacci-number/?envType=study-plan&id=dynamic-programming-i */
func tribonacci(n int) int {
	memo := make([]int, n+3)

	memo[0] = 0
	memo[1] = 1
	memo[2] = 1

	for i := 3; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2] + memo[i-3]
	}

	return memo[n]
}
