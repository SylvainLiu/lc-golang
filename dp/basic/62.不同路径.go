package basic

func uniquePaths(m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[1][0] = 1

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m][n]
}
