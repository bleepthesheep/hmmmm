package exam

// HiddenP q
func HiddenP(s1, s2 string) bool {

	for _, c := range s1 {
		x := indexOf(s2, c)

		if x < 0 {
			return false
		}

		// fmt.Println(s2)

		s2 = s2[x+1:]
	}

	return true
}

func indexOf(s string, r rune) int {

	for i, c := range s {
		if c == r {
			return i
		}
	}

	return -1
}
