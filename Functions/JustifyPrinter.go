package justify

import (
	"fmt"
	"os"
	"strings"
)

func JustifyPrinter(s string, font string, align string, numofsp int) {
	fontArray, err := GetFont(font)
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(0)
	}
	charArray := initializeLines(s)
	for i := 0; i < len(s); i++ {
		if s[i] == 32 {
			for linePos, _ := range GetCharacter(rune(s[i]), fontArray) {
				charArray[linePos+len(charArray)-8] += strings.Repeat(" ", numofsp)
			}
		} else if s[i] != '\n' && s[i] > 32 && s[i] <= 126 {
			for linePos, line := range GetCharacter(rune(s[i]), fontArray) {
				charArray[linePos+len(charArray)-8] += line
			}
		} else if s[i] == '\n' {
			if i == 0 || i == len(s)-1 || s[i+1] == '\n' {
				charArray = append(charArray, make([]string, 1)...)
			} else {
				charArray = append(charArray, make([]string, 8)...)
			}
		} else {
			fmt.Println("ERROR: Invalid character")
			os.Exit(0)
		}
	}
	for _, line := range charArray {
		fmt.Println(line)
	}
}
