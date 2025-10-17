package queue

import (
	"errors"
)

type Queue struct {
	elements []int
}

func (q *Queue) Size() int {
	return len(q.elements)
}

func (q *Queue) Enqueue(value int) {
	q.elements = append(q.elements, value)
}

func (q *Queue) Dequeue() (int, error) {

	if q.Size() == 0 {
		return 0, errors.New("Queue is empty")
	}
	value := q.elements[0]
	q.elements = q.elements[1:]

	return value, nil
}

func (q *Queue) Peek() (int, error) {

	if q.Size() == 0 {
		return 0, errors.New("You cannot read a value form an empty queue")
	}

	return q.elements[0], nil
}
