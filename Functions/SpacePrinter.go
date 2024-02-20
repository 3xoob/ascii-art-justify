package justify

func SpacePrinter(num int) string {
	a := ""
	for i := 1; i <= num; i++ {
		a += " "
	}
	return a
}
