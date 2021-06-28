package main

import "fmt"

func VisitTheIslandAndMeasureMap(i, j int) {
	if i < 0 || i >= len(Map) || j < 0 || j >= len(Map[0]) || Map[i][j] == 0 {
		return
	}
	Map[i][j] = 0
	IslandArea++
	VisitTheIslandAndMeasureMap(i+1, j)
	VisitTheIslandAndMeasureMap(i-1, j)
	VisitTheIslandAndMeasureMap(i, j+1)
	VisitTheIslandAndMeasureMap(i, j-1)
}

var Map [][]int

var MaxArea int
var IslandArea int

func maxAreaOfIsland(grid [][]int) int {

	MaxArea = 0

	Map = grid

	for i, _ := range grid {
		for j, _ := range grid[0] {
			IslandArea = 0
			//fmt.Println(i, j)
			if grid[i][j] == 1 {
				VisitTheIslandAndMeasureMap(i, j)
				if IslandArea > MaxArea {
					MaxArea = IslandArea
				}
			}
		}
	}

	return MaxArea
}

func main() {
	//in := [][]byte{{'1','1','0','0','0'},{'1','1','0','0','0'},{'0','0','1','0','0'},{'0','0','0','1','1'}}
	in := [][]int{{1},{1}}
	fmt.Println(maxAreaOfIsland(in))
}
