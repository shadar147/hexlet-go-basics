package ex14

func Remove(nums []int, i int) []int {
	if (i < 0) || (i > len(nums)-1) {
		return nums
	}

	return append(nums[:i], nums[i+1:]...)
}
