package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {

	previousChar := ""
	var isChar bool
	var w strings.Builder

	for _, r := range input {
		current := string(r)
		if current == "" {
			return "", nil
		}
		if _, err := strconv.Atoi(current); err == nil && previousChar == "" {
			return "", ErrInvalidString
		}
		if i, err := strconv.Atoi(current); err == nil {
			if !isChar {
				return "", ErrInvalidString
			} else {
				w.WriteString(strings.Repeat(previousChar, i))
				isChar = false
			}
		} else {
			if previousChar == "" {
				previousChar = current
			} else if isChar {
				w.WriteString(previousChar)
			}
			previousChar = current
			isChar = true
		}
	}
	w.WriteString(previousChar)
	return w.String(), nil
}
