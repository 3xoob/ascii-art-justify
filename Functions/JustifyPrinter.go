package justify

import (
	"fmt"
	"strings"
)

func JustifyPrinter(words string, filename string) {
	terminalWidth := GetTerminalSize()
	number1 := len(words)
	number2 := len(words)
	checkSuffix := false
	checkPrefix := false
	for !checkSuffix {
		if number1 > 0 && words[number1-1:number1] == " " {
			words = words[:number1-1]
			if number1 == number2 {
				fmt.Println("ERROR: string cannot end with space for justify alignment, removing trailing spaces")
			}
			number1--
		} else {
			checkSuffix = true
		}
	}
	number3 := len(words)
	for !checkPrefix {
		if strings.HasPrefix(words, " ") {
			words = words[1:]
			if len(words) == number3-1 {
				fmt.Println("ERROR: string cannot start with space for justify alignment, adjusting text")
			}
		} else {
			checkPrefix = true
		}
	}

	sws := RemoveExtraSpaces(words)
	ar := make([][]string, len(sws))
	j := 0
	container := ""
	for i := 0; i < 8; i++ {
		j = 0
		for _, letter := range words {
			if letter != ' ' {
				container += GetAscii(filename, 1+int(letter-' ')*9+i)
			} else if letter == ' ' && container != "" {
				ar[j] = append(ar[j], container)
				container = ""
				j++
			}
		}
		ar[j] = append(ar[j], container)
		container = ""
	}
	textLen := 0
	for _, arrOfStr := range ar {
		if len(arrOfStr) != 0 {
			textLen += len(arrOfStr[0])
		} 
	}
	numSpaces := (terminalWidth - textLen) / (len(sws) - 1)

	for i := 0; i < 8; i++ {
		for k, v := range ar {
			if len(v) == 0 && i == 0 {
				fmt.Print("\n")
			} else if len(v) != 0 {
			fmt.Print(v[i])
			if k != len(ar)-1 {
				fmt.Print(SpacePrinter(numSpaces))
			}
		}
		}
		fmt.Println()
	}

}
