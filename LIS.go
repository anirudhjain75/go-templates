package main

import "fmt"

func GetMaxOfArr(a []int) int {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
	}
	return max
}

func lengthOfLIS(nums []int) int {
	// Create DP Array
	dp := make([]int, len(nums))
	for i, _ := range dp {
		dp[i] = 1
	}

	for i, _ := range nums {
		arr := make([]int, i+1)
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				arr = append(arr, dp[j])
			}
		}
		dp[i] = 1 + GetMaxOfArr(arr)
	}
	return GetMaxOfArr(dp)
}

func main() {
	fmt.Println(lengthOfLIS([]int{7,7,7,7,7,7,7}))
}
