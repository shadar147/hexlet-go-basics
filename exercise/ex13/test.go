package ex13

import "fmt"

func Test() {
	nums := [5]int{1, 2, 3, 4, 5}

	fmt.Println(SafeWrite(nums, 0, 10))
	fmt.Println(SafeWrite(nums, 6, 10))
}
