package ex15

import "fmt"

func Test() {
	strs := []string{"a", "b", "c"}

	res := Map(strs, func(s string) string {
		return s + "!"
	})

	fmt.Println(res)
}
