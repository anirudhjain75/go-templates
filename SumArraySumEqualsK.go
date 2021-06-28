package main

import "fmt"

func subarraySum(nums []int, k int) int {
	count := 0;
	sum := 0;

	m := make(map[int]int, len(nums))
	m[0] = 1

	for _, v := range nums {
		sum += v
		if _, ok := m[sum-k];ok {
			count += m[sum-k]
		}
		val, ok := m[sum]
		if ok  {
			m[sum] = val+1
		} else {
			m[sum] = 1
		}
	}

	return count
}
func main() {
	fmt.Println("Hi")
}