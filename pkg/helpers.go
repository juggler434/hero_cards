package helpers

//Min returns the smaller of two values
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//Max returns the higher of two values
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
