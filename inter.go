package exam

// Inter q
func Inter(s1, s2 string) string {
	res := ""

	for _, c := range s1 {
		if !Contains(res, c) && Contains(s2, c) {
			res += string(c)
		}
	}

	for _, c := range s2 {
		if !Contains(res, c) && Contains(s1, c) {
			res += string(c)
		}
	}

	return res
}
