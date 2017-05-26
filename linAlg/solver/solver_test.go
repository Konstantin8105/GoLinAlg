package solver_test

import (
	"math/rand"
	"testing"

	"github.com/Konstantin8105/GoLinAlg/linAlg"
	"github.com/Konstantin8105/GoLinAlg/linAlg/solver"
)

func TestMatrix(t *testing.T) {
	m := linAlg.NewMatrix64bySize(10, 10)
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
	result := s.Solve(A, b)

	if !result.IsEqual(x) {
		t.Errorf("Not correct LU solver")
	}
}

// getTest - test simple test data
func getTest() (A linAlg.Matrix64, x linAlg.Matrix64, b linAlg.Matrix64) {
	n := 20
	A = linAlg.NewMatrix64bySize(n, n)
	b = linAlg.NewMatrix64bySize(n, 1)
	x = linAlg.NewMatrix64bySize(n, 1)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			A.Set(i, j, 4.0*rand.Float64()*float64(j-i+n*2))
		}
		x.Set(i, 0, 7.55*float64(i))
	}

	b = A.Times(x)

	//fmt.Println("A = ", A)
	//fmt.Println("b = ", b)
	//fmt.Println("x = ", x)

	return
}
