package ex27

import (
	"fmt"
)

// В Go нет классов, но существуют структуры с методами. Метод — это функция с дополнительным аргументом, который указывается в скобках между func и названием функции:
type Dog struct {
	IsBarked bool
}

// сначала объявляется дополнительный аргумент "(d Dog)", а следом идет обычное описание функции
func (d Dog) Bark() {
	fmt.Println("woof!")
}

func bark() {
	d := Dog{}
	d.Bark() // woof!
}

// В примере выше структура Dog передается по значению, то есть копируется. Если изменятся любые свойства внутри метода Bark, они останутся неизменными в исходной структуре:
func (d Dog) BarkAgain() {
	fmt.Println("woof!")
	d.IsBarked = true
}

func barkAgain() {
	d := Dog{}
	d.BarkAgain() // woof!

	fmt.Println(d.IsBarked) // false
}

// Если есть необходимость в изменении состояния, структура должна передаваться указателем:
func (d *Dog) BarkAndChange() {
	fmt.Println("woof!")
	d.IsBarked = true
}

func main() {
	d := &Dog{}
	d.BarkAndChange() // woof!

	fmt.Println(d.IsBarked) // true
}
