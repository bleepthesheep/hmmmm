package main

import (
	"fmt"

	e ".."
)

func main() {
	num1 := &e.NodeAddL{Num: 5}
	num1 = e.PushBack(num1, 1)
	num1 = e.PushBack(num1, 3)
	num1 = e.PushBack(num1, 1)
	num1 = e.PushBack(num1, 3)

	result := num1
	for tmp := result; tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Num)
		if tmp.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()

	result = e.Sortll(num1)
	for tmp := result; tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Num)
		if tmp.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
