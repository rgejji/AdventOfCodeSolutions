package graphutil

import (
	"fmt"

	"gonum.org/v1/gonum/graph"
	path "gonum.org/v1/gonum/graph/path"
	simple "gonum.org/v1/gonum/graph/simple"
)

//GetID returns the id of the grid graph
func GetID(grid [][]string, i, j int64) int64 {
	v := int64(len(grid))
	n := int64(len(grid[0]))
	if v < n {
		v = n
	}
	return v*i + j
}

//GetCoordFromID returns the y,x coords of the grid graph
func GetCoordFromID(grid [][]string, id int64) (int64, int64) {
	v := int64(len(grid))
	n := int64(len(grid[0]))
	if v < n {
		v = n
	}
	j := id % v
	i := (id - j) / v
	return i, j
}

//GetGraph returns the graph
func GetGraph(grid [][]string, validCharList []string) *simple.UndirectedGraph {
	return GetGraphUsingFunction(grid, validCharList, isValidCharUsingValidList)
}

//GetGraphInvalidList returns the graph
func GetGraphInvalidList(grid [][]string, invalidCharList []string) *simple.UndirectedGraph {
	return GetGraphUsingFunction(grid, invalidCharList, isValidCharUsingInvalidList)
}

//GetGraphUsingFunction gets the graph using a validity check function
func GetGraphUsingFunction(grid [][]string, charList []string, isValidChar func(string, []string) bool) *simple.UndirectedGraph {
	g := simple.NewUndirectedGraph()
	v := len(grid)
	n := len(grid[0])
	if v < n {
		v = n
	}

	//Add nodes
	nodes := []graph.Node{}
	for i := 0; i < v*v+v; i++ {
		nodes = append(nodes, g.NewNode())
		g.AddNode(nodes[i])
	}

	//Get rid of invalid nodes
	for n := 0; n < v*v+v; n++ {
		j := n % v
		i := (n - j) / v
		if i >= len(grid) || j >= len(grid[0]) || !isValidChar(grid[i][j], charList) {
			g.RemoveNode(int64(n))
		}
	}

	//Add edges
	for i, row := range grid {
		for j, char := range row {
			if isValidChar(char, charList) {
				if i > 0 && isValidChar(grid[i-1][j], charList) {
					e := g.NewEdge(g.Node(int64(v*i+j)), g.Node(int64(v*(i-1)+j)))
					g.SetEdge(e)
				}
				if i < len(grid)-1 && isValidChar(grid[i+1][j], charList) {
					e := g.NewEdge(g.Node(int64(v*i+j)), g.Node(int64(v*(i+1)+j)))
					g.SetEdge(e)
				}
				if j > 0 && isValidChar(grid[i][j-1], charList) {
					e := g.NewEdge(g.Node(int64(v*i+j)), g.Node(int64(v*i+j-1)))
					g.SetEdge(e)
				}
				if j < len(grid[i])-1 && isValidChar(grid[i][j+1], charList) {
					e := g.NewEdge(g.Node(int64(v*i+j)), g.Node(int64(v*i+j+1)))
					g.SetEdge(e)
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
func isValidCharUsingValidList(char string, validCharList []string) bool {
	for _, v := range validCharList {
		if v == char {
			return true
		}
	}
	return false
}

//PrintPathCoordinates takes in a path and prints the coordinates
func PrintPathCoordinates(grid [][]string, p []graph.Node) {
	fmt.Printf("[")
	for i, n := range p {
		if i != 0 {
			fmt.Printf(", ")
		}
		y, x := GetCoordFromID(grid, n.ID())
		fmt.Printf("(%d,%d)", y, x)
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
