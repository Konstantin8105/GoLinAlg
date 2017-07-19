package matrix

import (
	"runtime"
	"sync"
)

const (
	sizeof  = 8.
	memory  = 3000000. / sizeof
	floatL1 = memory / 200
)

// matrix a with size [m,n]
// matrix b with size [n,h]
// matrix c with size [m,h]
func timesAlgorithm(a, b, c *[][]float64, m, n, h int) {
	sizeSummaryFloats := m*n + n*h + m*h
	if sizeSummaryFloats < floatL1 {
		var sum float64
		for j := 0; j < h; j++ {
			for i := 0; i < m; i++ {
				sum = 0
				for k := 0; k < n; k++ {
					sum += (*a)[i][k] * (*b)[k][j]
				}
				(*c)[i][j] = sum
			}
		}
		return
	}
	// Found amount allowable parallelism
	threads := runtime.GOMAXPROCS(0)
	if threads > runtime.NumCPU() {
		threads = runtime.NumCPU()
	}
	memPerTh := memory / float64(threads)

	// parallel algo
	alpha := (memPerTh - float64(n)) / (float64(n) + 1.)
	if alpha > float64(m) {
		alpha = float64(m)
	}
	betta := (memPerTh - 2.*alpha) / (float64(n) + 2.*alpha)
	if betta > float64(h) {
		betta = float64(h)
	}
	// amount rows of [A]
	//iAlpha := int(alpha)
	// amount columns of [B]
	//iBetta := int(betta)

	// Create workgroup
	if alpha > betta {

		var wg sync.WaitGroup
		// Run calculation in goroutines
		step := int(float64(h) / float64(threads))
		from := 0
		var to int
		for t := 0; t < threads; t++ {
			to = from + step
			if to > h {
				to = h
			}
			// Add one goroutine in workgroup
			wg.Add(1)
			go func(from, to int) {
				// Change waitgroup after work done
				defer wg.Done()
				var sum float64
				buf := make([]float64, n, n)
				for j := from; j < to; j++ {
					for k := 0; k < n; k++ {
						buf[k] = (*b)[k][j]
					}
					for i := 0; i < m; i++ {
						sum = 0
						for k := 0; k < n; k++ {
							sum += (*a)[i][k] * buf[k]
						}
						(*c)[i][j] = sum
					}
				}
			}(from, to)
			from = to
		}
		wg.Wait()
		return
	}

	// add strategy for many buffers
	var sum float64
	for j := 0; j < h; j++ {
		for i := 0; i < m; i++ {
			sum = 0
			for k := 0; k < n; k++ {
				sum += (*a)[i][k] * (*b)[k][j]
			}
			(*c)[i][j] = sum
		}
	}
}
