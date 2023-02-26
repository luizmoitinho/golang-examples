package commmon

import (
	"errors"
	"unicode/utf8"
)

func Reverse(str string) (string, error) {
	if !utf8.ValidString(str) {
		return str, errors.New("input string is not valid utf8")
	}
	value := []rune(str)

	for i, j := 0, len(value)-1; i < len(value)/2; i, j = i+1, j-1 {
		value[i], value[j] = value[j], value[i]
	}
	return string(value), nil
}
