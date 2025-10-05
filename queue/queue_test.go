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

func TestEnqueue(t *testing.T) {
	q := Queue{}
	q.Enqueue(1)
	if q.Size() != 1 {
		t.Error("Somthing went wrong trying to enqueue and item")
	}
}
