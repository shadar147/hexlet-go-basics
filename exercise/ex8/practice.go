package ex8

import (
	"strings"
)

func Greetings(name string) string {
	return strings.Title("привет, " + strings.ToLower(strings.Trim(name, " ")) + "!")
}
