package piscine

func IsUpper(s string) bool {
	result := ""
	for i:=0 ; i<= len(s)-1 ; i++ {
		if s[i]>= 'A' && s[i]<='Z' {
			result+=string(s[i])
		}
	}
	if result == (s) {
		return true
	}
	return false
}