package piscine

func ToLower(s string) string {
	result :=""
	for i:=0 ; i<len(s) ;i++ {
		if s[i]>='A'&&s[i]<='Z'{
			result+=string(s[i]+32)
		}else if s[i]>='a'&&s[i]<='z'{
			result+= string(s[i])
		}else {
			result+=string(s[i])
		}

	}
	return result
}