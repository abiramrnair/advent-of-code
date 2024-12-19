package utils

import (
	"container/heap"
)

type Coord struct {
	Row int
	Col int
}

var Grid4Directions = []Coord{
	{Row: -1, Col: 0},
	{Row: 1, Col: 0},
	{Row: 0, Col: -1},
	{Row: 0, Col: 1},
}

type Queue struct {
	Elements []any
}

func (q *Queue) Enqueue(element any) {
	q.Elements = append(q.Elements, element)
}

func (q *Queue) Dequeue() any {
	element := q.Elements[0]
	if q.Size() == 1 {
		q.Elements = nil
		return element
	}
	q.Elements = q.Elements[1:]
	return element
}

func (q Queue) Size() int {
	return len(q.Elements)
}

func (q Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Clear() *Queue {
	q.Elements = nil
	return q
}

type PqItem struct {
	Value interface{}
	Priority int
	Index int
}

type PriorityQueue []*PqItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { // Change this to toggle highest/lowest priority pop.
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*PqItem)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *PqItem, Value string, Priority int) {
	item.Value = Value
	item.Priority = Priority
	heap.Fix(pq, item.Index)
}
