package justify

import (
	"fmt"
	"os"
)

func AlignedText(lines []string, align string, s string) int {
	cols := GetTerminalSize()
	if len(lines) > cols {
		fmt.Println("ERROR: Terminal size and lines can't fit")
		os.Exit(0)
	}
	numSpaces := len(s) - 1
	var spaceAdder int

	switch align {
	case "left":
		spaceAdder = 0
	case "right":
		spaceAdder = cols - len(s) - 1
	case "center":
		spaceAdder = (cols - len(s)) / 2
	case "justify":
		spaceAdder = ((cols - len(s)) / numSpaces) * 4
	}

	return spaceAdder
}
