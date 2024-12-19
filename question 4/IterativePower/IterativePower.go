package piscine

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if power == 1 {
		return nb
	}
	result := 1
	for i := power ; i > 0 ; i--{
			result *=nb 
	}
	return result
}