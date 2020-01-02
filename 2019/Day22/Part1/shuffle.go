package main

import (
	"fmt"
	"strconv"
	"strings"
)

func newStack(deck []int) []int {
	for i := 0; i < len(deck)/2; i++ {
		op := len(deck) - 1 - i
		deck[i], deck[op] = deck[op], deck[i]
	}
	return deck
}

func cut(deck []int, N int) []int {
	L := len(deck)
	for N >= L {
		N -= L
	}
	for N < 0 {
		N += L
	}
	deck = append(deck[N:], deck[0:N]...)
	return deck
}

func deal(deck []int, inc int) []int {
	L := len(deck)
	newDeck := make([]int, L, L)
	loc := 0
	for i := 0; i < L; i++ {
		newDeck[loc] = deck[i]
		loc = (loc + inc) % L
	}
	return newDeck
}

func findPosInDeck(deck []int, value int) int {
	for i := 0; i < len(deck); i++ {
		if deck[i] == value {
			return i
		}
	}
	return -1
}

func main() {
	//deckSize := 10007
	deckSize := 10007
	deck := make([]int, deckSize, deckSize)
	for i := 0; i < deckSize; i++ {
		deck[i] = i
	}

	for _, row := range strings.Split(inputStr, "\n") {

		vals := strings.Split(row, " ")
		instruction := vals[0]
		modifier := vals[len(vals)-1]

		switch instruction {
		case "deal":
			if modifier == "stack" {
				deck = newStack(deck)
			} else {
				v, _ := strconv.Atoi(modifier)
				deck = deal(deck, v)
			}
		case "cut":
			v, _ := strconv.Atoi(modifier)
			deck = cut(deck, v)
		default:
			fmt.Printf("Unknown instruction line %s\n", row)
			return
		}
		//fmt.Printf("Current result: %v\n", deck)

	}
	//fmt.Printf("Result: %v\n", deck)
	fmt.Printf("Have position %d\n", findPosInDeck(deck, 2019))
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
