package justify

import (
	"bufio"
	"fmt"
	"os"
)

func GetAscii(filename string, num int) string {
	file, e := os.Open(filename)
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(0)
	}
	defer file.Close()
	content := bufio.NewScanner(file)
	lineNum := 0
	line := ""
	for content.Scan() {
		if lineNum == num {
			line = content.Text()
		}
		lineNum++
	}
	return line
}
