package exam

// FirstRune return first rune of a string
func FirstRune(s string) rune {
	copy := []rune(s)

	l := StrLen(s)

	if l > 0 {
		return copy[0]
	}

	return 0
}
