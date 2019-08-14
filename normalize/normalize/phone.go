package normalize

import (
	"regexp"
	"strings"
)

// Normalize accepts a number string and normalizes it
func Normalize(number string) string {

	// filter removes chars other than digits
	filter := func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}

	return strings.Map(filter, number)
}

// RegexNormalize accepts a number string and normalizes it
// uses Regular Expressions
func RegexNormalize(number string) string {
	re := regexp.MustCompile("\\D")
	return re.ReplaceAllString(number, "")
}
