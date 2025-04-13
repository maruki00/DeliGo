package pkgUtils

import (
	"math"
)

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func LevenshteinDistance(a, b string) int {
	m, n := len(a), len(b)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					dp[i-1][j],
					dp[i][j-1],
					dp[i-1][j-1],
				) + 1
			}
		}
	}
	return dp[m][n]
}

func MatchPercentage(a, b string) float64 {
	distance := LevenshteinDistance(a, b)
	maxLength := math.Max(float64(len(a)), float64(len(b)))
	if maxLength == 0 {
		return 100.0
	}
	return (1 - float64(distance)/maxLength) * 100
}
