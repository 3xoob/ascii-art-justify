package main

import (
	"fmt"
	"os"
	"strings"

	Justify "Justify/Functions"
)

func main() {
	checkalign := false
	filename := "Fonts/standard.txt"
	align := "left"
	if len(os.Args) == 1 {
		fmt.Println("ERROR:")
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --align=right \"something\" standard")
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("Usage: go run . [STRING]")
		os.Exit(0)
	}
	if strings.HasPrefix(os.Args[1], "Ascii Art Justify Team") {
		Justify.R01()
		os.Exit(0)
	}
	if len(os.Args[1:]) == 3 && strings.HasPrefix(os.Args[1], "--align=") {
		align = strings.TrimPrefix(os.Args[1], "--align=")
		align = strings.ToLower(align)
		if align != "left" && align != "right" && align != "center" && align != "justify" {
			fmt.Println("ERROR: Invalid alignment option: ", align)
			os.Exit(0)
		}
		filename = Justify.FontChecker(os.Args[3])
		str := strings.ReplaceAll(strings.ReplaceAll(os.Args[2], "\\t", "    "), `\n`, "\n")
		if align == "right" || align == "center" || align == "left" {
			checkalign = true
		}
		if align == "justify" && len(Justify.RemoveExtraSpaces(str)) > 1 {
			Justify.JustifyPrinter(str, filename)
			os.Exit(0)
		}
		Justify.Printer(str, filename, checkalign, align)
		os.Exit(0)
	} else if len(os.Args[1:]) == 2 {
		if strings.HasPrefix(os.Args[1], "--align=") {
			align = strings.TrimPrefix(os.Args[1], "--align=")
			align = strings.ToLower(align)
			if align != "left" && align != "right" && align != "center" && align != "justify" {
				fmt.Println("ERROR: Invalid alignment option: ", align)
				os.Exit(0)
			}
			if align == "right" || align == "center" || align == "left" {
				checkalign = true
			}
			str := strings.ReplaceAll(strings.ReplaceAll(os.Args[2], "\\t", "    "), `\n`, "\n")
			if align == "justify" && len(Justify.RemoveExtraSpaces(str)) > 1 {
				Justify.JustifyPrinter(str, filename)
				os.Exit(0)
			}
			Justify.Printer(str, filename, checkalign, align)
			os.Exit(0)
		} else if os.Args[2] == "standard" || os.Args[2] == "shadow" || os.Args[2] == "thinkertoy" {
			filename = Justify.FontChecker(os.Args[2])
			str := strings.ReplaceAll(strings.ReplaceAll(os.Args[1], "\\t", "    "), `\n`, "\n")
			Justify.Printer(str, filename, checkalign, align)
			os.Exit(0)
		} else {
			fmt.Println("ERROR:")
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --align=right \"something\" standard")
			fmt.Println("Usage: go run . [STRING] [BANNER]")
			fmt.Println("Usage: go run . [STRING]")
			os.Exit(0)
		}

	} else if len(os.Args[1:]) == 1 {
		str := strings.ReplaceAll(strings.ReplaceAll(os.Args[1], "\\t", "    "), `\n`, "\n")
		Justify.Printer(str, filename, checkalign, align)
		os.Exit(0)
	} else {
		fmt.Println("ERROR:")
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --align=right \"something\" standard")
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("Usage: go run . [STRING]")
		os.Exit(0)
	}
}
