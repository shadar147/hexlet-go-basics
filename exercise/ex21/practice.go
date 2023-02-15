package ex21

func ShiftASCII(s string, step int) string {
	res := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		res[i] = s[i] + byte(step)
	}

	return string(res)
}
