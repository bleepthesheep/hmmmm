package exam

// LastRune returns last rune of a string
func LastRune(s string) rune {
	copy := []rune(s)

	l := StrLen(s)
	if l > 0 {
		return copy[l-1]
	}

	return 0
}
