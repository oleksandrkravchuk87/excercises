package raindrops

import (
	"bytes"
	"strconv"
)

func Convert(n int) string {
	var buf bytes.Buffer

	for i := 1; i <= n; i++ {
		if n%i != 0 {
			continue
		}
		switch i {
		case 3:
			buf.WriteString("Pling")
		case 5:
			buf.WriteString("Plang")
		case 7:
			buf.WriteString("Plong")

		}
	}
	if buf.Len() == 0 {
		return strconv.Itoa(n)
	}
	return buf.String()
}
