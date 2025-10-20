package wordwrap

import "fmt"

func Wrap(text string, width int) ([]string, error) {
	var result []string

	if len(text) < width {
		result = append(result, text)
	} else {
		init := 0
		end := init + width
		for init < len(text) {
			if end > len(text) {
				end = len(text)
			}
			result = append(result, text[init:end])
			init = end
			end = init + width
		}
	}

	fmt.Print(result)

	return result, nil

}
