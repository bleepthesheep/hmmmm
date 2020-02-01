package exam

import "fmt"

// PrintMemory q
func PrintMemory(arr [10]int) {
	for i, n := range arr {

		s := getHex(n)
		fmt.Print(s, " ")
		if (i+1)%4 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()

	for _, n := range arr {
		if n >= 32 {
			fmt.Print(string(n))
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func getHex(n int) string {
	base := "0123456789abcdef"

	res := ""
	for n != 0 {
		i := n % 16
		res = string(base[i]) + res
		n /= 16
	}

	res = extend(res)

	return res
}

func extend(s string) string {
	for len(s) <= 8 {
		if len(s) == 4 {
			s += " "
		}
		s += "0"
	}

	return s + " "
}
