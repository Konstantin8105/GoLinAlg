package solver

import (
	"fmt"
	"math"

	"github.com/Konstantin8105/GoLinAlg/matrix"
)

// LU64 - LU Decomposition
// For an m-by-n matrix A with m >= n, the LU decomposition is an m-by-n
// unit lower triangular matrix L, an n-by-n upper triangular matrix U,
// and a permutation vector piv of length m so that A(piv,:) = L*U.
// If m < n, then L is m-by-m and U is m-by-n.
// The LU decompostion with pivoting always exists, even if the matrix is
// singular, so the constructor will never fail.  The primary use of the
// LU decomposition is in the solution of square systems of simultaneous
// linear equations.  This will fail if isNonsingular() returns false.
type LU64 struct {
	lu         matrix.T64 // Array for internal storage of decomposition
	sizeRow    int        // Size of rows
	sizeColumn int        // Size of columns
	pivsign    int        // Pivot sign
	piv        []int      // Internal storage of pivot vector
}

// NewLUsolver - constructor
func NewLUsolver(A matrix.T64) (s LU64) {
	s.pivsign = 1

	s.lu = matrix.NewMatrix64byMatrix64(A)
	m := s.lu.GetRowSize()
	n := s.lu.GetColumnSize()
	s.piv = make([]int, m)
	for i := 0; i < m; i++ {
		s.piv[i] = i
	}

	LUcolj := make([]float64, m)

	// Outer loop.

	for j := 0; j < n; j++ {

		// Make a copy of the j-th column to localize references.
		for i := 0; i < m; i++ {
			LUcolj[i] = s.lu.Get(i, j)
		}

		// Apply previous transformations.
		for i := 0; i < m; i++ {
			// kmax = min(i,j)
			kmax := j
			if i < j {
				kmax = i
			}

			// Most of the time is spent in the following dot product.
			sum := 0.0
			for k := 0; k < kmax; k++ {
				sum += s.lu.Get(i, k) * LUcolj[k]
			}
			LUcolj[i] -= sum
			s.lu.Set(i, j, LUcolj[i])
		}

		// Find pivot and exchange if necessary.
		p := j
		for i := j + 1; i < m; i++ {
			if math.Abs(LUcolj[i]) > math.Abs(LUcolj[p]) {
				p = i
			}
		}
		if p != j {
			for k := 0; k < n; k++ {
				t := s.lu.Get(p, k)
				s.lu.Set(p, k, s.lu.Get(j, k))
				s.lu.Set(j, k, t)
			}
			s.piv[p], s.piv[j] = s.piv[j], s.piv[p]
			s.pivsign = -s.pivsign
		}

		// Compute multipliers.
		if j < m && s.lu.Get(j, j) != 0.0 {
			for i := j + 1; i < m; i++ {
				s.lu.Set(i, j, s.lu.Get(i, j)/s.lu.Get(j, j))
			}
		}
	}
	return s
}

// Is the matrix nonsingular?
// return true if U, and hence A, is nonsingular.
func (s *LU64) isNonsingular() bool {
	for j := 0; j < s.lu.GetColumnSize(); j++ {
		if s.lu.Get(j, j) == 0.0 {
			return false
		}
	}
	return true
}

// GetL - Return matrix lower triangular factor - L
func (s *LU64) GetL() matrix.T64 {
	x := matrix.NewMatrix64bySize(s.lu.GetRowSize(), s.lu.GetColumnSize())
	m := x.GetRowSize()
	n := x.GetColumnSize()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > j {
				x.Set(i, j, s.lu.Get(i, j))
			} else if i == j {
				x.Set(i, j, 1.0)
			} else {
				x.Set(i, j, 0.0)
			}
		}
	}
	return x
}

// GetU - Return upper triangular factor - U
func (s *LU64) GetU() matrix.T64 {
	x := matrix.NewMatrix64bySize(s.lu.GetColumnSize(), s.lu.GetColumnSize())
	n := x.GetColumnSize()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i <= j {
				x.Set(i, j, s.lu.Get(i, j))
			} else {
				x.Set(i, j, 0.0)
			}
		}
	}
	return x
}

// GetPivot -  Return pivot permutation vector piv
func (s *LU64) GetPivot() []int {
	p := make([]int, s.lu.GetRowSize())
	m := len(p)
	for i := 0; i < m; i++ {
		p[i] = s.piv[i]
	}
	return p
}

// GetPivotFloat64 - Return pivot permutation vector as a one-dimensional float64 array
func (s *LU64) GetPivotFloat64() []float64 {
	p := make([]float64, s.lu.GetRowSize())
	m := len(p)
	for i := 0; i < m; i++ {
		p[i] = (float64)(s.piv[i])
	}
	return p
}

// Det - determinant det(A) or error  if Matrix must be square
func (s *LU64) Det() float64 {
	if s.lu.GetRowSize() != s.lu.GetColumnSize() {
		panic(fmt.Errorf("Matrix must be square"))
	}
	d := (float64)(s.pivsign)
	n := s.lu.GetColumnSize()
	for j := 0; j < n; j++ {
		d *= s.lu.Get(j, j)
	}
	return d
}

// Solve - solving A * x = b
// A Matrix with as many rows as A and any number of columns.
// X so that L*U*X = B(piv,:)
// error - Matrix row dimensions must agree. or Matrix is singular.
func (s *LU64) Solve(b matrix.T64) (x matrix.T64, err error) {

	if b.GetRowSize() != s.lu.GetRowSize() {
		return x, fmt.Errorf("Matrix row dimensions must agree")
	}
	if !(s.isNonsingular()) {
		return x, fmt.Errorf("Matrix is singular")
	}

	// Copy right hand side with pivoting
	nx := b.GetColumnSize()
	x = b.GetSubMatrix(s.piv, 0, nx-1)

	n := s.lu.GetColumnSize()
	// Solve L*Y = B(piv,:)
	for k := 0; k < n; k++ {
		for i := k + 1; i < n; i++ {
			for j := 0; j < nx; j++ {
				x.Set(i, j, x.Get(i, j)-x.Get(k, j)*s.lu.Get(i, k))
			}
		}
	}
	// Solve U*X = Y;
	for k := n - 1; k >= 0; k-- {
		for j := 0; j < nx; j++ {
			x.Set(k, j, x.Get(k, j)/s.lu.Get(k, k))
		}
		for i := 0; i < k; i++ {
			for j := 0; j < nx; j++ {
				x.Set(i, j, x.Get(i, j)-x.Get(k, j)*s.lu.Get(i, k))
			}
		}
	}
	return x, nil
}
