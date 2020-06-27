package main

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) == 2 {
		return true
	}
	m := float64(coordinates[1][1] - coordinates[0][1])/float64(coordinates[1][0] - coordinates[0][0])
	for i := 1 ; i < len(coordinates); i++ {
		tempM := float64(coordinates[i][1] - coordinates[i-1][1])/float64(coordinates[i][0] - coordinates[i-1][0])
		if tempM != m {
			return false
		}
	}
	return true
}

// Depth First Search
func fill(image [][]int, sr int, sc int, newColor int, originColor int) {
	if sr<0 || sr>=len(image) || sc<0 || sc>=len(image[0]) || isDone[sr][sc] {

	} else {
		isDone[sr][sc] = true
		image[sr][sc] = newColor

		if sc < len(image[0])-1 {
			if image[sr][sc+1] == originColor {
				fill(image, sr, sc+1, newColor, originColor)
			}
		}
		if sc > 0 {
			if image[sr][sc-1] == originColor {
				fill(image, sr, sc-1, newColor, originColor)
			}
		}
		if sr < len(image)-1 {
			if image[sr+1][sc] == originColor {
				fill(image, sr+1, sc, newColor, originColor)
			}
		}
		if sr > 0 {
			if image[sr-1][sc] == originColor {
				fill(image, sr-1, sc, newColor, originColor)
			}
		}
	}
}

var isDone [][]bool

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	// Create isDone Arr for DFS
	isDone = make([][]bool, len(image))
	for i, _ := range isDone {
		isDone[i] = make([]bool, len(image[0]))
	}

	fill(image, sr, sc, newColor, image[sr][sc])
	return image
}

func floodFill2(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] != newColor {
		helper(image, sr, sc, image[sr][sc], newColor)
	}
	return image
}

// DFS
func helper(image [][]int, sr, sc, oldColor, newColor int) {
	if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) || image[sr][sc] != oldColor {
		return
	}
	image[sr][sc] = newColor
	helper(image, sr-1, sc, oldColor, newColor)
	helper(image, sr+1, sc, oldColor, newColor)
	helper(image, sr, sc-1, oldColor, newColor)
	helper(image, sr, sc+1, oldColor, newColor)
}