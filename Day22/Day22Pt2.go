package main

import (
	"fmt"
	"log"
	"strings"
)

const (
	up = iota
	right
	down
	left
)

const (
	clean    = '.'
	weakened = 'W'
	infected = '#'
	flagged  = 'F'
)

type runegrid [][]rune

func (rg runegrid) String() string {
	p := ""
	for i := 0; i < len(rg); i++ {
		if i != 0 {
			p += "\n"
		}
		for j := 0; j < len(rg[i]); j++ {
			if rg[i][j] == 0 {
				p += "."
				continue
			}
			p += string(rg[i][j])
		}
	}
	return p

}

func (rg runegrid) Count() int {
	count := 0
	for _, line := range rg {
		for _, char := range line {
			if char == '#' {
				count++
			}
		}
	}
	return count
}

func createGridFromInput(patternString string, splitStr string) (runegrid, int, int) {
	lines := strings.Split(patternString, splitStr)
	grid := make(runegrid, len(lines), len(lines))
	for i, line := range lines {
		cleanLine := strings.TrimSpace(line)
		grid[i] = make([]rune, len(cleanLine), len(cleanLine))
		for j, char := range cleanLine {
			grid[i][j] = char
		}
	}
	return grid, len(grid) / 2, len(grid[0]) / 2

}

func main() {
	numSteps := 10000000
	//grid, i, j := createGridFromInput(startA, "\n")
	grid, i, j := createGridFromInput(start, "\n")
	dir := up
	infectionCount := 0
	for t := 0; t < numSteps; t++ {
		switch grid[i][j] {
		case 0:
			fallthrough
		case clean:
			dir = turnLeft(dir)
			grid[i][j] = weakened
		case weakened:
			grid[i][j] = infected
			infectionCount++
		case infected:
			dir = turnRight(dir)
			grid[i][j] = flagged
		case flagged:
			dir = turnRight(turnRight(dir))
			grid[i][j] = clean
		default:
			log.Fatalf("Encountered unknown grid character %v\n", grid[i][j])
		}
		grid, i, j = expandGrid(grid, i, j, dir)
		i, j = moveNode(grid, i, j, dir)
		if t%1000000 == 0 {
			fmt.Printf("Took %d steps\n", t)
		}
	}
	//fmt.Printf("Have grid\n%+v\n", grid)
	fmt.Printf("Have %d infections\n", infectionCount)

}

func moveNode(grid runegrid, i, j, dir int) (int, int) {
	switch dir {
	case down:
		return i + 1, j
	case up:
		return i - 1, j
	case left:
		return i, j - 1
	case right:
		return i, j + 1
	}
	log.Fatalf("Error: UNKNOWN DIRECTION %v", dir)
	return 0, 0
}

func turnLeft(dir int) int {
	return (dir - 1 + 4) % 4
}
func turnRight(dir int) int {
	return (dir + 1 + 4) % 4
}

//Note expanding sets the runes to 0, the default. here . or 0 are clean
func expandGrid(grid runegrid, i, j, dir int) (runegrid, int, int) {
	numRows := len(grid)
	numCols := len(grid[0])
	switch dir {
	case down:
		if i != numRows-1 {
			return grid, i, j
		}
		fmt.Printf("Expanding grid of size %dx%d downward\n", numRows, numCols)

		newGrid := allocatePattern(2*numRows, numCols)
		for i := 0; i < numRows; i++ {
			newGrid[i] = grid[i]
		}
		return newGrid, i, j
	case up:
		if i != 0 {
			return grid, i, j
		}
		fmt.Printf("Expanding grid of size %dx%d upward\n", numRows, numCols)
		newGrid := allocatePattern(2*numRows, numCols)
		for i := 0; i < numRows; i++ {
			newGrid[i+numRows] = grid[i]
		}
		return newGrid, i + numRows, j
	case right:
		if j != numCols-1 {
			return grid, i, j
		}
		fmt.Printf("Expanding grid of size %dx%d to the right\n", numRows, numCols)
		newGrid := allocatePattern(numRows, 2*numCols)
		for i := 0; i < numRows; i++ {
			tmp := newGrid[i][numCols:]
			newGrid[i] = append(grid[i], tmp...)
		}
		return newGrid, i, j
	case left:
		if j != 0 {
			return grid, i, j
		}
		fmt.Printf("Expanding grid of size %dx%d to the left\n", numRows, numCols)
		newGrid := allocatePattern(numRows, 2*numCols)
		for i := 0; i < numRows; i++ {
			tmp := newGrid[i][:numCols]
			newGrid[i] = append(tmp, grid[i]...)
		}
		return newGrid, i, j + numCols
	}
	log.Fatalf("Unknown direction %v", dir)
	return grid, 0, 0
}

func allocatePattern(numRows int, numCols int) runegrid {
	pattern := make(runegrid, numRows, numRows)
	for i := 0; i < numRows; i++ {
		pattern[i] = make([]rune, numCols, numCols)
	}
	return pattern
}

const (
	startA = `..#
#..
...`

	start = `#.....##.####.#.#########
.###..#..#..####.##....#.
..#########...###...####.
.##.#.##..#.#..#.#....###
...##....###..#.#..#.###.
###..#...######.####.#.#.
#..###..###..###.###.##..
.#.#.###.#.#...####..#...
##........##.####..##...#
.#.##..#.#....##.##.##..#
###......#..##.####.###.#
....#..###..#######.#...#
#####.....#.##.#..#..####
.#.###.#.###..##.#..####.
..#..##.###...#######....
.#.##.#.#.#.#...###.#.#..
##.###.#.#.###.#......#..
###..##.#...#....#..####.
.#.#.....#..#....##..#..#
#####.#.##..#...##..#....
##..#.#.#.####.#.##...##.
..#..#.#.####...#........
###.###.##.#..#.##.....#.
.##..##.##...#..#..#.#..#
#...####.#.##...#..#.#.##`
)
