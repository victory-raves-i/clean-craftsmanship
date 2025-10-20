package wordwrap

import "testing"

func TestEmptyString(t *testing.T) {
	w, err := Wrap("", 8)

	if err != nil {
		t.Error("Function does not exist or it is not declared")
	}

	if w[0] != "" {
		t.Error("Function does not work with empty string")
	}
}

func TestTextSmallerThanWidth(t *testing.T) {
	w, _ := Wrap("the", 4)

	if w[0] != "the" {
		t.Error("Width checker for 1 word is not working")
	}

	x, _ := Wrap("the King", 10)

	if x[0] != "the King" {
		t.Error("Width checker for two words is not working")
	}

	y, _ := Wrap("the King of the world", 22)

	if y[0] != "the King of the world" {
		t.Error("Width checker for more than two words is not working")
	}
}

func TestTextBiggerThanWidth(t *testing.T) {
	w, _ := Wrap("the", 2)

	if w[0] != "th" || w[1] != "e" {
		t.Error("Not able to wrap a single word")
	}

	x, _ := Wrap("the King", 2)

	if x[0] != "th" || x[1] != "e " || x[2] != "Ki" || x[3] != "ng" {
		t.Error("Not able to wrap text with two words without taking into consideration spaces")
	}

}
