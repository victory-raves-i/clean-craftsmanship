package wordwrap

func Wrap(text string, width int) ([]string, error) {
	var s []string

	if len(text) < width {
		s = append(s, text)
		return s, nil
	} else {
		s = append(s, text[:width])
		s = append(s, text[width:])

	}

	return s, nil

}
