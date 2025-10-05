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
