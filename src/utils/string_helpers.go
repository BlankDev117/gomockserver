package utils

// SubstrLength substring a given string at the given specified index a specified number of characters
func SubstrLength(input string, start int, length int) string {
	if start < 0 {
		panic("start index can not be negative")
	}

	asRunes := []rune(input)

	if start >= len(asRunes) {
		panic("start can not be greater than length of string")
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}
