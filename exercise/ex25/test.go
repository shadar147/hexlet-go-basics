package ex25

import (
	"fmt"
)

func Test() {
	fmt.Println(MergeNumberLists([]int{1, 2, 3}, []int{4}, []int{5, 6, 7, 8}, []int{9, 10}))
	fmt.Println(MergeNumberLists([]int{}))
	fmt.Println(MergeNumberLists(nil, nil))
	fmt.Println(MergeNumberLists([]int{10, 20}, nil, []int{50, 60}))
}
