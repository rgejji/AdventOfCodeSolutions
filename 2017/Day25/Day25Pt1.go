package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type state struct {
	label             string
	zeroWrite         uint8
	zeroMove          int
	zeroNextState     *state
	zeroNextStateChar string
	oneWrite          uint8
	oneMove           int
	oneNextState      *state
	oneNextStateChar  string
}

var (
	tape   []uint8
	states map[string]*state
)

func main() {
	//numSteps := 6
	numSteps := 12172063
	//tape := []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0}
	tape := make([]uint8, 2, 2)
	loc := len(tape) / 2
	//readStates(bluePrintA)
	readStates(bluePrint)

	currState := states["A"]
	for t := 0; t < numSteps; t++ {
		if loc == 0 {
			newtape := make([]uint8, len(tape), len(tape))
			loc = loc + len(tape)
			tape = append(newtape, tape...)
		}
		if loc == len(tape)-1 {
			newtape := make([]uint8, len(tape), len(tape))
			tape = append(tape, newtape...)
		}
		if tape[loc] == 0 {
			tape[loc] = currState.zeroWrite
			loc += currState.zeroMove
			currState = currState.zeroNextState
		} else {
			tape[loc] = currState.oneWrite
			loc += currState.oneMove
			currState = currState.oneNextState
		}
		if t%1000000 == 0 {
			fmt.Printf("Performed %d steps\n", t)
		}
	}
	if len(tape) < 20 {
		fmt.Printf("Have tape %v\n", tape)
	}

	//perform checksum
	total := 0
	for _, val := range tape {
		total += int(val)
	}
	fmt.Printf("Have total %d\n", total)

}

func getDirFromString(dir string) int {
	if dir == "right" {
		return 1
	}
	if dir == "left" {
		return -1
	}
	log.Fatalf("Undecipherable direction %s", dir)
	return 0
}

func readStates(input string) {
	cnt := 0
	lines := strings.Split(input, "\n")
	states = make(map[string]*state)
	for i := 3; i < len(lines); i += 10 {
		stateVal := getLastString(lines[i])
		zeroWrite := getLastString(lines[i+2])
		zeroMove := getLastString(lines[i+3])
		zeroNext := getLastString(lines[i+4])
		oneWrite := getLastString(lines[i+6])
		oneMove := getLastString(lines[i+7])
		oneNext := getLastString(lines[i+8])

		newState := &state{
			label:             stateVal,
			zeroWrite:         getUintFromString(zeroWrite),
			zeroMove:          getDirFromString(zeroMove),
			zeroNextStateChar: zeroNext,
			oneWrite:          getUintFromString(oneWrite),
			oneMove:           getDirFromString(oneMove),
			oneNextStateChar:  oneNext,
		}
		states[stateVal] = newState
		cnt++
	}
	ok := true
	for _, v := range states {
		if v.zeroNextState, ok = states[v.zeroNextStateChar]; !ok {
			log.Fatalf("Encountered unknown state char %v\n", v.zeroNextStateChar)
		}
		if v.oneNextState, ok = states[v.oneNextStateChar]; !ok {
			log.Fatalf("Encountered unknown state char %v\n", v.oneNextStateChar)
		}

	}

	fmt.Printf("Made %d states\n", cnt)
}

func getLastString(val string) string {
	slice := strings.Fields(val)
	fmt.Printf("SLICE HAS %v\n", slice)
	last := slice[len(slice)-1]
	last = strings.Trim(last, ":")
	last = strings.Trim(last, ".")
	return last
}

func getIntFromString(tmpStr string) int {
	tmp, err := strconv.ParseInt(tmpStr, 10, 64)
	if err != nil {
		fmt.Printf("Unable to parse %s\n", err.Error())
	}
	return int(tmp)
}
func getUintFromString(tmpStr string) uint8 {
	tmp, err := strconv.ParseInt(tmpStr, 10, 64)
	if err != nil {
		fmt.Printf("Unable to parse %s\n", err.Error())
	}
	return uint8(tmp)
}

const (
	bluePrintA = `Begin in state A.
	Perform a diagnostic checksum after 6 steps.
	
	In state A:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state B.
  If the current value is 1:
    - Write the value 0.
    - Move one slot to the left.
    - Continue with state B.

In state B:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the left.
    - Continue with state A.
  If the current value is 1:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state A.`

	bluePrint = `Begin in state A.
Perform a diagnostic checksum after 12172063 steps.

In state A:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state B.
  If the current value is 1:
    - Write the value 0.
    - Move one slot to the left.
    - Continue with state C.

In state B:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the left.
    - Continue with state A.
  If the current value is 1:
    - Write the value 1.
    - Move one slot to the left.
    - Continue with state D.

In state C:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state D.
  If the current value is 1:
    - Write the value 0.
    - Move one slot to the right.
    - Continue with state C.

In state D:
  If the current value is 0:
    - Write the value 0.
    - Move one slot to the left.
    - Continue with state B.
  If the current value is 1:
    - Write the value 0.
    - Move one slot to the right.
    - Continue with state E.

In state E:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state C.
  If the current value is 1:
    - Write the value 1.
    - Move one slot to the left.
    - Continue with state F.

In state F:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the left.
    - Continue with state E.
  If the current value is 1:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state A.`
)
