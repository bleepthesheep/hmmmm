package main

import (
	e ".."
)

func main() {
	var root *e.TNode
	root = e.InsertTNode(root, 5)
	root = e.InsertTNode(root, 3)
	root = e.InsertTNode(root, 7)
	root = e.InsertTNode(root, 1)
	root = e.InsertTNode(root, 4)
	root = e.InsertTNode(root, 6)
	root = e.InsertTNode(root, 8)
	e.Print2DTree(root, 0)
}
