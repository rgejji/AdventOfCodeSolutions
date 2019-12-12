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
	//fmt.Printf("Test 3: %d\n", gridValue(grid, 101, 153))

	//Calculate total of three
	best := -999999
	bestX := -1
	bestY := -1
	bestS := -1
	for size := 1; size <= numY; size++ {
		for x := 1; x <= numX-size+1; x++ {
			value := 0
			//Evaluate case where y=1
			for n := 0; n < size; n++ {
				for m := 0; m < size; m++ {
					value += gridValue(grid, x+n, 1+m)
				}
			}
			if value > best {
				best = value
				bestX = x
				bestY = 1
				bestS = size
			}
			//Perform a rolling update for y>1
			for y := 2; y <= numY-size+1; y++ {
				//remove old y-row values
				for n := 0; n < size; n++ {
					value -= gridValue(grid, x+n, y-1)
				}

				//add new values
				for n := 0; n < size; n++ {
					value += gridValue(grid, x+n, y+size-1)
				}
				if value > best {
					best = value
					bestX = x
					bestY = y
					bestS = size
				}
			}

		}
	}

	fmt.Printf("Best score is %d at x=%d and y=%d and size=%d\n", best, bestX, bestY, bestS)
}
