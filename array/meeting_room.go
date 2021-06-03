package array

import (
	"container/heap"
	"sort"
)

//type minHeap []int
//
//func (h minHeap) Len() int {
//	return len(h)
//}
//func (h minHeap) Less(i, j int) bool {
//	return h[i] < h[j]
//}
//func (h minHeap) Swap(i, j int) {
//	h[i], h[j] = h[j], h[i]
//}
//func (h *minHeap) Push(v interface{}) {
//	*h = append(*h, v.(int))
//}
//func (h *minHeap) Pop() interface{} {
//	v := (*h)[len(*h)-1]
//	*h = (*h)[:len(*h)-1]
//	return v
//}

func minMeetingRooms(ivs [][]int) int {

	if len(ivs) == 0 {
		return 0
	}

	// Slice sorts the slice ivs given the provided less function. It panics if x is not a slice.
	// sort by start time
	sort.Slice(ivs, func(i, j int) bool {
		return ivs[i][0] < ivs[j][0]
	})

	// insert end time of first schedule into heap
	h := minHeap{ivs[0][1]}

	// iterate from second schedule to the end
	for _, iv := range ivs[1:] {
		// heap's time is less than equal to iv's start time
		// pop the heap
		if h[0] <= iv[0] {
			heap.Pop(&h)
		}
		// insert iv's end time
		heap.Push(&h, iv[1])
	}

	// count of the heap
	return h.Len()

}
