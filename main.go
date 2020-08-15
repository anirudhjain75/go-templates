package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	//input := []string{"dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"}
	//fmt.Println(reorderLogFiles(input))
	// input := []int{1, 0, 0, 1, 0, 0, 1, 0}
	//input := []string{"leet", "code"}
	//fmt.Println(wordBreak("leetcode", input))
	preorderTraversal(nil)
}

func wordBreak(s string, wordDict []string) bool {

	// Remove non existent sub strings from wordDict
	for i, v := range wordDict {
		if !strings.Contains(s, v) {
			wordDict = append(wordDict[:i], wordDict[i+1:]...)
		}
	}

	// Brute force approach

	fmt.Println(wordDict)

	return false
}

func prisonAfterNDays(cells []int, N int) []int {
	if N > 14 {
		N = N%14 + 14
	} else {
		N = N % 14
	}
	for i := 0; i < N; i++ {
		temp := make([]int, len(cells))
		for j := 1; j < len(temp)-1; j++ {
			if cells[j-1] == cells[j+1] {
				temp[j] = 1
			} else {
				temp[j] = 0
			}
		}
		cells = temp
	}
	return cells
}

func reorderLogFiles(logs []string) []string {
	digitLogs := []string{}
	letterLogs := []string{}
	for _, val := range logs {
		idx := strings.Index(val, " ")
		if val[idx+1] >= 'a' && val[idx+1] <= 'z' {
			letterLogs = append(letterLogs, val)
		} else {
			digitLogs = append(digitLogs, val)
		}
	}
	sortLetterLogs(letterLogs)
	result := []string{}
	result = append(result, letterLogs...)
	result = append(result, digitLogs...)
	return result
}

//  this sorting is done in place
func sortLetterLogs(letterLogs []string) {
	sort.Slice(letterLogs, func(i, j int) bool {
		this := letterLogs[i]
		other := letterLogs[j]
		iIdx := strings.Index(this, " ")
		jIdx := strings.Index(other, " ")
		iRest := this[iIdx+1:]
		jRest := other[jIdx+1:]
		if iRest == jRest {
			return this < other
		}
		return iRest < jRest
	})
}

func mostCommonWord(paragraph string, banned []string) string {

	paragraph = strings.Replace(paragraph, "!", " ", -1)
	paragraph = strings.Replace(paragraph, "?", " ", -1)
	paragraph = strings.Replace(paragraph, "'", " ", -1)
	paragraph = strings.Replace(paragraph, ",", " ", -1)
	paragraph = strings.Replace(paragraph, ";", " ", -1)
	paragraph = strings.Replace(paragraph, ".", " ", -1)

	paragraph = strings.ToLower(paragraph)

	// Create bannedMap
	bannedMap := make(map[string]bool)
	for _, v := range banned {
		_, ok := bannedMap[v]
		if !ok {
			bannedMap[v] = true
		}
	}

	// Create normal paragraph map
	m := make(map[string]int)
	for _, v := range strings.Split(paragraph, " ") {
		if v == "" {
			continue
		}
		_, ok := bannedMap[v]
		if !ok {
			val, ok := m[v]
			if !ok {
				m[v] = 1
			} else {
				m[v] = val + 1
			}
		} else {

		}
	}

	maxStr := ""
	max := 0

	for k, v := range m {
		if v > max {
			max = v
			maxStr = k
		}
	}

	return maxStr
}

func kClosest(points [][]int, K int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		a := points[i]
		b := points[j]

		aV := math.Sqrt(math.Pow(float64(a[0]), float64(2)) + math.Pow(float64(a[1]), float64(2)))
		bV := math.Sqrt(math.Pow(float64(b[0]), float64(2)) + math.Pow(float64(b[1]), float64(2)))
		//fmt.Println(a, b, aV, bV)

		return aV < bV
	})
	return points[:K]
}

func singleNonDuplicate(nums []int) int {

	// If only 1 value, return it
	if len(nums) == 1 {
		return nums[0]
	}

	// Check Every next two element if they are different, return the first element
	for i := 0; i < len(nums)-1; i = i + 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}

	// Return last element incase of Odd length Arr and All even elements pass the check
	return nums[len(nums)-1]
}

func findJudge2(N int, trust [][]int) int {
	a := make([]int, N+1)
	for _, p := range trust {
		a[p[0]]--
		a[p[1]]++
	}
	for i := 1; i < N+1; i++ {
		if a[i] == N-1 {
			return i
		}
	}
	return -1
}

func findJudge(N int, trust [][]int) int {

	// Create the trust 2D array
	Arr := make([][]int, N)
	for _, v := range trust {
		Arr[v[0]-1] = append(Arr[v[0]-1], v[1])
	}

	// Create Candidates for Judge using 1st Condition
	candidate := 0
	for i, v := range Arr {
		if len(v) == 0 {
			candidate = i + 1
			break
		}
	}

	// Confirm the candidacy
	vote := 0
	for _, v := range Arr {
		for _, cV := range v {
			if cV == candidate {
				vote++
			}
		}
	}

	if vote == N-1 {
		return candidate
	}

	return -1
}

func firstUniqChar(s string) int {
	m := make([]int, 26)
	sb := []byte(s)
	for i := 0; i < len(sb); i++ {
		m[sb[i]-97]++
	}
	for i := 0; i < len(sb); i++ {
		if m[sb[i]-97] == 1 {
			return i
		}
	}
	return -1
}

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		val, ok := m[v]
		if ok {
			m[v] = val + 1
		} else {
			m[v] = 1
		}
	}
	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}
	return -1
}
