package main

import (
	"fmt"
)

func main() {
	b := 79
	c := b
	b *= 100
	b += 100000
	c = b + 17000
	h := 0
	for {
		f := 1
		for d := 2; d < b; {
			if b%d == 0 {
				f = 0
				break
			}
			d++
		}
		if f == 0 {
			h++
		}
		if b == c {
			break
		}
		b += 17
	}

	fmt.Printf("The value in h is %v\n", h)
}

//In pseudo code the program roughly translated to
/*
b=7900
b=b+100000
c=b
c=c+17000

do{
	f=1
	d=2
	do {
		e=2
		//The loop ends if g=b
		//The following sets f to be 0 if de-b==0 where e<be
		//Basically it checks if d|b or b%d==0 and if so, f=0
		do{
			//g=de-b
			g=d
			g=g*e
			g=g-b
			if g==0{
				f=0
			}

			e=e+1
			g=e
			g=g-b
		}while g!= 0
		d=d+1
	}while d!= b
	//if f is 0, increase h
	if f==0{
		h=h+1
	}
	//End program if b=c, otherwise increase b by 17
	g=b
	g=g-c
	b=b+17
} while g != 0
*/
