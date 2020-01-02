package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	//"math"
	//"sync"
)

const (
	//positionMode indicates we use the position
	positionMode int = 0
)

var errEnd = errors.New("Exit")
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

func (p *computer) update() (int, error) {
	var value int
	var err error
	var accum, op, modeA, modeB, modeC, a, b, c int
	values := p.state
	i := p.loc
	for i < len(values) {
		if values[i] == 99 {
			value = 0
			err = errEnd
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
				value = -1
				err = fmt.Errorf("Unaccounted for value mode for modeC: %v", values)
				break
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
			input := p.inputs[0]
			p.inputs = p.inputs[1:]
			values[values[i+1]] = input
			fmt.Printf("Input: loc %d is now %d\n", values[i+1], input)
			i += 2
		case 4:
			fmt.Printf("> At i=%d value=%d\n", i, a)
			i += 2
			value = a
			break
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
			value = -1
			err = fmt.Errorf("Invalid operation %d at loc %d", op, c)
			break
		}
	}
	if i >= len(values) {
		err = fmt.Errorf("Unexpected end of program")
	}
	if err == nil {
		p.lastOutput = value
	}
	p.loc = i
	return value, err
}

type computer struct {
	loc        int
	state      []int
	inputs     []int
	lastOutput int
}

func (p *computer) initialize() {
	p.loc = 0
	p.state = readInput()
	p.inputs = []int{}
	p.lastOutput = -1
}

func main() {
	bestVal := -999999999999999999
	bestPhase := []int{}

	//Should use permutation iteration alg like Fisher Yaters, but we have small perms to iterate over
	/*for i := 5; i < 10; i++ {
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
	}*/
	bestPhase = []int{9, 8, 7, 6, 5}
	bestVal = runSim(bestPhase)

	fmt.Printf("FINAL ANSWER: %d with phase %v\n", bestVal, bestPhase)

}

//phaseSeq := []int{3, 1, 2, 4, 0}
//phaseSeq := []int{4, 3, 2, 1, 0}
//phaseSeq := []int{0, 1, 2, 3, 4}
//phaseSeq := []int{1, 0, 4, 3, 2}

func runSim(phaseSeq []int) int {
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

	A.inputs = append(A.inputs, phaseSeq[0])
	B.inputs = append(B.inputs, phaseSeq[1])
	C.inputs = append(C.inputs, phaseSeq[2])
	D.inputs = append(D.inputs, phaseSeq[3])
	E.inputs = append(E.inputs, phaseSeq[4])
	A.inputs = append(A.inputs, 0)

	done := false
	for !done {
		if _, err = A.update(); err != nil {
			done = true
			if err != errEnd {
				fmt.Printf("Error: Phase A update 1: %s", err.Error())
			}
		}
		B.inputs = append(B.inputs, A.lastOutput)
		if _, err = B.update(); err != nil {
			done = true
			if err != errEnd {
				fmt.Printf("Error: Phase B update: %s", err.Error())
			}
		}
		C.inputs = append(C.inputs, B.lastOutput)
		if _, err = C.update(); err != nil {
			done = true
			if err != errEnd {
				fmt.Printf("Error: Phase C update: %s", err.Error())
			}
		}
		D.inputs = append(D.inputs, C.lastOutput)
		if _, err = D.update(); err != nil {
			done = true
			if err != errEnd {
				fmt.Printf("Error: Phase D update: %s", err.Error())
			}
		}
		E.inputs = append(E.inputs, D.lastOutput)
		if _, err = E.update(); err != nil {
			done = true
			if err != errEnd {
				fmt.Printf("Error: Phase E update: %s", err.Error())
			}
		}
		A.inputs = append(A.inputs, E.lastOutput)
		fmt.Printf("LAST OUTPUTS: [%d, %d, %d, %d, %d]\n", A.lastOutput, B.lastOutput, C.lastOutput, D.lastOutput, E.lastOutput)
	}

	return E.lastOutput
}

const inputStr = `3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5`

//const inputStr = `3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0`
//const inputStr = `3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0`
//const inputStr = `3,8,1001,8,10,8,105,1,0,0,21,42,67,84,109,122,203,284,365,446,99999,3,9,1002,9,3,9,1001,9,5,9,102,4,9,9,1001,9,3,9,4,9,99,3,9,1001,9,5,9,1002,9,3,9,1001,9,4,9,102,3,9,9,101,3,9,9,4,9,99,3,9,101,5,9,9,1002,9,3,9,101,5,9,9,4,9,99,3,9,102,5,9,9,101,5,9,9,102,3,9,9,101,3,9,9,102,2,9,9,4,9,99,3,9,101,2,9,9,1002,9,3,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,99`
