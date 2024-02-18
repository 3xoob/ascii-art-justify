package justify

import (
	"fmt"
	"os"
)

func AlignedText(lines string, align string, s string) int {
	cols := GetTerminalSize()
	if len(lines) > cols {
		fmt.Println("ERROR: Terminal size and lines can't fit")
		os.Exit(0)
	}
	if align == "justify" {
		s = removeExtraSpaces(s)
	}
	countSpaces := 0
	for _, ch := range s {
		if rune(ch) == 32 {
			countSpaces++
		}
	}
	var spaceAdder int

	switch align {
	case "left":
		spaceAdder = 0
	case "right":
		spaceAdder = cols - len(s) - 1
	case "center":
		spaceAdder = (cols - len(s)) / 2
	case "justify":
		if countSpaces > 0 {
			spaceAdder = ((cols - len(s)) / countSpaces)
		} else {
			spaceAdder = 0
		}
	}

	return spaceAdder
}
