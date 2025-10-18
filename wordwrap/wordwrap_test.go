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
