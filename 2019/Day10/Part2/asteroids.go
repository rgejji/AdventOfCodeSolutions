package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

const pi = 3.14159265358979323846
const pi2 = 1.57079632679489661923

var grid [][]string
var asteroids []Point

//Point is a point
type Point struct {
	x int
	y int
}

func readInput() {
	slice := strings.Split(inputStr, "\n")
	grid = make([][]string, 0)
	asteroids = make([]Point, 0)
	for y, row := range slice {
		vals := make([]string, 0)
		for x, r := range row {
			vals = append(vals, string(r))
			if r == '#' {
				asteroids = append(asteroids, Point{x: x, y: y})
			}

		}
		grid = append(grid, vals)
	}
}

func printAsteroids() {
	for _, a := range asteroids {
		fmt.Printf("x: %d, y: %d\n", a.x, a.y)
	}
}

func storeVisibleFromAsteroid(a Point) map[float64][]Point {
	foundAngles := make(map[float64][]Point)
	var angle float64
	for _, p := range asteroids {
		//Check vertical edge case
		if p.x == a.x {
			if p.y == a.y {
				continue
			}
			if p.y > a.y {
				angle = pi2
			} else {
				angle = -pi2
			}
		} else {
			opp := float64(p.y - a.y)
			adj := float64(p.x - a.x)
			angle = math.Atan2(opp, adj)
		}
		angleRnd := calculateOrderedAngle(angle)
		if p.x == a.x && p.y > a.y {
			fmt.Printf("Have case here and angle is %v\n", angleRnd)
		}

		//angleStr := fmt.Sprintf("%.6f", angle)
		if _, ok := foundAngles[angleRnd]; !ok {
			foundAngles[angleRnd] = []Point{p}
		} else {
			foundAngles[angleRnd] = append(foundAngles[angleRnd], p)
		}
	}
	return foundAngles
}

//Adjust angle so pi/2 is 2pi and the angle decreases going clockwise
func calculateOrderedAngle(angle float64) float64 {
	eps := 0.0000001
	angle = angle - pi2
	for angle < pi {
		angle += 2 * pi
	}
	return round(angle, eps)
}

func round(x, u float64) float64 {
	return math.Round(x/u) * u
}

func printGrid() {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%s", grid[i][j])
		}
		fmt.Printf("\n")
	}
}

type asteroidAngle struct {
	angle     float64
	asteroids []Point
}
type aSlice []asteroidAngle

func (a aSlice) Len() int           { return len(a) }
func (a aSlice) Less(i, j int) bool { return a[i].angle < a[j].angle }
func (a aSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	//a := Point{x: 8, y: 3}
	a := Point{x: 31, y: 20}
	exitNum := 200

	readInput()
	//read and sort asteroid order
	aMap := storeVisibleFromAsteroid(a)
	i := 0
	angleSlice := make(aSlice, len(aMap))
	for k, v := range aMap {
		angleSlice[i] = asteroidAngle{
			angle:     k,
			asteroids: v,
		}
		i++
	}
	sort.Sort(angleSlice)
	//fmt.Printf("Have angle slice: %+v", angleSlice)
	//Go through array
	numDestroyedPrev := -1
	numDestroyedTotal := 0
	printGrid()
	pt := Point{}
	for numDestroyedPrev != numDestroyedTotal && numDestroyedTotal < exitNum {
		numDestroyedPrev = numDestroyedTotal
		for i, aa := range angleSlice {
			//fmt.Printf("Looking at angle %v\n", aa.angle)
			if len(aa.asteroids) == 0 {
				continue
			}
			angleSlice[i].asteroids, pt = shootShortestDistance(a, aa.asteroids)
			grid[pt.y][pt.x] = fmt.Sprintf("%d", i%10)
			numDestroyedTotal++
			//fmt.Printf("Destroyed %v at cnt %d\n", pt, numDestroyedTotal)
			if numDestroyedTotal == exitNum {
				fmt.Printf("Destroyed (x,y)=(%d,%d) at cnt %d\n", pt.x, pt.y, numDestroyedTotal)
				break
			}
		}
	}
	printGrid()
}

func shootShortestDistance(start Point, objs []Point) ([]Point, Point) {
	minDist := 999999999.0
	loc := -1
	for i, o := range objs {
		dist := math.Abs(float64(o.x-start.x)) + math.Abs(float64(o.y-start.y))
		if dist < minDist {
			minDist = dist
			loc = i
		}
	}
	p := objs[loc]
	objs[loc] = objs[len(objs)-1]
	objs = objs[:len(objs)-1]

	return objs, p
}

/*
const inputStr = `.#....#####...#..
##...##.#####..##
##...#...#.#####.
..#.....#...###..
..#.#.....#....##`

*/
const inputStr = `##.###.#.......#.#....#....#..........#.
....#..#..#.....#.##.............#......
...#.#..###..#..#.....#........#......#.
#......#.....#.##.#.##.##...#...#......#
.............#....#.....#.#......#.#....
..##.....#..#..#.#.#....##.......#.....#
.#........#...#...#.#.....#.....#.#..#.#
...#...........#....#..#.#..#...##.#.#..
#.##.#.#...#..#...........#..........#..
........#.#..#..##.#.##......##.........
................#.##.#....##.......#....
#............#.........###...#...#.....#
#....#..#....##.#....#...#.....#......#.
.........#...#.#....#.#.....#...#...#...
.............###.....#.#...##...........
...#...#.......#....#.#...#....#...#....
.....#..#...#.#.........##....#...#.....
....##.........#......#...#...#....#..#.
#...#..#..#.#...##.#..#.............#.##
.....#...##..#....#.#.##..##.....#....#.
..#....#..#........#.#.......#.##..###..
...#....#..#.#.#........##..#..#..##....
.......#.##.....#.#.....#...#...........
........#.......#.#...........#..###..##
...#.....#..#.#.......##.###.###...#....
...............#..#....#.#....#....#.#..
#......#...#.....#.#........##.##.#.....
###.......#............#....#..#.#......
..###.#.#....##..#.......#.............#
##.#.#...#.#..........##.#..#...##......
..#......#..........#.#..#....##........
......##.##.#....#....#..........#...#..
#.#..#..#.#...........#..#.......#..#.#.
#.....#.#.........#............#.#..##.#
.....##....#.##....#.....#..##....#..#..
.#.......#......#.......#....#....#..#..
...#........#.#.##..#.#..#..#........#..
#........#.#......#..###....##..#......#
...#....#...#.....#.....#.##.#..#...#...
#.#.....##....#...........#.....#...#...`
