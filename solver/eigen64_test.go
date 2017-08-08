package solver_test

import (
	"math"
	"testing"

	"github.com/Konstantin8105/GoLinAlg/matrix"
	"github.com/Konstantin8105/GoLinAlg/solver"
)

func TestEigen(t *testing.T) {
	A := matrix.NewMatrix64bySize(2, 2)
	A.Set(0, 0, -1.0)
	A.Set(1, 0, 2.0)

	A.Set(0, 1, -6.0)
	A.Set(1, 1, 6.0)

	e := solver.NewEigen(A)

	r := e.GetRealEigenvalues()
	if !isSame(r[0], 2.0) || !isSame(r[1], 3.0) {
		t.Errorf("Case 1 in not Ok")
	}

	im := e.GetImagEigenvalues()
	if !isSame(im[0], 0.0) || !isSame(im[1], 0.0) {
		t.Errorf("Case 2 is not Ok")
	}

	v := e.GetV()

	vector1 := []float64{-2.0, 1.0}
	factor1 := v.Get(0, 0) / vector1[0]
	if !isSame(v.Get(1, 0), factor1*vector1[1]) {
		t.Errorf("Case 3 is not Ok")
	}

	vector2 := []float64{-1.5, 1.0}
	factor2 := v.Get(0, 1) / vector2[0]
	if !isSame(v.Get(1, 1), factor2*vector2[1]) {
		t.Errorf("Case 4 is not Ok")
	}
}

func isSame(var1, var2 float64) bool {
	eps := 1e-7
	if math.Abs(var1) > eps {
		if math.Abs((var1-var2)/var1) < eps {
			return true
		}
	} else {
		if math.Abs(var1-var2) < eps {
			return true
		}
	}
	return false
}
