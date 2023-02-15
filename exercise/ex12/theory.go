package ex12

import "fmt"

const StatusOkay int = 200
const defaultStatus = 200

const (
	// публичная константа, которую можно использовать во внешних пакетах
	StatusOk       = 200
	StatusNotFound = 404
	// приватная константа, доступная только в рамках текущего пакета
	statusInternalError = 500
)

const (
	zero = iota
	one
	two
	three
)

type Person struct {
}

func main() {
	// такие константы допустимы
	const (
		num     = 20
		str     = "hey"
		isValid = true
	)

	// нельзя объявить структуру как константу
	//const p = Person{} // ошибка компиляции: const initializer Person{} is not a constant

	const status = 404

	fmt.Println("default status:", defaultStatus) // default status: 200
	fmt.Println("current status:", status)        // current status: 404

	fmt.Println(zero, one, two, three) // 0 1 2 3
}
