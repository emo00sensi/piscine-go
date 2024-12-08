package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	c := 0
	for i := '0'; i <= '7'; i++ {
		for s := i + 1; s <= '8'; s++ {
			for t := s + 1; t <= '9'; t++ {
				 c++
				if i < s && s < t{
					if c!=1 {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
					z01.PrintRune(i)
					z01.PrintRune(s)
					z01.PrintRune(t)
				}
			}
	}
}
z01.PrintRune('\n')
}