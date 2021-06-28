/*
// Sample code to perform I/O:

fmt.Scanf("%s", &myname)            // Reading input from STDIN
fmt.Println("Hello", myname)        // Writing output to STDOUT

// Warning: Printing unwanted or ill-formatted data to output will cause the test cases to fail
*/

// Write your code here
package main

import (
	"fmt"
)

func main() {
	noOfTestcases := 0
	fmt.Scanf("%d", &noOfTestcases)
	for test := 0; test < noOfTestcases; test++ {
		size := 0
		shiftSize := 0
		fmt.Scanf("%d", &size)
		fmt.Scanf("%d", &shiftSize)

		shiftSize = shiftSize % size

		input := make([]int, size)

		for i, _ := range(input) {
			fmt.Scanf("%d", &input[i])
		}

		if shiftSize == 0 {
			for _, v := range(input) {
				fmt.Printf("%v ", v)
			}
			fmt.Println()
		} else {
			// Perform Shift
			for i, _ := range(input) {
				// Size-1-Shift+i
				if size-1-shiftSize+i < size {
					fmt.Printf("%d ", input[size-1-shiftSize+i])
				} else {
					fmt.Printf("%d ", input[i-shiftSize-1])
				}
			}
			fmt.Println()
		}
	}
}