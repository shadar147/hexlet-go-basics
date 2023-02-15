package ex9

import (
	"fmt"
	"strings"
)

func statusByName(name string) string {
	// функция проверяет, что строка name начинается с подстроки "Mr."
	if strings.HasPrefix(name, "Mr.") {
		return "married man"
	} else if strings.HasPrefix(name, "Mrs.") {
		return "married woman"
	} else {
		return "single person"
	}
}

func main() {
	//if "hi" { // non-bool "hi" (type string) used as if condition
	//}

	n := "Mr. Doe"
	fmt.Println(n + " is a " + statusByName(n)) // Mr. Doe is a married man

	n = "Mrs. Berry"
	fmt.Println(n + " is a " + statusByName(n)) // Mrs. Berry is a married woman

	n = "Karl"
	fmt.Println(n + " is a " + statusByName(n)) // Karl is a single person
}
