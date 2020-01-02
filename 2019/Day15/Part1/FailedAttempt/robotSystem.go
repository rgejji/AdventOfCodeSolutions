package main

import (
	"fmt"

	intcode "github.com/RichardGejji/AdventOfCode2019/Intcode"
)

const (
	north int64 = 1
	south int64 = 2
	west  int64 = 3
	east  int64 = 4
)

type exploreLoc struct {
	pt        point
	history   []int64
	direction int64
}

type point struct {
	x int64
	y int64
}

func explore(s *intcode.Computer, e *exploreLoc) {
	var gridVal int64
	maxCnt := 10

	fmt.Printf("Exploring location %v\n", e.pt)
	//Perform previous command to catchup
	if len(e.history) > 0 {
		fmt.Printf("History is %+v", e.history)
		for _, instruction := range e.history {
			s.Input <- instruction
			gridVal = <-s.Output
			fmt.Printf(">Fast direction %d, got output %d\n", instruction, gridVal)
		}
		if inGrid(e.pt) {
			printGrid()
			panic(fmt.Sprintf("ERROR: exploration revisiting already visited %v\nValue is %s\n", e.pt, getGridValue(e.pt)))
		}
		writeGrid(gridVal, e.pt)
	}
	currHistory = make([]int64, len(e.history))
	copy(currHistory, e.history)
	fmt.Printf("Copied history is %+v\n", currHistory)
	curr = e.pt

	//Eliminate the already explored options.
	//Pick one option to explore (most likely previous direction)
	//Store the history of options we are skipping
	//Abort if we have no options left
	var testFirst int64
	var testPt point
	for cnt := 0; cnt < maxCnt; cnt++ {
		if len(currHistory) == 0 {
			testFirst = 0
		} else {
			testFirst = currHistory[len(currHistory)-1]
		}

		immediateLocs := []exploreLoc{}
		for i := int64(0); i < int64(4); i++ {
			val := (i+testFirst)%4 + int64(1)
			switch val {
			case north:
				testPt = point{curr.x, curr.y - 1}
			case south:
				testPt = point{curr.x, curr.y + 1}
			case west:
				testPt = point{curr.x - 1, curr.y}
			case east:
				testPt = point{curr.x + 1, curr.y}

			}
			if !inGrid(testPt) {
				copyHistory := make([]int64, len(currHistory))
				copy(copyHistory, currHistory)
				copyHistory = append(copyHistory, val)
				immediateLocs = append(immediateLocs, exploreLoc{pt: testPt, history: copyHistory, direction: val})
			}
		}
		//Test each immediate location, and put the remaining immediate locations on the explore later queue
		i := 0
		for i < len(immediateLocs) {
			e := immediateLocs[i]
			val := testMove(s, &e)
			if val != 0 {
				break
			}
			i++
		}
		//eplore late the places we haven't explored yet
		for i < len(immediateLocs) {
			e := immediateLocs[i]
			hash := getHash(e.pt)
			if _, ok := toExplore[hash]; !ok {
				toExplore[hash] = &e
			}
			i++
		}
		if s.LastOutput <= 0 {
			//fmt.Printf("Stuck at a wall, escaping\n")
			s.Input <- intcode.Abort
			return
		}
	}
	s.Input <- intcode.Abort
	return
}

func inGrid(p point) bool {
	val := grid[p.y][p.x]
	if val == " " {
		return false
	}
	return true
}

func testMove(s *intcode.Computer, e *exploreLoc) int64 {
	var output int64

	s.Input <- e.direction
	output = <-s.Output
	fmt.Printf(">Gave direction %d, got output %d\n", e.direction, output)

	//Check that given grid value doesn't change
	checkA := getGridValue(e.pt)
	checkB := writeGrid(output, e.pt)
	if checkA != " " && checkB != checkA {
		panic(fmt.Sprintf("Error: grid %v changed value from %s to %s\n", e.pt, checkA, checkB))
	}

	//Move robot if appropriate output from move
	if output != 0 {
		moveRobot(e)
		if output == 2 {
			fuel = e.pt
			fmt.Printf("FOUND FUEL %v", e.pt)
		}
	}
	return output
}

