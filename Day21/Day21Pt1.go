package main

import (
	"fmt"
	"log"
	"strings"
)

var (
	startingPatternString = `.#.
	..#
	###`

	rulesMap map[string]runegrid
)

type runegrid [][]rune

func (rg runegrid) String() string {
	return getPatternString(rg)
}

func (rg runegrid) Count() int {
	count := 0
	for _, line := range rg {
		for _, char := range line {
			if char == '#' {
				count++
			}
		}
	}
	return count
}

func main() {
	//readRules(rulesA)
	readRules(rules)
	for k, v := range rulesMap {
		fmt.Printf("Have rule %v goes to %v\n", k, v)
	}

	size := 3
	currentPattern := allocatePattern(size)
	readPattern(currentPattern, startingPatternString, "\n")

	numIterations := 5
	for i := 0; i < numIterations; i++ {
		if size%2 == 0 {
			currentPattern, size = twoByTwoUpdate(currentPattern, size)
		} else {
			currentPattern, size = threeByThreeUpdate(currentPattern, size)
		}
		fmt.Printf("size is %d and count is %d\n", size, currentPattern.Count())
	}

	fmt.Printf("Final size is %d and count is %d\n", size, currentPattern.Count())
}

func readRules(rules string) {
	rulesMap = make(map[string]runegrid)
	//Here we read the rules as a hash, and store the enhancements for flips and rotateions
	lines := strings.Split(rules, "\n")
	for _, line := range lines {
		lineSplit := strings.Split(line, " => ")
		inputSplit := strings.Split(lineSplit[0], "/")
		numRows := len(inputSplit)

		//Read input and output pattern
		inputPattern := allocatePattern(numRows)
		readPattern(inputPattern, strings.TrimSpace(lineSplit[0]), "/")

		outputPattern := allocatePattern(numRows + 1)
		readPattern(outputPattern, strings.TrimSpace(lineSplit[1]), "/")

		//Hash input pattern
		//Add d4 symmetries of order 8
		fmt.Printf("ADDING BASE RULE WITH INPUT: %v\n", inputPattern)
		addRule(inputPattern, outputPattern)
		//Hash 3-rotations of input pattern
		rot1 := rotate(inputPattern)
		fmt.Printf("ADDING rot RULE WITH INPUT: %v\n", rot1)
		addRule(rot1, outputPattern)
		rot2 := rotate(rot1)
		fmt.Printf("ADDING rot RULE WITH INPUT: %v\n", rot2)
		addRule(rot2, outputPattern)
		rot3 := rotate(rot2)
		fmt.Printf("ADDING rot RULE WITH INPUT: %v\n", rot3)
		addRule(rot3, outputPattern)

		//Hash flip of input pattern
		//flipT := flipTop(inputPattern)
		//fmt.Printf("ADDING flip RULE WITH INPUT: %v\n", flipT)
		//addRule(flipT, outputPattern)
		flipS := flipSide(inputPattern)
		fmt.Printf("ADDING flip RULE WITH INPUT: %v\n", flipS)
		addRule(flipS, outputPattern)

		rot1flipS := rotate(flipS)
		fmt.Printf("ADDING rot-flip RULE WITH INPUT: %v\n", rot1flipS)
		addRule(rot1flipS, outputPattern)

		rot2flipS := rotate(rot1flipS)
		fmt.Printf("ADDING rot-flip RULE WITH INPUT: %v\n", rot2flipS)
		addRule(rot2flipS, outputPattern)

		rot3flipS := rotate(rot2flipS)
		fmt.Printf("ADDING rot-flip RULE WITH INPUT: %v\n", rot3flipS)
		addRule(rot3flipS, outputPattern)
	}

}

func flipTop(pattern runegrid) runegrid {
	size := len(pattern)
	flipped := allocatePattern(size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			newi := size - 1 - i
			flipped[newi][j] = pattern[i][j]
		}
	}
	return flipped
}

func flipSide(pattern runegrid) runegrid {
	size := len(pattern)
	flipped := allocatePattern(size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			newj := size - 1 - j
			flipped[i][newj] = pattern[i][j]
		}
	}
	return flipped
}

func addRule(input runegrid, output runegrid) {
	in := getPatternString(input)
	if val, ok := rulesMap[in]; ok {
		if val.String() != output.String() {
			fmt.Printf("Warning: input rule %s already encountered with different outputs:%+v\n%+v\n", in, val, output)
		}
	}
	rulesMap[in] = output
	return

}

