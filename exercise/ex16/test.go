package ex16

import "fmt"

func Test() {
	src := []int{1, 2, 3, 4, 5}

	fmt.Println(IntsCopy(src, 0))
	fmt.Println(IntsCopy(src, 3))
	fmt.Println(IntsCopy(src, 7))
}
