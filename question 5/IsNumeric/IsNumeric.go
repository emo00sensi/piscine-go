package piscine

func IsNumeric(s string) bool {
	result := ""
	for i:=0 ; i< len(s) ; i++ {
		if s[i]>= '0' && s[i]<='9' {
			result+=string(s[i])
		}
	}
	if result == (s) {
		return true
	}
	return false
}