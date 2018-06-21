package twofer

import "fmt"
import "testing"

// ExampleShareWith() is an Example function. Examples are testable snippets of
// Go code that are used for documenting and verifying the package API.
// They may be present in some exercises to demonstrate the expected use of the
// exercise API and can be run as part of a package's test suite.
//
// When an Example test is run the data that is written to standard output is
// compared to the data that comes after the "Output: " comment.
//
// Below the result of ShareWith() is passed to standard output
// using fmt.Println, and this is compared against the expected output.
// If they are equal, the test passes.
func ExampleShareWith() {
	h := ShareWith("")
	fmt.Println(h)
	// Output: One for you, one for me.
}

// Define a function ShareWith(string) string.

func TestShareWith(t *testing.T) {
	tests := []struct {
		name, expected string
	}{
		{"", "One for you, one for me."},
		{"Alice", "One for Alice, one for me."},
		{"Bob", "One for Bob, one for me."},
	}
	for _, test := range tests {
		if observed := ShareWith(test.name); observed != test.expected {
			t.Fatalf("ShareWith(%s) = %v, want %v", test.name, observed, test.expected)
		}
	}
}
