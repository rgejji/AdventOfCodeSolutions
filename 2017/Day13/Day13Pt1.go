package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := strings.Split(inputStr, "\n")

	//lastLine := lines[len(lines)-1]
	//lastLayer := strings.TrimSpace(strings.Split(lastLine, ':')[0])

	//Construct layer array
	layerArray := []int{}
	currLoc := 0
	for _, line := range lines {
		lineSplit := strings.Split(line, ":")

		layerStr := strings.TrimSpace(lineSplit[0])
		layer := getIntFromString(layerStr)

		maxDepthStr := strings.TrimSpace(lineSplit[1])
		maxDepth := getIntFromString(maxDepthStr)
		for layer > currLoc {
			layerArray = append(layerArray, 0)
			currLoc++
		}
		layerArray = append(layerArray, maxDepth)
		currLoc++
	}

	//Perform simulation where i = time
	severity := 0
	for i, maxDepth := range layerArray {
		if 0 == getDepth(maxDepth, i) {
			severity += i * maxDepth
		}

	}
	fmt.Printf("Severity is %d\n", severity)
}
func getIntFromString(tmpStr string) int {
	tmp, err := strconv.ParseInt(tmpStr, 10, 64)
	if err != nil {
		fmt.Printf("Unable to parse %s\n", err.Error())
	}
	return int(tmp)
}

func getDepth(maxDepth int, time int) int {
	if maxDepth == 0 {
		return -1
	}
	tmp := time % (maxDepth*2 - 2)
	if tmp >= maxDepth {
		return maxDepth*2 - 2 - tmp
	}
	return tmp
}

const (
	/*inputStr = `0: 3
	1: 2
	4: 4
	6: 4`*/
	inputStr = `0: 3
	1: 2
	2: 5
	4: 4
	6: 6
	8: 4
	10: 8
	12: 8
	14: 6
	16: 8
	18: 6
	20: 6
	22: 8
	24: 12
	26: 12
	28: 8
	30: 12
	32: 12
	34: 8
	36: 10
	38: 9
	40: 12
	42: 10
	44: 12
	46: 14
	48: 14
	50: 12
	52: 14
	56: 12
	58: 12
	60: 14
	62: 14
	64: 12
	66: 14
	68: 14
	70: 14
	74: 24
	76: 14
	80: 18
	82: 14
	84: 14
	90: 14
	94: 17`
)
