package justify

import (
	"fmt"
	"os"
	"strings"
)

func AsciiPrint(s string, font string, align string) []string {
	fontArray, err := GetFont(font)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return nil
	}
	if align == "justify" {
		s = removeExtraSpaces(s)
	}
	charArray := initializeLines(s)
	for i := 0; i < len(s); i++ {
		if s[i] != '\n' && s[i] >= 32 && s[i] <= 126 {
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
			return nil
		}
	}
	for _, line := range charArray {
		alignner := AlignedText(line, align, s)
		if align != "" {
			if align == "justify" {
				JustifyPrinter(s, font, align, alignner)
				os.Exit(0)
			} else {
				result := strings.Repeat(" ", alignner) + line
				fmt.Println(result)
			}

		} else {
			fmt.Println(line)
		}
	}
	return charArray
}
