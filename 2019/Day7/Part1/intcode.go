package main

import (
	"fmt"
	"strconv"
	"strings"
	//"math"
)

const (
	//positionMode indicates we use the position
	positionMode int = 0
)

var numParams = []int{1000, 4, 4, 2, 2, 3, 3, 4, 4}

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

func calculate(values []int, i int, input chan int) (int, int, error) {
	var accum, op, modeA, modeB, modeC, a, b, c int

	for i < len(values) {
		if values[i] == 99 {
			break
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
				return -1, i, fmt.Errorf("Unaccounted for value mode for modeC: %v", values)
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
			values[values[i+1]] = <-input
			//fmt.Printf("Input: loc %d is now %d\n", values[i+1], values[values[i+1]])
			i += 2
		case 4:
			//fmt.Printf("> At i=%d value=%d\n", i, a)
			i += 2
			return a, i, nil
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
			return -1, i, fmt.Errorf("Invalid operation %d at loc %d", op, c)
		}

	}
	return 0, i, fmt.Errorf("Unexpected end of program")
}

type computer struct {
	loc   int
	state []int
	input chan int
}

func (c *computer) update() (int, error) {
	var value int
	var err error
	value, c.loc, err = calculate(c.state, c.loc, c.input)
	return value, err
}

func (c *computer) initialize() {
	c.loc = 0
	c.state = readInput()
	c.input = make(chan int)
}

func main() {
	bestVal := 0
	bestPhase := []int{}

	//Should use permutation iteration alg like Fisher Yaters, but we have small perms to iterate over
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("TICK\n")
			if i == j {
				continue
			}
			for k := 0; k < 5; k++ {
				if k == i || k == j {
					continue
				}
				for l := 0; l < 5; l++ {
					if l == i || l == j || l == k {
						continue
					}
					for m := 0; m < 5; m++ {
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

	fmt.Printf("FINAL ANSWER: %d with phase %v\n", bestVal, bestPhase)

}

//phaseSeq := []int{3, 1, 2, 4, 0}
//phaseSeq := []int{4, 3, 2, 1, 0}
//phaseSeq := []int{0, 1, 2, 3, 4}
//phaseSeq := []int{1, 0, 4, 3, 2}

func runSim(phaseSeq []int) int {
	var val int
	var err error

	A := computer{}
	A.initialize()
	B := computer{}
	B.initialize()
	C := computer{}
	C.initialize()
	D := computer{}
	D.initialize()
	E := computer{}
	E.initialize()

	go func() {
		A.input <- phaseSeq[0]
		A.input <- 0
	}()
	if val, err = A.update(); err != nil {
		fmt.Printf("Error: Phase A update 1: %s", err.Error())
	}
	go func() {
		B.input <- phaseSeq[1]
		B.input <- val
	}()
	if val, err = B.update(); err != nil {
		fmt.Printf("Error: Phase B update: %s", err.Error())
	}
	go func() {
		C.input <- phaseSeq[2]
		C.input <- val
	}()
	if val, err = C.update(); err != nil {
		fmt.Printf("Error: Phase C update: %s", err.Error())
	}
	go func() {
		D.input <- phaseSeq[3]
		D.input <- val
	}()
	if val, err = D.update(); err != nil {
		fmt.Printf("Error: Phase D update: %s", err.Error())
	}
	go func() {
		E.input <- phaseSeq[4]
		E.input <- val
	}()
	if val, err = E.update(); err != nil {
		fmt.Printf("Error: Phase E update: %s", err.Error())
	}
	return val
}

//const inputStr = `3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0`
//const inputStr = `3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0`
//const inputStr = `3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0`
const inputStr = `3,8,1001,8,10,8,105,1,0,0,21,42,67,84,109,122,203,284,365,446,99999,3,9,1002,9,3,9,1001,9,5,9,102,4,9,9,1001,9,3,9,4,9,99,3,9,1001,9,5,9,1002,9,3,9,1001,9,4,9,102,3,9,9,101,3,9,9,4,9,99,3,9,101,5,9,9,1002,9,3,9,101,5,9,9,4,9,99,3,9,102,5,9,9,101,5,9,9,102,3,9,9,101,3,9,9,102,2,9,9,4,9,99,3,9,101,2,9,9,1002,9,3,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,99`
