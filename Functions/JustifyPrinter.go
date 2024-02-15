package justify

import (
	"fmt"
	"os"
	"strings"
)

func JustifyPrinter(s string, font string, align string) {
	fontArray, err := GetFont(font)
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(0)
	}
	firstarra := AsciiPrint(s, font, "")
	alignner := AlignedTextPrinter(firstarra, align, s)
	array := strings.Split(s, " ")
	for i, word := range array {
		charArray := initializeLines(word)
		for i := 0; i < len(array); i++ {
			if word[i] != '\n' && word[i] >= 32 && word[i] <= 126 {
				for linePos, line := range GetCharacter(rune(word[i]), fontArray) {
					charArray[linePos+len(charArray)-8] += line
				}
			} else if word[i] == '\n' {
				if i == 0 || i == len(word)-1 || word[i+1] == '\n' {
					charArray = append(charArray, make([]string, 1)...)
				} else {
					charArray = append(charArray, make([]string, 8)...)
				}
			} else {
				fmt.Println("ERROR: Invalid character")
				os.Exit(0)
			}
		}
		if word[i] == 32 {
			for _, line := range charArray {
				
				result := line + alignner
				fmt.Println(result)
			}
		}

	}

}
