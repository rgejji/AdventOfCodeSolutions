package main

import (
	"fmt"
	"strconv"
	"strings"
	//"math"
	"sync"
)

const (
	//positionMode indicates we use the position
	positionMode int = 0
)

var numParams = []int{1000, 4, 4, 2, 2, 3, 3, 4, 4}
var wg sync.WaitGroup

func readInput() []int {
	values := []int{}
	slice := strings.Split(inputStr, ",")
	for _, value := range slice {
		v, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("Unable to parse value from '%s': %s\n", value, err.Error())
			return values
		}
		values = append(values, v)
	}
	return values
}

func modifyInput(values []int, noun, verb int) {
	values[1] = noun
	values[2] = verb
}

type computer struct {
	id         int
	loc        int
	state      []int
	input      chan int
	lastOutput int
}

func (s *computer) update(output chan int) error {
	var accum, op, modeA, modeB, modeC, a, b, c int
	values := s.state
	i := s.loc

	for i < len(values) {
		if values[i] == 99 {
			return nil
		}

		accum = values[i]
		op = accum % 100
		accum = accum / 100
		modeA = accum % 10
		accum = accum / 10
		modeB = accum % 10
		accum = accum / 10
		modeC = accum % 10

		a = values[i+1]
		if modeA == positionMode {
			a = values[a]
		}

		if numParams[op] >= 3 {
			b = values[i+2]
			if modeB == positionMode {
				b = values[b]
			}
		}
		if numParams[op] >= 4 {
			c = values[i+3]
			if modeC != positionMode {
				return fmt.Errorf("Unaccounted for value mode for modeC: %v", values)
			}
		}

		switch op {
		case 1:
			values[c] = a + b
			i += 4
		case 2:
			values[c] = a * b
			i += 4
		case 3:
			values[values[i+1]] = <-s.input
			//fmt.Printf("Input: loc %d is now %d\n", values[i+1], values[values[i+1]])
			i += 2
		case 4:
			//fmt.Printf("> At i=%d value=%d\n", i, a)
			s.lastOutput = a
			output <- a
			i += 2
		case 5:
			i += 3
			if a != 0 {
				i = b
			}
		case 6:
			i += 3
			if a == 0 {
				i = b
			}
		case 7:
			if a < b {
				values[c] = 1
			} else {
				values[c] = 0
			}
			i += 4
		case 8:
			if a == b {
				values[c] = 1
			} else {
				values[c] = 0
			}
			i += 4
		default:
			return fmt.Errorf("Invalid operation %d at loc %d", op, c)
		}
	}
	return fmt.Errorf("Unexpected end of program")
}

func (s *computer) initialize() {
	s.loc = 0
	s.state = readInput()
	s.input = make(chan int, 10)
	s.lastOutput = -1
}

func main() {
	bestVal := -999999999999999999
	bestPhase := []int{}

	//Should use permutation iteration alg like Fisher Yaters, but we have small perms to iterate over
	for i := 5; i < 10; i++ {
		for j := 5; j < 10; j++ {
			fmt.Printf("TICK\n")
			if i == j {
				continue
			}
			for k := 5; k < 10; k++ {
				if k == i || k == j {
					continue
				}
				for l := 5; l < 10; l++ {
					if l == i || l == j || l == k {
						continue
					}
					for m := 5; m < 10; m++ {
						if m == i || m == j || m == k || m == l {
							continue
						}
						phase := []int{i, j, k, l, m}
						currVal := runSim(phase)
						if currVal > bestVal {
							bestVal = currVal
							bestPhase = phase
						}
					}
				}
			}
		}
	}
	//bestPhase = []int{9, 8, 7, 6, 5}
	//bestVal = runSim(bestPhase)

	fmt.Printf("FINAL ANSWER: %d with phase %v\n", bestVal, bestPhase)

}

func pipeOutput(inComp *computer, outComp *computer) {
	if err := inComp.update(outComp.input); err != nil {
		fmt.Printf("Error: Could not run current computer %d: %s", inComp.id, err.Error())
	}

	//fmt.Printf("Phase %d is complete. Exiting\n", inComp.id)
	wg.Done()
}

func runSim(phaseSeq []int) int {
	A := computer{id: 0}
	A.initialize()
	B := computer{id: 1}
	B.initialize()
	C := computer{id: 2}
	C.initialize()
	D := computer{id: 3}
	D.initialize()
	E := computer{id: 4}
	E.initialize()

	wg.Add(1)
	go func() {
		A.input <- phaseSeq[0]
		B.input <- phaseSeq[1]
		C.input <- phaseSeq[2]
		D.input <- phaseSeq[3]
		E.input <- phaseSeq[4]
		A.input <- 0
		wg.Done()
	}()
	wg.Wait()
	wg.Add(5)

	//Can truly parallelize this by putting updates in individual threads and passing channels
	go pipeOutput(&A, &B)
	go pipeOutput(&B, &C)
	go pipeOutput(&C, &D)
	go pipeOutput(&D, &E)
	go pipeOutput(&E, &A)

	wg.Wait()
	//fmt.Printf("LAST OUTPUTS: [%d, %d, %d, %d, %d]\n", A.lastOutput, B.lastOutput, C.lastOutput, D.lastOutput, E.lastOutput)

	return E.lastOutput
}

//const inputStr = `3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5`
//const inputStr = `3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10`

const inputStr = `3,8,1001,8,10,8,105,1,0,0,21,42,67,84,109,122,203,284,365,446,99999,3,9,1002,9,3,9,1001,9,5,9,102,4,9,9,1001,9,3,9,4,9,99,3,9,1001,9,5,9,1002,9,3,9,1001,9,4,9,102,3,9,9,101,3,9,9,4,9,99,3,9,101,5,9,9,1002,9,3,9,101,5,9,9,4,9,99,3,9,102,5,9,9,101,5,9,9,102,3,9,9,101,3,9,9,102,2,9,9,4,9,99,3,9,101,2,9,9,1002,9,3,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,99`
