package ex28

import "fmt"

// В Go можно объявить алиас на существующий тип данных для выразительности и абстракции. Например, тип byte из модуля "строки" — это алиас uint8. Алиас объявляется через ключевое слово type:
type NumCount int
type errorCode string
type counter int

// Также у алиасов могут быть методы. Объявление метода происходит так же, как и со структурами:
func (c *counter) inc() {
	*c++
}

func main() {
	nc := NumCount(len([]int{1, 2, 3}))

	fmt.Println(nc) // 3

	// Алиас можно конвертировать в оригинальный тип и обратно:
	ec := errorCode("internal")

	fmt.Println(ec) // internal

	fmt.Println(string(ec)) // internal

	c := counter(0)
	(&c).inc() // передается указатель на счетчик &c, так как функция "inc()" работает с указателями
	(&c).inc()

	fmt.Println(c) // 2
}
