package aoc2017

import (
	"fmt"
)

const (
	numRounds   = 64
	numElements = 256
)

//GetKnotHash calculates the knot hash
func GetKnotHash(inputStr string) string {
	elements := []byte{}
	for i := 0; i < numElements; i++ {
		elements = append(elements, byte(i))
	}
	input := getInput(inputStr)
	//Append Addition
	input = append(input, 17, 31, 73, 47, 23)

	currPosition := 0
	skipSize := 0
	for round := 0; round < numRounds; round++ {
		for _, instruction := range input {
			elements, currPosition, skipSize = processList(elements, int(instruction), currPosition, numElements, skipSize)
		}
	}
	result := ""
	for resultLoc := 0; resultLoc < 16; resultLoc++ {
		var val byte
		val = 0x00
		for loc := resultLoc * 16; loc < (resultLoc+1)*16; loc++ {
			val = val ^ elements[loc]
		}
		char := fmt.Sprintf("%.2x", val)
		result += char
	}
	return result

}

func processList(elements []byte, instruction, currPosition, N, skipSize int) ([]byte, int, int) {
	updatedElements := reverseInPlace(elements, instruction, currPosition, N)
	newPosition := (currPosition + instruction + skipSize) % N
	skipSize = (skipSize + 1) % N
	return updatedElements, newPosition, skipSize

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
