package ex24

import "fmt"

// В предыдущих уроках мы использовали пакет fmt для вывода переменных или результатов функций:
func sayHello() {
	s := "hello world"

	// печатает вывод на следующей строке
	fmt.Println(s) // "hello world"
}

// Пакет fmt так же используется для форматирования строк. Плейсхолдеры разных типов данных в основном не отличаются от других языков:
func strFormat() {
	name := "Andy"

	// подставляем строку
	fmt.Sprintf("hello %s", name) // "hello Andy"

	// число
	fmt.Sprintf("there are %d kittens", 10) // "there are 10 kittens"

	// логический тип
	fmt.Sprintf("your story is %t", true) // "your story is true"
}

// Так же существуют специализированные плейсхолдеры, которые преобразуют сложные структуры:
type Person struct {
	Name string
	Age  int
}

func structFormat() {
	p := Person{Name: "Andy", Age: 18}

	// вывод значений структуры
	fmt.Println("simple struct:", p)

	// вывод названий полей и их значений
	fmt.Printf("detailed struct: %+v\n", p)

	// вывод названий полей и их значений в виде инициализации
	fmt.Printf("Golang struct: %#v\n", p)
}

// Вывод:
// simple struct: {Andy 18}
// detailed struct: {Name:Andy Age:18}
// Golang struct: main.Person{Name:"Andy", Age:18}
