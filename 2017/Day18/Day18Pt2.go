package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	msgQueueA       = []int{}
	msgQueueB       = []int{}
	numProgOneSends = 0
)

func main() {
	instructions := strings.Split(inputStr, "\n")
	instructionA := 0
	instructionB := 0
	deadlockA := false
	deadlockB := false
	registersA := make(map[string]int)
	registersB := make(map[string]int)
	registersA["p"] = 0
	registersB["p"] = 1
	cnt := 0
	for {
		instructionA, deadlockA = playMusicLine(registersA, instructions, instructionA, "0")
		instructionB, deadlockB = playMusicLine(registersB, instructions, instructionB, "1")
		if deadlockA && deadlockB {
			break
		}
		cnt++
		if cnt%1000000 == 0 {
			fmt.Printf("Encountered %d steps, msgA has %d entries, msgB has %d entries, and counts are %d\n", cnt, len(msgQueueA), len(msgQueueB), numProgOneSends)
		}
		if instructionA < 0 || instructionA >= len(instructions) || instructionB < 0 || instructionB >= len(instructions) {
			break
		}
		//fmt.Printf("Register A is: %+v\n", registersA)
		//fmt.Printf("Register B is: %+v\n", registersB)
	}
	fmt.Printf("Program 1 send %d msgs\n", numProgOneSends)
}

func playMusicLine(registers map[string]int, instructions []string, i int, id string) (int, bool) {
	line := instructions[i]
	lineSplit := strings.Fields(line)
	//fmt.Printf("Line number is %v, line is %v for program %v\n", i+1, lineSplit, id)

	switch lineSplit[0] {
	case "snd":
		val := registers[lineSplit[1]]
		switch id {
		case "0":
			msgQueueB = append(msgQueueB, val)
		case "1":
			msgQueueA = append(msgQueueA, val)
			numProgOneSends++
		default:
			fmt.Printf("Unidentified id %s", id)
		}
	case "set":
		val, err := getIntFromString(lineSplit[2])
		if err != nil {
			val = registers[lineSplit[2]]
		}
		registers[lineSplit[1]] = val
	case "add":
		val, err := getIntFromString(lineSplit[2])
		if err != nil {
			val = registers[lineSplit[2]]
		}
		registers[lineSplit[1]] += val
	case "mul":
		val, err := getIntFromString(lineSplit[2])
		if err != nil {
			val = registers[lineSplit[2]]
		}
		registers[lineSplit[1]] *= val
	case "mod":
		val, err := getIntFromString(lineSplit[2])
		if err != nil {
			val = registers[lineSplit[2]]
		}
		registers[lineSplit[1]] %= val
	case "jgz":
		check, err := getIntFromString(lineSplit[1])
		if err != nil {
			check = registers[lineSplit[1]]
		}
		if check > 0 {
			jump, err := getIntFromString(lineSplit[2])
			if err != nil {
				jump = registers[lineSplit[2]]
			}
			return i + jump, false
		}
	case "rcv":
		val := 0
		switch id {
		case "0":
			if len(msgQueueA) == 0 {
				return i, true
			}
			val, msgQueueA = msgQueueA[0], msgQueueA[1:]
			registers[lineSplit[1]] = val
		case "1":
			if len(msgQueueB) == 0 {
				return i, true
			}
			val, msgQueueB = msgQueueB[0], msgQueueB[1:]
			registers[lineSplit[1]] = val
		default:
			fmt.Printf("Unidentified ID %s\n", id)
		}
	default:
		fmt.Printf("UNKNOWN INSTRUCTION %s", line)
	}
	//fmt.Printf("Have registers: %+v\n", registers)
	return i + 1, false
}

func getIntFromString(tmpStr string) (int, error) {
	tmp, err := strconv.ParseInt(tmpStr, 10, 64)
	if err != nil {
		return -1, fmt.Errorf("%s", err.Error())
	}
	return int(tmp), nil
}

const (
	/*inputStr = `snd 1
	  snd 2
	  snd p
	  rcv a
	  rcv b
	  rcv c
	  rcv d`*/
	/*	inputStr = `set a 1
		add a 2
		mul a a
		mod a 5
		snd a
		set a 0
		rcv a
		jgz a -1
		set a 1
		jgz a -2`*/

	inputStr = `set i 31
set a 1
mul p 17
jgz p p
mul a 2
add i -1
jgz i -2
add a -1
set i 127
set p 464
mul p 8505
mod p a
mul p 129749
add p 12345
mod p a
set b p
mod b 10000
snd b
add i -1
jgz i -9
jgz a 3
rcv b
jgz b -1
set f 0
set i 126
rcv a
rcv b
set p a
mul p -1
add p b
jgz p 4
snd a
set a b
jgz 1 3
snd b
set f 1
add i -1
jgz i -11
snd a
jgz f -16
jgz a -19`
)
