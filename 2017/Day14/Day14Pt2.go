package main

import (
	"fmt"

	aoc "github.com/RichardGejji/AdventOfCode2017/Library"
)

var (
	byteFromChar = map[rune]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'a': 10,
		'b': 11,
		'c': 12,
		'd': 13,
		'e': 14,
		'f': 15,
	}
)

const (
	numRows = 128
)

func main() {
	mask := 8
	count := 0

	connectionMap := make(map[string][]string)
	grid := [numRows][numRows]string{}

	for i := 0; i < numRows; i++ {
		input := fmt.Sprintf("%s-%d", inputStr, i)
		hash := aoc.GetKnotHash(input)
		fmt.Printf("Input %s and output %s\n", input, hash)
		for byteNum, b := range hash {
			tmp := byteFromChar[b]
			for j := 0; j < 4; j++ {
				if tmp&mask == mask {
					grid[i][byteNum*4+j] = "#"
					count++
				} else {
					grid[i][byteNum*4+j] = "."
				}

				tmp = tmp << 1
			}
		}
	}
	fmt.Printf("Have total number of valid spots %d\n", count)

	//Calculate connection map
	for i := 0; i < numRows; i++ {
		for j := 0; j < numRows; j++ {
			//Skip . grid connections
			if grid[i][j] != "#" {
				continue
			}

			nodeStr := intToStr(i*numRows + j)
			currList := []string{}
			if i > 0 && grid[i-1][j] == "#" {
				currList = append(currList, intToStr((i-1)*numRows+j))
			}
			if i < numRows-1 && grid[i+1][j] == "#" {
				currList = append(currList, intToStr((i+1)*numRows+j))
			}
			if j > 0 && grid[i][j-1] == "#" {
				currList = append(currList, intToStr(i*numRows+j-1))
			}
			if j < numRows-1 && grid[i][j+1] == "#" {
				currList = append(currList, intToStr(i*numRows+j+1))
			}

			connectionMap[nodeStr] = currList
		}
	}
	numClusters := aoc.CountNumberOfClusters(connectionMap)
	fmt.Printf("Have %d clusters\n", numClusters)
}

func intToStr(val int) string {
	return fmt.Sprintf("%d", val)
}

const (
	//inputStr = "flqrgnkx"
	inputStr = "ljoxqyyw"
)
