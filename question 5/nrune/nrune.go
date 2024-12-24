package piscine

func NRune(s string, n int) rune {
	if len(s) <= 0 {
		return 0
	}
	for i := 0; i <= len(s)-1; i++ {
		if i == n-1 {
			return rune(s[i])
		}
	}
	return 0
}
