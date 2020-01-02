package main

import (
	"fmt"
	"strconv"
	"strings"
	//"math"
)

const (
	//positionMode indicates we use the position
	positionMode     int64 = 0
	intermediateMode int64 = 1
	relativeMode     int64 = 2
)

var numParams = []int64{1000, 4, 4, 2, 2, 3, 3, 4, 4, 2}

type computer struct {
	id           int64
	loc          int
	state        []int64
	lastOutput   int64
	relativeBase int
}

func readInput() []int64 {
	values := []int64{}
	slice := strings.Split(inputStr, ",")
	for _, value := range slice {
		v, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("Unable to parse value from '%s': %s\n", value, err.Error())
			return values
		}
		values = append(values, int64(v))
	}
	return values
}

func (s *computer) run() {
	var accum, a, b, op, modeA, modeB, modeC int64
	var aloc, bloc, cloc int
	i := s.loc
	for i < len(s.state) {
		if s.state[i] == 99 {
			break
		}

		accum = int64(s.state[i])
		op = accum % 100
		accum = accum / 100
		modeA = accum % 10
		accum = accum / 10
		modeB = accum % 10
		accum = accum / 10
		modeC = accum % 10

		aloc = s.getLocation(i+1, modeA)
		a = s.readSafe(aloc)
		if numParams[op] >= 3 {
			bloc = s.getLocation(i+2, modeB)
			b = s.readSafe(bloc)
		}
		if numParams[op] >= 4 {
			cloc = s.getLocation(i+3, modeC)
		}

		switch op {
		case 1:
			s.writeSafe(cloc, a+b)
			i += 4
		case 2:
			s.writeSafe(cloc, a*b)
			i += 4
		case 3:
			input := readIntcodeInput()
			s.writeSafe(int(aloc), input)
			//fmt.Printf("Input: loc %d is now %d\n", aloc, input)
			i += 2
		case 4:
			//fmt.Printf("> At i=%d value=%v\n", i, a)
			writeIntcodeOutput(a)
			i += 2
		case 5:
			i += 3
			if a != 0 {
				i = int(b)
			}
		case 6:
			i += 3
			if a == 0 {
				i = int(b)
			}
		case 7:
			if a < b {
				s.writeSafe(cloc, 1)
			} else {
				s.writeSafe(cloc, 0)
			}
			i += 4
		case 8:
			if a == b {
				s.writeSafe(cloc, 1)
			} else {
				s.writeSafe(cloc, 0)
			}
			i += 4
		case 9:
			s.relativeBase += int(a)
			i += 2
		default:
			fmt.Printf("Error: Invalid operation %d at loc %d\n", op, cloc)
			return
		}

	}
}

func (s *computer) getLocation(start int, mode int64) int {
	loc := start
	switch mode {
	case positionMode:
		loc = int(s.readSafe(loc))
	case relativeMode:
		loc = int(s.readSafe(loc)) + s.relativeBase
	}
	return loc
}

func (s *computer) readSafe(loc int) int64 {
	if loc >= len(s.state) {
		return 0
	}
	return s.state[loc]
}

func (s *computer) writeSafe(loc int, value int64) {
	if loc < 0 {
		panic(fmt.Sprintf("Unable to write to location %d", loc))
	}

	for loc >= len(s.state) {
		s.state = append(s.state, 0)
	}
	s.state[loc] = value
}

func (s *computer) initialize() {
	s.loc = 0
	s.state = readInput()
	s.relativeBase = 0
	s.lastOutput = -1
}

func readIntcodeInput() int64 {
	//fmt.Printf("Reading (%d,%d)\n", locX, locY)
	if grid[locY][locX] == "#" {
		return int64(1)
	}
	return int64(0)
}

func writeIntcodeOutput(o int64) {
	//fmt.Printf("Have output %d\n", o)
	if toggle {
		writeIntcodeOutputDir(o)
	} else {
		writeIntcodeOutputGrid(o)
	}
	toggle = !toggle
}

func writeIntcodeOutputDir(o int64) {
	if o == 0 {
		dir = (dir + 4 - 1) % 4
	} else if o == 1 {
		dir = (dir + 1) % 4
	} else {
		panic(fmt.Sprintf("ERROR: INVALID OUTPUT %d", o))
	}
	switch dir {
	case 0:
		locY--
	case 1:
		locX++
	case 2:
		locY++
	case 3:
		locX--
	default:
		panic(fmt.Sprintf("ERROR: INVALID DIRECTION %d", dir))
	}

}

func writeIntcodeOutputGrid(o int64) {
	if o == 0 {
		grid[locY][locX] = "."
		painted[fmt.Sprintf("%d,%d", locX, locY)] = true
	} else if o == 1 {
		grid[locY][locX] = "#"
		painted[fmt.Sprintf("%d,%d", locX, locY)] = true
	} else {
		panic(fmt.Sprintf("ERROR: INVALID OUTPUT %d", o))
	}

}

