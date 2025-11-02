package wordwrap

import (
	"fmt"
	s "strings"
)

func Wrap(text string, width int) ([]string, error) {
	var result []string

	length := len(text)

	if length < width {
		result = append(result, text)
	} else {
		init := 0
		end := init + width
		for init < length {
			if end > length {
				end = length
			}

			lastSpace := s.LastIndex(text[init:end], " ")
			if lastSpace != -1 {
				end = init + lastSpace + 1
			}

			result = append(result, text[init:end])
			init = end
			end = init + width
		}
	}

	fmt.Print(s.Repeat("-", 20) + "\n")
	text = s.Join(result, "\n")
	fmt.Println(text)
	fmt.Print(s.Repeat("-", 20))

	return result, nil

}
