package padding

import (
	"errors"
	"unicode/utf8"
)

type PadType int

const (
	PadLeft PadType = iota
	PadRight
)

// Pad pad string to your string
func Pad(source string, length int, padStr string, padType PadType) (string, error) {
	var str string
	str = source

	if isAlreadyReachedLength(source, length) {
		return source, nil
	}

	switch padType {
	case PadLeft:
		for {
			str = padStr + str

			if utf8.RuneCountInString(str) == length {
				break
			} else if utf8.RuneCountInString(str) > length {
				var exceeded = utf8.RuneCountInString(str) - length
				str = str[exceeded : length+exceeded]
				break
			}
		}
	case PadRight:
		for {
			str = str + padStr

			if utf8.RuneCountInString(str) == length {
				break
			} else if utf8.RuneCountInString(str) > length {
				str = str[:length]
				break
			}
		}
	default:
		return "", errors.New("unknown padding type is specified")
	}

	return str, nil
}

// isAlreadyReachedLength checks if the length of the source has already exceeded the specified length
func isAlreadyReachedLength(source string, length int) bool {
	return utf8.RuneCountInString(source) >= length
}
