package ex20

import (
	"fmt"
	"reflect"
)

// Строки в Go — это иммутабельные массивы байт. Для стандартного компилятора Go внутренняя структура строки описана как:
type _string struct {
	elements *byte // байты
	len      int   // кол-во байт
}

// После инициализации строку нельзя изменить и такая иммутабельность позволяет избежать побочных эффектов в коде.
//s := "hello"
//s[4] = "" // ошибка компиляции: cannot assign to s[4] (strings are immutable)

// Стоит отметить, что тип данных byte — это алиас к типу uint8 (0-255). Во-первых, потому что нужно абстрактно отличать типы в коде. Во-вторых, байты представляют ASCII символы, а в кодовой таблице ASCII символов 256 кодов:
func printHey() {
	s := "hey"

	fmt.Println(s[0], s[1], s[2]) // 104 101 121

	fmt.Println(string(s[0]), string(s[1]), string(s[2])) // h e y
}

// Большинство библиотечных функций работают со слайсами байт []byte для производительности. Конвертация строки в слайс байт описывается в коде явно:
func createString() {
	s := "hey"
	bs := []byte(s)

	fmt.Println([]byte(s)) // [104 101 121]

	fmt.Println(string(bs)) // hey
}

// Отдельные ASCII символы можно объявлять сразу с типом byte. Для этого нужно обернуть символ в одинарные кавычки и указать тип byte:
func strFromBytes() {
	asciiCh := byte('Z')
	asciiChStr := string(asciiCh)

	fmt.Println(reflect.TypeOf(asciiCh), asciiCh) // uint8 90

	fmt.Println(reflect.TypeOf(asciiChStr), asciiChStr) // string Z
}
