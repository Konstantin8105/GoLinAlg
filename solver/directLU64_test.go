package solver_test

import (
	"fmt"
	"math"
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

func TestLUcheck(t *testing.T) {
	A := matrix.NewMatrix64bySize(3, 3)
	A.Set(0, 0, 1.0)
	A.Set(0, 1, 2.0)
	A.Set(0, 2, 3.0)

	A.Set(1, 0, 4.0)
	A.Set(1, 1, 5.0)
	A.Set(1, 2, 6.0)

	A.Set(2, 0, 7.0)
	A.Set(2, 1, 8.0)
	A.Set(2, 2, 9.0)

	fmt.Println("A = ", A)

	lu := solver.NewLUsolver(A)

	expectedL := matrix.NewMatrix64bySize(3, 3)
	expectedL.Set(0, 0, 1.0)
	expectedL.Set(0, 1, 0.0)
	expectedL.Set(0, 2, 0.0)

	expectedL.Set(1, 0, 4.0)
	expectedL.Set(1, 1, 1.0)
	expectedL.Set(1, 2, 0.0)

	expectedL.Set(2, 0, 7.0)
	expectedL.Set(2, 1, 2.0)
	expectedL.Set(2, 2, 1.0)

	expectedU := matrix.NewMatrix64bySize(3, 3)
	expectedU.Set(0, 0, 1.0)
	expectedU.Set(0, 1, 2.0)
	expectedU.Set(0, 2, 3.0)

	expectedU.Set(1, 0, 0.0)
	expectedU.Set(1, 1, -3.0)
	expectedU.Set(1, 2, -6.0)

	expectedU.Set(2, 0, 0.0)
	expectedU.Set(2, 1, 0.0)
	expectedU.Set(2, 2, 0.0)

	actualL := lu.GetL()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if math.Abs(expectedL.Get(i, j)-actualL.Get(i, j)) > 1e-8 {
				t.Errorf("Not Ok for element [%v,%v]", i, j)
			}
		}
	}

	actualU := lu.GetU()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if math.Abs(expectedU.Get(i, j)-actualU.Get(i, j)) > 1e-8 {
				t.Errorf("Not Ok for element [%v,%v]", i, j)
			}
		}
	}

	actualA := actualL.Times(&actualU)
	fmt.Println("actual A = ", actualA)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if math.Abs(A.Get(i, j)-actualA.Get(i, j)) > 1e-8 {
				t.Errorf("Not Ok for element [%v,%v]", i, j)
			}
		}
	}

	expectedA := expectedL.Times(&expectedU)
	fmt.Println("expected A = ", expectedA)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if math.Abs(A.Get(i, j)-expectedA.Get(i, j)) > 1e-8 {
				t.Errorf("Not Ok for element [%v,%v]", i, j)
			}
		}
	}

	fmt.Println("pivot -->", lu.GetPivot())
	fmt.Println("pivot2 ->", lu.GetPivotFloat64())
}
