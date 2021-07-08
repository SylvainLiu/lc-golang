package robber

func rob2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		if nums[1] > nums[0] {
			return nums[1]
		} else {
			return nums[0]
		}
	}

	res1 := robrange(nums, 0, n-2)
	res2 := robrange(nums, 1, n-1)

	if res1 > res2 {
		return res1
	} else {
		return res2
	}
}

func robrange(nums []int, start, end int) int {
	dp := make([]int, len(nums))
	dp[start] = nums[start]

	if nums[start] > nums[start+1] {
		dp[start+1] = nums[start]
	} else {
		dp[start+1] = nums[start+1]
	}

	for i := start + 2; i <= end; i++ {
		c1 := dp[i-2] + nums[i]
		c2 := dp[i-1]
		if c1 > c2 {
			dp[i] = c1
		} else {
			dp[i] = c2
		}
	}

	return dp[end]
}
