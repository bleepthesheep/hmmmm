package exam

// Union q
func Union(s1, s2 string) string {
	res := ""

	theUltimateString := Join([]string{s1, s2}, "")

	for _, c1 := range theUltimateString {
		if !Contains(res, c1) {
			res += string(c1)
		}
	}

	return res
}

// Join q
func Join(args []string, divider string) string {
	res := ""
	for i, arg := range args {
		if i != 0 {
			res += divider
		}

		res += arg
	}

	return res
}

// Contains q
func Contains(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}

	return false
}
