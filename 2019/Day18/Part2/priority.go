package main

import (
	"math/bits"
)

type path struct {
	keys        uint32
	keysOrdered []byte
	currNodeID  []int64
	numMoves    int64
	moveHistory []byte
}

func abs(v int64) int64 {
	if v >= 0 {
		return v
	}
	return -v
}

func (p *path) Score() int64 {
	//return p.numKeys() / (1 + p.numMoves)
	return p.numMoves
}

func (p *path) numKeys() int64 {
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
