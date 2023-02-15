// Определение пакета main
package main

import (
	"fmt"
	"hexlet/go-basics/exercise/ex7"
)

// Определение функции main
func main() {
	res := ex7.IsValid(225, "hello world")

	fmt.Println(res)
}
