package ex23

import (
	"strings"
	"unicode"
)

func LatinLetters(s string) string {
	res := strings.Builder{}

	for _, symbol := range s {
		if unicode.Is(unicode.Latin, symbol) {
			res.WriteString(string(symbol))
		}
	}

	return res.String()
}
