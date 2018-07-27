package main

import (
	"reflect"
	"testing"
)

// In the separate file frequency.go, you are given a function, Frequency(),
// to sequentially count letter frequencies in a single text.
//
// Perform this exercise on parallelism using Go concurrency features.
// Make concurrent calls to Frequency and combine results to obtain the answer.

func TestConcurrentFrequency(t *testing.T) {
	seq := OriginalFrequency(euro + dutch + us)
	con := ConcurrentFrequency([]string{euro, dutch, us}...)
	//	fmt.Println(seq)
	//	fmt.Println(con)
	if !reflect.DeepEqual(con, seq) {
		t.Fatal("ConcurrentFrequency wrong result")
	}

}

func TestSequentialFrequency(t *testing.T) {
	oSeq := OriginalFrequency(euro + dutch + us)
	seq := OriginalFrequency(euro + dutch + us)
	if !reflect.DeepEqual(oSeq, seq) {
		t.Fatal("Frequency wrong result")
	}
}

func BenchmarkSequentialFrequency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OriginalFrequency(euro + dutch + us)
	}
}

func BenchmarkConcurrentFrequency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcurrentFrequency([]string{euro, dutch, us}...)
	}
}
