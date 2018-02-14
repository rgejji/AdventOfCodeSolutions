package main

import (
	"fmt"
	"math"
)

func main() {
	val := float64(12.0)
	fmt.Printf("%v\n", getNormOnNumericSpiral(val))
	val = float64(23.0)
	fmt.Printf("%v\n", getNormOnNumericSpiral(val))
	val = float64(1024.0)
	fmt.Printf("%v\n", getNormOnNumericSpiral(val))
	val = float64(277678.0)
	fmt.Printf("%v\n", getNormOnNumericSpiral(val))

}

func getNormOnNumericSpiral(loc float64) float64 {

	//Level 0 has 1 = 1^2
	//Level 1 is 1 < loc <= 9 = 3^2
	//Level 2 is 9 < loc <= 25=5^2

	//level_lower_val := (2*level-1)^2
	//level_upper_val := (2*level+1)^2
	//llv <val < luv
	//level-1/2=sqrt(llv)/2<sqrt(val)/2 < sqrt(llu)/2 = level+1/2
	//level = round(sqrt(llv)/2)

	level := math.Floor(math.Sqrt(loc)/2.0 + 0.5)
	n := level
	fmt.Printf("Loc is %v and Level is %v\n", loc, level)

	lowerRightPrev := (2*n - 1) * (2*n - 1)
	upperRight := 4*n*n + 1 - 2*n
	upperLeft := 4*n*n + 1
	lowerLeft := 4*n*n + 1 + 2*n
	lowerRight := (2*n + 1) * (2*n + 1)

	fmt.Printf("One before level is %v and corners are %v, %v, %v, and %v\n", lowerRightPrev, upperRight, upperLeft, lowerLeft, lowerRight)

	switch {
	case loc > lowerLeft:
		mid := (lowerLeft + lowerRight) / 2
		fmt.Printf("On the bottom with mid %v\n", mid)
		return math.Abs(mid-loc) + level
	case loc > upperLeft:
		mid := (upperLeft + lowerLeft) / 2
		fmt.Printf("On the left with mid %v\n", mid)
		return math.Abs(mid-loc) + level
	case loc > upperRight:
		mid := (upperRight + upperLeft) / 2
		fmt.Printf("On the top with mid %v\n", mid)
		return math.Abs(mid-loc) + level
	}
	mid := (lowerRightPrev + upperRight) / 2
	fmt.Printf("On the right with mid %v\n", mid)
	return math.Abs(mid-loc) + level

}
