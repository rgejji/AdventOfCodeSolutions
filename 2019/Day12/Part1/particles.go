package main

import (
	"fmt"
	"strconv"
	"strings"
)

type point struct {
	x  int
	y  int
	z  int
	vx int
	vy int
	vz int
}

var points []*point

func readInput() {
	points = []*point{}
	rowSlice := strings.Split(inputStr, "\n")
	for _, row := range rowSlice {
		pt := point{}
		row = strings.TrimLeft(row, "<")
		row = strings.TrimRight(row, ">")
		valSplit := strings.Split(row, ", ")
		x := strings.Split(valSplit[0], "=")[1]
		y := strings.Split(valSplit[1], "=")[1]
		z := strings.Split(valSplit[2], "=")[1]
		pt.x, _ = strconv.Atoi(x)
		pt.y, _ = strconv.Atoi(y)
		pt.z, _ = strconv.Atoi(z)

		points = append(points, &pt)
	}
	return
}

func updatePositions() {
	for _, p := range points {
		p.x = p.x + p.vx
		p.y = p.y + p.vy
		p.z = p.z + p.vz
	}
}

func updateVelocities() {
	for i, p := range points {
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			if p.x < p2.x {
				p.vx++
				p2.vx--
			}
			if p.x > p2.x {
				p.vx--
				p2.vx++
			}
			if p.y < p2.y {
				p.vy++
				p2.vy--
			}
			if p.y > p2.y {
				p.vy--
				p2.vy++
			}
			if p.z < p2.z {
				p.vz++
				p2.vz--
			}
			if p.z > p2.z {
				p.vz--
				p2.vz++
			}
		}
	}
}

func getPotentialEnergy(p *point) int {
	return abs(p.x) + abs(p.y) + abs(p.z)
}

func getKineticEnergy(p *point) int {
	return abs(p.vx) + abs(p.vy) + abs(p.vz)
}

func getTotalEnergy() int {
	sum := 0
	for _, p := range points {
		sum += getPotentialEnergy(p) * getKineticEnergy(p)
	}
	return sum
}

func printLocs() {
	for _, p := range points {
		fmt.Printf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>\n", p.x, p.y, p.z, p.vx, p.vy, p.vz)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	maxWait := 1000
	readInput()
	printLocs()
	fmt.Println("")
	for i := 0; i < maxWait; i++ {
		updateVelocities()
		updatePositions()
	}
	printLocs()
	fmt.Printf("Total Energery: %d\n", getTotalEnergy())

}

const inputStr = `<x=-10, y=-10, z=-13>
<x=5, y=5, z=-9>
<x=3, y=8, z=-16>
<x=1, y=3, z=-3>`

/*
const inputStr = `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`
*/

/*const inputStr = `<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>`*/
