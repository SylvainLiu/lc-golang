package basic

func integerBreak(n int) int {
	dp := make([]int, n)
	dp[0] = 1

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if (j+1)*(i-j) > dp[j]*(i-j) && (j+1)*(i-j) > dp[i] {
				dp[i] = (j + 1) * (i - j)
			} else if dp[j]*(i-j) > dp[i] {
				dp[i] = dp[j] * (i - j)
			}
		}
	}

	return dp[n-1]
}
