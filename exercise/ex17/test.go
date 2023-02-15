package ex17

import "fmt"

func Test() {
	var userIDs []int64
	fmt.Println(UniqueSortedUserIDs(userIDs))

	userIDs = []int64{10}
	fmt.Println(UniqueSortedUserIDs(userIDs))

	userIDs = []int64{55, 55}
	fmt.Println(UniqueSortedUserIDs(userIDs))

	userIDs = []int64{55, 55, 33, 22}
	fmt.Println(UniqueSortedUserIDs(userIDs))

	userIDs = []int64{55, 2, 88, 33, 2, 2, 55, 103, 33, 88}
	fmt.Println(UniqueSortedUserIDs(userIDs))
}
