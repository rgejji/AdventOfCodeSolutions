package main

import (
	"container/heap"
	"fmt"
	"strings"
	"unicode"
	//graphutil "github.com/RichardGejji/AdventOfCode2019/GraphUtil"
	//intcode "github.com/RichardGejji/AdventOfCode2019/Intcode"
)

var maxHeapSize = 500000

//var maxHeapSize = 250000
//var maxHeapSize = 30000

//var maxHeapSize = 10000
var maxNumIter = 5000

const (
	north int64 = 0
	east  int64 = 1
	south int64 = 2
	west  int64 = 3
)

const (
	aASCII uint = 97
	zASCII uint = 122
)

var grid [][]rune

var start point
var paths *PathHeap
var totalKeys int64

func keyToLoc(a rune) uint {
	return uint(a) - aASCII
}

func readInput() {
	rowSlice := strings.Split(inputStr, "\n")
	grid = [][]rune{}
	for i, row := range rowSlice {
		gridRow := []rune{}
		for j, val := range row {
			if val == '@' {
				start = point{y: int64(i), x: int64(j)}
			}
			if uint(val) >= aASCII && uint(val) <= zASCII {
				totalKeys++
			}

			gridRow = append(gridRow, val)
		}
		grid = append(grid, gridRow)
	}
	//fill grid dead ends
	filledRecent := true
	for filledRecent {
		filledRecent = false

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid); j++ {
				cnt := 0
				if grid[i][j] != '.' {
					continue
				}
				if i-1 >= 0 && grid[i-1][j] == '#' {
					cnt++
				}
				if i+1 < len(grid) && grid[i+1][j] == '#' {
					cnt++
				}
				if j-1 >= 0 && grid[i][j-1] == '#' {
					cnt++
				}
				if j+1 < len(grid) && grid[i][j+1] == '#' {
					cnt++
				}
				if cnt >= 3 || (cnt == 2 && (i == 0 || j == 0 || i == len(grid)-1 || j == len(grid)-1)) {
					grid[i][j] = '#'
					filledRecent = true
				}

			}
		}

	}
}

func printGrid() {
	for i := 0; i < len(grid); i++ {
		fmt.Printf("%s\n", string(grid[i]))
	}
}

func iterateAllPaths() {
	var found bool
	for i := 0; i < maxNumIter; i++ {
		maxKeys := int64(0)
		nextRound := []*path{}
		pathIDs := make(map[string]int64)

		//Update each path 1x
		for paths.Len() > 0 {
			curr := heap.Pop(paths).(path)
			//fmt.Printf("Score at %v is %d with keys %s, %v and num moves %d\n", curr.pt, curr.Score(), string(curr.keysOrdered), curr.moveHistory, curr.numMoves)

			cnumKeys := curr.numKeys()
			if totalKeys == cnumKeys {
				fmt.Printf("Found path %s with move history %v and num moves: %d. %d paths remaining\n", string(curr.keysOrdered), curr.moveHistory, curr.numMoves, paths.Len())
				found = true
			}
			nextRound = append(nextRound, getUpdatedPaths(&curr)...)
			if cnumKeys > maxKeys {
				maxKeys = cnumKeys
			}
		}
		if found {
			return
		}

		//Keep best paths with same key history
		for _, p := range nextRound {
			s := string(p.keysOrdered)
			currScore, ok := pathIDs[s]
			if !ok || p.Score() > currScore {
				pathIDs[s] = p.Score()
			}
		}
		justAdd := len(nextRound) < maxHeapSize
		for _, p := range nextRound {
			if justAdd || p.Score() == pathIDs[string(p.keysOrdered)] {
				heap.Push(paths, *p)
			} /*else {
				fmt.Printf("Removing %s with score %d when we want score %d\n", string(p.keysOrdered), p.Score(), pathIDs[string(p.keysOrdered)])
			}*/
		}

		//If we still need to  remove, more history, remove.
		//Cull heaps with poor key scores or repeat scores
		fmt.Printf(">i=%d, have Heap Size %d and max keys found %d\n", i, paths.Len(), maxKeys)
		if paths.Len() == 0 {
			return
		}
		for paths.Len() > maxHeapSize {
			heap.Pop(paths)
			/*rem := heap.Pop(paths).(path)
			//if strings.HasPrefix(string(rem.keysOrdered), "acfidgb") {
			if strings.HasPrefix(string(rem.keysOrdered), "afbj") {

				fmt.Printf("Score at %v is %d with keys %s, %v and num moves %d\n", rem.pt, rem.Score(), string(rem.keysOrdered), rem.moveHistory, rem.numMoves)
				return
			}*/
		}
	}
	fmt.Printf("Out of paths\n")
}

