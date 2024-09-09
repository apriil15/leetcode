package main

func main() {

}

// time: O(len(s) * len(wordDict))
// space: O(len(s))
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[len(dp)-1] = true

	for i := len(dp) - 1; i >= 0; i-- {
		for _, word := range wordDict {
			if i+len(word)-1 <= len(s)-1 && s[i:i+len(word)] == word {
				dp[i] = dp[i+len(word)]
			}
			if dp[i] {
				break
			}
		}
	}
	return dp[0]
}
