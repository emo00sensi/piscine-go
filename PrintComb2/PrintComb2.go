package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	c := 0
	for i := '0'; i <= '9'; i++ {
		for v := '0'; v <= '9'; v++ {
			for t := '0'; t <= '9'; t++ {
				for  n := '0'; n <= '9'; n++ {
					if v != n && v < n{
						c++
						if c!=1 {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
						z01.PrintRune(i)
						z01.PrintRune(v)
						z01.PrintRune(' ')
						z01.PrintRune(t)
						z01.PrintRune(n)
					}
				}
			}
	}
}
z01.PrintRune('\n')
}
