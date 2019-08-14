package normalize

import (
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
