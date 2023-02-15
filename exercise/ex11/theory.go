package ex11

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Person struct { // структура публична
	// [название поля] [тип данных]
	Name   string // поле публично
	Age    int
	wallet wallet // поле приватно: можно обращаться только внутри текущего пакета
}

type wallet struct { // структура приватна: можно инициализировать только внутри текущего пакета
	id          string
	moneyAmount float64
}

type User struct {
	ID        int64  `json:"id" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required"`
}

func main() {
	p := Person{Name: "John", Age: 25}

	fmt.Println(p.Name) // "John"
	fmt.Println(p.Age)  // "25"

	p = Person{}

	fmt.Println(p.Name) // ""
	fmt.Println(p.Age)  // "0"

	u := User{}
	u.ID = 22
	u.Email = "test@test.com"
	u.FirstName = "John"

	bs, _ := json.Marshal(u)

	fmt.Println(string(bs)) // {"id":22,"email":"test@test.com","first_name":"John"}

	// создали пустую структуру, чтобы проверить валидацию
	u = User{}

	// создаем валидатор
	v := validator.New()

	// метод Struct валидирует переданную структуру и возвращает ошибку `error`, если какое-то поле некорректно
	fmt.Println(v.Struct(u))
}
