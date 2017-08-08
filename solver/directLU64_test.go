package solver_test

import (
	"math/rand"
	"testing"

	"github.com/Konstantin8105/GoLinAlg/matrix"
	"github.com/Konstantin8105/GoLinAlg/solver"
)

func TestMatrix(t *testing.T) {
	m := matrix.NewMatrix64bySize(10, 10)
	if m.GetRowSize() != 10 {
		t.Errorf("Not correct size")
	}
	if m.GetColumnSize() != 10 {
		t.Errorf("Not correct size")
	}
}

func TestLUSolver(t *testing.T) {
	A, x, b := getTest()
	s := solver.NewLUsolver(A)
	result, _ := s.Solve(b)

	if !result.IsEqual(x) {
		t.Errorf("Not correct LU solver")
	}
}

// getTest - test simple test data
func getTest() (A matrix.T64, x matrix.T64, b matrix.T64) {
	n := 20
	A = matrix.NewMatrix64bySize(n, n)
	b = matrix.NewMatrix64bySize(n, 1)
	x = matrix.NewMatrix64bySize(n, 1)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			A.Set(i, j, 4.0*rand.Float64()*float64(j-i+n*2))
		}
		x.Set(i, 0, 7.55*float64(i))
	}

	b = A.Times(&x)
	return
}

func TestSingular(t *testing.T) {
	A := matrix.NewMatrix64bySize(2, 2)
	A.Set(0, 0, -1.0)
	A.Set(1, 0, 2.0)

	A.Set(0, 1, -6.0)
	A.Set(1, 1, 12.0)

	e := solver.NewLUsolver(A)

	B := matrix.NewMatrix64bySize(2, 1)
	B.Set(0, 0, 1.0)
	B.Set(1, 0, 2.0)

	_, err := e.Solve(B)

	if err == nil {
		t.Errorf("Have not checking singular")
	}
}

func TestSizes(t *testing.T) {
	A := matrix.NewMatrix64bySize(2, 2)
	e := solver.NewLUsolver(A)
	B := matrix.NewMatrix64bySize(3, 1)
	_, err := e.Solve(B)
	if err == nil {
		t.Errorf("Have not checking of size. Input matrix is square, only")
	}
}
