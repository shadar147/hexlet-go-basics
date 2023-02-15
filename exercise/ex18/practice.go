package ex18

func UniqueUserIDs(userIDs []int64) []int64 {
	uniqueUserIDs := make(map[int64]bool)

	for _, userID := range userIDs {
		uniqueUserIDs[userID] = true
	}

	res := make([]int64, 0, len(uniqueUserIDs))
	for _, userID := range userIDs {
		if uniqueUserIDs[userID] {
			res = append(res, userID)
			uniqueUserIDs[userID] = false
		}
	}

	return res
}
