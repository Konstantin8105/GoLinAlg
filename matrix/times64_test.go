package matrix

import (
	"fmt"
	"testing"
)

func BenchmarkTimes64(b *testing.B) {
	benchmarks := []struct {
		m, n, h int
	}{
		{1, 1, 1},

		{5, 1, 1},
		{1, 5, 1},
		{1, 1, 5},
		{5, 5, 1},
		{1, 5, 5},
		{5, 1, 5},
		{5, 5, 5},

		{1000, 1, 1},
		{1, 1000, 1},
		{1, 1, 1000},
		{1000, 1000, 1},
		{1, 1000, 1000},
		{1000, 1, 1000},
		{1000, 1000, 1000},

		{2000, 1, 1},
		{1, 2000, 1},
		{1, 1, 2000},
		{2000, 2000, 1},
		{1, 2000, 2000},
		{2000, 1, 2000},
		{2000, 2000, 2000},
	}
	for _, bm := range benchmarks {
		b.Run(fmt.Sprintf("%5v%5v%5v", bm.m, bm.n, bm.h), func(b *testing.B) {
			b.StopTimer()

			A := make([][]float64, bm.m, bm.m)
			for r := 0; r < bm.m; r++ {
				A[r] = make([]float64, bm.n, bm.n)
			}

			B := make([][]float64, bm.n, bm.n)
			for r := 0; r < bm.n; r++ {
				B[r] = make([]float64, bm.h, bm.h)
			}

			C := make([][]float64, bm.m, bm.m)
			for r := 0; r < bm.m; r++ {
				C[r] = make([]float64, bm.h, bm.h)
			}

			b.StartTimer()
			for i := 0; i < b.N; i++ {
				timesAlgorithm(&A, &B, &C, bm.m, bm.n, bm.h)
			}
		})
	}
}

/*
func TestTimes(t *testing.T) {
	benchmarks := []struct {
		m, n, h int
	}{
		{1, 1, 1},

		{5, 1, 1},
		{1, 5, 1},
		{1, 1, 5},
		{5, 5, 1},
		{1, 5, 5},
		{5, 1, 5},
		{5, 5, 5},

		{1000, 3, 3},
		{3, 1000, 3},
		{3, 3, 1000},
		{1000, 1000, 3},
		{3, 1000, 1000},
		{1000, 3, 1000},
		{1000, 1000, 1000},

		{2000, 3, 3},
		{3, 2000, 3},
		{3, 3, 2000},
		{2000, 2000, 3},
		{3, 2000, 2000},
		{2000, 3, 2000},
		{2000, 2000, 2000},
	}
	for _, bm := range benchmarks {
		t.Run(fmt.Sprintf("%5v%5v%5v", bm.m, bm.n, bm.h), func(t *testing.T) {
			A := make([][]float64, bm.m, bm.m)
			for r := 0; r < bm.m; r++ {
				A[r] = make([]float64, bm.n, bm.n)
			}

			B := make([][]float64, bm.n, bm.n)
			for r := 0; r < bm.n; r++ {
				B[r] = make([]float64, bm.h, bm.h)
			}

			C := make([][]float64, bm.m, bm.m)
			for r := 0; r < bm.m; r++ {
				C[r] = make([]float64, bm.h, bm.h)
			}
			timesAlgorithm(&A, &B, &C, bm.m, bm.n, bm.h)
		})
	}
}
*/
