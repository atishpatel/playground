package main

import "testing"

func BenchmarkInPlace(b *testing.B) {
	v := make([]int, 10000)
	for i := 0; i < b.N; i++ {
		shiftInPlace(v, 100)
	}
}

func BenchmarkAlloc(b *testing.B) {
	v := make([]int, 10000)
	for i := 0; i < b.N; i++ {
		shiftNewAlloc(v, 100)
	}
}

func BenchmarkCopyArray(b *testing.B) {
	v := make([]int, 10000)
	for i := 0; i < b.N; i++ {
		shiftCopyArray(v, 100)
	}
}
