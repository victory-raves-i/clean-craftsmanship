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
}

func TestEnqueueMultipleItems(t *testing.T) {
	q := Queue{}

	q.Enqueue(1)
	q.Enqueue(2)

	if q.Size() != 2 {
		t.Error("Enqueue multiple values is not working")
	}
}

func TestDequeueOneItem(t *testing.T) {
	q := Queue{}

	q.Enqueue(10)
	value, _ := q.Dequeue()

	if q.Size() != 0 || value != 10 {
		t.Error("Dequeue one item is not working")
	}
}

func TestDequeueMultipleValues(t *testing.T) {
	q := Queue{}

	q.Enqueue(233)
	q.Enqueue(43)

	value, _ := q.Dequeue()

	if value != 233 || q.Size() != 1 {
		t.Error("Multiple values dequeue is not working")
	}

	value, _ = q.Dequeue()

	if value != 43 || q.Size() != 0 {
		t.Error("Multiple values dequeue is not working")
	}
}

func TestDequeueFromEmptyQueue(t *testing.T) {
	q := Queue{}

	_, error := q.Dequeue()

	if error == nil {
		t.Error("Dequeue from empty Queue it not working")
	}

}
