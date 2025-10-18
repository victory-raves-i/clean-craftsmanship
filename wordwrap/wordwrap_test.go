package wordwrap

import "testing"

func TestEmptyString(t *testing.T) {
	w, err := Wrap("", 8)

	if err != nil {
		t.Error("Function does not exist or it is not declared")
	}

	if w != "" {
		t.Error("Function does not work with empty string")
	}
}

func TestTextSmallerThanWidth(t *testing.T) {
	w, _ := Wrap("the", 4)

	if w != "the" {
		t.Error("Width checker is not working")
	}

	x, _ := Wrap("the King", 10)

	if x != "the King" {
		t.Error("Width checker is not working")
	}
}
