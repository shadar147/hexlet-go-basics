package ex11

import "strings"

type UserCreateRequest struct {
	FirstName string // не может быть пустым; не может содержать пробелы
	Age       int    // не может быть 0 или отрицательным; не может быть больше 150
}

func Validate(req UserCreateRequest) string {
	if isValidName(req.FirstName) && isValidAge(req.Age) {
		return ""
	}

	return "invalid request"
}

func isValidName(name string) bool {
	return !(name == "" || strings.Contains(name, " "))
}

func isValidAge(age int) bool {
	return !(age == 0 || age < 0 || age > 150)
}
