package padding

import (
	"errors"
	"unicode/utf8"
)

type PadType int

const (
	// PadLeft is padding type to pad into left side.
	PadLeft PadType = iota
	// PadRight is padding type to pad into right side.
	PadRight
)

// Pad pad string to your string.
func Pad(source string, length int, padStr string, padType PadType) (string, error) {
	if isAlreadyReachedLength(source, length) {
		return source, nil
	}

	str := source

	switch padType {
	case PadLeft:
		for {
			str = padStr + str

			if utf8.RuneCountInString(str) == length {
				break
			}

			if utf8.RuneCountInString(str) > length {
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
			}

			if utf8.RuneCountInString(str) > length {
				str = str[:length]
				break
			}
		}
	default:
		return "", errors.New("unknown padding type is specified")
	}

	return str, nil
}

// isAlreadyReachedLength checks if the length of the source has already exceeded the specified length.
func isAlreadyReachedLength(source string, length int) bool {
	return utf8.RuneCountInString(source) >= length
}
