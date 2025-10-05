package queue

type Queue struct {
	elements []int
}

func (q *Queue) Size() int {
	return len(q.elements)
}

func (q *Queue) Enqueue(value int) {
	q.elements = append(q.elements, value)
}

func (q *Queue) Dequeue() int {
	value := q.elements[0]
	q.elements = q.elements[1:]

	return value
}
