package exam

// Sortll q
func Sortll(node *NodeAddL) *NodeAddL {
	copy := node

	for x := copy; x.Next != nil; x = x.Next {
		for it := copy; it.Next != nil; it = it.Next {
			if it.Num < it.Next.Num {
				it.Num, it.Next.Num = it.Next.Num, it.Num
			}
		}
	}

	return copy
}
