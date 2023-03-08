package ex33

import "time"

func MaxSum(nums1, nums2 []int) []int {
	var sum1, sum2 int

	go func(nums []int) {
		sum1 = sum(nums)
	}(nums1)
	go func(nums []int) {
		sum2 = sum(nums)
	}(nums2)

	time.Sleep(100 * time.Millisecond)

	if sum1 < sum2 {
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
