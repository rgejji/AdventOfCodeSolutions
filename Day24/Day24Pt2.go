package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var (
	componentCounts map[string]int
	portMap         map[string][]port
)

//Sorted values
type port struct {
	A int
	B int
}

func (p port) String() string {
	return fmt.Sprintf("%d/%d", p.A, p.B)
}

func main() {
	//readComponents(inputA)
	readComponents(input)
	finalStrength, finalLength := recurseBridge([]port{}, 0, 0)
	fmt.Printf("found final strength %d and length %d\n", finalStrength, finalLength)
}

//recurse bridge returns the stength and length of the longest bridge with ties going to strongest
func recurseBridge(currentBridge []port, currStrength int, openLink int) (int, int) {
	//find open link
	possiblePorts := portMap[string(openLink)]
	bestStrength := currStrength
	bestLength := len(currentBridge)
	for _, port := range possiblePorts {
		//Check count
		portString := port.String()
		count, ok := componentCounts[portString]
		if count <= 0 || !ok {
			if !ok {
				log.Fatalf("Unknown value pair %s", port)
			}
			continue
		}

		//Perform pre-recursion Update
		currentBridge = append(currentBridge, port)
		componentCounts[portString] = count - 1
		newOpenLink := port.A
		if newOpenLink == openLink {
			newOpenLink = port.B
		}
		//fmt.Printf("Adding port %s ounts are %+v\n", portString, componentCounts)
		//Recurse
		foundStrength, foundLength := recurseBridge(currentBridge, currStrength+port.A+port.B, newOpenLink)
		if foundLength > bestLength || (foundLength == bestLength && foundStrength > bestStrength) {
			bestStrength, bestLength = foundStrength, foundLength
		}

		//Perform post-recursion cleanup
		componentCounts[portString] = count
		currentBridge = currentBridge[:len(currentBridge)-1]
	}
	/*if currStrength == bestStrength {
		fmt.Printf("Found terminal bridge: -")
		for _, val := range currentBridge {
			fmt.Printf("%s--", val)
		}
		fmt.Printf("\n")
	}*/

	return bestStrength, bestLength
}

func contains(s []port, test port) bool {
	for _, val := range s {
		if val.A == test.A && val.B == test.B {
			return true
		}
	}
	return false
}

func readComponents(input string) {
	componentCounts = make(map[string]int)
	portMap = make(map[string][]port)
	for _, line := range strings.Split(input, "\n") {
		vals := strings.Split(line, "/")
		valA := getIntFromString(vals[0])
		valB := getIntFromString(vals[1])
		pair := port{A: valA, B: valB}
		if valA > valB {
			pair.A = valB
			pair.B = valA
		}
		pairStr := pair.String()
		cnt, ok := componentCounts[pairStr]
		if !ok {
			cnt = 0
		}
		componentCounts[pairStr] = cnt + 1

		sliceA, ok := portMap[string(valA)]
		if !ok {
			sliceA = []port{}
		}
		if !contains(sliceA, pair) {
			portMap[string(valA)] = append(sliceA, pair)
		}

		sliceB, ok := portMap[string(valB)]
		if !ok {
			sliceB = []port{}
		}
		if !contains(sliceB, pair) {
			portMap[string(valB)] = append(sliceB, pair)
		}
	}

}
func getIntFromString(tmpStr string) int {
	tmp, err := strconv.ParseInt(tmpStr, 10, 64)
	if err != nil {
		fmt.Printf("Unable to parse %s\n", err.Error())
	}
	return int(tmp)
}

const (
	inputA = `0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10`

	input = `14/42
2/3
6/44
4/10
23/49
35/39
46/46
5/29
13/20
33/9
24/50
0/30
9/10
41/44
35/50
44/50
5/11
21/24
7/39
46/31
38/38
22/26
8/9
16/4
23/39
26/5
40/40
29/29
5/20
3/32
42/11
16/14
27/49
36/20
18/39
49/41
16/6
24/46
44/48
36/4
6/6
13/6
42/12
29/41
39/39
9/3
30/2
25/20
15/6
15/23
28/40
8/7
26/23
48/10
28/28
2/13
48/14`
)
