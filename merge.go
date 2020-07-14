package main

import "fmt"

func main() {
	fmt.Println("Hi")
	input := []int32{2, 1, 3, 1, 2}
	fmt.Println(countInversions(input))
}

func countInversions(arr []int32) int64 {

	_, totalInv := SortAndCountInversions(arr, 0)

	return totalInv
}

func SortAndCountInversions(s []int32, inv int64) ([]int32, int64) {
	if len(s) == 1 {
		return s, inv
	} else {
		s1, s2 := SplitSlice(s)

		leftS, leftInv := SortAndCountInversions(s1, inv)
		rightS, rightInv := SortAndCountInversions(s2, inv)
		mergedS, mergedInv := MergeAndCountInversions(leftS, rightS)

		totalInv := leftInv + rightInv + mergedInv

		return mergedS, totalInv
	}
}

func MergeAndCountInversions(leftS, rightS []int32) ([]int32, int64) {
	var mergedS []int32
	mergedInv := 0
	leftI, rightI := 0, 0

	for i := 0; i < len(leftS)+len(rightS); i++ {
		if leftI == len(leftS) {
			mergedS = append(mergedS, rightS[rightI])
			rightI++
		} else if rightI == len(rightS) {
			mergedS = append(mergedS, leftS[leftI])
			leftI++
		} else if (leftI != len(leftS) && rightI != len(rightS)) && leftS[leftI] < rightS[rightI] {
			mergedS = append(mergedS, leftS[leftI])
			leftI++
		} else if (leftI != len(leftS) && rightI != len(rightS)) && leftS[leftI] > rightS[rightI] {
			mergedS = append(mergedS, rightS[rightI])
			for _, v := range leftS {
				rv := rightS[rightI]
				if v > rv {
					mergedInv++
				}
			}
			rightI++
		}
	}
	return mergedS, int64(mergedInv)
}

func SplitSlice(s []int32) (s1, s2 []int32) {
	if len(s)%2 == 0 {
		s1 = s[:len(s)/2]
		s2 = s[len(s)/2:]
	} else {
		s1 = s[:(len(s)+1)/2]
		s2 = s[(len(s)+1)/2:]
	}
	return
}

// var count int

// func mergeAndCount(in []int, begin int, middle int, end int) int {
// 	left := in[begin:middle]
// 	right := in[middle:end]
// 	fmt.Println("merging", left, "and", right)
// 	swaps := 0
// 	i := 0
// 	j := 0
// 	k := begin
// 	for i < len(left) && j < len(right) {
// 		if left[i] > right[j] {
// 			in[k] = right[j]
// 			j++
// 			swaps += middle - len(left)
// 		} else {
// 			in[k] = left[i]
// 			i++
// 		}
// 		k++
// 	}
// 	for i < len(left) {
// 		in = append(in, left[i])
// 		i++
// 	}
// 	for j < len(right) {
// 		in = append(in, right[j])
// 		j++
// 	}

// 	fmt.Println("resulting to", in[begin:end])

// 	return swaps
// }

// func mergeSort(in []int, begin int, end int) int {

// 	middle := int((begin + end) / 2)

// 	count = 0

// 	if begin < end {
// 		count += mergeSort(in, begin, middle)
// 		count += mergeSort(in, middle+1, end)
// 		count += mergeAndCount(in, begin, middle, end)
// 	}

// 	return count

// }
