package wordwrap

import (
	s "strings"
)

func Wrap(text string, width int) ([]string, error) {
	var result []string
	if len(text) < width {
		result = append(result, text)
	} else if !s.Contains(text, " ") {
		result = append(result, text[:width])
		result = append(result, text[width:])
	}

	return result, nil

}
