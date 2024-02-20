package justify

import "fmt"

func AsciiPrint(res string, align string) {
	terminalWidth := GetTerminalSize()
	if align == "right" {
		numberOfSpaces := terminalWidth - len(res)
		for i := 0; i < numberOfSpaces; i++ {
			fmt.Print(" ")
		}
		fmt.Print(res + "\n")
	} else if align == "center" {
		numberOfSpaces := terminalWidth - len(res)
		for i := 0; i < numberOfSpaces/2; i++ {
			fmt.Print(" ")
		}
		fmt.Print(res + "\n")
	}
}
