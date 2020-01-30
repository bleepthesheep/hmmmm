package exam

// Reverse q
func Reverse(node *NodeAddL) *NodeAddL {
	var copy *NodeAddL
	for it := node; it != nil; it = it.Next {
		copy = PushFront(copy, it.Num)
	}
	return node
}

// PushFront q
func PushFront(node *NodeAddL, data int) *NodeAddL {
	return &NodeAddL{Num: data, Next: node}
}
