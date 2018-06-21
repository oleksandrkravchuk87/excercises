package main

import (
	"math"
	"testing"
)

var testCases = []struct {
	description string
	planet      Planet
	seconds     float64
	expected    float64
}{
	{
		description: "age on Earth",
		planet:      "Earth",
		seconds:     1000000000,
		expected:    31.69,
	},
	{
		description: "age on Mercury",
		planet:      "Mercury",
		seconds:     2134835688,
		expected:    280.88,
	},
	{
		description: "age on Venus",
		planet:      "Venus",
		seconds:     189839836,
		expected:    9.78,
	},
	{
		description: "age on Mars",
		planet:      "Mars",
		seconds:     2329871239,
		expected:    39.25,
	},
	{
		description: "age on Jupiter",
		planet:      "Jupiter",
		seconds:     901876382,
		expected:    2.41,
	},
	{
		description: "age on Saturn",
		planet:      "Saturn",
		seconds:     3000000000,
		expected:    3.23,
	},
	{
		description: "age on Uranus",
		planet:      "Uranus",
		seconds:     3210123456,
		expected:    1.21,
	},
	{
		description: "age on Neptune",
		planet:      "Neptune",
		seconds:     8210123456,
		expected:    1.58,
	},
}

func TestAge(t *testing.T) {
	const precision = 0.2
	for _, tc := range testCases {
		actual := Age(tc.seconds, tc.planet)
		if math.IsNaN(actual) || math.Abs(actual-tc.expected) > precision {
			t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", tc.description, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkAge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Age(tc.seconds, tc.planet)
		}
	}
}
