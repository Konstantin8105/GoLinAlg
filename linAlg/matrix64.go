package linAlg

import (
	"fmt"
	"math"
)

// Matrix64 - matrix with values "float64"
type Matrix64 struct {
	values     [][]float64 // Internal storage of values
	sizeRow    int         // Size of rows
	sizeColumn int         // Size of columns
}

// NewMatrix64bySize - constructor type of Matrix64 by sizes
func NewMatrix64bySize(rows, columns int) (m Matrix64) {
	m.sizeRow = rows
	m.sizeColumn = columns
	m.values = make([][]float64, rows, rows)
	for row := 0; row < rows; row++ {
		m.values[row] = make([]float64, columns, columns)
	}
	return
}

// NewMatrix64byMatrix64 - constructor type of Matrix64 by other Matrix64
func NewMatrix64byMatrix64(in Matrix64) (m Matrix64) {
	m = NewMatrix64bySize(in.sizeRow, in.sizeColumn)
	for row := 0; row < in.sizeRow; row++ {
		for column := 0; column < in.sizeColumn; column++ {
			m.values[row][column] = in.values[row][column]
		}
	}
	return
}

// GetRowSize - return size of row
func (m *Matrix64) GetRowSize() int {
	return m.sizeRow
}

// GetColumnSize - return size of column
func (m *Matrix64) GetColumnSize() int {
	return m.sizeColumn
}

// Get - return value of matrix
func (m *Matrix64) Get(i, j int) float64 {
	if i < 0 || i >= m.sizeRow || j < 0 || j >= m.sizeColumn {
		panic(fmt.Errorf("Not correct algoritm for [%v,%v] - out of matrix", i, j))
	}
	return m.values[i][j]
}

// Set - insert into matrix
func (m *Matrix64) Set(i, j int, value float64) {
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
func (m *Matrix64) GetSubMatrix(r []int, j0, j1 int) Matrix64 {
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
func (m *Matrix64) IsEqual(m2 Matrix64) bool {
	if m.sizeRow != m2.sizeRow {
		return false
	}
	if m.sizeColumn != m2.sizeColumn {
		return false
	}
	eps := 1e-7
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
func (m Matrix64) Times(B Matrix64) (result Matrix64) {
	if B.GetRowSize() != m.GetColumnSize() {
		panic(fmt.Errorf("Matrix inner dimensions must agree"))
	}
	x := NewMatrix64bySize(m.GetRowSize(), B.GetColumnSize())
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

func (m Matrix64) String() (s string) {
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
