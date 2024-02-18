package justify

import (
	"regexp"
)

func removeExtraSpaces(input string) string {
	regex := regexp.MustCompile(`\s+`)
	result := regex.ReplaceAllString(input, " ")
	return result
}
