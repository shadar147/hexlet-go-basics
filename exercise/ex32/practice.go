package ex32

func MaxSum(nums1, nums2 []int) []int {
	if sum(nums1) < sum(nums2) {
		return nums2
	}

	return nums1
}

func sum(nums []int) int {
	s := 0

	for _, v := range nums {
		s += v
	}

	return s
}