//check around current loc of paths to goto
func getUpdatedPaths(curr *path) []*path {
	newPaths := []*path{}
	northPt := getPtFromDirection(curr.pt, north)
	southPt := getPtFromDirection(curr.pt, south)
	westPt := getPtFromDirection(curr.pt, west)
	eastPt := getPtFromDirection(curr.pt, east)
	if isValidPoint(northPt, curr.keys) && (curr.lastDir != south || isNewKey(curr)) {
		newPaths = append(newPaths, newPathAtPt(curr, northPt, north))
	}
	if isValidPoint(southPt, curr.keys) && (curr.lastDir != north || isNewKey(curr)) {
		newPaths = append(newPaths, newPathAtPt(curr, southPt, south))
	}
	if isValidPoint(westPt, curr.keys) && (curr.lastDir != east || isNewKey(curr)) {
		newPaths = append(newPaths, newPathAtPt(curr, westPt, west))
	}
	if isValidPoint(eastPt, curr.keys) && (curr.lastDir != west || isNewKey(curr)) {
		newPaths = append(newPaths, newPathAtPt(curr, eastPt, east))
	}
	/*//If we can't go anywhere, backup.
	if len(newPaths) == 0 {
		moveToLast := (curr.lastDir + 2) % 4
		prevPt := getPtFromDirection(curr.pt, moveToLast)
		newPaths = append(newPaths, newPathAtPt(curr, prevPt, moveToLast))
	}
	//fmt.Printf("Have %d new paths\n", len(newPaths))
	//for _, z := range newPaths {
	//	fmt.Printf("Pt: (y,x)=(%d,%d)\n", z.pt.y, z.pt.x)
	//}
	*/
	return newPaths
}

func isNewKey(curr *path) bool {
	//if len(curr.keysOrdered) > 0 && uint(curr.keysOrdered[len(curr.keysOrdered)-1]) == uint(grid[curr.pt.y][curr.pt.x]) {
	if len(curr.keysOrdered) > 0 && int64(curr.moveHistory[len(curr.moveHistory)-1]) == curr.numMoves && uint(curr.keysOrdered[len(curr.keysOrdered)-1]) == uint(grid[curr.pt.y][curr.pt.x]) {
		return true
	}
	return false
}

func isValidPoint(p point, keys uint32) bool {
	val := grid[p.y][p.x]
	if val == '#' {
		return false
	}
	//Check if there is a door, see if it open
	lowerVal := unicode.ToLower(val)
	if val != '.' && val != '@' && (val != lowerVal) {
		//fmt.Printf("Found value %s\n", string(val))
		loc := keyToLoc(lowerVal)
		if (keys>>loc)&1 == 0 {
			return false
		}
		return true
	}
	return true
}

