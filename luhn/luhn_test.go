package luhn

import "testing"

func TestValid(t *testing.T) {
	for _, test := range testCases {
		if ok := Valid(test.input); ok != test.ok {
			t.Fatalf("Valid(%s): %s\n\t Expected: %t\n\t Got: %t", test.input, test.description, test.ok, ok)
		}
	}
}

func BenchmarkValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Valid("2323 2005 7766 3554")
	}
}

// Source: exercism/problem-specifications
// Commit: e0358d3 luhn: Apply new "input" policy (#1054)
// Problem Specifications Version: 1.1.0

var testCases = []struct {
	description string
	input       string
	ok          bool
}{
	{
		"single digit strings can not be valid",
		"1",
		false,
	},
	{
		"a single zero is invalid",
		"0",
		false,
	},
	{
		"a simple valid SIN that remains valid if reversed",
		"059",
		true,
	},
	{
		"a simple valid SIN that becomes invalid if reversed",
		"59",
		true,
	},
	{
		"a valid Canadian SIN",
		"055 444 285",
		true,
	},
	{
		"invalid Canadian SIN",
		"055 444 286",
		false,
	},
	{
		"invalid credit card",
		"8273 1232 7352 0569",
		false,
	},
	{
		"valid strings with a non-digit included become invalid",
		"055a 444 285",
		false,
	},
	{
		"valid strings with punctuation included become invalid",
		"055-444-285",
		false,
	},
	{
		"valid strings with symbols included become invalid",
		"055£ 444$ 285",
		false,
	},
	{
		"single zero with space is invalid",
		" 0",
		false,
	},
	{
		"more than a single zero is valid",
		"0000 0",
		true,
	},
	{
		"input digit 9 is correctly converted to output digit 9",
		"091",
		true,
	},
}
