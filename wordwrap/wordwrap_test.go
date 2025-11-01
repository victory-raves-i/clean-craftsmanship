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

	w, _ = Wrap("the King", 10)

	if w[0] != "the King" {
		t.Error("Width checker for two words is not working")
	}

	w, _ = Wrap("the King of the world", 22)

	if w[0] != "the King of the world" {
		t.Error("Width checker for more than two words is not working")
	}
}

func TestTextBiggerThanWidth(t *testing.T) {
	w, _ := Wrap("the", 2)

	if w[0] != "th" || w[1] != "e" {
		t.Error("Not able to wrap a single word")
	}

	w, _ = Wrap("the King", 2)

	if w[0] != "th" || w[1] != "e " || w[2] != "Ki" || w[3] != "ng" {
		t.Error("Not able to wrap text with two words without taking into consideration spaces")
	}

	w, _ = Wrap("the King", 3)
	if w[0] != "the" || w[1] != " Ki" || w[2] != "ng" {
		t.Error("Not able to wrap text with three characters without taking into consideration spaces")
	}

}
