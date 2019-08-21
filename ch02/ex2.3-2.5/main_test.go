package main

import (
	"testing"

	"github.com/tesujiro/TheGoProgrammingLanguage/ch02/ex2.3-2.5/popcount"
)

func BenchmarkMain(b *testing.B) {

	b.ResetTimer()

	b.Run("OneStep", func(b *testing.B) {
		benchmark(b, popcount.PopCount)
	})

	b.Run("Loop", func(b *testing.B) {
		benchmark(b, popcount.PopCount2)
	})

	b.Run("OneBitLoop", func(b *testing.B) {
		benchmark(b, popcount.PopCount3)
	})

	b.Run("X&(X-1) Loop", func(b *testing.B) {
		benchmark(b, popcount.PopCount4)
	})

	return
}

func benchmark(b *testing.B, fn func(uint64) int) {
	b.ResetTimer()
	for i := 1; i < b.N; i++ {
		fn(uint64(i))
	}
}
