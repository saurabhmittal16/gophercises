package hackerrank

func camelcase(s string) int32 {
	var count int32
	count = 0
	for i := range s {
		if s[i] >= 65 && s[i] <= 90 {
			count++
		}
	}

	return count + 1
}