func moveRobot(e *exploreLoc) {
	//set current to next
	curr = e.pt
	currHistory = append(currHistory, e.direction)
	//clear location from to explore
	key := getHash(curr)
	delete(toExplore, key)
}

func getHash(p point) string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)

}

func getGridValue(loc point) string {
	return grid[loc.y][loc.x]
}

func writeGrid(o int64, loc point) string {
	switch o {
	case 0:
		grid[loc.y][loc.x] = "#"
	case 1:
		grid[loc.y][loc.x] = "."
	case 2:
		grid[loc.y][loc.x] = "X"
	case 3:
		grid[loc.y][loc.x] = "@"
	default:
		panic(fmt.Sprintf("ERROR: Received unexpected output %d\n", o))
	}
	return grid[loc.y][loc.x]
}

func printGrid() {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%s", grid[i][j])
		}
		fmt.Printf("\n")
	}

}

var grid [][]string
var toExplore map[string]*exploreLoc
var curr, next, fuel point
var currHistory []int64

func main() {
	maxCnt := 2
	//Initialize Grid:
	gridSize := 41
	grid = [][]string{}
	for i := 0; i < gridSize; i++ {
		row := []string{}
		for j := 0; j < gridSize; j++ {
			row = append(row, " ")
		}
		grid = append(grid, row)

	}
	//Setup first location in locations to explore from
	start := point{x: int64((gridSize - 1) / 2), y: int64((gridSize - 1) / 2)}
	writeGrid(3, start)
	toExplore = make(map[string]*exploreLoc)
	e := &exploreLoc{pt: start, history: []int64{}}
	toExplore[getHash(start)] = e

	cnt := 0
	for len(toExplore) > 0 && cnt < maxCnt {
		curr = start
		key := ""
		var val *exploreLoc
		//get an element from the map
		for key, val = range toExplore {
			e = val
			break
		}

		A := intcode.Computer{}
		A.Initialize(inputStr)
		A.Input = make(chan int64, 10)
		A.Output = make(chan int64, 10)
		go func() {
			A.Run()
		}()
		fmt.Sprintf("Deleting toExplore loc %v\n", e)
		explore(&A, e)
		delete(toExplore, key)

		if cnt%100 == 0 && cnt != 0 {
			fmt.Printf("> i=%d\n", cnt)
			fmt.Printf("\n")
		}
		cnt++
	}
	printGrid()
}

