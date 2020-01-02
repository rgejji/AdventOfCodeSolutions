package main

import (
	"fmt"
	"math"
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
		vals := make([]string, len(slice), len(slice))
		for x, r := range row {
			vals[x] = string(r)
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

func findMaxCount() {
	maxCnt := -1
	maxAsteroid := Point{}
	for _, a := range asteroids {
		cnt := numVisibleFromAsteroid(a)
		if cnt > maxCnt {
			maxCnt = cnt
			maxAsteroid = a
		}
	}
	fmt.Printf("Best asteroid is %+v with count %d\n", maxAsteroid, maxCnt)
}

func numVisibleFromAsteroid(a Point) int {
	foundAngles := make(map[string]bool)
	var angle float64
	count := 0
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
			//if p.x < a.x {
			//	angle = (pi - angle) * sgn(opp)
			//}
		}
		angleStr := fmt.Sprintf("%.6f", angle)
		if _, ok := foundAngles[angleStr]; !ok {
			foundAngles[angleStr] = true
			count++
		}
	}
	return count
}

/*
func sgn(f float64) float64 {
	if f < 0 {
		return -1.0
	}
	return 1.0
}
*/
func main() {
	readInput()
	printAsteroids()
	findMaxCount()
}

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

/*
const inputStr = `#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`


const inputStr = `.#..#
.....
#####
....#
...##`
*/
