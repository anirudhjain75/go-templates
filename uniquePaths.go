package main

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i:=0; i<m; i++ {
		dp[i] = make([]int, n)
	}
	return memoize(dp, m-1, n-1)
}

func memoize(dp [][]int, x int, y int) int {
	result := 0
	if x == 0 || y == 0 {
		dp[x][y] = 1
		result = dp[x][y]
	} else if dp[x][y] != 0 {
		result = dp[x][y]
	} else {
		dp[x][y] = memoize(dp, x-1, y) + memoize(dp, x, y-1)
		result = dp[x][y]
	}
	return result
}

