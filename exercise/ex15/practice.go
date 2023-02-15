package ex15

func Map(strs []string, mapFunc func(s string) string) []string {
	res := make([]string, len(strs))

	for i, str := range strs {
		res[i] = mapFunc(str)
	}

	return res
}
