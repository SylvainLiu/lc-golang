package basic

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 1 {
		return cost[0]
	}

	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		if dp[i-1] < dp[i-2] {
			dp[i] = dp[i-1] + cost[i]
		} else {
			dp[i] = dp[i-2] + cost[i]
		}
	}

	if dp[len(dp)-1] < dp[len(dp)-2] {
		return dp[len(dp)-1]
	} else {
		return dp[len(dp)-2]
	}
}
