package ex18

func UniqueUserIDs(userIDs []int64) []int64 {
	userIDsMap := make(map[int64]bool)

	for _, userID := range userIDs {
		userIDsMap[userID] = true
	}

	res := make([]int64, 0, len(userIDsMap))
	for userID := range userIDsMap {
		res = append(res, userID)
	}

	return res
}
