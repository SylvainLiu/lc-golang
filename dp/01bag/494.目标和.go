package zeroonebag

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if (sum+target)%2 != 0 {
		return 0
	}
	tar := (sum + target) / 2

	dp := make([]int, tar+1)
	dp[0] = 1

	for i := 0; i < len(nums); i++ {
		for j := tar; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[tar]
}
