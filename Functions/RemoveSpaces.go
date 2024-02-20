package justify

func RemoveExtraSpaces(str string) []string {
	nbw := 0
	var prel rune
	prel = ' '
	word := ""

	for _, v := range str {
		if (v != ' ' && v != '\t' && v != '\n') && (prel == ' ' || prel == '\t' || prel == '\n') {
			nbw++
		}
		prel = v
	}

	ar := make([]string, 0)

	for _, v := range str {
		if v != ' ' && v != '\t' && v != '\n' {
			word = word + string(v)
		}
		if (v == ' ' || v == '\t' || v == '\n') && word != "" {
			ar = append(ar, word)
			word = ""
		}
	}

	if word != "" {
		ar = append(ar, word)
	}

	return ar
}
