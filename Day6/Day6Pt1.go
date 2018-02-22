package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	seenVals := make(map[string][]int64)
	currVal := convertInputStr(inputStr)
	currStep := 0
	for {
		if currStep%1000 == 0 {
			fmt.Printf("On step %d\n", currStep)
		}

		currStr := intSliceToString(currVal)
		if _, ok := seenVals[currStr]; ok {
			fmt.Printf("Found Val on Step %d\n", currStep)
			break
		}
		//fmt.Printf("Have string %s\n", intSliceToString(currVal))
		seenVals[currStr] = currVal

		currVal = iterate(currVal)
		currStep++
	}

	fmt.Println("vim-go")
}

//Function taken from snowman's one liner transform int to string stackexchange.
func intSliceToString(intSlice []int64) string {
	delim := ","
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(intSlice)), delim), "[]")
}

func iterate(input []int64) []int64 {
	N := len(input)
	if N < 1 {
		fmt.Println("ERROR, need a non-empty input")
		return []int64{}
	}

	output := make([]int64, len(input))
	max := input[0]
	maxLoc := 0
	//copy and find max
	for k, v := range input {
		output[k] = v
		if v > max {
			max = v
			maxLoc = k
		}
	}

	//clear value
	bucket := max
	output[maxLoc] = 0

	//spread out
	for currentLoc := (maxLoc + 1) % N; bucket > 0; currentLoc = (currentLoc + 1) % N {
		output[currentLoc]++
		bucket--
	}
	return output
}

func convertInputStr(inputStr string) []int64 {
	vals := []int64{}
	strNums := strings.Fields(inputStr)
	for _, strNum := range strNums {
		currVal, err := strconv.ParseInt(strNum, 10, 64)
		if err != nil {
			fmt.Printf("ERROR: ENCOUNTERED UNPARSABLE STR %s:%s\n", strNum, err.Error())
		}
		vals = append(vals, currVal)
	}
	return vals
}

//Debrecated in factor of one liner intslice to string method
/*
func checkSame(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if v != b[k] {
			return false
		}
	}
	return true
}*/

const (
	//inputStr = `0 2 7 0`
	inputStr = `4	10	4	1	8	4	9	14	5	1	14	15	0	15	3	5`
)
