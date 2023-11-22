package stringutil

import (
	"unicode/utf8"
)

// Unique returns the deduplicated slice of string.
func Unique(slice []string) (result []string) {
	smap := make(map[string]bool)
	for _, s := range slice {
		smap[s] = true
	}
	for s := range smap {
		result = append(result, s)
	}
	return result
}

// Return the index of string in the array.
// If it does not exist in the array, return -1.
func FindString(array []string, str string) int {
	for index, s := range array {
		if str == s {
			return index
		}
	}
	return -1
}

// StringIn determine if string is in the target array.
func StringIn(array []string, str string) bool {
	return FindString(array, str) > -1
}

// Reverse return the flipped string.
func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}
