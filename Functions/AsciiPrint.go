package justify

import (
	"fmt"
	"os"
)

func AsciiPrint(res string, align string) {
	terminalWidth := GetTerminalSize()
	if align == "right" {
		numberOfSpaces := terminalWidth - len(res)
		if numberOfSpaces < 0 {
			fmt.Println("ERROR: The text is longer than the terminal width.")
			os.Exit(0)
		}
		for i := 0; i < numberOfSpaces; i++ {
			fmt.Print(" ")
		}
		fmt.Print(res + "\n")
	} else if align == "center" {
		numberOfSpaces := terminalWidth - len(res)
		if numberOfSpaces < 0 {
			fmt.Println("ERROR: The text is longer than the terminal width.")
			os.Exit(0)
		}
		for i := 0; i < numberOfSpaces/2; i++ {
			fmt.Print(" ")
		}
		fmt.Print(res + "\n")
	} else if align == "left" {
		numberOfSpaces := terminalWidth - len(res)
		if numberOfSpaces < 0 {
			fmt.Println("ERROR: The text is longer than the terminal width.")
			os.Exit(0)
		}
		fmt.Print(res + "\n")
	}
}
