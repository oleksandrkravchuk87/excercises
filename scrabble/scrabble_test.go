package scrabble

import "testing"

func TestScore(t *testing.T) {
	for _, test := range scrabbleScoreTests {
		if actual := Score(test.input); actual != test.expected {
			t.Errorf("Score(%q) expected %d, Actual %d", test.input, test.expected, actual)
		}
	}
}

func BenchmarkScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range scrabbleScoreTests {
			Score(test.input)
		}
	}
}

type scrabbleTest struct {
	input    string
	expected int
}

var scrabbleScoreTests = []scrabbleTest{
	{"a", 1},                // lowercase letter
	{"A", 1},                // uppercase letter
	{"f", 4},                // valuable letter
	{"at", 2},               // short word
	{"zoo", 12},             // short, valuable word
	{"street", 6},           // medium word
	{"quirky", 22},          // medium, valuable word
	{"OxyphenButazone", 41}, // long, mixed-case word
	{"pinata", 8},           // english-like word
	{"", 0},                 // empty input
	{"abcdefghijklmnopqrstuvwxyz", 87}, // entire alphabet available
}
