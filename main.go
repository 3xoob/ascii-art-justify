package main

import (
	Justify "Justify/Functions"
	"fmt"
	"os"
	"strings"
)

func main() {
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
		Justify.AsciiPrint(strings.ReplaceAll(os.Args[2], `\n`, "\n"), os.Args[3], align)
	} else if len(os.Args[1:]) == 2 {
		if strings.HasPrefix(os.Args[1], "--align=") {
			align := strings.TrimPrefix(os.Args[1], "--align=")
			Justify.AsciiPrint(strings.ReplaceAll(os.Args[2], `\n`, "\n"), "standard", align)
		} else if os.Args[2] == "standard" || os.Args[2] == "shadow" || os.Args[2] == "thinkertoy" {
			Justify.AsciiPrint(strings.ReplaceAll(os.Args[1], `\n`, "\n"), os.Args[2], "")
		} else {
			fmt.Println("ERROR:")
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --align=right \"something\" standard")
			fmt.Println("Usage: go run . [STRING] [BANNER]")
			fmt.Println("Usage: go run . [STRING]")
			os.Exit(0)
		}
	} else if len(os.Args[1:]) == 1 {
		Justify.AsciiPrint(strings.ReplaceAll(os.Args[1], `\n`, "\n"), "standard", "")
	} else {
		fmt.Println("ERROR:")
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --align=right \"something\" standard")
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("Usage: go run . [STRING]")
		os.Exit(0)
	}

}
