package utils

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
