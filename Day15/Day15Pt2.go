package main

import (
	"fmt"
)

const (
	factorA = 16807
	factorB = 48271
)

const (
	//initialA = 65
	//initialB = 8921
	initialA = 634
	initialB = 301
)

func main() {
	numSteps := 5000000
	mask := 0xFFFF
	genAMaskCheck := 3
	genBMaskCheck := 7
	genA := initialA
	genB := initialB
	N := 2147483647 // THis number is 2^31-1, so is safe for signed and unsigned 32 bit ints
	count := 0
	for i := 0; i < numSteps; i++ {
		genA = (genA * factorA) % N
		for genA&genAMaskCheck != 0 {
			genA = (genA * factorA) % N
		}

		genB = (genB * factorB) % N
		for genB&genBMaskCheck != 0 {
			genB = (genB * factorB) % N
		}

		if (mask & genA) == (mask & genB) {
			count++
		}
		if i < 10 {
			fmt.Printf("A is %d B is %d\n", genA, genB)
		}
	}

	fmt.Printf("Total count is %d\n", count)
}
