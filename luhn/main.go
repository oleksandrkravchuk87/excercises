package luhn

import (
	"unicode"
)

func Valid(val string) bool {
	if len(val) < 2 {
		return false
	}
	var sum, c int
L:
	for i := len(val) - 1; i >= 0; i-- {
		switch {
		case !unicode.IsSpace(rune(val[i])) && !unicode.IsNumber(rune(val[i])):
			{
				return false
			}
		case unicode.IsSpace(rune(val[i])):
			{
				continue L
			}
		case c%2 != 0:
			{
				dob := int(val[i]-'0') * 2

				if dob > 9 {
					dob = dob - 9
				}

				sum += dob
				c++
			}
		case c%2 == 0:
			{
				n := int(val[i] - '0')

				sum += n
				c++
			}

		}
	}
	return sum%10 == 0 && c > 1
}
