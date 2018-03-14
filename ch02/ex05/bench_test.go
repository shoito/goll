package main

import (
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClear(0x1234567890ABCDEF)
	}
}
