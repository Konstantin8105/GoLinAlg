package matrix

// matrix a with size [m,n]
// matrix b with size [n,h]
// matrix c with size [m,h]
func timesAlgorithm(a, b, c *[][]float64, m, n, h int) {
	for i := 0; i < m; i++ {
		for j := 0; j < h; j++ {
			for k := 0; k < n; k++ {
				(*c)[i][j] += (*a)[i][k] * (*b)[k][j]
			}
		}
	}
}

/*
	MinSizeForParallel := 100
	if MinSizeForParallel < B.GetRowSize() {
		// Found amount allowable parallelism
		threads := runtime.GOMAXPROCS(0)
		if threads > runtime.NumCPU() {
			threads = runtime.NumCPU()
		}
		// Create workgroup
		var wg sync.WaitGroup
		// Run calculation in goroutines
		for t := 0; t < threads; t++ {
			// Add one goroutine in workgroup
			wg.Add(1)
			// The value "init" is a number of thread
			// that created for offset of loop
			go func(init int) {
				// Change waitgroup after work done
				defer wg.Done()
				Bcolj := make([]float64, m.GetColumnSize(), m.GetColumnSize())
				for j := init; j < B.GetColumnSize(); j += threads {
					for k := 0; k < m.GetColumnSize(); k++ {
						Bcolj[k] = B.Get(k, j)
					}
					for i := 0; i < m.GetRowSize(); i++ {
						sum := 0.0
						for k := 0; k < m.GetColumnSize(); k++ {
							sum += m.values[i][k] * Bcolj[k]
						}
						x.Set(i, j, sum)
					}
				}
			}(t)
		}
		wg.Wait()
		return x
	}
*/
/*
	Bcolj := make([]float64, m.GetColumnSize(), m.GetColumnSize())
	for j := 0; j < B.GetColumnSize(); j++ {
		for k := 0; k < m.GetColumnSize(); k++ {
			Bcolj[k] = B.Get(k, j)
		}
		for i := 0; i < m.GetRowSize(); i++ {
			sum := 0.0
			for k := 0; k < m.GetColumnSize(); k++ {
				sum += m.values[i][k] * Bcolj[k]
			}
			x.Set(i, j, sum)
		}
	}
*/
/*
	for j := 0; j < B.GetColumnSize(); j++ {
		for i := 0; i < m.GetRowSize(); i++ {
			sum := 0.0
			for k := 0; k < m.GetColumnSize(); k++ {
				sum += m.values[i][k] * B.Get(k, j)
			}
			x.Set(i, j, sum)
		}
	}
*/
/*
	for j := 0; j < B.GetColumnSize(); j++ {
		for i := 0; i < m.GetRowSize(); i++ {
			for k := 0; k < m.GetColumnSize(); k++ {
				x.Set(i, j, x.Get(i, j)+m.values[i][k]*B.Get(k, j))
			}
		}
	}
*/
