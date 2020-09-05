package main

// Leetcode 904. Fruit Into Baskets

// backTrackFruitArr returns the fruit id and its number
func backTrackFruitArr(tree []int, pos int) (int, int) {
	elem := -1
	no := 0
	for i := pos-1; i >= 0; i-- {
		if elem == -1 {
			elem = tree[i]
			no++
		} else if elem == tree[i] {
			no++
		} else if elem != tree[i] {
			break
		}
	}
	return elem, no
}

func totalFruit(tree []int) int {
	fruit1 := -1
	fruit2 := -1
	noOfFruit1 := 0
	noOfFruit2 := 0
	maxNoOfFruits := -1

	if len(tree) == 0 {
		return 0
	}

	for i, v := range tree {
		if fruit1 == -1 {
			fruit1 = v
			noOfFruit1 = 1
		} else if fruit2 == -1 && fruit1 != v {
			fruit2 = v
			noOfFruit2 = 1
		} else {
			if fruit1 == v {
				noOfFruit1++
			} else if fruit2 == v {
				noOfFruit2++
			} else {
				elem, no := backTrackFruitArr(tree, i)
				if fruit1 == elem {
					fruit2 = v
					noOfFruit2 = 1
					noOfFruit1 = no
				} else if fruit2 == elem {
					fruit1 = v
					noOfFruit1 = 1
					noOfFruit2 = no
				}
			}
		}

		if maxNoOfFruits < noOfFruit1 + noOfFruit2 {
			maxNoOfFruits = noOfFruit1 + noOfFruit2
		}

	}
	return maxNoOfFruits
}