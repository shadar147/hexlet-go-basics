package ex25

func MergeNumberLists(numberLists ...[]int) []int {
	result := []int{}
	for _, numberList := range numberLists {
		result = append(result, numberList...)
	}
	return result
}
