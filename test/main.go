package main

import (
	"fmt"

	e ".."
)

func main() {
	var num1 *e.NodeAddL
	// num1 = e.PushFront(num1, 5)
	// num1 = e.PushFront(num1, 1)
	// num1 = e.PushFront(num1, 3)

	// 5 -> 9 -> 2
	var num2 *e.NodeAddL
	// num2 = e.PushFront(num2, 2)
	// num2 = e.PushFront(num2, 9)
	// num2 = e.PushFront(num2, 5)

	// 9 -> 0 -> 7
	result := e.AddLinkedNumbers(num1, num2)
	for tmp := result; tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Num)
		if tmp.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
