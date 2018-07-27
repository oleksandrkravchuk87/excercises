package functions

import "testing"

func BenchmarkTpointer(b *testing.B) {

	b.ReportAllocs()

	teststr := some{a: "astring", b: "bstring", c: 1000}

	for i := 0; i < b.N; i++ {
		tpointer(&teststr)
	}
}

func BenchmarkTvalue(b *testing.B) {
	b.ReportAllocs()

	teststr := some{a: "astring", b: "bstring", c: 1000}

	for i := 0; i < b.N; i++ {
		tvalue(&teststr)
	}
}
