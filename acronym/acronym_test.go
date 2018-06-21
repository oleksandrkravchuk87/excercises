package acronym

import (
	"testing"
)

func TestAcronym(t *testing.T) {
	for _, test := range stringTestCases {
		actual := Abbreviate(test.input)
		if actual != test.expected {
			t.Errorf("Acronym test [%s], expected [%s], actual [%s]", test.input, test.expected, actual)
		}
	}
}

func BenchmarkAcronym(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range stringTestCases {
			Abbreviate(test.input)
		}
	}
}

// Source: exercism/problem-specifications
// Commit: 5ae1dba Acronym canonical-data: Remove redundant test case
// Problem Specifications Version: 1.3.0

type acronymTest struct {
	input    string
	expected string
}

var stringTestCases = []acronymTest{
	{
		input:    "Portable Network Graphics",
		expected: "PNG",
	},
	{
		input:    "Ruby on Rails",
		expected: "ROR",
	},
	{
		input:    "First In, First Out",
		expected: "FIFO",
	},
	{
		input:    "GNU Image Manipulation Program",
		expected: "GIMP",
	},
	{
		input:    "Complementary metal-oxide semiconductor",
		expected: "CMOS",
	},
}
