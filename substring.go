package main

import "fmt"

func main() {
	fmt.Println("Substring Question")
	fmt.Println(countSubstrings("aaa"))
}

// Leetcode 647. Palindromic Substrings
//Given a string, your task is to count how many palindromic substrings in this string.
//
//The substrings with different start indexes or end indexes are counted as different substrings even they consist of same characters.

func isPalindrome(input string) bool {
	isPalin := true
	if len(input) <= 1 {
		return true
	}
	for i := 0; i < len(input) / 2; i++ {
		if input[i] != input[len(input)-1-i] {
			isPalin = false
			break
		}
	}
	return isPalin
}

func countSubstrings(s string) int {
	count := 0
	for slidingWindowLen := 1; slidingWindowLen <= len(s); slidingWindowLen++ {
		for i := 0; i < len(s) - slidingWindowLen + 1; i++ {
			if isPalindrome(s[i:i+slidingWindowLen]) {
				count++
			}
		}
	}
	return count
}