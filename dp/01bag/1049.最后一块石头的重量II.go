package zeroonebag

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, s := range stones {
		sum += s
	}
	target := sum / 2

	dp := make([]int, target+1)
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			if dp[j] < dp[j-stones[i]]+stones[i] {
				dp[j] = dp[j-stones[i]] + stones[i]
			}
		}
	}

	return (sum - dp[target]) - dp[target]
}
