package functions

import "fmt"

type some struct {
	a string
	b string
	c int
}

func main() {

}

func tvalue(s *some) {
	wrap(s.a, s.b, s.c)
}

func tpointer(s *some) {
	wrap(&s.a, &s.b, &s.c)
}

func wrap(val ...interface{}) {
	fmt.Sprint(val)
	return
}
