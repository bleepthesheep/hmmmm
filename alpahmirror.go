package exam

import "fmt"

func AlphaMirror(s string) {

	result := ""
	for _, c := range s {
		out := 'a' + ('a' + 25 - c)
		result += string(out)
	}

	fmt.Println(result)
}
