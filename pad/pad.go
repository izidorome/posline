package pad

import (
	"strings"
)

// Right pads string to the right
func Right(str string, lgt int, pad string) string {
	content := []rune(str)

	if len(content) >= lgt {
		return string(content[:lgt])
	}

	return str + strings.Repeat(pad, lgt-len(content))
}

// Left pads string to the left
func Left(str string, lgt int, pad string) string {
	content := []rune(str)

	if len(content) >= lgt {
		return string(content[:lgt])
	}

	return strings.Repeat(pad, lgt-len(content)) + str
}
