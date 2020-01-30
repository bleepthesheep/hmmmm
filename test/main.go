package main

import (
	"fmt"

	e ".."
)

func main() {
	case1 := []int{1, 2, 3, 4}
	out := e.TwoSum(case1, 5)
	fmt.Println(out)
}
