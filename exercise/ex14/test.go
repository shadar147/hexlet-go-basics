package ex14

import "fmt"

func Test() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(Remove(nums, 12))
	fmt.Println(Remove(nums, 0))
}
