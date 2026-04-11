package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func processMultiplier(result *strings.Builder, prevChar string, count int) {
	if count == 0 {
		current := result.String()
		result.Reset()
		result.WriteString(current[:len(current)-1])
		return
	}
	result.WriteString(strings.Repeat(prevChar, count-1))
}

func Unpack(text string) (string, error) {
	var result strings.Builder
	var prevStringChar string
	var hasStringPrevChar bool
	var prevEscaped bool

	for _, runeValue := range text {
		// обработка символа экранирования
		if !prevEscaped && runeValue == '\\' {
			prevEscaped = true
			continue
		}

		// обработка цифры как множителя (только если не экранирована)
		if unicode.IsDigit(runeValue) && !prevEscaped {
			if !hasStringPrevChar {
				return "", ErrInvalidString
			}

			count, err := strconv.Atoi(string(runeValue))
			if err != nil {
				return "", errors.New("error get int from char")
			}

			processMultiplier(&result, prevStringChar, count)
			hasStringPrevChar = false
			continue
		}

		// обычный символ (включая экранированные цифры)
		hasStringPrevChar = true
		prevStringChar = string(runeValue)
		result.WriteRune(runeValue)
		prevEscaped = false
	}
	return result.String(), nil
}
