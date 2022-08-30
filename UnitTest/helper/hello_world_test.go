package helper

import "testing"

// Table Benchmark
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Anggi",
			request: "Anggi",
		},
		{
			name:    "Eko",
			request: "Eko",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

// Sub Benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("Anggi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Anggi")
		}
	})
}

// Benchmark
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Anggi")
	}
}
