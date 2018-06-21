package acronym

import (
	"strings"

	"fmt"
	"unicode"
)

func Abbreviate(text string) string {
	sl := strings.FieldsFunc(text, Split)
	var abr []rune
	for _, v := range sl {
		abr = append(abr, unicode.ToUpper(rune(v[0])))
	}
	fmt.Println(string(abr))
	return string(abr)
}

func Split(r rune) bool {
	return r == ' ' || r == '-'
}
