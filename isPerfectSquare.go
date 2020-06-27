package main

func isPerfectSquare(num int) bool {
	// l -> Lower, r -> higher
	l, r := 1, num

	// Binary Search Propagation
	for l<r {
		mid := (l + r) / 2
		if mid*mid == num {
			return true
		} else if mid * mid < num {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l*l == num
}
