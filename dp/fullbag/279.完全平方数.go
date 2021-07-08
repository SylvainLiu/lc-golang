package fullbag

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i*i <= n; i++ {
		for j := i * i; j < n+1; j++ {
			if dp[j] > dp[j-i*i]+1 {
				dp[j] = dp[j-i*i] + 1
			}
		}
	}

	return dp[n]
}
