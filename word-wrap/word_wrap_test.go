package wordwrap

import "testing"

func TestCreateEmptyTest(t *testing.T) {
	t.Error("It Fails")
}

func TestEmptyString(t *testing.T) {
	w := Wrap(text, width)

	if w == nil {
		t.Error("Function does not exist or it is not declared")
	}
}
