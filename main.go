package main

import (
	"fmt"
	"os"
	"strings"

	Justify "Justify/Functions"
)

func main() {
	// calling all the function to work
	if len(os.Args) == 1 {
		fmt.Println("ERROR:")
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --align=right \"something\" standard")
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("Usage: go run . [STRING]")
		os.Exit(0)
	}
	if strings.HasPrefix(os.Args[1], "Ascii Art Justify Team") {
		Justify.R01()
		return
	}
	if len(os.Args[1:]) == 3 && strings.HasPrefix(os.Args[1], "--align=") {
		align := strings.TrimPrefix(os.Args[1], "--align=")
		align = strings.ToLower(align)
		if align != "left" && align != "right" && align != "center" && align != "justify" {
			fmt.Println("ERROR: Invalid alignment option: ", align)
			os.Exit(0)
		}
		if align == "justify" {
			Justify.JustifyPrinter(strings.ReplaceAll(strings.ReplaceAll(os.Args[2], "\\t", "    "), `\n`, "\n"), os.Args[3], align)
		} else {
			Justify.AsciiPrint(strings.ReplaceAll(strings.ReplaceAll(os.Args[2], "\\t", "    "), `\n`, "\n"), os.Args[3], align)
		}
	} else if len(os.Args[1:]) == 2 {
		if strings.HasPrefix(os.Args[1], "--align=") {
			align := strings.TrimPrefix(os.Args[1], "--align=")
			align = strings.ToLower(align)
			if align != "left" && align != "right" && align != "center" && align != "justify" {
				fmt.Println("ERROR: Invalid alignment option: ", align)
				os.Exit(0)
			}
			if align == "justify" {
				Justify.JustifyPrinter(strings.ReplaceAll(strings.ReplaceAll(os.Args[2], "\\t", "    "), `\n`, "\n"), "standard", align)
			} else {
				Justify.AsciiPrint(strings.ReplaceAll(strings.ReplaceAll(os.Args[2], "\\t", "    "), `\n`, "\n"), "standard", align)
			}
		} else if os.Args[2] == "standard" || os.Args[2] == "shadow" || os.Args[2] == "thinkertoy" {
			Justify.AsciiPrint(strings.ReplaceAll(strings.ReplaceAll(os.Args[1], "\\t", "    "), `\n`, "\n"), os.Args[2], "")
		} else {
			fmt.Println("ERROR:")
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --align=right \"something\" standard")
			fmt.Println("Usage: go run . [STRING] [BANNER]")
			fmt.Println("Usage: go run . [STRING]")
			os.Exit(0)
		}
	} else if len(os.Args[1:]) == 1 {
		Justify.AsciiPrint(strings.ReplaceAll(strings.ReplaceAll(os.Args[1], "\\t", "    "), `\n`, "\n"), "standard", "")
	} else {
		fmt.Println("ERROR:")
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --align=right \"something\" standard")
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("Usage: go run . [STRING]")
		os.Exit(0)
	}
}

// note for future s = strings.ReplaceAll(s, "\\t", "    ")
