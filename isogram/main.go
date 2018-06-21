package isogram

import "unicode"

func IsIsogram(text string) bool {
	if text == "" {
		return true
	}
	runes := []rune(text)
	for i := 0; i < len(runes)-1; i++ {
		for j := i + 1; j <= len(runes)-1; j++ {
			if unicode.ToLower(runes[i]) == unicode.ToLower(runes[j]) && unicode.IsLetter(runes[i]) {
				return false
			}
		}
	}
	return true
}
