package exam

import "fmt"

// TNode q
type TNode struct {
	Val   int
	Left  *TNode
	Right *TNode
}

// InsertTNode q
func InsertTNode(root *TNode, data int) *TNode {
	newNode := &TNode{Val: data}

	if root == nil {
		return newNode
	}

	if data < root.Val {
		root.Left = InsertTNode(root.Left, data)
	} else {
		root.Right = InsertTNode(root.Right, data)
	}

	return root
}

// PrintTree q
func PrintTree(root *TNode) {
	if root != nil {
		PrintTree(root.Left)
		fmt.Println(root.Val)
		PrintTree(root.Right)
	}
}

// InvertTree q
func InvertTree(root *TNode) *TNode {

	if root == nil {
		return root
	}

	root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)

	return root
}