func rotate(pattern runegrid) runegrid {
	size := len(pattern)
	shift := float64(size-1) / 2.0
	rotated := allocatePattern(size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			yloc := float64(i) - shift
			xloc := float64(j) - shift
			newxloc := int(yloc + shift + 0.5)
			newyloc := int(-xloc + shift + 0.5)
			rotated[newyloc][newxloc] = pattern[i][j]
			//fmt.Printf("Adding (%d,%d) from (%d, %d)\n", newxloc, newyloc, i, j)
		}
	}
	return rotated
}

func getPatternString(pattern runegrid) string {
	p := ""
	for i := 0; i < len(pattern); i++ {
		if i != 0 {
			p += "/"
		}
		for j := 0; j < len(pattern[i]); j++ {
			p += string(pattern[i][j])
		}
	}
	return p
}

func applySubPattern(newPattern, subPattern runegrid, blockI, blockJ int) {
	n := len(subPattern)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			newPattern[blockI*n+i][blockJ*n+j] = subPattern[i][j]
		}
	}

}

func enhance(pattern runegrid) runegrid {
	patternStr := getPatternString(pattern)
	val, ok := rulesMap[patternStr]
	if !ok {
		log.Fatalf("Error: Unknown rule for %s", patternStr)
	}
	return val
}

func extract(pattern runegrid, i, j, size int) runegrid {
	retPattern := allocatePattern(size)
	tmp := pattern[i : i+size]
	for i, line := range tmp {
		retPattern[i] = line[j : j+size]
	}
	return retPattern
}

func twoByTwoUpdate(pattern runegrid, size int) (runegrid, int) {
	newSize := 3 * (size >> 1)
	newPattern := allocatePattern(newSize)

	for i := 0; i < size/2; i++ {
		for j := 0; j < size/2; j++ {
			extractedPattern := extract(pattern, i*2, j*2, 2)
			newSubPattern := enhance(extractedPattern)
			applySubPattern(newPattern, newSubPattern, i, j)
		}
	}
	return newPattern, newSize
}
func threeByThreeUpdate(pattern runegrid, size int) (runegrid, int) {
	newSize := size / 3
	newSize *= 4
	newPattern := allocatePattern(newSize)
	for i := 0; i < size/3; i++ {
		for j := 0; j < size/3; j++ {
			extractedPattern := extract(pattern, i*3, j*3, 3)
			newSubPattern := enhance(extractedPattern)
			applySubPattern(newPattern, newSubPattern, i, j)
		}
	}
	return newPattern, newSize
}

func readPattern(pattern runegrid, patternString string, splitStr string) {
	for i, line := range strings.Split(patternString, splitStr) {
		cleanLine := strings.TrimSpace(line)
		for j, char := range cleanLine {
			pattern[i][j] = char
		}
	}
}

func allocatePattern(size int) runegrid {
	pattern := make(runegrid, size, size)
	for i := 0; i < size; i++ {
		pattern[i] = make([]rune, size, size)
	}
	return pattern
}

