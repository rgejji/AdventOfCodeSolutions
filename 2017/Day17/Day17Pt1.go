package main

import (
	"container/ring"
	"fmt"
)

func main() {
	//numSteps := 3
	numSteps := 363
	maxVal := 2017
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
	}

	fmt.Printf("After the last inserted record is %v\n", circBuff.Value)

}

func printRing(r *ring.Ring) {
	fmt.Printf("Values: ")
	for i := 0; i < r.Len(); i++ {
		r = r.Next()
		fmt.Printf("%v ", r.Value)

	}
	fmt.Printf("\n")
}
