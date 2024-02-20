package justify

import (
	"fmt"
	"strings"
)

func Printer(str string, fileTouse string, alignCheck bool, align string) {
	Argument := strings.Contains(str, "\\n")
	res := ""
	if Argument {
		words := strings.Split(str, "\\n")
		check := true
		for _, word := range words {
			if word != "" {
				check = false
			}
		}
		if check {
			words = words[1:]
		}
		for _, word := range words {
			for i := 0; i < 8; i++ {
				for _, letter := range word {
					res += GetAscii(fileTouse, 1+int(letter-' ')*9+i)
				}
				if alignCheck {
					AsciiPrint(res, align)
				} else {
					fmt.Println(res)
				}
				res = ""
			}
		}
	} else {
		res := ""
		for i := 0; i < 8; i++ {
			for _, letter := range str {
				res += GetAscii(fileTouse, 1+int(letter-' ')*9+i)
			}
			if alignCheck {
				AsciiPrint(res, align)
			} else {
				fmt.Println(res)
			}
			res = ""
		}
	}
}
