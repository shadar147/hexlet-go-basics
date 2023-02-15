package ex5

import (
	"errors"
	"fmt"
)

// публичная функция, можно вызвать извне как math.Multiply(5,7)
func Multiply(x int, y int) (res int) { // Использовать именованные возвращаемые аргументы — плохая практика.
	return x * y
}

// приватная функция, извне не вызвать
func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("cannot divide on zero")
	}

	return x / y, nil
}

func myPrint(msg string) {
	// пакет.функция
	fmt.Println(msg)
}
