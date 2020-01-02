package intcode

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	//positionMode indicates we use the position
	positionMode     int64 = 0
	intermediateMode int64 = 1
	relativeMode     int64 = 2
	//Abort constant input to tell machine to abort
	Abort int64 = -999999999999
)

var numParams = []int64{1000, 4, 4, 2, 2, 3, 3, 4, 4, 2}

//Computer is the int code Computer
type Computer struct {
	ID           int64
	loc          int
	State        []int64
	LastOutput   int64
	relativeBase int
	Input        chan int64
	Output       chan int64
}

func readInput(inputStr string) []int64 {
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

//Run will execute the computer
func (s *Computer) Run() {
	var accum, a, b, op, modeA, modeB, modeC int64
	var aloc, bloc, cloc int
	i := s.loc
	for i < len(s.State) {
		if s.State[i] == 99 {
			break
		}

		accum = int64(s.State[i])
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
			//fmt.Printf("Waiting for input ... \n")
			input := s.readIntcodeInput()
			//fmt.Printf("Recieved input %v\n", input)
			if input == Abort {
				return
			}
			s.writeSafe(int(aloc), input)
			//fmt.Printf("Input: loc %d is now %d\n", aloc, input)
			i += 2
		case 4:
			//fmt.Printf("Constructing output... \n")
			s.writeIntcodeOutput(a)
			//fmt.Printf("Received output at i=%d value=%v\n", i, a)
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

func (s *Computer) getLocation(start int, mode int64) int {
	loc := start
	switch mode {
	case positionMode:
		loc = int(s.readSafe(loc))
	case relativeMode:
		loc = int(s.readSafe(loc)) + s.relativeBase
	}
	return loc
}

func (s *Computer) readSafe(loc int) int64 {
	if loc >= len(s.State) {
		return 0
	}
	return s.State[loc]
}

func (s *Computer) writeSafe(loc int, value int64) {
	if loc < 0 {
		panic(fmt.Sprintf("Unable to write to location %d", loc))
	}

	for loc >= len(s.State) {
		s.State = append(s.State, 0)
	}
	s.State[loc] = value
}

//Initialize will setup the computer
func (s *Computer) Initialize(input string) {
	s.loc = 0
	s.State = readInput(input)
	s.relativeBase = 0
	s.LastOutput = -1
}

//Send next robot command
func (s *Computer) readIntcodeInput() int64 {
	return <-s.Input
}

func (s *Computer) writeIntcodeOutput(o int64) {
	s.LastOutput = o
	s.Output <- o
}

//GetASCIIFromOutput returns the ASCII string representation for the int
//e.g. 35 turns into a #
func GetASCIIFromOutput(o int64) string {
	return string([]byte{byte(o)})
}

//InputString will input an Ascii string
func (s *Computer) InputString(input string) {
	for _, r := range input {
		s.Input <- GetIntFromRune(r)
	}
}

//GetIntFromRune returns the int64
func GetIntFromRune(s rune) int64 {
	return int64(s)
}
