package exam

// RobotToOrigin q
func RobotToOrigin(s string) bool {
	x, y := 0, 0

	for _, c := range s {
		switch c {
		case 'U':
			y++
		case 'D':
			y--
		case 'R':
			x++
		case 'L':
			x--

		// error
		default:
			return false
		}
	}

	if x == 0 && y == 0 {
		return true
	}

	return false
}
