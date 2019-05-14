package main

import (
	"testing"
)

func BenchmarkMain(b *testing.B) {

	/*
		_, outw, _ := os.Pipe()
		orgStdout := os.Stdout
		os.Stdout = outw
	*/

	b.ResetTimer()

	b.Run("Slow", func(b *testing.B) {
		benchmark(b, printArgs)
	})

	b.Run("Fast", func(b *testing.B) {
		benchmark(b, printArgs2)
	})

	return
}

func benchmark(b *testing.B, fn func()) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn()
	}
}
