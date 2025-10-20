package wordwrap

import "fmt"

func Wrap(text string, width int) ([]string, error) {
	var result []string
	var length int // Length of the text

	length = len(text)

	if length < width {
		result = append(result, text)
	} else {
		init := 0
		end := init + width
		for init < length {
			if end > length {
				end = length
			}
			result = append(result, text[init:end])
			init = end
			end = init + width
		}
	}

	fmt.Print(result)

	return result, nil

}
