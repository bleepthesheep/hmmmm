package exam

// Lcm q
func Lcm(first, second int) int {

	return first * second / Mcd(first, second)
}

// Mcd i
func Mcd(first, second int) int {

	least := first
	if second < first {
		least = second
	}

	for i := 2; i <= least; i++ {
		if second%i == 0 && first%i == 0 {
			return i * Mcd(first/i, second/i)
		}
	}

	return 1
}
