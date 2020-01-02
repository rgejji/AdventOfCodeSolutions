package main

import (
	"math/bits"
)

type path struct {
	keys        uint32
	keysOrdered []byte
	moveHistory []int
	pt          point
	numMoves    int64
	lastDir     int64
	prevKeyLoc  point
}

type point struct {
	x int64
	y int64
}

func abs(v int64) int64 {
	if v >= 0 {
		return v
	}
	return -v
}

func (p path) Score() int64 {
	sum := int64(0)
	L := len(p.moveHistory)
	/*if L >= 3 {
		sum += int64(p.moveHistory[L-2]) - int64(p.moveHistory[L-3])
	}*/
	/*if L >= 2 {
		sum += int64(p.moveHistory[L-1]) - int64(p.moveHistory[L-2])
	}*/
	if L >= 1 {
		sum += p.numMoves - int64(p.moveHistory[L-1])
	}
	sum -= (abs(p.prevKeyLoc.x-p.pt.x) + abs(p.prevKeyLoc.y-p.pt.y))

	return int64(10000)*p.numKeys()/2 - sum
}

func (p path) numKeys() int64 {
	return int64(bits.OnesCount32(p.keys))
}

// PathHeap is a min-heap of paths
type PathHeap []path

func (h PathHeap) Len() int           { return len(h) }
func (h PathHeap) Less(i, j int) bool { return h[i].Score() < h[j].Score() }
func (h PathHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

//Push will add an element
func (h *PathHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(path))
}

//Pop will remove an element
func (h *PathHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
