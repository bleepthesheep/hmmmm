package main

import (
	"fmt"

	e ".."
)

func main() {
	var num1 *e.NodeAddL
	num1 = e.PushFront(num1, 1)
	num1 = e.PushFront(num1, 2)
	num1 = e.PushFront(num1, 3)
	num1 = e.PushFront(num1, 4)
	num1 = e.PushFront(num1, 5)

	result := e.Reverse(num1)
	for tmp := result; tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Num)
		if tmp.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
