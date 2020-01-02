package main

import (
	"fmt"
	"strings"
)

type grid [][]int

var hashTile map[string]bool
var grids []grid
var numLevels = 402
var midLevel = 200

func readInput() {
	rowSlice := strings.Split(inputStr, "\n")
	grids = []grid{}
	for n := 0; n < numLevels; n++ {
		g := grid{}
		for i := 0; i < 5; i++ {
			row := make([]int, 5, 5)
			g = append(g, row)
		}
		grids = append(grids, g)
	}
	fmt.Printf("Grids first grid has len %d\n", len(grids[0]))
	fmt.Printf("Grids first grid row has len %d\n", len(grids[0][0]))

	for i, row := range rowSlice {
		for j, v := range row {
			val := 0
			if v == '#' {
				val = 1
			}
			grids[midLevel][i][j] = val
		}
	}
	fmt.Printf("Grids has len %d\n", len(grids))
}

func iterateGrid() {
	newGrids := []grid{}
	for n := 0; n < numLevels; n++ {
		g := grid{}
		for i := 0; i < 5; i++ {
			row := make([]int, 5, 5)
			g = append(g, row)
		}
		newGrids = append(newGrids, g)
	}

	for n := 0; n < numLevels; n++ {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				//Skip the center
				if i == 2 && j == 2 {
					continue
				}
				//Check neighbors for non-center elements
				sum := 0
				if i > 0 {
					sum += grids[n][i-1][j]
				} else if n > 0 {
					sum += grids[n-1][1][2]
				}
				if j > 0 {
					sum += grids[n][i][j-1]
				} else if n > 0 {
					sum += grids[n-1][2][1]
				}
				if i < 4 {
					sum += grids[n][i+1][j]
				} else if n > 0 {
					sum += grids[n-1][3][2]
				}
				if j < 4 {
					sum += grids[n][i][j+1]
				} else if n > 0 {
					sum += grids[n-1][2][3]
				}

				//perform near center updates
				if i == 1 && j == 2 && n+1 < len(grids) {
					for k := 0; k < 5; k++ {
						sum += grids[n+1][0][k]
					}
				}
				if i == 2 && j == 1 && n+1 < len(grids) {
					for k := 0; k < 5; k++ {
						sum += grids[n+1][k][0]
					}
				}
				if i == 3 && j == 2 && n+1 < len(grids) {
					for k := 0; k < 5; k++ {
						sum += grids[n+1][4][k]
					}
				}
				if i == 2 && j == 3 && n+1 < len(grids) {
					for k := 0; k < 5; k++ {
						sum += grids[n+1][k][4]
					}
				}

				val := grids[n][i][j]
				if val == 0 && (sum == 1 || sum == 2) {
					newGrids[n][i][j] = 1
				}
				if val == 1 && sum == 1 {
					newGrids[n][i][j] = 1
				}

			}
		}
	}
	grids = newGrids
}

func printGrid(level int) {
	fmt.Printf("Over Here for level %d\n", level)
	for i := 0; i < len(grids[level]); i++ {
		for j := 0; j < len(grids[level][i]); j++ {
			fmt.Printf("%d", grids[level][i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func run(numRounds int) {
	for i := 0; i < numRounds; i++ {
		iterateGrid()
	}

}
func countBugs() {
	sum := int64(0)
	for n := 0; n < len(grids); n++ {
		for i := 0; i < len(grids[n]); i++ {
			for j := 0; j < len(grids[n][i]); j++ {
				if grids[n][i][j] == 1 {
					sum++
				}
			}
		}
	}
	fmt.Printf("Found a total of %d bugs\n", sum)

}

func printClose() {
	printGrid(197)
	printGrid(198)
	printGrid(199)
	printGrid(200)
	printGrid(201)
	printGrid(202)
}

func main() {
	readInput()
	printClose()
	run(200)
	fmt.Printf("RAN!\n")
	printClose()
	countBugs()
}

/*const inputStr = `#.###
.....
#..#.
##.##
..#.#`*/

/*
const inputStr = `....#
#..#.
#.?##
..#..
#....`
*/

/*
const inputStr = `....#
#..#.
#..##
..#..
#....`*/

const inputStr = `#.###
.....
#..#.
##.##
..#.#`
