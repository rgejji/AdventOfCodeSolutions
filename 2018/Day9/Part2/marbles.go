package main

import "fmt"

//CircleListNode is a circular linked list of nodes
type CircleListNode struct {
	prev  *CircleListNode
	next  *CircleListNode
	value int64
}

//AddNode adds a node to the circular linked list right after the given node
func (cln *CircleListNode) AddNode(value int64) *CircleListNode {
	if cln == nil || value == 0 {
		newNode := &CircleListNode{
			value: 0,
		}
		newNode.next = newNode
		newNode.prev = newNode
		return newNode
	}
	nextNode := cln.next
	newNode := &CircleListNode{value: value}
	newNode.prev = cln
	newNode.next = nextNode
	cln.next.prev = newNode
	cln.next = newNode

	return newNode
}

//RemoveNode removes a node frome the circular linked list and returns the next node in the list. If the list becomes empty after the last node is removed, return nil
func (cln *CircleListNode) RemoveNode() (*CircleListNode, int64) {
	if cln == nil || cln.next == cln {
		return nil, 0
	}

	returnValue := cln.value
	prevNode := cln.prev
	nextNode := cln.next
	prevNode.next = nextNode
	nextNode.prev = prevNode

	return nextNode, returnValue
}

func (cln *CircleListNode) String() string {
	startValue := cln.value
	returnStr := fmt.Sprintf("%d", startValue)
	currNode := cln.next
	for currNode.value != startValue {
		returnStr += fmt.Sprintf(", %d", currNode.value)
		currNode = currNode.next
	}
	return returnStr
}

func getMaxScore(scores []int64) (int64, int64) {
	var currMax int64
	var maxPlayer int64

	for i, score := range scores {
		if currMax < score {
			currMax = score
			maxPlayer = int64(i)
		}
	}
	return maxPlayer, currMax
}

func main() {
	var cln *CircleListNode

	//setup initial marble
	cln = cln.AddNode(0)
	//fmt.Printf("%d %d %d\n", cln.value, cln.next.value, cln.prev.value)

	//setup player score
	playerScores := make([]int64, numPlayers, numPlayers)
	currPlayer := 1
	tmpScore := int64(0)

	//Go through each marble
	for i := int64(1); i <= lastMarble; i++ {
		//fmt.Printf("i=%d:\n", i)
		//place marble according to rules
		if i%23 == 0 {
			//update score
			playerScores[currPlayer] += i
			for j := 0; j < 7; j++ {
				cln = cln.prev
			}
			cln, tmpScore = cln.RemoveNode()
			playerScores[currPlayer] += tmpScore
		} else {
			cln = cln.next
			cln = cln.AddNode(i)
		}
		//fmt.Printf("%s\n", cln.String())
		//change player
		currPlayer = (currPlayer + 1) % numPlayers
	}
	player, score := getMaxScore(playerScores)
	fmt.Printf("Player %d earned the max score of %d\n", player, score)

}

const (
	//	numPlayers = 10
	//	lastMarble = 1618
	// numPlayers = 30
	// lastMarble = 5807
	numPlayers = 430
	lastMarble = 7158800
)
