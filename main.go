// Определение пакета main
package main

import (
	"fmt"
	"hexlet/go-basics/exercise/ex9"
)

// Определение функции main
func main() {
	res := ex9.DomainForLocale("github.com", "")

	fmt.Println(res)
}
