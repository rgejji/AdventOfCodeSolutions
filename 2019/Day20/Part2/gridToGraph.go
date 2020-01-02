package main

import (
	"fmt"

	"gonum.org/v1/gonum/graph"
	path "gonum.org/v1/gonum/graph/path"
	simple "gonum.org/v1/gonum/graph/simple"
)

//GetID returns the id of the grid graph
func GetID(grid [][]string, l, i, j int64) int64 {
	v := int64(len(grid))
	n := int64(len(grid[0]))
	if v < n {
		v = n
	}
	return v*v*l + v*i + j
}

//GetCoordFromID returns the y,x coords of the grid graph
func GetCoordFromID(grid [][]string, id int64) (int64, int64, int64) {
	v := int64(len(grid))
	n := int64(len(grid[0]))
	if v < n {
		v = n
	}
	j := id % v
	i := ((id - j) / v) % v
	level := (id - v*i - j) / (v * v)
	return level, i, j
}

//GetMultilevelGraphInvalidList returns the graph
func GetMultilevelGraphInvalidList(grid [][]string, numLevels int64, invalidCharList []string) *simple.UndirectedGraph {
	return GetMultilevelGraphUsingFunction(grid, numLevels, invalidCharList, isValidCharUsingInvalidList)
}

//GetMultilevelGraphUsingFunction gets the graph using a validity check function
func GetMultilevelGraphUsingFunction(grid [][]string, numLevels int64, charList []string, isValidChar func(string, []string) bool) *simple.UndirectedGraph {
	g := simple.NewUndirectedGraph()
	m := int64(len(grid))
	v := m
	n := int64(len(grid[0]))
	if v < n {
		v = n
	}

	//Add nodes
	nodes := []graph.Node{}
	for i := int64(0); i < numLevels*v*v; i++ {
		nodes = append(nodes, g.NewNode())
		g.AddNode(nodes[i])
	}

	//Get rid of invalid nodes
	for n := int64(0); n < numLevels*v*v; n++ {
		_, i, j := GetCoordFromID(grid, n)
		if i >= int64(len(grid)) || j >= int64(len(grid[0])) || !isValidChar(grid[i][j], charList) {
			g.RemoveNode(n)
		}
	}

	//Add edges
	for level := int64(0); level < numLevels; level++ {
		for i, row := range grid {
			for j, char := range row {
				if isValidChar(char, charList) {
					currID := GetID(grid, level, int64(i), int64(j))
					upID := GetID(grid, level, int64(i-1), int64(j))
					downID := GetID(grid, level, int64(i+1), int64(j))
					leftID := GetID(grid, level, int64(i), int64(j-1))
					rightID := GetID(grid, level, int64(i), int64(j+1))

					if i > 0 && isValidChar(grid[i-1][j], charList) {
						e := g.NewEdge(g.Node(currID), g.Node(upID))
						g.SetEdge(e)
					}
					if i < len(grid)-1 && isValidChar(grid[i+1][j], charList) {
						e := g.NewEdge(g.Node(currID), g.Node(downID))
						g.SetEdge(e)
					}
					if j > 0 && isValidChar(grid[i][j-1], charList) {
						e := g.NewEdge(g.Node(currID), g.Node(leftID))
						g.SetEdge(e)
					}
					if j < len(grid[i])-1 && isValidChar(grid[i][j+1], charList) {
						e := g.NewEdge(g.Node(currID), g.Node(rightID))
						g.SetEdge(e)
					}
				}
			}
		}
	}
	fmt.Printf("Finished reading in grid\n")
	return g

}

func isValidCharUsingInvalidList(char string, invalidCharList []string) bool {
	for _, v := range invalidCharList {
		if v == char {
			return false
		}
	}
	return true
}

//PrintPathCoordinates takes in a path and prints the coordinates
func PrintPathCoordinates(grid [][]string, p []graph.Node) {
	currLevel := int64(0)
	fmt.Printf("[")
	for i, n := range p {
		if i != 0 {
			fmt.Printf(", ")
		}
		level, y, x := GetCoordFromID(grid, n.ID())
		if level != currLevel {
			currLevel = level
			fmt.Printf("\n")
		}

		fmt.Printf("(%d,%d,%d)", level, y, x)
	}
	fmt.Printf("]\n")
}

//GetAllShortestPaths returns the shortest path tree
func GetAllShortestPaths(g *simple.UndirectedGraph, from int64) path.Shortest {
	return path.DijkstraFrom(g.Node(from), g)
}

//GetShortestPath returns the shortest path
func GetShortestPath(g *simple.UndirectedGraph, from, to int64) []graph.Node {
	pathTree := GetAllShortestPaths(g, from)
	p, _ := pathTree.To(to)
	return p

	/*allpaths, ok := path.FloydWarshall(g)
	if !ok {
		fmt.Printf("INvalid path, floydwarshall returned negative\n")
	}
	fmt.Printf("HAVE PATHS! \n")
	p, _, _ := allpaths.Between(from, to)
	fmt.Printf("HAVE NETWEEN! %v\n", p)
	return p
	*/

}

//GetShortestPathDistance returns the shortest path dist
func GetShortestPathDistance(g *simple.UndirectedGraph, from, to int64) int64 {
	return int64(len(GetShortestPath(g, from, to)) - 1)
}
