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

func ptXString(p *point) string {
	return fmt.Sprintf("%d,%d", p.x, p.vx)
}
func ptYString(p *point) string {
	return fmt.Sprintf("%d,%d", p.y, p.vy)
}
func ptZString(p *point) string {
	return fmt.Sprintf("%d,%d", p.z, p.vz)
}

func main() {
	//Make keys for each state, calculate when repeat for each state occurs
	xMap := make(map[string]int)
	yMap := make(map[string]int)
	zMap := make(map[string]int)

	repeatXLocA := -1
	repeatYLocA := -1
	repeatZLocA := -1
	repeatXLocB := -1
	repeatYLocB := -1
	repeatZLocB := -1

	readInput()
	//We use that x,y, and z are indep
	i := 0
	for {
		if i%1000 == 0 && i != 0 {
			fmt.Printf("> i=%d\n", i)
		}

		if repeatXLocA < 0 {
			xKey := ""
			for _, p := range points {
				xKey += fmt.Sprintf("%s/", ptXString(p))
			}

			if _, ok := xMap[xKey]; ok {
				repeatXLocA = xMap[xKey]
				repeatXLocB = i
				fmt.Printf("Found x-repeat at i=%d and i=%d\n", repeatXLocA, repeatXLocB)
			} else {
				xMap[xKey] = i
			}
		}
		if repeatYLocA < 0 {
			yKey := ""
			for _, p := range points {
				yKey += fmt.Sprintf("%s/", ptYString(p))
			}

			if _, ok := yMap[yKey]; ok {
				repeatYLocA = yMap[yKey]
				repeatYLocB = i
				fmt.Printf("Found y-repeat at i=%d and i=%d\n", repeatYLocA, repeatYLocB)
			} else {
				yMap[yKey] = i
			}
		}
		if repeatZLocA < 0 {
			zKey := ""
			for _, p := range points {
				zKey += fmt.Sprintf("%s/", ptZString(p))
			}

			if _, ok := zMap[zKey]; ok {
				repeatZLocA = zMap[zKey]
				repeatZLocB = i
				fmt.Printf("Found z-repeat at i=%d and i=%d\n", repeatZLocA, repeatZLocB)
			} else {
				zMap[zKey] = i
			}
		}

		if repeatXLocA >= 0 && repeatYLocA >= 0 && repeatZLocA >= 0 {
			break
		}

		updateVelocities()
		updatePositions()
		i++
	}
	printLocs()

	cycleX := repeatXLocB - repeatXLocA
	cycleY := repeatYLocB - repeatYLocA
	cycleZ := repeatZLocB - repeatZLocA

	max := repeatXLocA
	if repeatYLocA > max {
		max = repeatYLocA
	}
	if repeatZLocA > max {
		max = repeatZLocA
	}

	fmt.Printf("Full repeat is %d + lcm of %d, %d, %d\n", max, cycleX, cycleY, cycleZ)
	fmt.Printf("Answer is: %d\n", int64(max)+LCM(int64(cycleX), int64(cycleY), int64(cycleZ)))

}

// GCD greatest common divisor (GCD) via Euclidean algorithm
//Code copied from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

//LCM find Least Common Multiple (LCM) via GCD
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
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
/*
const inputStr = `<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>`*/
