package main

import (
	"fmt"
)

var (
	skipSize  = 0
	numRounds = 64
)

func main() {
	elements := []byte{}
	for i := 0; i < numElements; i++ {
		elements = append(elements, byte(i))
	}

	input := getInput(inputStr)
	//Append Addition
	input = append(input, 17, 31, 73, 47, 23)
	fmt.Printf("HAVE INPUT ARRAY %v\n", input)

	currPosition := 0
	for round := 0; round < numRounds; round++ {
		for _, instruction := range input {
			elements, currPosition = processList(elements, int(instruction), currPosition, numElements)
		}
	}
	fmt.Printf("Final list %v\n", elements)
	fmt.Printf("Answer is: ")
	for resultLoc := 0; resultLoc < 16; resultLoc++ {
		var val byte
		val = 0x00
		for loc := resultLoc * 16; loc < (resultLoc+1)*16; loc++ {
			val = val ^ elements[loc]
		}
		fmt.Printf("%.2x", val)
	}
	fmt.Println("\n")
}

func processList(elements []byte, instruction, currPosition, N int) ([]byte, int) {
	updatedElements := reverseInPlace(elements, instruction, currPosition, N)
	newPosition := (currPosition + instruction + skipSize) % N
	skipSize = (skipSize + 1) % N
	return updatedElements, newPosition

}

func reverseInPlace(elements []byte, subArraySize, loc, N int) []byte {
	for count := 0; count < subArraySize/2; count++ {
		locA := (loc + count) % N
		locB := (loc + subArraySize - 1 - count + N) % N
		tmp := elements[locA]
		elements[locA] = elements[locB]
		elements[locB] = tmp
	}
	return elements
}
func getInput(inputStr string) []byte {
	return []byte(inputStr)
}

const (
	numElements = 256
	//inputStr    = `1,2,3`
	//inputStr = ``
	//inputStr = `1,2,4`
	inputStr = `187,254,0,81,169,219,1,190,19,102,255,56,46,32,2,216`
)