const (
	rules = `../.. => ..#/#../.#.
#./.. => #../#../...
##/.. => ###/#.#/#..
.#/#. => ###/##./.#.
##/#. => .../.#./..#
##/## => ##./#.#/###
.../.../... => ##../.#../#.#./....
#../.../... => ..../##.#/...#/##.#
.#./.../... => ###./####/#.../#..#
##./.../... => ###./.##./...#/..##
#.#/.../... => .###/.##./#.../#.##
###/.../... => ##.#/#..#/#.#./#.##
.#./#../... => #.#./.###/#.../#.##
##./#../... => #.../####/#.##/....
..#/#../... => #.##/..#./...#/...#
#.#/#../... => #.##/####/.#.#/#.#.
.##/#../... => #.../##../##.#/.##.
###/#../... => ..../#.#./.###/#...
.../.#./... => .#.#/#..#/##../#.##
#../.#./... => ###./.###/.#.#/..#.
.#./.#./... => ..##/.##./..##/.#.#
##./.#./... => ..#./##../###./...#
#.#/.#./... => ..##/.##./.###/###.
###/.#./... => ..#./.###/###./#.##
.#./##./... => ###./..../.#../#...
##./##./... => .#.#/##../##.#/...#
..#/##./... => ##.#/.##./.###/..##
#.#/##./... => .###/..#./#.##/####
.##/##./... => ##.#/..#./..##/###.
###/##./... => ..../.#.#/.#../#...
.../#.#/... => ###./.#.#/.#../#.##
#../#.#/... => ####/#..#/..../....
.#./#.#/... => #.../..##/#.##/#.#.
##./#.#/... => #.#./###./##../#.#.
#.#/#.#/... => ...#/.##./.##./.#..
###/#.#/... => ..../.##./####/#.#.
.../###/... => .###/.#../.###/#.##
#../###/... => ..##/..##/.##./##..
.#./###/... => .#.#/..#./..##/##.#
##./###/... => ...#/#.##/#.#./##.#
#.#/###/... => #.##/.##./...#/###.
###/###/... => ##../...#/..##/####
..#/.../#.. => #.##/#.../.#../#.#.
#.#/.../#.. => .##./.##./.#.#/.##.
.##/.../#.. => .#.#/#.##/...#/##.#
###/.../#.. => ##../..#./...#/##..
.##/#../#.. => ##../..##/#..#/#..#
###/#../#.. => ##../..#./#.#./....
..#/.#./#.. => .##./##.#/##../####
#.#/.#./#.. => ####/...#/.#.#/..#.
.##/.#./#.. => .#.#/..#./##.#/.#..
###/.#./#.. => #.../#.##/..../##.#
.##/##./#.. => #.#./#.#./#.##/#.#.
###/##./#.. => ...#/###./.##./.#.#
#../..#/#.. => ####/####/..../.##.
.#./..#/#.. => #.##/...#/..#./####
##./..#/#.. => ..#./#.../..##/####
#.#/..#/#.. => #.../#.##/#.##/..##
.##/..#/#.. => ####/..../##../####
###/..#/#.. => ..../##.#/.##./####
#../#.#/#.. => ...#/..##/###./#..#
.#./#.#/#.. => #..#/..#./.###/##.#
##./#.#/#.. => ###./####/#.##/..#.
..#/#.#/#.. => ##../##.#/..##/.##.
#.#/#.#/#.. => .#.#/.##./#.../##.#
.##/#.#/#.. => .#.#/#..#/.##./..#.
###/#.#/#.. => ...#/.#../.##./##.#
#../.##/#.. => ###./##../#.#./####
.#./.##/#.. => .#../##../#.#./.#.#
##./.##/#.. => ##.#/.#../.#.#/####
#.#/.##/#.. => ####/.#.#/..../....
.##/.##/#.. => ####/##../#..#/####
###/.##/#.. => .###/##.#/.#../#.##
#../###/#.. => #..#/###./####/.#.#
.#./###/#.. => ..##/##../##.#/.#.#
##./###/#.. => #..#/.#../####/...#
..#/###/#.. => ##../##.#/...#/#..#
#.#/###/#.. => ..#./.##./#..#/....
.##/###/#.. => #..#/#.../..../.#..
###/###/#.. => ..#./#.##/.##./#...
.#./#.#/.#. => .#.#/.##./##.#/.##.
##./#.#/.#. => #..#/.###/.#.#/.##.
#.#/#.#/.#. => #.../##../#.../.###
###/#.#/.#. => ###./.###/###./....
.#./###/.#. => .#../####/...#/##..
##./###/.#. => ####/###./..../....
#.#/###/.#. => ...#/.###/..../####
###/###/.#. => ..../#.../..#./.###
#.#/..#/##. => #.#./#.../####/#.##
###/..#/##. => .#.#/#..#/.###/#...
.##/#.#/##. => ..##/..#./..../##..
###/#.#/##. => #.#./##.#/####/#..#
#.#/.##/##. => ..../.#../#.#./##.#
###/.##/##. => ..../..../.#../##.#
.##/###/##. => #.#./.###/#.#./#.##
###/###/##. => ##.#/##.#/.###/..#.
#.#/.../#.# => #..#/.#../#.../...#
###/.../#.# => ##../.#../##.#/..#.
###/#../#.# => ..##/#.#./####/.#..
#.#/.#./#.# => ...#/...#/#..#/#.#.
###/.#./#.# => ..../####/.##./.#.#
###/##./#.# => #..#/.#.#/..##/####
#.#/#.#/#.# => #.#./..#./...#/.#..
###/#.#/#.# => ...#/##.#/.###/.#..
#.#/###/#.# => .#.#/###./.#../.##.
###/###/#.# => ...#/.###/.#.#/###.
###/#.#/### => #.##/.#.#/...#/.#..
###/###/### => ..##/.#../#.#./.#..`

	rulesA = `../.#  =>  ##./#../...
.#./..#/### => #..#/..../..../#..#`
)
