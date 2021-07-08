package fullbag

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 0; i < len(s)+1; i++ {
		for j := 0; j < len(wordDict); j++ {
			if i < len(wordDict[j]) {
				continue
			}

			w := s[i-len(wordDict[j]) : i]
			if w == wordDict[j] && dp[i-len(wordDict[j])] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
