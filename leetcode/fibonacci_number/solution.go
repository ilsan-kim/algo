package fibonacci_number

/* https://leetcode.com/problems/fibonacci-number/?envType=study-plan&id=dynamic-programming-i */
func fib(n int) int {
	memo := make([]int, n+2)

	memo[0] = 0
	memo[1] = 1

	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}

	return memo[n]
}
