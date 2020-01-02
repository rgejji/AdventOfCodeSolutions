package main

import (
	"fmt"
	"strconv"
	"strings"
	//"math"
)

var values []int

func readInput() {
	slice := strings.Split(inputStr, ",")
	for _, value := range slice {
		v, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("Unable to parse value from '%s': %s\n", value, err.Error())
			return
		}
		values = append(values, v)
	}
}

func modifyInput() {
	values[1] = 12
	values[2] = 2
}

func calculate() {
	for i := 0; i+3 < len(values); i += 4 {
		if values[i] == 99 {
			break
		}
		op := values[i]
		a := values[i+1]
		b := values[i+2]
		c := values[i+3]
		if a < 0 || a >= len(values) {
			fmt.Printf("Error: Read loc A at i=%d is beyond bdry", i+1)
			break
		}
		if b < 0 || b >= len(values) {
			fmt.Printf("Error: Read loc B at i=%d is beyond bdry", i+2)
			break
		}

		if c < 0 || c >= len(values) {
			fmt.Printf("Error: Write location at i=%d is beyond bdry", i+3)
			break
		}

		switch op {
		case 1:
			values[c] = values[a] + values[b]
		case 2:
			values[c] = values[a] * values[b]
		default:
			fmt.Printf("Error: Invalid operation %d at loc %d", c, op)
		}

	}
}

func main() {
	readInput()
	modifyInput()
	calculate()
	fmt.Printf("Initial Value is %d\n", values[0])
}

//const inputStr = `1,9,10,3,2,3,11,0,99,30,40,50`
const inputStr = `1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,10,1,19,1,19,9,23,1,23,6,27,2,27,13,31,1,10,31,35,1,10,35,39,2,39,6,43,1,43,5,47,2,10,47,51,1,5,51,55,1,55,13,59,1,59,9,63,2,9,63,67,1,6,67,71,1,71,13,75,1,75,10,79,1,5,79,83,1,10,83,87,1,5,87,91,1,91,9,95,2,13,95,99,1,5,99,103,2,103,9,107,1,5,107,111,2,111,9,115,1,115,6,119,2,13,119,123,1,123,5,127,1,127,9,131,1,131,10,135,1,13,135,139,2,9,139,143,1,5,143,147,1,13,147,151,1,151,2,155,1,10,155,0,99,2,14,0,0`