func printGrid() {
	for _, row := range grid {
		for _, v := range row {
			fmt.Printf("%s", v)
		}
		fmt.Printf("\n")
	}
}

var grid [][]string
var painted map[string]bool
var area, locX, locY, dir int
var toggle bool

func main() {

	//Make the Grid
	//gridSize := 7
	gridSize := 100
	grid = make([][]string, 0)
	for i := 0; i < gridSize; i++ {
		row := make([]string, gridSize, gridSize)
		for j := range row {
			row[j] = "."
		}
		grid = append(grid, row)
	}
	locX = (gridSize - 1) / 2
	locY = (gridSize - 1) / 2
	grid[locY][locX] = "#"
	painted = make(map[string]bool)

	A := computer{}
	A.initialize()

	A.run()
	printGrid()
	area = len(painted)

	fmt.Printf("Final Loc (%d,%d). Area: %d\n", locX, locY, area)
	//fmt.Printf("\nEntries: %v\n", A.state)
}

const inputStr = `3,8,1005,8,319,1106,0,11,0,0,0,104,1,104,0,3,8,1002,8,-1,10,101,1,10,10,4,10,108,1,8,10,4,10,1001,8,0,28,2,1008,7,10,2,4,17,10,3,8,102,-1,8,10,101,1,10,10,4,10,1008,8,0,10,4,10,1002,8,1,59,3,8,1002,8,-1,10,101,1,10,10,4,10,1008,8,0,10,4,10,1001,8,0,81,1006,0,24,3,8,1002,8,-1,10,101,1,10,10,4,10,108,0,8,10,4,10,102,1,8,105,2,6,13,10,1006,0,5,3,8,1002,8,-1,10,101,1,10,10,4,10,108,0,8,10,4,10,1002,8,1,134,2,1007,0,10,2,1102,20,10,2,1106,4,10,1,3,1,10,3,8,102,-1,8,10,101,1,10,10,4,10,108,1,8,10,4,10,1002,8,1,172,3,8,1002,8,-1,10,1001,10,1,10,4,10,108,1,8,10,4,10,101,0,8,194,1,103,7,10,1006,0,3,1,4,0,10,3,8,1002,8,-1,10,1001,10,1,10,4,10,1008,8,1,10,4,10,101,0,8,228,2,109,0,10,1,101,17,10,1006,0,79,3,8,1002,8,-1,10,1001,10,1,10,4,10,108,0,8,10,4,10,1002,8,1,260,2,1008,16,10,1,1105,20,10,1,3,17,10,3,8,1002,8,-1,10,1001,10,1,10,4,10,1008,8,1,10,4,10,1002,8,1,295,1,1002,16,10,101,1,9,9,1007,9,1081,10,1005,10,15,99,109,641,104,0,104,1,21101,387365733012,0,1,21102,1,336,0,1105,1,440,21102,937263735552,1,1,21101,0,347,0,1106,0,440,3,10,104,0,104,1,3,10,104,0,104,0,3,10,104,0,104,1,3,10,104,0,104,1,3,10,104,0,104,0,3,10,104,0,104,1,21102,3451034715,1,1,21101,0,394,0,1105,1,440,21102,3224595675,1,1,21101,0,405,0,1106,0,440,3,10,104,0,104,0,3,10,104,0,104,0,21101,0,838337454440,1,21102,428,1,0,1105,1,440,21101,0,825460798308,1,21101,439,0,0,1105,1,440,99,109,2,22101,0,-1,1,21102,1,40,2,21101,0,471,3,21101,461,0,0,1106,0,504,109,-2,2106,0,0,0,1,0,0,1,109,2,3,10,204,-1,1001,466,467,482,4,0,1001,466,1,466,108,4,466,10,1006,10,498,1102,1,0,466,109,-2,2105,1,0,0,109,4,2101,0,-1,503,1207,-3,0,10,1006,10,521,21101,0,0,-3,21202,-3,1,1,22102,1,-2,2,21101,1,0,3,21102,540,1,0,1105,1,545,109,-4,2105,1,0,109,5,1207,-3,1,10,1006,10,568,2207,-4,-2,10,1006,10,568,22102,1,-4,-4,1106,0,636,22102,1,-4,1,21201,-3,-1,2,21202,-2,2,3,21102,587,1,0,1105,1,545,21201,1,0,-4,21101,0,1,-1,2207,-4,-2,10,1006,10,606,21102,0,1,-1,22202,-2,-1,-2,2107,0,-3,10,1006,10,628,22102,1,-1,1,21102,1,628,0,105,1,503,21202,-2,-1,-2,22201,-4,-2,-4,109,-5,2106,0,0`

//const inputStr = `3,200,104,1,104,0,3,200,104,0,104,0,3,200,104,1,104,0,104,1,104,0,3,200,104,0,104,1,104,1,104,0,104,1,104,0,99`
