package functions

import (
	"regexp"
	"strings"
)

func LevenshteinDistance(s1, s2 string) int {
	m := len(s1)
	n := len(s2)

	// Create a 2D matrix to store the distances
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// Initialize the first row and column
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// Calculate the distances
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}

	// Return the distance
	return dp[m][n]
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	} else {
		return c
	}
}

func CountIncludedWords(s1, s2 string) int {
	count := 0
	words := strings.Fields(s1)
	for _, word := range words {
		if strings.Contains(s2, word) {
			count++
		}
	}
	return count
}

func ExactPhrase(s1, s2 string) bool {
	regex := regexp.MustCompile(`\b` + regexp.QuoteMeta(s1) + `\b`)
	return regex.MatchString(s2)
}
