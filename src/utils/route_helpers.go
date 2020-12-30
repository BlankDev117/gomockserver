package utils

import "strings"

// SanitizeRoute mutates the given url path to remove excess '/' characters, etc.
func SanitizeRoute(path string) string {
	return strings.Trim(path, "/")
}
