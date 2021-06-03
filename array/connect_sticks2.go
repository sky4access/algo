package array

import "container/heap"

type minHeap []int

func (m *minHeap) Len() int { return len(*m) }

func (m *minHeap) Less(i, j int) bool { return (*m)[i] < (*m)[j] }

func (m *minHeap) Swap(i, j int) { (*m)[i], (*m)[j] = (*m)[j], (*m)[i] }

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *minHeap) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func connectSticks2(sticks []int) int {

	var minH minHeap

	heap.Init(&minH)

	for i := 0; i < len(sticks); i++ {
		heap.Push(&minH, sticks[i])
	}

	totalCost := 0

	for len(minH) != 1 {
		s1 := heap.Pop(&minH).(int)
		s2 := heap.Pop(&minH).(int)

		sum := s1 + s2
		totalCost += sum
		heap.Push(&minH, sum)
	}

	return totalCost
}
