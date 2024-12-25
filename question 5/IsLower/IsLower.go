package piscine

func IsLower(s string) bool {
	result := ""
	for i:=0 ; i< len(s) ; i++ {
		if s[i]>= 'a' && s[i]<='z' {
			result+=string(s[i])
		}
	}
	if result == (s) {
		return true
	}
	return false

}