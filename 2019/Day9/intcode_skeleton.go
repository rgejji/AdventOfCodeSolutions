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
	return int64(0)
}

func writeIntcodeOutput(o int64) {
}

func main() {

	A := computer{}
	A.initialize()

	A.run()
}
