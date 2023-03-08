package ex33

import "fmt"

func Test() {
	fmt.Println(MaxSum(nil, nil))                          // []
	fmt.Println(MaxSum([]int{1, 2, 3}, []int{3, 4, 5, 6})) // [3 4 5 6]
	fmt.Println(MaxSum([]int{1, 2, 3}, []int{3, 2, 1}))    // [1 2 3]
	fmt.Println(MaxSum([]int{10, 20}, []int{4, 6}))        // [10 20]
}