func newPathAtPt(curr *path, newPt point, dir int64) *path {
	val := grid[newPt.y][newPt.x]

	newKeys := curr.keys
	newKeysOrdered := make([]byte, len(curr.keysOrdered))
	copy(newKeysOrdered, curr.keysOrdered)
	newMoveHistory := make([]int, len(curr.moveHistory))
	copy(newMoveHistory, curr.moveHistory)
	prevKeyLoc := curr.prevKeyLoc
	if val != '.' && val != '@' && unicode.ToLower(val) == val && ((newKeys>>keyToLoc(val))&1) == 0 {
		newKeys = newKeys | 1<<keyToLoc(val)
		newKeysOrdered = append(newKeysOrdered, byte(val))
		newMoveHistory = append(newMoveHistory, int(curr.numMoves)+1)
		prevKeyLoc = newPt
		//fmt.Printf("FOund key %s and changing keys from %s to %s\n", string(val), string(curr.keysOrdered), string(newKeysOrdered))
	}

	p := path{
		keys:        newKeys,
		keysOrdered: newKeysOrdered,
		moveHistory: newMoveHistory,
		pt:          newPt,
		numMoves:    curr.numMoves + 1,
		lastDir:     dir,
		prevKeyLoc:  prevKeyLoc,
	}
	return &p

}

func getPtFromDirection(p point, dir int64) point {
	switch dir {
	case north:
		return point{x: p.x, y: p.y - 1}
	case south:
		return point{x: p.x, y: p.y + 1}
	case east:
		return point{x: p.x + 1, y: p.y}
	case west:
		return point{x: p.x - 1, y: p.y}
	}
	panic(fmt.Sprintf("Error at point %v with direction %d\n", p, dir))
}

func main() {
	readInput()
	printGrid()
	p := path{
		pt:         start,
		numMoves:   0,
		lastDir:    -1,
		prevKeyLoc: start,
	}
	paths = &PathHeap{}
	heap.Init(paths)
	heap.Push(paths, p)
	fmt.Printf("Start is at x=%d, y=%d\n", start.x, start.y)
	iterateAllPaths()

}

/*
const inputStr = `#########
#b.A.@.a#
#########`
*/
/*
const inputStr = `########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################`
*/
/*
const inputStr = `########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################`
*/
/*
const inputStr = `#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################`
*/
/*
const inputStr = `########################
#@..............ac.GI.b#
###d#e#f################
###A#B#C################
###g#h#i################
########################`
*/

