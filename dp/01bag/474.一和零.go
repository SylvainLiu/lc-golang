package zeroonebag

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		zero := 0
		one := 0

		for _, c := range str {
			if c == '0' {
				zero++
			} else {
				one++
			}
		}

		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				if dp[i-zero][j-one]+1 > dp[i][j] {
					dp[i][j] = dp[i-zero][j-one] + 1
				}
			}
		}
	}

	return dp[m][n]
}
