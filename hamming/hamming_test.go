package hamming

import "testing"

func TestHamming(t *testing.T) {
	for _, tc := range testCases {
		got, err := Distance(tc.s1, tc.s2)
		if tc.want < 0 {
			// check if err is of error type
			var _ error = err

			// we expect error
			if err == nil {
				t.Fatalf("Distance(%q, %q). error is nil.",
					tc.s1, tc.s2)
			}
		} else {
			if got != tc.want {
				t.Fatalf("Distance(%q, %q) = %d, want %d.",
					tc.s1, tc.s2, got, tc.want)
			}

			// we do not expect error
			if err != nil {
				t.Fatalf("Distance(%q, %q) returned error: %v when expecting none.",
					tc.s1, tc.s2, err)
			}
		}
	}
}

func BenchmarkHamming(b *testing.B) {
	// bench combined time to run through all test cases
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Distance(tc.s1, tc.s2)
		}
	}
}

var testCases = []struct {
	s1   string
	s2   string
	want int
}{
	{ // empty strands
		"",
		"",
		0,
	},
	{ // identical strands
		"A",
		"A",
		0,
	},
	{ // long identical strands
		"GGACTGA",
		"GGACTGA",
		0,
	},
	{ // complete distance in single nucleotide strands
		"A",
		"G",
		1,
	},
	{ // complete distance in small strands
		"AG",
		"CT",
		2,
	},
	{ // small distance in small strands
		"AT",
		"CT",
		1,
	},
	{ // small distance
		"GGACG",
		"GGTCG",
		1,
	},
	{ // small distance in long strands
		"ACCAGGG",
		"ACTATGG",
		2,
	},
	{ // non-unique character in first strand
		"AAG",
		"AAA",
		1,
	},
	{ // non-unique character in second strand
		"AAA",
		"AAG",
		1,
	},
	{ // same nucleotides in different positions
		"TAG",
		"GAT",
		2,
	},
	{ // large distance
		"GATACA",
		"GCATAA",
		4,
	},
	{ // large distance in off-by-one strand
		"GGACGGATTCTG",
		"AGGACGGATTCT",
		9,
	},
	{ // disallow first strand longer
		"AATG",
		"AAA",
		-1,
	},
	{ // disallow second strand longer
		"ATA",
		"AGTG",
		-1,
	},
}
