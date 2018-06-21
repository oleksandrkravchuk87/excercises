package isogram

import "testing"

func TestIsIsogram(t *testing.T) {
	for _, c := range testCases {
		if IsIsogram(c.input) != c.expected {
			t.Fatalf("FAIL: %s\nWord %q, expected %t, got %t", c.description, c.input, c.expected, !c.expected)
		}

		t.Logf("PASS: Word %q", c.input)
	}
}

func BenchmarkIsIsogram(b *testing.B) {
	for i := 0; i < b.N; i++ {

		for _, c := range testCases {
			IsIsogram(c.input)
		}

	}
}

var testCases = []struct {
	description string
	input       string
	expected    bool
}{
	{
		description: "empty string",
		input:       "",
		expected:    true,
	},
	{
		description: "isogram with only lower case characters",
		input:       "isogram",
		expected:    true,
	},
	{
		description: "word with one duplicated character",
		input:       "eleven",
		expected:    false,
	},
	{
		description: "longest reported english isogram",
		input:       "subdermatoglyphic",
		expected:    true,
	},
	{
		description: "word with duplicated character in mixed case",
		input:       "Alphabet",
		expected:    false,
	},
	{
		description: "hypothetical isogrammic word with hyphen",
		input:       "thumbscrew-japingly",
		expected:    true,
	},
	{
		description: "isogram with duplicated hyphen",
		input:       "six-year-old",
		expected:    true,
	},
	{
		description: "made-up name that is an isogram",
		input:       "Emily Jung Schwartzkopf",
		expected:    true,
	},
	{
		description: "duplicated character in the middle",
		input:       "accentor",
		expected:    false,
	},
}
