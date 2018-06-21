package grains

import (
	"testing"
)

func TestSquare(t *testing.T) {
	for _, test := range squareTests {
		actualVal, actualErr := Square(test.input)

		// check actualVal only if no error expected
		if !test.expectError && actualVal != test.expectedVal {
			t.Fatalf("FAIL: %s\nSquare(%d) expected %d, Actual %d", test.description, test.input, test.expectedVal, actualVal)
		}

		// if we expect an error and there isn't one
		if test.expectError && actualErr == nil {
			t.Fatalf("FAIL: %s\nSquare(%d) expected an error, but error is nil", test.description, test.input)
		}
		// if we don't expect an error and there is one
		if !test.expectError && actualErr != nil {
			var _ error = actualErr
			t.Fatalf("FAIL: %s\nSquare(%d) expected no error, but error is: %s", test.description, test.input, actualErr)
		}
		t.Logf("PASS: %s", test.description)
	}
}

func TestTotal(t *testing.T) {
	var expected uint64 = 18446744073709551615
	if actual := Total(); actual != expected {
		t.Errorf("Total() expected %d, Actual %d", expected, actual)
	}
}

func BenchmarkSquare(b *testing.B) {

	for i := 0; i < b.N; i++ {

		for _, test := range squareTests {
			Square(test.input)
		}

	}
}

func BenchmarkTotal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Total()
	}
}

// Source: exercism/problem-specifications
// Commit: f079c2d grains: Move input (square) to input object (#1191)
// Problem Specifications Version: 1.1.0

// returns the number of grains on the square
var squareTests = []struct {
	description string
	input       int
	expectedVal uint64
	expectError bool
}{
	{
		description: "1",
		input:       1,
		expectedVal: 1,
	},
	{
		description: "2",
		input:       2,
		expectedVal: 2,
	},
	{
		description: "3",
		input:       3,
		expectedVal: 4,
	},
	{
		description: "4",
		input:       4,
		expectedVal: 8,
	},
	{
		description: "16",
		input:       16,
		expectedVal: 32768,
	},
	{
		description: "32",
		input:       32,
		expectedVal: 2147483648,
	},
	{
		description: "64",
		input:       64,
		expectedVal: 9223372036854775808,
	},
	{
		description: "square 0 returns an error",
		input:       0,
		expectError: true,
	},
	{
		description: "negative square returns an error",
		input:       -1,
		expectError: true,
	},
	{
		description: "square greater than 64 returns an error",
		input:       65,
		expectError: true,
	},
}
