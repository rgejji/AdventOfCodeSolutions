package main

import (
	"fmt"
)

const (
	right = iota
	up
	left
	down
)

const (
	input = 277678
)

func main() {

	//numRows := 7
	numRows := 200
	createGrid(numRows)
	/*
		grid := createGrid(numRows)

			for i := 0; i < numRows; i++ {
				for j := 0; j < numRows; j++ {
					fmt.Printf("%v ", grid[i][j])
				}
				fmt.Printf("\n")
			}*/
}

func createGrid(numRows int) [][]int {
	mid := int(numRows / 2)
	grid := make([][]int, numRows)
	for i := range grid {
		grid[i] = make([]int, numRows)
	}
	xloc := mid
	yloc := mid

	dir := right
	for cnt := 1; cnt < 20000; {
		if xloc >= numRows || yloc >= numRows {
			return grid
		}
		//fmt.Printf("Adding to col=%d, row=%d, %d Dir is %v\n.", yloc, xloc, cnt, dir)
		val := calcVal(grid, numRows, cnt, xloc, yloc)
		grid[yloc][xloc] = val
		if xloc == mid && yloc == mid {
			grid[yloc][xloc] = 1
		}
		if val >= input {
			fmt.Printf("FOUND FIRST VALUE %v\n", val)
			return grid
		}

		cnt++
		switch dir {
		case right:
			xloc++
			if xloc >= yloc+1 {
				dir = up
			}
			continue
		case left:
			xloc--
			if mid-xloc >= mid-yloc {
				dir = down
			}
			continue
		case up:
			yloc--
			if mid-yloc >= xloc-mid {
				dir = left
			}
			continue
		case down:
			yloc++
			if yloc-mid >= mid-xloc {
				dir = right
			}
			continue
		}
	}
	return grid
}

func calcVal(grid [][]int, numRows, cnt, xloc, yloc int) int {
	if xloc-1 < 0 || xloc+1 >= numRows || yloc-1 < 0 || yloc+1 >= numRows {
		fmt.Println("ERROR: Exceed dimensions")
		return 0
	}
	sum := 0
	sum += grid[yloc-1][xloc-1] + grid[yloc-1][xloc] + grid[yloc-1][xloc+1]
	sum += grid[yloc][xloc-1] + grid[yloc][xloc+1]
	sum += grid[yloc+1][xloc-1] + grid[yloc+1][xloc] + grid[yloc+1][xloc+1]
	//return cnt
	return sum
}
