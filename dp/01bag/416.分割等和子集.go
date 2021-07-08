package zeroonebag

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 0 {
		sum /= 2
	} else {
		return false
	}

	dp := make([]int, sum+1)
	for i := 0; i < len(nums); i++ {
		for j := sum; j >= nums[i]; j-- {
			if dp[j] < dp[j-nums[i]]+nums[i] {
				dp[j] = dp[j-nums[i]] + nums[i]
			}

			if dp[j] == sum {
				return true
			}
		}
	}
	return false
}
