package hackerrank

func caesarCipher(s string, k int32) string {
	k = k % 26
	res := []byte{}
	for i := range s {
		if s[i] >= 65 && s[i] <= 90 {
			if s[i]+byte(k) > 90 {
				res = append(res, s[i]+byte(k)-26)
			} else {
				res = append(res, s[i]+byte(k))
			}
		} else if s[i] >= 97 && s[i] <= 122 {
			if s[i]+byte(k) > 122 {
				res = append(res, s[i]+byte(k)-26)
			} else {
				res = append(res, s[i]+byte(k))
			}
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}
