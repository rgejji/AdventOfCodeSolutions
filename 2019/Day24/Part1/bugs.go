package main

import (
	"fmt"
	"strings"
)

var hashTile map[string]bool
var grid [][]int

func readInput() {
	hashTile = make(map[string]bool)
	grid = [][]int{}
	rowSlice := strings.Split(inputStr, "\n")
	for _, row := range rowSlice {
		gridRow := []int{}
		for _, v := range row {
			val := 0
			if v == '#' {
				val = 1
			}
			gridRow = append(gridRow, val)
		}
		grid = append(grid, gridRow)
	}
}

func addHashReturnCollision() bool {
	s := ""
	for _, row := range grid {
		for _, val := range row {
			s += fmt.Sprintf("%s", val)
		}
	}
	if _, ok := hashTile[s]; ok {
		return true
	}
	hashTile[s] = true
	return false
}

func iterateGrid() {
	newGrid := [][]int{}
	for i, row := range grid {
		newRow := make([]int, len(row), len(row))
		for j, val := range row {
			sum := 0
			if i > 0 {
				sum += grid[i-1][j]
			}
			if j > 0 {
				sum += grid[i][j-1]
			}
			if i < len(grid)-1 {
				sum += grid[i+1][j]
			}
			if j < len(grid)-1 {
				sum += grid[i][j+1]
			}
			if val == 0 && (sum == 1 || sum == 2) {
				newRow[j] = 1
			}
			if val == 1 && sum == 1 {
				newRow[j] = 1
			}

		}
		newGrid = append(newGrid, newRow)
	}
	grid = newGrid
}

func printGrid() {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%d", grid[i][j])
		}
		fmt.Printf("\n")
	}
}

func runTillMatch() {
	cnt := 0
	for !addHashReturnCollision() {
		iterateGrid()
		cnt++
		if cnt%100 == 0 {
			fmt.Printf("> n=%d\n", cnt)
		}

	}

}
func calculateDiversity() {
	var sum int64
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				sum += 1 << uint((i*len(grid[i]) + j))
			}
		}
	}
	fmt.Printf("Diversity is %d\n", sum)
}

func main() {
	readInput()
	runTillMatch()
	printGrid()
	calculateDiversity()
}

const inputStr = `#.###
.....
#..#.
##.##
..#.#`

/*
const inputStr = `....#
#..#.
#..##
..#..
#....`*/
