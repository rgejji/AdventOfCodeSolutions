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

func main() {
	mask := 8
	numRows := 128
	count := 0
	for i := 0; i < numRows; i++ {
		input := fmt.Sprintf("%s-%d", inputStr, i)
		hash := aoc.GetKnotHash(input)
		fmt.Printf("Input %s and output %s\n", input, hash)
		for _, b := range hash {
			tmp := byteFromChar[b]
			for j := 0; j < 4; j++ {
				if tmp&mask == mask {
					//fmt.Printf("#")
					count++
				} else {
					//fmt.Printf(".")
				}

				tmp = tmp << 1
			}
		}
		fmt.Printf("\n")

	}
	fmt.Printf("Have total count %d\n", count)

}

const (
	//inputStr = "flqrgnkx"
	inputStr = "ljoxqyyw"
)
