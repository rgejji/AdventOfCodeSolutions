package main

import (
	"container/heap"
	"fmt"
	"strings"

	graphutil "github.com/RichardGejji/AdventOfCode2019/GraphUtil"
	simple "gonum.org/v1/gonum/graph/simple"
)

const (
	aASCII    uint = 97
	zASCII    uint = 122
	capAASCII uint = 65
	capZASCII uint = 90
)

const large float64 = 999999999.0

//var maxHeapSize = 500000
var grid [][]string
var paths *PathHeap

func keyToLoc(a rune) uint {
	return uint(a) - aASCII
}

func readGrid() {
	rowSlice := strings.Split(inputStr, "\n")
	grid = [][]string{}
	for _, row := range rowSlice {
		gridRow := []string{}
		for _, val := range row {
			gridRow = append(gridRow, string(val))
		}
		grid = append(grid, gridRow)
	}
}

func printGrid() {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%s", string(grid[i][j]))
		}
		fmt.Printf("\n")
	}
}

//getValues of interest with keys first
func getValuesOfInterest() ([]string, []int64) {
	values := []string{}
	ids := []int64{}
	for i := int64(0); i < int64(len(grid)); i++ {
		for j := int64(0); j < int64(len(grid[i])); j++ {
			val := grid[i][j]
			if valueIsKey(val[0]) {
				values = append(values, val)
				ids = append(ids, graphutil.GetID(grid, i, j))
			}
		}
	}
	for i := int64(0); i < int64(len(grid)); i++ {
		for j := int64(0); j < int64(len(grid[i])); j++ {
			val := grid[i][j]
			if val != " " && val != "#" && val != "." && !valueIsKey(val[0]) {
				values = append(values, val)
				ids = append(ids, graphutil.GetID(grid, i, j))
			}

		}
	}
	return values, ids
}

func getStart(values []string) []int64 {
	startLocs := []int64{}
	for i := 0; i < len(values); i++ {
		if values[i] == "@" {
			startLocs = append(startLocs, int64(i))
		}
	}
	return startLocs
}

func getShortenedGraph(g *simple.UndirectedGraph, values []string, ids []int64) *simple.WeightedUndirectedGraph {
	N := int64(len(ids))
	newG := simple.NewWeightedUndirectedGraph(large, large)
	for i := int64(0); i < N; i++ {
		node := newG.NewNode()
		newG.AddNode(node)
	}
	for i := int64(0); i < N; i++ {
		pathTree := graphutil.GetAllShortestPaths(g, ids[i])
		for j := int64(i) + 1; j < N; j++ {
			p, _ := pathTree.To(ids[j])
			weight := float64(len(p) - 1)
			e := newG.NewWeightedEdge(newG.Node(i), newG.Node(j), weight)
			newG.SetWeightedEdge(e)
		}
	}
	return newG
}

func trimShortenedGraph(sg *simple.WeightedUndirectedGraph, values []string) {
	//If we have a triple A,B, C where B is a Gate
	//and the weight around the gate is the same as the weight through the gate,
	//we remove edges from A to C.
	N := int64(len(values))
	for B := int64(0); B < N; B++ {
		nodeVal := values[B]
		if strings.ToUpper(nodeVal) == nodeVal && nodeVal != "@" {
			fmt.Printf("Trimming edges that go through: %s\n", nodeVal)
			for A := int64(0); A < N; A++ {
				if A == B {
					continue
				}
				for C := int64(0); C < N; C++ {
					if A == C || B == C {
						continue
					}
					wAB, ok := sg.Weight(A, B)
					if !ok {
						continue
					}
					wBC, ok := sg.Weight(B, C)
					if !ok {
						continue
					}
					wAC, ok := sg.Weight(A, C)
					if !ok {
						continue
					}
					if wAB+wBC == wAC {
						fmt.Printf("Removing edge (%s,%s) going around %s\n", values[A], values[C], values[B])
						sg.RemoveEdge(A, C)
					}

				}

			}
		}

	}

	//Remove edges with large weight
	for A := int64(0); A < N; A++ {
		for B := A + 1; B < N; B++ {
			if wAB, ok := sg.Weight(A, B); ok {
				//fmt.Printf("Looking at edge (%s,%s) with weight %v\n", values[A], values[B], wAB)
				if wAB <= 0 {
					fmt.Printf("Removing edge (%s,%s) for not connecting\n", values[A], values[B])
					sg.RemoveEdge(A, B)
				}
			}
		}
	}

}

func iterateAllPaths(sg *simple.WeightedUndirectedGraph, values []string) {
	totalKeys := int64(0)
	for _, v := range values {
		if valueIsKey(v[0]) {
			totalKeys++
		}
	}

	fmt.Printf("Have values array: %v\n\n", values)

	//Update paths of lowest priority in heap
	cnt := 0
	for {
		curr := heap.Pop(paths).(path)
		//fmt.Printf("Examining path with history: %s\n", string(curr.moveHistory))
		cnumKeys := curr.numKeys()
		if totalKeys == cnumKeys {
			fmt.Printf("Found path %s with move history %v and num moves: %d.\n", string(curr.keysOrdered), string(curr.moveHistory), curr.numMoves)
			return
		}

		nbrsToExplore := getUpdatedPaths(sg, values, &curr)
		maxNumKeys := addNeightborsToHeap(nbrsToExplore)

		cnt++
		if cnt%1000 == 0 {
			fmt.Printf("Took %d steps. Max num keys is %d\n", cnt, maxNumKeys)
		}
	}

}

