package main

import "fmt"

func VisitTheIsland(i, j int) {
	if i < 0 || i >= len(Area) || j < 0 || j >= len(Area[0]) || Area[i][j] == '0' {
		return
	}
	Area[i][j] = '0'
	VisitTheIsland(i+1, j)
	VisitTheIsland(i-1, j)
	VisitTheIsland(i, j+1)
	VisitTheIsland(i, j-1)
}

var Area [][]byte

func numIslands(grid [][]byte) int {
	noOfIslands := 0

	Area = grid

	for i, _ := range grid {
		for j, _ := range grid[0] {
			//fmt.Println(i, j)
			if grid[i][j] == '1' {
				noOfIslands++
				VisitTheIsland(i, j)
			}
		}
	}

	return noOfIslands
}

func main() {
	//in := [][]byte{{'1','1','0','0','0'},{'1','1','0','0','0'},{'0','0','1','0','0'},{'0','0','0','1','1'}}
	in := [][]byte{{'1'},{'1'}}
	fmt.Println(numIslands(in))
}
