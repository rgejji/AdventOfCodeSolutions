package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func getNumberFromString(s string) int {
	rule := 0
	for i, char := range s {
		if i >= 5 {
			continue
		}
		rule = rule << 1
		if char == '#' {
			rule++
		}

	}
	return rule
}

func readInput() ([]int, map[int]bool) {
	rowSlice := strings.Split(inputStr, "\n")

	currState := []int{0, 0, 0, 0}
	rules := make(map[int]bool)
	for _, char := range rowSlice[1] {
		switch char {
		case '#':
			currState = append(currState, 1)
		case '.':
			currState = append(currState, 0)
		}
	}
	for i := 0; i < 4; i++ {
		currState = append(currState, 0)
	}

	//Read Rules
	for i := 2; i < len(rowSlice); i++ {
		row := rowSlice[i]
		rvec := strings.Split(row, "=")
		if len(rvec) != 2 {
			continue
		}
		if rvec[1] != "> #" {
			continue
		}

		rules[getNumberFromString(rvec[0])] = true
	}
	fmt.Printf("Have starting state: %+v\n", currState)

	return currState, rules
}

var firstPot = -4

func updateState(currState []int, rules map[int]bool) []int {
	//Default next state to a bunch of zeros of the same length as current state
	nextState := make([]int, len(currState))

	index := currState[0]<<4 + currState[1]<<3 + currState[2]<<2 + currState[3]<<1 + currState[4]
	for leftPlant := 0; leftPlant < len(currState)-4; leftPlant++ {
		if leftPlant != 0 {
			index = ((index & 0xF) << 1) + currState[leftPlant+4]
		}

		if rules[index] {
			nextState[leftPlant+2] = 1
		}
	}

	//Expand state as necessary
	for nextState[0]+nextState[1]+nextState[2]+nextState[3] > 0 {
		nextState = insert(nextState)
		firstPot--
	}

	L := len(nextState)
	for nextState[L-1]+nextState[L-2]+nextState[L-3]+nextState[L-4] > 0 {
		nextState = append(nextState, 0)
		L++
	}
	return nextState
}

func insert(nextState []int) []int {
	nextState = append(nextState, 0)
	copy(nextState[1:], nextState[0:])
	nextState[0] = 0
	return nextState
}

func scorePots(currState []int) int {
	sum := 0
	for loc, value := range currState {
		sum += value * (firstPot + loc)
	}
	return sum
}

func printState(iteration int, currState []int) {
	fmt.Printf("%d: ", iteration)
	for _, value := range currState {
		if value == 0 {
			fmt.Printf(".")
		} else {
			fmt.Printf("#")
		}
	}
	fmt.Printf("\n")
}

func main() {
	var numSteps int64
	numSteps = 10000
	currState, rules := readInput()
	startTime := time.Now()
	printState(0, currState)
	prev := 0
	for n := int64(1); n <= numSteps; n++ {
		currState = updateState(currState, rules)
		//fmt.Printf("%d Resulting Sum is %d\n", n, scorePots(currState))

		//Observe a pattern with period 15 (period found since there are 5 squares that alternate, it should be 2^k-1
		//Period is found to be length 15, with value 57- when large enough.
		//use y=mx+b to find answer at 50 billion
		//if n%(1<<4-1) == 0 {
		if float64(n) > math.Floor(0.9*float64(numSteps)) {
			tmp := scorePots(currState)
			fmt.Printf("%d: Resulting Sum is %d %d \n", n, tmp, tmp-prev)
			prev = tmp
		}
	}

	fmt.Printf("Time elapsed: %v\n", time.Since(startTime))
	fmt.Printf("Resulting Sum is %d\n", scorePots(currState))
}

const inputStr = `initial state:
##.#.#.##..#....######..#..#...#.#..#.#.#..###.#.#.#..#..###.##.#..#.##.##.#.####..##...##..#..##.#.

...## => #
#.#.# => #
.###. => #
#.#.. => .
.#..# => #
#..#. => #
..##. => .
....# => .
#.... => .
###.. => #
.#### => #
###.# => .
#..## => #
..... => .
##.## => #
####. => .
##.#. => .
#...# => .
##### => .
..#.. => .
.#.#. => .
#.### => .
.##.# => .
..#.# => .
.#.## => #
...#. => .
##... => #
##..# => #
.##.. => .
.#... => #
#.##. => #
..### => .`

/*
const inputStr = `initial state:
#..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #`
*/
