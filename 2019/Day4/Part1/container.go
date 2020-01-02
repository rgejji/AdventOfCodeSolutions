package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	inputMin := 156218
	inputMax := 652527

	rA, err := regexp.Compile("^1*2*3*4*5*6*7*8*9*$")
	if err != nil {
		fmt.Printf("Unable to compile regex: %s\n", err.Error())
	}

	cnt := 0
	for i := inputMin; i <= inputMax; i++ {
		s := strconv.Itoa(i)

		if !rA.MatchString(s) {
			continue
		}
		for j := 1; j < 6; j++ {
			if s[j-1] == s[j] {
				cnt++
				break
			}
		}

	}

	fmt.Printf("Have %d matches\n", cnt)
}
