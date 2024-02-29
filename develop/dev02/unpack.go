package unpack

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrIncorrectString = errors.New("некорректная строка")

func Unpack(str string) (string, error) {
	var (
		res      strings.Builder
		prevChar rune
	)

	if len(str) == 0 {
		return str, ErrIncorrectString
	}

	for i, char := range str {
		if unicode.IsDigit(char) {
			if i == 0 || unicode.IsDigit(prevChar) {
				return str, ErrIncorrectString
			}

			number, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Println(err)
				return str, ErrIncorrectString
			}

			for i := 0; i < number-1; i++ {
				res.WriteRune(prevChar)
			}
			prevChar = char
			continue
		}

		res.WriteRune(char)
		prevChar = char
	}

	return res.String(), nil
}
