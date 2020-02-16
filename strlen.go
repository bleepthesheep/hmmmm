package exam

// StrLen count lenght of a string
func StrLen(str string) int {
	l := 0
	for range str {
		l++
	}
	return l
}
