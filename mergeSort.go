package main

import "fmt"

func main() {
	fmt.Println("Merge Sort")
	input := []int32{2, 1, 3, 1, 2}
	fmt.Println(countInversions(input))
}

var inversion int64

func countInversions(arr []int32) int64 {
	inversion = 0
	mergeSort(arr, 0, len(arr))
	return inversion
}

func mergeSort(input []int32, low int, high int) {

	// fmt.Println(len(input[low:high]), input[low:high])
	if high-low > 1 {
		middle := int((high + low + 1) / 2)
		mergeSort(input, low, middle)
		mergeSort(input, middle, high)
		// fmt.Println("running for", low, middle, high)
		merge(input, low, middle, high)
	}
}

func merge(input []int32, low int, middle int, high int) {

	var (
		left  = make([]int32, 0)
		right = make([]int32, 0)
	)

	for t := low; t < middle; t++ {
		left = append(left, input[t])
	}

	for t := middle; t < high; t++ {
		right = append(right, input[t])
	}

	// fmt.Println("merging", left, right, input)
	// fmt.Println("low", low, "middle", middle, "high", high)

	index := low

	for len(left) > 0 && len(right) > 0 {
		if left[0] > right[0] {
			input[index] = right[0]
			inversion += int64(len(left))
			right = right[1:]
		} else {
			input[index] = left[0]
			left = left[1:]
		}
		index++
	}

	for len(left) > 0 {
		input[index] = left[0]
		left = left[1:]
		index++
	}

	for len(right) > 0 {
		input[index] = right[0]
		right = right[1:]
		index++
	}

	// fmt.Println(input)

}