const inputStr = `#################################################################################
#.........#.......#...#...#..v..#.....#.#.........#...#....g#...L...#......z#...#
#.#.#####.#.#####.###.#.#.#.#.###I#.#.#.#.#.#######E#.#.###.#.#####.#.#####.#.###
#.#.....#...#.........#.#...#.....#.#...#.#.#..n....#.#.#.#...#...#.#...#...#..p#
#.#####.#####.#######.#.###########.###.#.#.#.#.#####.#.#.#####.#.#.###.#O###.#.#
#.#...#.#.....#...#...#...........#.#...#.#.#.#.#...#.#.#.......#.#.....#.#...#.#
#.###.#.#######.#.#.###########.###.#####.#.#.###.#.#.#.#.#####.#########.#####.#
#...#.#..y..#...#.#...#.........#...#...#.#.#.....#.#...#.#...#.........#.#.....#
###.#.#####.#.###.#####.#########.###.#.#.#.#######.#####.#.#.###.#.#####.#.#####
#...#....c#...#.#.......#.......#.#...#.#.#.......#....x#.#.#.#...#.#.....#.....#
#.###.#.#######.#########.#####.#.#.###.#.#####.#.#####.###.#.#.#####.#####.###.#
#.#...#.................#.....#.#.....#.#...#...#...#.#...#.#.#.#.....#...#.#...#
#.#.#############.#.#####.###.#.#.#####.#####.#####.#.###.#.#.#.#.#####.#.###.#.#
#.#...#.....#...#.#.#...#...#.#.#.#.....#.....#...#.....#.#.#...#.....#.#.#...#.#
#.###.#.#.#.#.###.###.#.###.#.#.#.#.###.#.###.#.#.#####.#.#.#########.#.#.#.###.#
#...#.#.#.#k..#...#...#...#.#.#.#.#...#.#.#...#.#.#.....#.#.........#.#.#...#.#.#
###.#.###F#####.###.#####.###.#.#####.#.#.#####.#.#.#####.#.#####.###.#.#####.#.#
#.#.#...#...#...#...#...#.....#.....#.#.#.......#.#.....#.#.....#.....#.#.....#.#
#.#.###.###.#.###.#.#.#.###########.#.###.#######.#####.#.#######.#####.#.#####.#
#.P.#.#.....#.#...#.#.#.....#...#...#...#...#...#.#...#.#.....H.#...#...#.....#.#
#.###.#.#####.#.###.#.#####.#.###.#####.#.###.#.#.#.#.#########.###.###.#.###.#.#
#...#...#.....#.#...#.#.....#.....#.....#.#...#...#.#.......#.#.#.#.#...#.#.#...#
###.#####.#####.#######.#####.#####.#.#.#.#.#######.#####.#.#.#.#.#.#.###.#.#####
#.#.#...#.....#.....#...#...#...#.#.#.#.#.#...#.....#.....#.#...#.#.#...#...#...#
#.#.#.#.#####.#####.#.###.#.###.#.#.#.###.###.#.#####.#.#########.#.###U#####.#.#
#.#...#...#...#.....#.#.#.#...#.#...#...#...#...#.#...#.#.....#.....#...#.....#.#
#.#######.#.#.#.#####.#.#.###.#.#.#####.###.#####.#.#####.###.#.#####.###.#####.#
#.....#...#.#.#...#...#...#.#...#...#...#.#.#...#.#.....#.#.#.#.#.....#...#.....#
#.#####.###.#.###.#.###.###.#####.###W#.#.#.#.#.#.#####.#.#.#.#.#.#.###.###.###.#
#...#...#...#.#.#...#.#.....#.....#...#.#.#.#.#.......#...#.#...#.#...#.#.#...#.#
#.#.#.###.###.#.#####.#####.#####.#.#####.#.#.#######.#####.#######.#.#T#.###.#.#
#.#.#.#.....#.....#.....#...#...#.#.....#...#.#.....#.#...........#.#.....#...#.#
#.#.#.#.###.#####.#.#.###.###.#.#.#####.#.###.#.###.###.#####.###.#########.###.#
#.#...#...#...#.#.#.#.#...#...#.#.#.....#...#.....#.......#...#.#.#.....#...#...#
#.#######.###.#.#.#.###.###.###.###.###.###.#############.#.###.#.#.###.#.###.###
#.#.......#.#.#.....#...#...J.#.....#...#.#...#...#...#...#...#...#...#...#...#.#
#.#######.#.#.#######.#######.#######.###.###.#.#.#.#.#######.#.###.#.#####.###.#
#.......#...#.#.......#.....#.....#.#.#.#...#...#...#...#.....#...#.#...#.#...#.#
#######.#####.#.#######.###.#####.#.#.#.#.#############.#.#######.#####.#.###.#.#
#.............#.........#.........#.......................#.............#.......#
#######################################.@.#######################################
#.........#...#.........#.....#...#...........#...........#..q....#.....#.......#
#####.###.#.###D#####.###.#.#.#.###.#.#.#.###.#######.#.#.#####.#.###.#.#.#####.#
#.....#.....#...#.#...#...#.#...#...#.#.#...#.......#.#.#j#...#.#.#...#d#.....#.#
#.#.#######.#.###.#.#.#.###.###.#.###.#.#.#.#######.###.#.#.#.#.#.#.###.###.#.###
#t#.#.....#.#...#.#.#.#...#...#.#.#.#.#.#.#...#....w....#...#...#...#.#...#.#...#
#.###.###.#.###.#.#.#####.###.###.#.#.#.#####.#######.###############.###.#.###.#
#.....#.#.#.#...#.#...Q.#...#.......#.#.#.....#...#...#.......#.#.....#...#...#.#
#.#####.#.###.###.#####.###.#######.#.###.###.#.#.#####.#####.#.#.###.#.###.###.#
#.#.....#...#.#.......#.....#...#...#...#.#...#.#.#.......#...#...#.#...#...#...#
#.#####.###.#.#R#####.#######.#.#######.#.#####.#.#.#######.###.###.#######.#.###
#.....#...#...#.#...#...#.....#.....#...#...#...#...#.......#.........#...#.#...#
#####.#.#.#######.#.#.#.#.#########.#.#.###.#.#######.###############.#.#.#.###.#
#.#...#.#...#.....#.#.#.#...#...#...#.#.#.#...#...#...#.............#...#.#...#.#
#.#.###.###.#S#####.#.#####.#.###.###.#.#.#####.#.#.###.###########.#####.#####.#
#...#...#.#...#...#...#.....#...#.#...#.#....m#.#.#...#.......#.....#...#.....#.#
#.#####.#.#######.###.#.#####.#.#.###.#.#.###.#.#.###.###.#####.#####.#.#####.#.#
#u......#.........#...#...#.#.#.#...#.#.#.#...#.#...#.#...#.....#.....#.....#.#.#
#########.#########.#####.#.#.#.###.#.#.#.#####.###.#.#.###.#######.#####.###.#.#
#.M.....#.......#...#.....#.#.#...#.#.#.#.........#.#.#.#.#.......#..s#.#.#...#.#
#####.#.#.#####.#.###.#####.#.###.#.#.###.#######.#.#.#.#.#######.###.#.#.#.###.#
#.....#.#.....#.#...#.#.....#.#.#...#...#...#...#.#.#...#...#.....#.#.#.#...#...#
#.#########B###.###.#######.#.#.#######.#####.#.###.#######.#######.#.#.#####.#.#
#.#.A...#...#.....#.....#...#.#.#.....#.#.....#...#...#.......#.....#.#.......#.#
#.#.###.#.###.#########.#.###.#.#.#.#.#.#.#######.#.#.#.#####.#.###.#.#######.#.#
#.....#.#...#.........#.#.#...#...#.#...#.....#.....#...#a....#.#...#.....#...#.#
#####.#.#.###########.#.#.#.###.###.###.#.###.#.###############.#.#.#####.#.###.#
#...#.#...#...........#...#...#...#.#...#.#...#.#.......#.......#.#.#.....#.#...#
#.#.#######.#############.###.#####.#####.#.###.#.#####.#.#######.###.#####.#.###
#.#.........#.............#.#.V...#.....#.#.C.#.#.....#.#.#.....#...#.#r....#...#
#.#########.#.#############.#####.#####.#####.#.#####.#.#.#####.###.#.#.#######.#
#.#.........#...#...#.G...#.....#.#...#.#...#.#.#.#...#...#b..#.#...#.Y.#.....#.#
#.#############.#.#.#####.#.#.###.#.#.#.#.#.#.#.#.#.#######.#.#.#.#######.#.###.#
#.#...........#...#.....#...#.#...#.#.#.#.#...#...#.#...K...#...#.#.......#.#...#
#.#X#######.###########.#####.#.###.#.#.#.#########.#####.#######.#.#.#######.###
#.#...#...#...N.#.#...#.#l..#.#...#.#...#.......#...#..f#.......#...#.#......o#.#
#.###.#.#.###.#.#.#.#.#.#.#.#.###.#.###.#.#####.#.###.#.#.#####.#####.#.#######.#
#.#...#.#.....#...#.#.#...#...#...#..i#.#...#.#...#...#.#.#.....#.#.Z.#...#...#.#
#.#.###.###########.#.#######.#.#####.#.###.#.#######.#.###.###.#.#.#####.#.#.#.#
#h....#..........e..#.........#.......#.#.............#.....#.....#.........#...#
#################################################################################`

//const inputStr = ``
