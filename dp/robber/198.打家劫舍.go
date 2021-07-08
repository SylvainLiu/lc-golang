package robber

func rob1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]

	if nums[0] > nums[1] {
		dp[1] = nums[0]
	} else {
		dp[1] = nums[1]
	}

	for i := 2; i < len(nums); i++ {
		c1 := dp[i-2] + nums[i]
		c2 := dp[i-1]
		if c1 > c2 {
			dp[i] = c1
		} else {
			dp[i] = c2
		}
	}

	return dp[len(dp)-1]
}
