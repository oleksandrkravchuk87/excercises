package hamming

import "errors"

func Distance(s1, s2 string) (int, error) {
	if len(s1) != len(s2) {
		return -1, errors.New("error")
	}

	var c int
	for k := range s1 {
		if s1[k] != s2[k] {
			c++
		}
	}
	return c, nil
}
