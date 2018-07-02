package main

import (
	"fmt"
)

func main() {
	//numSteps := 3
	//maxVal := 9
	numSteps := 363
	maxVal := 50000001
	currLoc := 0
	afterZero := 0
	for i := 1; i <= maxVal; i++ {
		//fmt.Printf("Current loc is %d for length %d. Next Calc is %v\n", currLoc, i, (currLoc+numSteps+1)%i)
		currLoc = (currLoc + numSteps) % i
		currLoc = (currLoc + 1) % (i + 1)
		if currLoc == 1 {
			afterZero = i
			fmt.Printf("After zero is now %d\n", i)
		}

	}
	fmt.Printf("After zero is %d\n", afterZero)
}
