package main

import (
	"testing"
)

var input uint64 = 0x1234567890ABCDEF
var output int

func BenchmarkPopCount(b *testing.B) {
	var c int
	for i := 0; i < b.N; i++ {
		c = PopCount(input)
	}
	output = c
}

func BenchmarkPopCountByClear(b *testing.B) {
	var c int
	for i := 0; i < b.N; i++ {
		PopCountByClear(input)
	}
	output = c
}
