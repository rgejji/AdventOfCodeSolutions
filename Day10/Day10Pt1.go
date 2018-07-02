package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	skipSize = 0
)

func main() {
	elements := []int{}
	for i := 0; i < numElements; i++ {
		elements = append(elements, i)
	}

	input := getInput(inputStr)

	currPosition := 0
	for _, instruction := range input {
		elements, currPosition = processList(elements, instruction, currPosition, numElements)
	}
	fmt.Printf("Final list %v\n", elements)
	fmt.Printf("Answer is %d\n", elements[0]*elements[1])

	fmt.Println("vim-go")
}

func processList(elements []int, instruction, currPosition, N int) ([]int, int) {
	updatedElements := reverseInPlace(elements, instruction, currPosition, N)
	newPosition := (currPosition + instruction + skipSize) % N
	skipSize = (skipSize + 1) % N
	return updatedElements, newPosition

}

func reverseInPlace(elements []int, subArraySize, loc, N int) []int {
	for count := 0; count < subArraySize/2; count++ {
		locA := (loc + count) % N
		locB := (loc + subArraySize - 1 - count + N) % N
		tmp := elements[locA]
		elements[locA] = elements[locB]
		elements[locB] = tmp
	}
	return elements
}
func getInput(inputStr string) []int {
	input := []int{}
	inputSplit := strings.Split(inputStr, ",")

	for _, val := range inputSplit {
		cleanedVal := strings.TrimSpace(val)
		tmp, err := strconv.ParseInt(cleanedVal, 10, 64)
		if err != nil {
			fmt.Printf("Unable to parse %s\n", err.Error())
			return input
		}
		input = append(input, int(tmp))
	}
	return input
}

const (
	//numElements = 5
	//inputStr    = `3, 4, 1, 5`
	numElements = 256
	inputStr    = `187,254,0,81,169,219,1,190,19,102,255,56,46,32,2,216`
)
