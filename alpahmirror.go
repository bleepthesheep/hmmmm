package exam

// AlphaMirror q
func AlphaMirror(s string) string {

	result := ""
	for _, c := range s {
		out := 'a' + ('z' - c)
		result += string(out)
	}

	return result
}
