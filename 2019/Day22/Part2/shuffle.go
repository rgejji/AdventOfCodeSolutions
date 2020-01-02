package main

import (
	"fmt"
	"strconv"
	"strings"
)

func newStackPos(pos int64, L int64) int64 {
	return L - 1 - pos
}

func cutPos(pos int64, N int64, L int64) int64 {
	for N > L {
		N -= L
	}
	return (pos + L - N) % L
}

//This may be overflowing, need to redo in python or with bignum library
func dealPos(pos int64, inc int64, L int64) int64 {
	return (pos * inc) % L
}

func doShuffle(pos int64, L int64) int64 {
	for _, row := range strings.Split(inputStr, "\n") {
		vals := strings.Split(row, " ")
		instruction := vals[0]
		modifier := vals[len(vals)-1]

		switch instruction {
		case "deal":
			if modifier == "stack" {
				pos = newStackPos(pos, L)
			} else {
				v, _ := strconv.Atoi(modifier)
				pos = dealPos(pos, int64(v), L)
			}
		case "cut":
			v, _ := strconv.Atoi(modifier)
			pos = cutPos(pos, int64(v), L)
		default:
			fmt.Printf("Unknown instruction line %s\n", row)
			return 0
		}
		//fmt.Printf("Current result: %v\n", deck)

	}
	return pos

}

//L is the length

func main() {
	var L int64 = 119315717514047
	//L := int64(10007)
	//pos := int64(2019)
	//numReps := int64(101741582076661)
	start := int64(2020)
	posV := []int64{0, 0, 0}

	fmt.Printf("Have start %d\n", start)
	//The instructions is a linear change on the position, we run a couple of times to figure out linear coefficients
	for i := 0; i < 3; i++ {
		pos := start * int64(i)
		pos = doShuffle(pos, L)
		posV[i] = pos
		fmt.Printf("Have position %d\n", pos)
	}
	//Note this iteration is just ax+b%L for some a and b.
	//The first pos gives b. The second (pos -b)*inv(x,L) = a
	//The third pos is a check
	b := posV[0]
	fmt.Printf("b=%d\n", b)
	//fmt.Printf("for x=2020, ax=%d\n", (posV[1]-b+L)%L)
	xinv := int64(108506422313517)
	fmt.Printf("For x=start, check xinv=%v: %v\n", xinv, (xinv*start)%L)
	//a := (xinv * ((posV[1] - b + 2*L) % L)) % L
	//Use python to find a
	a := int64(88872671823520)
	fmt.Printf("a = %v, with 0 value check: %v\n", a, (b-posV[0]+L)%L)
	fmt.Printf("a = %v, with 0 value check: %v\n", a, (a*start+b-posV[1]+L)%L)
	fmt.Printf("a = %v, with 0 value check: %v\n", a, (a*start+a*start+b-posV[2]+L)%L)

	test := doShuffle(doShuffle(start, L), L)
	fmt.Printf("Value 0 check: %v\n", test-(a*((a*start+b)%L)+b)%L)

	//Answer is found repeating x := ax+b
	//Unfortunately, we overflow here, will switch to python
	fmt.Printf("a will overflow: a*a=%v\n", a*a)

}

const inputStr = `deal into new stack
cut 7990
deal into new stack
cut -5698
deal with increment 29
cut 1503
deal with increment 65
cut -9095
deal with increment 56
cut 9104
deal into new stack
deal with increment 5
cut -7708
deal with increment 20
cut 4813
deal with increment 2
cut 4728
deal into new stack
cut -5429
deal with increment 47
cut 1739
deal with increment 63
cut 6707
deal with increment 29
cut 4293
deal with increment 44
cut 8873
deal with increment 53
cut 6046
deal into new stack
cut 8054
deal into new stack
deal with increment 14
cut 2426
deal with increment 11
cut 4006
deal with increment 49
cut -6277
deal with increment 3
cut 2231
deal with increment 45
cut -5059
deal with increment 7
cut 4251
deal with increment 16
cut -6081
deal with increment 25
cut -4067
deal with increment 29
cut 7656
deal into new stack
cut 5091
deal with increment 57
deal into new stack
deal with increment 63
cut 4047
deal with increment 24
cut -8596
deal with increment 13
cut 1946
deal with increment 16
cut -1656
deal into new stack
deal with increment 15
cut -6557
deal with increment 10
cut 2378
deal with increment 24
cut -2162
deal with increment 7
deal into new stack
deal with increment 37
cut -4310
deal into new stack
deal with increment 48
cut 6842
deal with increment 13
cut 2960
deal into new stack
cut 7128
deal with increment 30
cut -2529
deal with increment 31
cut -2500
deal with increment 28
deal into new stack
deal with increment 37
cut -8133
deal with increment 74
cut -7823
deal with increment 42
cut 2092
deal with increment 41
cut -6752
deal with increment 56
cut -9577
deal into new stack
cut -4736
deal with increment 8
cut -3584`

//const inputStr = `deal with increment 3`
//const inputStr = `deal stack`
//const inputStr = `cut -4`

//const inputStr = `deal with increment 7
//deal into new stack
//deal into new stack`
/*
const inputStr = `deal into new stack
cut -2
deal with increment 7
cut 8
cut -4
deal with increment 7
cut 3
deal with increment 9
deal with increment 3
cut -1`
*/
