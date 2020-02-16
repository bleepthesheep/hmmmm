package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		z01.PrintRune('\n')
		return
	}

	printStr(union(args[0], args[1]) + "\n")
}

func printStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func union(s1, s2 string) string {
	res := ""

	theUltimateString := s1 + s2

	for _, c1 := range theUltimateString {
		if !contains(res, c1) {
			res += string(c1)
		}
	}

	return res
}

func contains(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}

	return false
}
