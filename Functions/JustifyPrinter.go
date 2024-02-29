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
				fmt.Println("ERROR: String cannot end with space in justify alignment.")
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
				fmt.Println("ERROR: String cannot start with space in justify alignment.")
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
		textLen += len(arrOfStr[0])
	}
	numSpaces := (terminalWidth - textLen) / (len(sws) - 1)

	for i := 0; i < 8; i++ {
		for k, v := range ar {
			fmt.Print(v[i])
			if k != len(ar)-1 {
				fmt.Print(SpacePrinter(numSpaces))
			}
		}
		fmt.Println()
	}

}
