package ex17

import "sort"

func UniqueSortedUserIDs(userIDs []int64) []int64 {
	sort.Slice(userIDs, func(i, j int) bool {
		return userIDs[i] < userIDs[j]
	})

	for i := 1; i < len(userIDs); i++ {
		if userIDs[i] == userIDs[i-1] {
			if i == len(userIDs)-1 {
				userIDs = userIDs[:i]
			} else {
				userIDs = append(userIDs[:i], userIDs[i+1:]...)
			}
			i--
		}
	}

	return userIDs
}
