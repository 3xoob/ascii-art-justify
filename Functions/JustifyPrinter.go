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
	// لما تكتب e   olllfhfhhdh\ndhdhxhdhdhdhhdhdhdhdhxhxbxbdbbxbdbxbxnzx مثلَا 
	// يرجع ليك  array طوله 3 
	// المفروض طوله 2 لأن في نيو لاين وحدة
	// fmt.Print(len(sws))

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
		// حطيتها عشان لو الأري فاضي ماتحسب عند الانديكس 0 وتطلع بانيك
		if len(arrOfStr) != 0 {
			textLen += len(arrOfStr[0])
		} 
	}

	numSpaces := (terminalWidth - textLen) / (len(sws) - 1)

		//			for i := 0; i < 8; i++ {  انت كنت حاط 
		//فقمت وقلبتها عشان يطبع نيولاين بين السطور لأن ماكان يطبع range حاطها قبل الثانية مالت ال
		
	for k, v := range ar {
		for i := 0; i < 8; i++ {
		//  يعد نفس الخردة عشان لو السطر فاضي ماتدخله وتطلع بانيك  
		if len(v) != 0 {
			fmt.Print(v[i])
			if k != len(ar)-1 {
				fmt.Print(SpacePrinter(numSpaces))
			}
			// عشان النيولاين بين السطور ، اذا وصلت أخر سطر ماتطبع عقبه
			if i != 7 {
				fmt.Println()
			}
			}
		}
					// عشان النيولاين بين السطور ، اذا وصلت أخر جملة ماتطبع عقبها

		if k != len(v)-1 {
			fmt.Println()
		}
	}

}
