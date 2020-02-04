package exam

// AlphaMirror q
func AlphaMirror(s string) string {

	result := ""
	for _, c := range s {
		out := c
		if c >= 'a' && c <= 'z' {
			out = 'a' + 'z' - c
		} else if c >= 'A' && c <= 'Z' {
			out = 'A' + 'Z' - c
		}

		result += string(out)
	}

	return result
}