const inputStr = `3,1033,1008,1033,1,1032,1005,1032,31,1008,1033,2,1032,1005,1032,58,1008,1033,3,1032,1005,1032,81,1008,1033,4,1032,1005,1032,104,99,102,1,1034,1039,1001,1036,0,1041,1001,1035,-1,1040,1008,1038,0,1043,102,-1,1043,1032,1,1037,1032,1042,1106,0,124,1001,1034,0,1039,102,1,1036,1041,1001,1035,1,1040,1008,1038,0,1043,1,1037,1038,1042,1106,0,124,1001,1034,-1,1039,1008,1036,0,1041,1002,1035,1,1040,102,1,1038,1043,1002,1037,1,1042,1106,0,124,1001,1034,1,1039,1008,1036,0,1041,1001,1035,0,1040,1001,1038,0,1043,1002,1037,1,1042,1006,1039,217,1006,1040,217,1008,1039,40,1032,1005,1032,217,1008,1040,40,1032,1005,1032,217,1008,1039,7,1032,1006,1032,165,1008,1040,33,1032,1006,1032,165,1101,2,0,1044,1105,1,224,2,1041,1043,1032,1006,1032,179,1102,1,1,1044,1105,1,224,1,1041,1043,1032,1006,1032,217,1,1042,1043,1032,1001,1032,-1,1032,1002,1032,39,1032,1,1032,1039,1032,101,-1,1032,1032,101,252,1032,211,1007,0,60,1044,1105,1,224,1101,0,0,1044,1106,0,224,1006,1044,247,101,0,1039,1034,101,0,1040,1035,1002,1041,1,1036,1002,1043,1,1038,101,0,1042,1037,4,1044,1105,1,0,92,17,17,33,88,37,85,63,23,14,79,46,37,69,8,6,63,55,61,21,86,19,37,78,49,15,54,28,54,94,91,14,11,40,56,96,20,20,82,28,12,91,68,43,18,63,16,82,71,8,83,88,25,79,67,26,55,33,51,74,68,59,64,58,78,30,65,64,9,48,87,26,85,32,82,92,21,34,99,1,20,66,34,85,65,58,87,12,21,13,51,90,54,19,12,85,3,88,47,31,93,95,49,70,95,55,7,67,2,92,42,80,88,42,24,91,2,59,41,41,70,89,42,83,43,92,44,93,62,26,63,99,81,35,98,70,71,79,8,90,26,66,94,22,47,55,90,93,6,87,92,88,40,73,40,97,14,73,90,31,92,16,35,93,36,27,69,57,97,80,34,58,42,95,34,9,93,22,94,45,79,32,33,90,72,77,58,29,63,56,95,37,61,58,51,57,8,25,86,75,25,63,64,93,57,7,79,85,57,53,97,16,63,40,71,52,23,33,75,13,56,65,90,26,12,66,93,26,36,64,30,10,75,18,77,76,86,33,98,4,23,52,64,66,82,38,90,17,63,94,24,97,20,92,70,63,80,19,73,8,74,93,16,98,77,52,38,90,46,49,76,84,53,50,22,93,19,16,61,47,54,67,56,78,21,77,52,88,4,64,91,90,10,97,10,51,89,15,57,97,22,79,59,92,17,84,71,30,96,58,82,52,93,48,20,62,4,89,64,53,85,37,92,52,89,43,80,86,2,41,81,53,53,82,77,31,66,92,31,44,81,14,49,96,66,42,91,2,61,82,36,32,90,8,61,32,67,52,25,81,15,63,27,59,61,1,15,88,87,62,10,85,47,75,24,46,63,24,77,34,73,34,45,71,10,96,46,43,75,31,23,72,37,87,57,88,63,30,6,86,91,16,53,16,89,81,11,32,75,22,82,69,50,88,53,67,50,65,67,26,81,83,20,14,23,89,98,57,64,3,79,7,69,89,57,1,61,65,14,52,76,66,83,3,57,90,82,53,13,72,94,37,26,97,77,32,53,43,78,22,36,65,83,98,55,82,58,48,24,68,92,18,22,90,65,28,81,33,63,79,3,31,65,92,53,46,74,7,80,37,79,79,83,42,82,84,33,21,79,79,21,81,55,4,95,10,53,84,14,25,86,65,24,74,53,26,61,47,19,66,86,58,99,37,83,35,46,3,11,89,27,66,53,33,67,8,95,44,45,70,71,65,59,49,77,25,3,56,83,39,91,3,52,86,67,57,99,86,40,39,3,99,25,69,94,93,62,36,37,91,17,26,80,98,77,15,5,90,25,40,69,11,85,66,56,40,83,61,10,85,33,28,86,26,41,61,4,86,78,20,71,78,47,94,39,92,26,61,91,52,69,20,47,45,99,38,96,39,98,76,58,28,94,27,47,97,2,45,54,64,94,98,27,69,54,23,72,89,96,22,58,21,16,79,28,45,55,78,75,15,92,67,10,81,80,64,61,13,30,98,65,57,35,4,22,96,72,92,47,51,87,33,78,26,83,20,5,93,22,73,83,68,24,17,61,69,39,62,53,20,95,84,53,83,36,48,99,33,13,42,90,97,87,9,55,64,34,94,7,78,62,42,43,83,54,82,57,24,36,98,95,54,63,75,52,15,40,92,87,77,5,13,93,48,82,71,65,97,96,1,3,68,49,97,9,77,88,99,25,78,4,84,97,77,4,92,91,76,53,71,58,64,55,68,97,96,48,99,2,86,51,69,15,72,42,72,44,86,55,73,0,0,21,21,1,10,1,0,0,0,0,0,0`
