package main

import (
	"fmt"
	"strings"
)

func readInput() ([]int, [][]int) {
	rowSlice := strings.Split(inputStr, "\n")

	currState := []int{0, 0, 0, 0}
	rules := [][]int{}
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

		rule := []int{0, 0, 0, 0, 0}
		for i, char := range rvec[0] {
			if i >= 5 {
				continue
			}
			switch char {
			case '#':
				rule[i] = 1
			case '.':
				rule[i] = 0
			}

		}
		rules = append(rules, rule)

	}
	fmt.Printf("Have rules: v\n")
	for _, r := range rules {
		printState(-1, r)
	}
	fmt.Printf("Have starting state: %+v\n", currState)

	return currState, rules
}

var firstPot = -4

func updateState(currState []int, rules [][]int) []int {
	//Default next state to a bunch of zeros of the same length as current state
	nextState := make([]int, len(currState))
	//Go through each rule
	for _, rule := range rules {
		//calculate update for each state for this rule.
		//Only do a update for the given rule
		for leftPlant := 0; leftPlant < len(currState)-4; leftPlant++ {
			ruleSaysHasPlant := true
			for i, value := range rule {
				if value != currState[i+leftPlant] {
					ruleSaysHasPlant = false
					break
				}
			}
			if ruleSaysHasPlant {
				nextState[leftPlant+2] = 1
			}
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
	numSteps := 20
	currState, rules := readInput()
	printState(0, currState)
	for n := 1; n <= numSteps; n++ {
		currState = updateState(currState, rules)
		printState(n, currState)
	}

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
