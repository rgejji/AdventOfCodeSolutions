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

func createGrid() {
	for _, loc := range startingClaims {
		if loc.x > maxX {
			maxX = loc.x
		}
		if loc.y > maxY {
			maxY = loc.y
		}
	}
	grid = make([][]Property, maxX+1, maxX+1)
	for i := 0; i < maxX+1; i++ {
		grid[i] = make([]Property, maxY+1, maxY+1)
	}

	addPointClaimsToGrid(startingClaims, 0)
}

//addPointClaimsToGrid adds the points and owners to the grid
func addPointClaimsToGrid(points []Claim, currDistance int) ([]Claim, int) {
	numChange := 0
	dedupedClaimsMap := make(map[string]Claim)
	for _, loc := range points {
		//set the owner to unclaimed areas
		if grid[loc.x][loc.y].owner == 0 {
			grid[loc.x][loc.y].owner = loc.owner
			grid[loc.x][loc.y].distance = currDistance
			dedupedClaimsMap[loc.String()] = loc
			numChange++
		} else {
			//if we have claim conflict, set owner to -1
			if grid[loc.x][loc.y].distance == currDistance && grid[loc.x][loc.y].owner != loc.owner {
				grid[loc.x][loc.y].owner = -1
				dedupedClaimsMap[loc.String()] = Claim{x: loc.x, y: loc.y, owner: -1}
			}
		}
	}
	//Dedup claims
	dedupedClaims := []Claim{}
	for _, c := range dedupedClaimsMap {
		dedupedClaims = append(dedupedClaims, c)
	}

	return dedupedClaims, numChange
}

func fillGrid() {
	numChange := len(startingClaims)
	currDistance := 1
	currentClaims := startingClaims
	//While we keep writing new distances
	for numChange > 0 {
		//Update point locs with possible locations
		currentClaims = getNewClaims(currentClaims)
		//add points to grid and get back the new points to search
		currentClaims, numChange = addPointClaimsToGrid(currentClaims, currDistance)
		currDistance++
	}
}

func printGrid() {
	for j := 0; j <= maxY; j++ {
		for i := 0; i <= maxX; i++ {
			if grid[i][j].owner < 0 {
				fmt.Printf(". ")
			} else {
				fmt.Printf("%d ", grid[i][j].owner)
			}
		}
		fmt.Printf("\n")
	}
}

//Get new locs explores the grid for valid new location claims
func getNewClaims(locs []Claim) []Claim {
	newClaims := []Claim{}
	for _, l := range locs {
		//up
		if l.y-1 >= 0 {
			newClaims = append(newClaims, Claim{x: l.x, y: l.y - 1, owner: l.owner})
		}
		//down
		if l.y+1 <= maxY {
			newClaims = append(newClaims, Claim{x: l.x, y: l.y + 1, owner: l.owner})
		}

		//right
		if l.x+1 <= maxX {
			newClaims = append(newClaims, Claim{x: l.x + 1, y: l.y, owner: l.owner})
		}
		//left
		if l.x-1 >= 0 {
			newClaims = append(newClaims, Claim{x: l.x - 1, y: l.y, owner: l.owner})
		}
	}
	return newClaims
}

func countOwners() map[int]int {
	ownerCounts := make(map[int]int)

	for i := 0; i <= maxX; i++ {
		for j := 0; j <= maxY; j++ {
			owner := grid[i][j].owner
			if owner > 0 {
				if _, ok := ownerCounts[owner]; !ok {
					ownerCounts[owner] = 1
				} else {
					ownerCounts[owner]++
				}
			}
		}
	}
	return ownerCounts
}

func removeBorderOwners(ownerCounts map[int]int) {
	for i := 0; i <= maxX; i++ {
		owner := grid[i][0].owner
		if owner > 0 {
			delete(ownerCounts, owner)
		}
	}
	for i := 0; i <= maxX; i++ {
		owner := grid[i][maxY].owner
		if owner > 0 {
			delete(ownerCounts, owner)
		}
	}
	for j := 0; j <= maxY; j++ {
		owner := grid[0][j].owner
		if owner > 0 {
			delete(ownerCounts, owner)
		}
	}
	for j := 0; j <= maxY; j++ {
		owner := grid[maxX][j].owner
		if owner > 0 {
			delete(ownerCounts, owner)
		}
	}
}

func findMaxArea(ownerCounts map[int]int) int {
	max := 0
	for _, val := range ownerCounts {
		if val > max {
			max = val
		}
	}
	return max
}

func main() {
	//read in data, add to grid and stack
	readRows()
	//Figure out size of grid and make grid
	createGrid()

	//Fill out grid
	fillGrid()
	printGrid()
	//Count ownership
	ownerCounts := countOwners()
	//Any owner of a border will have infinite area since they can extend out indefinitely
	removeBorderOwners(ownerCounts)
	//Find max Owner
	maxArea := findMaxArea(ownerCounts)
	fmt.Printf("Found an area of size %d\n", maxArea)
}

const (

	/*inputStr = `1,1
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
