package justify

import (
	"fmt"
	"os"
	"strings"
)

func AlignedTextPrinter(lines []string, align string, s string) string {
	cols := GetTerminalSize()
	if len(lines) > cols {
		fmt.Println("ERROR: Terminal size and lines can't fit")
		os.Exit(0)
	}
	numWords := len(s) - 1
	var spaceAdder string

	switch align {
	case "left":
		spaceAdder = ""
	case "right":
		spaceAdder = strings.Repeat(" ", cols-len(s)-1)
	case "center":
		spaceAdder = strings.Repeat(" ", (cols-len(s))/2)
	case "justify":
		spaceAdder = strings.Repeat(" ", (cols-len(s))/numWords)
	}

	return spaceAdder
}
