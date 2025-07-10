package utils

import (
	"slices"
	"strings"
)

var Numbers = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ".", "e", "-"}

func IsAlphaNumeric(c byte) bool {
	return (c <= 'z' && c >= 'a') || (c <= '9' && c >= '0') || (c <= 'Z' && c >= 'A')
}

func IsAllAlphaNumeric(str string) bool {
	for i := range len(str) {
		if !IsAlphaNumeric(str[i]) {
			return false
		}
	}
	return true
}

func IsNull(str string) bool {
	return str == "null"
}

func IsBool(str string) bool {
	return str == "false" || str == "true"
}

func IsNumber(str string) bool {
	for idx := range len(str) {
		if !slices.Contains(Numbers, string(str[idx])) {
			return false
		}
	}
	return true
}

func IsFloat(str string) bool {
	return strings.Contains(str, ".") || strings.Contains(str, "e")
}

func ParseArray(idx *int, Tokens []string) {
	//TO DO
}

func ParseObject(idx *int, Tokens []string) {
	// SHOULD BE RESCURSIVE

	// TODO
}
