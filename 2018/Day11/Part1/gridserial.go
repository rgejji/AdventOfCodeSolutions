package main

import (
	"fmt"
)

func takeHundred(n int) int {
	return n / 100
}

func getValueFromCurrent(n int) int {
	return takeHundred(n) - 5
}

func gridValue(grid [][]int, x, y int) int {
	return grid[x-1][y-1]
}

func main() {

	numX := 300
	numY := 300
	//serial := 8
	//serial := 57
	//serial := 39
	//serial := 71
	//serial := 42
	serial := 5177

	grid := make([][]int, 300, 300)
	for i := 0; i < numX; i++ {
		grid[i] = make([]int, 300, 300)
	}

	//Calculate power levels going left to right, top to bottom
	for x := 1; x <= numX; x++ {
		rackID := 10 + x
		curr := ((rackID*1 + serial) * rackID) % 1000
		grid[x-1][0] = getValueFromCurrent(curr)
		for y := 2; y <= numY; y++ {
			diff := rackID * rackID
			curr = (diff + curr) % 1000
			grid[x-1][y-1] = getValueFromCurrent(curr)
		}
	}

	//fmt.Printf("Test 0: %d\n", grid[3-1][5-1])
	//fmt.Printf("Test 1: %d\n", grid[122-1][79-1])
	//fmt.Printf("Test 2: %d\n", gridValue(grid, 217, 196))
	fmt.Printf("Test 3: %d\n", gridValue(grid, 101, 153))

	//Calculate total of three
	best := -999999
	bestX := -1
	bestY := -1

	for x := 1; x <= numX-2; x++ {
		value := gridValue(grid, x, 1) + gridValue(grid, x, 2) + gridValue(grid, x, 3) +
			gridValue(grid, x+1, 1) + gridValue(grid, x+1, 2) + gridValue(grid, x+1, 3) +
			gridValue(grid, x+2, 1) + gridValue(grid, x+2, 2) + gridValue(grid, x+2, 3)
		if value > best {
			best = value
			bestX = x
			bestY = 1
		}
		for y := 2; y <= numY-2; y++ {
			value = value - gridValue(grid, x, y-1) - gridValue(grid, x+1, y-1) - gridValue(grid, x+2, y-1) +
				gridValue(grid, x, y+2) + gridValue(grid, x+1, y+2) + gridValue(grid, x+2, y+2)
			if value > best {
				best = value
				bestX = x
				bestY = y
			}
		}

	}

	fmt.Printf("Best score is %d at x=%d and y=%d\n", best, bestX, bestY)
}
