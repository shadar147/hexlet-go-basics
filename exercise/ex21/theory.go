package ex21

import "fmt"

// Так как строка — это массив байт, ее можно обойти с помощью цикла for:
func strFor() {
	s := "hello"
	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
	}
}

// Вывод:
// h
// e
// l
// l
// o

// Таким способом можно обойти только строки, состоящие из ASCII символов. Если строка содержит мультибайтовые символы, вывод будет некорректен:
func main() {
	s := "привет"
	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
	}
}