//For getUpdated Paths, we get all the neighbors, and go to the next keys we haven't visited yet that we can reach
func getUpdatedPaths(sg *simple.WeightedUndirectedGraph, values []string, curr *path) []*path {
	newPaths := []*path{}
	for i, cid := range curr.currNodeID {
		//fmt.Printf("Out of neighbors %v, looking at section %v\n", curr.currNodeID, i)
		nodes := sg.From(cid)
		for nodes.Next() {
			nbr := nodes.Node()
			nbrID := nbr.ID()
			val := values[nbrID][0]
			newKeys := curr.keys
			newKeysOrdered := make([]byte, len(curr.keysOrdered))
			copy(newKeysOrdered, curr.keysOrdered)
			newMoveHistory := make([]byte, len(curr.moveHistory))
			copy(newMoveHistory, curr.moveHistory)
			newMoveHistory = append(newMoveHistory, val)

			//skip nodes for keys we already have.
			//We can skip @ since it is not a door
			if val == '@' {
				continue
			}
			if valueIsKey(val) {
				if currHasKey(curr, val) {
					continue
				}
				newKeys = newKeys | 1<<(uint(val)-aASCII)
				newKeysOrdered = append(newKeysOrdered, val)
			}

			//skip doors we don't have keys for
			if valueIsDoor(val) && !currHasKey(curr, strings.ToLower(string(val))[0]) {
				continue
			}

			weight, _ := sg.Weight(nbrID, cid)

			nbrIDSet := make([]int64, len(curr.currNodeID), len(curr.currNodeID))
			copy(nbrIDSet, curr.currNodeID)
			nbrIDSet[i] = nbrID
			//fmt.Printf("Have Weight: %v\n", weight)
			p := path{
				keys:        newKeys,
				keysOrdered: newKeysOrdered,
				moveHistory: newMoveHistory,
				numMoves:    curr.numMoves + int64(weight),
				currNodeID:  nbrIDSet,
			}
			//fmt.Printf("Will explore neightbor %s\n", newMoveHistory)
			newPaths = append(newPaths, &p)
		}
	}
	return newPaths
}

func currHasKey(curr *path, val byte) bool {
	if (curr.keys>>(uint(val)-aASCII))&1 == 0 {
		return false
	}
	return true
}

func valueIsKey(v byte) bool {
	if uint(v) >= uint(aASCII) && uint(zASCII) >= uint(v) {
		return true
	}
	return false
}

func valueIsDoor(v byte) bool {
	if uint(v) >= uint(capAASCII) && uint(capZASCII) >= uint(v) {
		return true
	}
	return false
}

func addNeightborsToHeap(nbrsToExplore []*path) int64 {
	pathsHeld := []*path{}
	currMaxNumKeys := int64(0)

	for paths.Len() > 0 {
		curr := heap.Pop(paths).(path)
		pathsHeld = append(pathsHeld, &curr)
	}
	//Trim by checking if there are duplicates
	for _, pA := range pathsHeld {
		addPA := true
		for j, pB := range nbrsToExplore {
			nodeSetsMatch := true
			for section := 0; section < len(pA.currNodeID); section++ {
				if pA.currNodeID[section] != pB.currNodeID[section] {
					nodeSetsMatch = false
					break
				}
			}

			if nodeSetsMatch {
				//Check if pB is better than pA.
				//e.g. pA keys <= pB keys,
				//they are in the same place
				//and the numMoves A >= numMoves B
				//If so, don't readd pA
				if pA.keys&pB.keys == pA.keys && pA.numMoves >= pB.numMoves {
					addPA = false
					break
				}
				//Check if pA is better than pB. If so, don't add pB
				if pA.keys&pB.keys == pB.keys && pB.numMoves >= pA.numMoves {
					nbrsToExplore = append(nbrsToExplore[:j], nbrsToExplore[j+1:]...)
					break
				}
			}
		}
		if addPA {
			cnumKeys := pA.numKeys()
			if cnumKeys > currMaxNumKeys {
				currMaxNumKeys = cnumKeys
			}

			heap.Push(paths, *pA)
		}
	}
	//Add valid neighors
	for _, pB := range nbrsToExplore {
		cnumKeys := pB.numKeys()
		if cnumKeys > currMaxNumKeys {
			currMaxNumKeys = cnumKeys
		}

		heap.Push(paths, *pB)
	}

	return currMaxNumKeys
}

//g.From gets the list of neighors

func main() {
	//Fill out grid
	readGrid()
	printGrid()

	//Create Grpah
	values, ids := getValuesOfInterest()
	fmt.Printf("Have ids %v\n", ids)
	fmt.Printf("Have valued %v\n", values)
	g := graphutil.GetGraphInvalidList(grid, []string{"#", " "})
	//fmt.Printf("%v\n", g)
	shortG := getShortenedGraph(g, values, ids)
	trimShortenedGraph(shortG, values)

	start := getStart(values)

	p := path{
		currNodeID: start,
		numMoves:   0,
	}
	paths = &PathHeap{}
	heap.Init(paths)
	heap.Push(paths, p)
	fmt.Printf("Start is at nodes %v of shortened graph\n", start)
	iterateAllPaths(shortG, values)

	//	fmt.Printf("%v\n", shortG)

	fmt.Printf("Done\n")

}

/*
const inputStr = `###############
#d.ABC.#.....a#
######@#@######
###############
######@#@######
#b.....#.....c#
###############`
*/
/*
const inputStr = `#############
#g#f.D#..h#l#
#F###e#E###.#
#dCba@#@BcIJ#
#############
#nK.L@#@G...#
#M###N#H###.#
#o#m..#i#jk.#
#############`
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
#.............#.........#.........#....@#@................#.............#.......#
#################################################################################
#.........#...#.........#.....#...#....@#@....#...........#..q....#.....#.......#
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

//End
