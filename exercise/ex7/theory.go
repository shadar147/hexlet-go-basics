package ex7

func main() {
	//true && false // false
	//false || true // true

	//var b bool = true

	// сокращенная запись
	//bs := false

	//true == false // false

	//false == false // true

	//true == "hello" // invalid operation: false == "hello" (mismatched types untyped bool and untyped string)

	//flag := true
	//text := "hello"

	// вариант не сработает, потому что нельзя конвертировать строку в bool
	//flag && bool(text) // cannot convert text (type string) to type bool

	// правильный вариант: если строка не пустая, то в ней есть текст
	//flag && text != "" // true
}
