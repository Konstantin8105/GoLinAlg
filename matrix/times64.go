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

	//if m > n && h > n {
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
	//	}
	/*

		// add strategy for many buffers
		var wg sync.WaitGroup
		// Run calculation in goroutines
		step := int(float64(h) / float64(threads))
		from := 0
		var to int
		for t := 0; t < threads; t++ {
			if from >= h {
				continue
			}
			to = from + step
			if to > h {
				to = h
			}
			// Add one goroutine in workgroup
			wg.Add(1)
			go func(from, to int) {
				// Change waitgroup after work done
				defer wg.Done()
				var sum0 float64
				var sum1 float64
				buf0 := make([]float64, n, n)
				buf1 := make([]float64, n, n)
				for j := from; j < to; j += 2 {
					for k := 0; k < n; k++ {
						buf0[k] = (*b)[k][j+0]
						buf1[k] = (*b)[k][j+1]
					}
					for i := 0; i < m; i++ {
						sum0 = 0
						sum1 = 0
						for k := 0; k < n; k++ {
							sum0 += (*a)[i][k] * buf0[k]
							sum1 += (*a)[i][k] * buf1[k]
						}
						(*c)[i][j+0] = sum0
						(*c)[i][j+1] = sum1
					}
				}
				if (to-from)/2 != int(float64(to-from)/2.) {
					j := to - 1
					for k := 0; k < n; k++ {
						buf0[k] = (*b)[k][j+0]
					}
					for i := 0; i < m; i++ {
						sum0 = 0
						for k := 0; k < n; k++ {
							sum0 += (*a)[i][k] * buf0[k]
						}
						(*c)[i][j+0] = sum0
					}
				}
			}(from, to)
			from = to
		}
		wg.Wait()
	*/
}
