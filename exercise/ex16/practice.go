package ex16

func IntsCopy(src []int, maxLen int) []int {
	if maxLen <= 0 {
		return make([]int, 0)
	}

	if maxLen > len(src) {
		maxLen = len(src)
	}

	res := make([]int, maxLen)

	copy(res, src)

	return res
}
