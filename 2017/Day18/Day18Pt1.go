package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	playMusic()
}

func playMusic() {
	registers := make(map[string]int)

	lastPlayed := -1
	instructions := strings.Split(inputStr, "\n")
	for i := 0; i < len(instructions); {
		line := instructions[i]
		lineSplit := strings.Fields(line)

		switch lineSplit[0] {
		case "snd":
			lastPlayed = registers[lineSplit[1]]
			fmt.Printf("Sound %d\n", lastPlayed)
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
				i += jump
				continue
			}
		case "rcv":
			if val := registers[lineSplit[1]]; val != 0 {
				fmt.Printf("Recovering last sound %d\n", lastPlayed)
				return
			}
		default:
			fmt.Printf("UNKNOWN INSTRUCTION %s", line)
		}
		//fmt.Printf("Have registers: %+v\n", registers)
		i++
	}
	fmt.Println("done\n")
	return
}

func getIntFromString(tmpStr string) (int, error) {
	tmp, err := strconv.ParseInt(tmpStr, 10, 64)
	if err != nil {
		return -1, fmt.Errorf("%s", err.Error())
	}
	return int(tmp), nil
}

const (
	/*inputStr = `set a 1
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
