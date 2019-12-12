package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

//Claim is location
type Claim struct {
	x     int
	y     int
	owner int
}

func (c Claim) String() string {
	return fmt.Sprintf("(%d,%d)", c.x, c.y)
}

//Property contains the owner and distance on the gird
type Property struct {
	owner    int
	distance int
}

var (
	maxX           int //max X coordinate value
	maxY           int //max Y coordinate value
	startingClaims []Claim
	grid           [][]Property
)

func readRows() {
	//Read in Data
	rowSlice := strings.Split(inputStr, "\n")
	for ownerCnt, row := range rowSlice {
		locs := strings.Split(row, ",")
		x, err := strconv.Atoi(locs[0])
		if err != nil {
			log.Fatalf("Unable to parse x in string locs: %s", err.Error())
		}
		y, err := strconv.Atoi(strings.Trim(locs[1], " "))
		if err != nil {
			log.Fatalf("Unable to parse y in string locs: %s", err.Error())
		}
		startingClaims = append(startingClaims, Claim{x: x, y: y, owner: ownerCnt + 1})
	}
	return
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

//const limit = 10000
const limit = 10000

func findMaxArea() int {
	count := 0
	for i := -limit; i < limit+1000; i++ {
		for j := -limit; j < limit+1000; j++ {
			found := true
			totalDist := 0
			for _, claim := range startingClaims {
				totalDist += absDiff(i, claim.x) + absDiff(j, claim.y)
				if totalDist >= limit {
					found = false
					break
				}
			}
			if found {
				count++
			}
		}
		if i%500 == 0 {
			fmt.Printf("At i=%d\n", i)
		}
	}
	return count
}

func main() {
	//read in data, add to grid and stack
	readRows()

	//Calculate manahattan distance for all points
	maxArea := findMaxArea()
	fmt.Printf("Found an area of size %d\n", maxArea)
}

const (
	/*inputStr = `1, 1
	1, 6
	8, 3
	3, 4
	5, 5
	8, 9`*/
	inputStr = `137, 140
318, 75
205, 290
104, 141
163, 104
169, 164
238, 324
180, 166
260, 198
189, 139
290, 49
51, 350
51, 299
73, 324
220, 171
146, 336
167, 286
51, 254
40, 135
103, 138
100, 271
104, 328
80, 67
199, 180
320, 262
215, 290
96, 142
314, 128
162, 106
214, 326
303, 267
340, 96
211, 278
335, 250
41, 194
229, 291
45, 97
304, 208
198, 214
250, 80
200, 51
287, 50
120, 234
106, 311
41, 116
359, 152
189, 207
300, 167
318, 315
296, 72`
)
