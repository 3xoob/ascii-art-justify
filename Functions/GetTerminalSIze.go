package justify

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetTerminalSize() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("ERROR: Size of the terminal")
		os.Exit(0)
	}

	size := strings.Split(strings.TrimSpace(string(out)), " ")
	width, err := strconv.Atoi(size[1])
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(0)
	}

	return width
}
