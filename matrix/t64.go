package matrix

import (
	"fmt"
	"math"
	"sort"
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
func (m *T64) Times(B *T64) (result T64) {
	if B.GetRowSize() != m.GetColumnSize() {
		panic(fmt.Errorf("Matrix inner dimensions must agree"))
	}
	X := NewMatrix64bySize(m.GetRowSize(), B.GetColumnSize())
	//timesAlgorithm(&m.values, &B.values, &x.values, m.GetRowSize(), m.GetColumnSize(), B.GetColumnSize())
	// matrix a with size [m,n]
	// matrix b with size [n,h]
	// matrix c with size [m,h]
	a := &m.values
	b := &B.values
	c := &X.values
	t := m.GetRowSize()
	n := m.GetColumnSize()
	h := B.GetColumnSize()
	//func timesAlgorithm(a, b, c *[][]float64, t, n, h int) {
	var sum float64
	for j := 0; j < h; j++ {
		for i := 0; i < t; i++ {
			sum = 0
			for k := 0; k < n; k++ {
				sum += (*a)[i][k] * (*b)[k][j]
			}
			(*c)[i][j] = sum
		}
	}
	//}
	return X

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
