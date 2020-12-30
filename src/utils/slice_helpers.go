package utils

// StringInSlice will determine if the provided string value is in the given string slice
func StringInSlice(a string, list []string) bool {
	if list == nil {
		return false
	}

	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}
