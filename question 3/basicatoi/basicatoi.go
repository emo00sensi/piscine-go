package piscine

func BasicAtoi(s string) int {
	result := 0
	multiplier := 1

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			result += int(s[i]-'0')  *multiplier
			multiplier *= 10
		}
	}
	return result
}