package nthprime

import "testing"

func TestNth(t *testing.T) {
	for _, test := range tests {
		switch p, ok := Nth(test.n); {
		case !ok:
			if test.ok {
				t.Fatalf("FAIL %s\nNth(%d) returned !ok.  Expecting ok.", test.description, test.n)
			}
		case !test.ok:
			t.Fatalf("FAIL %s\nNth(%d) = %d, ok = %t.  Expecting !ok.", test.description, test.n, p, ok)
		case p != test.p:
			t.Fatalf("FAIL %s\nNth(%d) = %d, want %d.", test.description, test.n, p, test.p)
		}
		t.Logf("PASS %s", test.description)
	}
}

func BenchmarkNth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Nth(10001)
	}
}

// Source: exercism/problem-specifications
// Commit: 4a3ba76 nth-prime: Apply new "input" policy
// Problem Specifications Version: 2.1.0

var tests = []struct {
	description string
	n           int
	p           int
	ok          bool
}{
	{
		"first prime",
		1,
		2,
		true,
	},
	{
		"second prime",
		2,
		3,
		true,
	},
	{
		"sixth prime",
		6,
		13,
		true,
	},
	{
		"big prime",
		10001,
		104743,
		true,
	},
	{
		"there is no zeroth prime",
		0,
		0,
		false,
	},
}
