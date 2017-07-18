package matrix

import (
	"fmt"
	"math"
	"runtime"
	"sort"
	"sync"
)

// T64 - matrix with values "float64"
// acceptable to use for vectors
type T64 struct {
	values         [][]float64 // Internal storage of values
	sizeRow        int         // Size of rows
	capacityRow    int         // Capacity of rows = real size of matrix
	sizeColumn     int         // Size of columns
	capacityColumn int         // Capacity of columns = real size of matrix
}

// NewMatrix64bySize - constructor type of Matrix64 by sizes
func NewMatrix64bySize(rows, columns int) (m T64) {
	m.sizeRow = rows
	m.capacityRow = rows
	m.sizeColumn = columns
	m.capacityColumn = columns
	m.values = make([][]float64, rows, rows)
	for row := 0; row < rows; row++ {
		m.values[row] = make([]float64, columns, columns)
	}
	return
}

// NewMatrix64byMatrix64 - constructor type of t64 by other t64
func NewMatrix64byMatrix64(in T64) (m T64) {
	m = NewMatrix64bySize(in.sizeRow, in.sizeColumn)
	for row := 0; row < in.sizeRow; row++ {
		for column := 0; column < in.sizeColumn; column++ {
			m.values[row][column] = in.values[row][column]
		}
	}
	return
}

// SetNewSize - resize a matrix with zero initialization of matrix
func (m *T64) SetNewSize(rows, columns int) {
	if rows > m.capacityRow || columns > m.capacityColumn {
		*m = NewMatrix64bySize(rows, columns)
		return
	}
	m.sizeRow = rows
	m.sizeColumn = columns
	for i := 0; i < m.sizeRow; i++ {
		for j := 0; j < m.sizeColumn; j++ {
			m.values[i][j] = 0.0
		}
	}
}

// GetRowSize - return size of row
func (m *T64) GetRowSize() int {
	return m.sizeRow
}

// GetColumnSize - return size of column
func (m *T64) GetColumnSize() int {
	return m.sizeColumn
}

// Get - return value of matrix
func (m *T64) Get(i, j int) float64 {
	if i < 0 || i >= m.sizeRow || j < 0 || j >= m.sizeColumn {
		panic(fmt.Errorf("Not correct algoritm for [%v,%v] - out of matrix", i, j))
	}
	return m.values[i][j]
}

// Set - insert into matrix
func (m *T64) Set(i, j int, value float64) {
	if i < 0 || i >= m.sizeRow || j < 0 || j >= m.sizeColumn {
		panic(fmt.Errorf("Not correct algoritm for [%v,%v] - out of matrix", i, j))
	}
	m.values[i][j] = value
}

// GetSubMatrix - Get a submatrix
// r  - Array of row indices.
// j0 - Initial column index
// j1 - Final column index
// return A(r(:),j0:j1)
// or error : Submatrix indices
func (m *T64) GetSubMatrix(r []int, j0, j1 int) T64 {
	x := NewMatrix64bySize(len(r), j1-j0+1)
	for i := 0; i < len(r); i++ {
		if r[i] < 0 || r[i] >= m.sizeRow {
			panic(fmt.Errorf("Index is outside of matrix. i = %v", i))
		}
		for j := j0; j <= j1; j++ {
			if j < 0 || j >= m.sizeColumn {
				panic(fmt.Errorf("Index is outside of matrix. j = %v", j))
			}
			x.values[i][j-j0] = m.values[r[i]][j]
		}
	}
	return x
}

// IsEqual - return true is matrix are equal
func (m *T64) IsEqual(m2 T64) bool {
	if m.sizeRow != m2.sizeRow {
		return false
	}
	if m.sizeColumn != m2.sizeColumn {
		return false
	}
	eps := 1e-8
	for i := 0; i < m.sizeRow; i++ {
		for j := 0; j < m.sizeColumn; j++ {
			e := math.Abs(m.values[i][j] - m2.values[i][j])
			if e > eps {
				return false
			}
		}
	}
	return true
}

// Times - Linear algebraic matrix multiplication, A * B or error
func (m T64) Times(B T64) (result T64) {
	if B.GetRowSize() != m.GetColumnSize() {
		panic(fmt.Errorf("Matrix inner dimensions must agree"))
	}
	x := NewMatrix64bySize(m.GetRowSize(), B.GetColumnSize())

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
	return x

}

// MultiplyTtKT - multiply matrix
// formula: T(transponse) * M * T
func (m *T64) MultiplyTtKT(t T64) T64 {
	if t.GetRowSize() != m.GetRowSize() {
		panic("Not correct algoritm")
	}

	buffer := NewMatrix64bySize(t.GetColumnSize(), m.GetColumnSize())

	for i := 0; i < buffer.GetRowSize(); i++ {
		for j := 0; j < buffer.GetColumnSize(); j++ {
			sum := 0.0
			for k := 0; k < t.GetRowSize(); k++ {
				sum += t.Get(k, i) * m.Get(k, j)
			}
			buffer.Set(i, j, sum)
		}
	}

	result := NewMatrix64bySize(buffer.GetRowSize(), t.GetColumnSize())
	for i := 0; i < result.GetRowSize(); i++ {
		for j := 0; j < result.GetColumnSize(); j++ {
			sum := 0.0
			for k := 0; k < buffer.GetColumnSize(); k++ {
				sum += buffer.Get(i, k) * t.Get(k, j)
			}
			result.Set(i, j, sum)
		}
	}

	return result
}

// RemoveRowAndColumn - remove rows and columns of matrix
// without reallocation matrix
func (m *T64) RemoveRowAndColumn(indexes ...int) {
	if len(indexes) == 0 {
		return
	}
	// sorting indexes for optimization of algoritm
	sort.Ints(indexes)
	// global checking indexes
	if indexes[0] < 0 {
		panic(fmt.Errorf("Index is outside of matrix indexes = %v", indexes))
	}
	if indexes[len(indexes)-1] >= m.sizeRow || indexes[len(indexes)-1] >= m.sizeColumn {
		panic(fmt.Errorf("indexes is outside of matrix. Indexes = %v", indexes))
	}
	// modify values of matrix
	positionIndexI := 0
	newPositionInMatrixI := 0
	for i := 0; i < m.sizeRow; i++ {
		if positionIndexI != len(indexes) && i == indexes[positionIndexI] {
			positionIndexI++
			continue
		}
		positionIndexJ := 0
		newPositionInMatrixJ := 0
		for j := 0; j < m.sizeColumn; j++ {
			if positionIndexJ != len(indexes) && j == indexes[positionIndexJ] {
				positionIndexJ++
				continue
			}
			m.Set(newPositionInMatrixI, newPositionInMatrixJ, m.Get(i, j))
			newPositionInMatrixJ++
		}
		newPositionInMatrixI++
	}
	m.sizeRow = m.sizeRow - len(indexes)
	m.sizeColumn = m.sizeColumn - len(indexes)
}

func (m T64) String() (s string) {
	s += fmt.Sprintf("\n")
	for i := 0; i < m.sizeRow; i++ {
		s += fmt.Sprintf("[")
		for j := 0; j < m.sizeColumn; j++ {
			s += fmt.Sprintf("%.4E", m.values[i][j])
			if j != m.sizeColumn-1 {
				s += fmt.Sprintf(",")
			}
		}
		s += fmt.Sprintf("]\n")
	}
	return
}
