package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var b strings.Builder

	emptyRune := rune(0)

	currentSymbol := emptyRune
	currentCount := 1
	previousIsNumber := true

	symbols := []rune(input)
	for _, symbol := range symbols {
		number, err := strconv.Atoi(string(symbol))
		if err == nil { // is number
			if previousIsNumber {
				return "", ErrInvalidString
			}
			currentCount = number
			previousIsNumber = true
		} else { // !number
			if currentSymbol != emptyRune {
				for i := 0; i < currentCount; i++ {
					b.WriteRune(currentSymbol)
				}
			}
			currentSymbol = symbol
			currentCount = 1
			previousIsNumber = false
		}
	}
	if !previousIsNumber {
		b.WriteRune(currentSymbol)
	}
	return b.String(), nil
}
