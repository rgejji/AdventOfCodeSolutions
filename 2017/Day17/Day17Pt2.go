package main

import (
	"container/ring"
	"fmt"
)

func main() {
	//numSteps := 3
	numSteps := 363
	maxVal := 50000000
	circBuff := ring.New(1)
	circBuff.Value = 0
	for newVal := 1; newVal <= maxVal; newVal++ {
		for j := 0; j < numSteps; j++ {
			circBuff = circBuff.Next()
		}
		newRing := ring.New(1)
		newRing.Value = newVal
		newRing.Link(circBuff)
		//printRing(circBuff)
		if newVal%1000000 == 0 {
			fmt.Printf("At start %d\n", newVal)
		}
	}

	//fmt.Printf("After the last inserted record is %v\n", circBuff.Value)
	findAfterZero(circBuff)
}

func findAfterZero(r *ring.Ring) {
	for i := 0; i < r.Len(); i++ {
		if r.Value == 0 {
			r = r.Next()
			fmt.Printf("Final answer is %d\n", r.Value)
			return
		}
		r = r.Next()
	}
	fmt.Printf("Failed to find 0\n")
	return

}

func printRing(r *ring.Ring) {
	fmt.Printf("Values: ")
	for i := 0; i < r.Len(); i++ {
		r = r.Next()
		fmt.Printf("%v ", r.Value)

	}
	fmt.Printf("\n")
}
