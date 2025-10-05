package queue

import (
	"testing"
)

func TestCreateEmptyQueue(t *testing.T) {
	q := Queue{}

	if q.Size() != 0 {
		t.Errorf("Queue doesn't exists")
	}
}

func TestEnqueueOneItem(t *testing.T) {
	q := Queue{}
	q.Enqueue(10)
	if q.Size() != 1 {
		t.Error("Something went wrong trying to enqueue and item")
	}
	if q.elements[q.Size()-1] != 10 {
		t.Error("Queue its no able to store one value correctly")
	}
}

func TestEnqueueMultipleItems(t *testing.T) {
	q := Queue{}

	q.Enqueue(1)
	q.Enqueue(2)

	if q.Size() != 2 {
		t.Error("Enqueue multiple values is not working")
	}
}
