package matrix_test

import (
	"testing"

	"github.com/Konstantin8105/GoLinAlg/matrix"
)

func TestString1(t *testing.T) {
	m := matrix.NewMatrix64bySize(1, 1)
	if len(m.String()) == 0 {
		t.Errorf("cannot convert to string")
	}
}

func TestString2(t *testing.T) {
	m := matrix.NewMatrix64bySize(2, 2)
	if len(m.String()) == 0 {
		t.Errorf("cannot convert to string")
	}
}

func TestRemoveRowAndColumn1(t *testing.T) {
	size := 2
	m := matrix.NewMatrix64bySize(size, size)
	if m.GetRowSize() != size {
		t.Errorf("Not correct amount of rows")
	}
	if m.GetColumnSize() != size {
		t.Errorf("Not correct amount of columns")
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			m.Set(i, j, float64(i+1)*10.+float64(j+1)*100.)
		}
	}
	if m.Get(0, 0) != 110. {
		t.Errorf("Cannot get correct result")
	}
	if m.Get(1, 0) != 120. {
		t.Errorf("Cannot get correct result")
	}
	if m.Get(1, 1) != 220. {
		t.Errorf("Cannot get correct result")
	}
	m.RemoveRowAndColumn()
	if m.GetRowSize() != 2 && m.GetColumnSize() != 2 {
		t.Errorf("Not correct amount of matrix")
	}

	m.RemoveRowAndColumn(0)
	if m.GetRowSize() != 1 && m.GetColumnSize() != 1 {
		t.Errorf("Not correct amount of matrix")
	}
	if m.Get(0, 0) != 220. {
		t.Errorf("Not correct removing row and column function")
	}
}

func TestRemoveRowAndColumn2(t *testing.T) {
	m := matrix.NewMatrix64bySize(1, 1)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	// Testing
	m.RemoveRowAndColumn(0, -1)
}

func TestRemoveRowAndColumn3(t *testing.T) {
	m := matrix.NewMatrix64bySize(1, 1)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	// Testing
	m.RemoveRowAndColumn(0, 1)
}

func TestRemoveRowAndColumn4(t *testing.T) {
	m := matrix.NewMatrix64bySize(5, 3)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	// Testing
	m.RemoveRowAndColumn(4)
}

func TestEqual1(t *testing.T) {
	m1 := matrix.NewMatrix64bySize(1, 1)
	m1.Set(0, 0, 42.)
	m2 := matrix.NewMatrix64bySize(1, 1)
	m2.Set(0, 0, 42.)
	if !m1.IsEqual(m2) {
		t.Errorf("Matrixes is not equal")
	}
}

func TestEqual2(t *testing.T) {
	m1 := matrix.NewMatrix64bySize(1, 1)
	m1.Set(0, 0, 42.)
	m2 := matrix.NewMatrix64bySize(2, 1)
	m2.Set(0, 0, 42.)
	if m1.IsEqual(m2) {
		t.Errorf("Matrixes is equal")
	}
}

func TestEqual3(t *testing.T) {
	m1 := matrix.NewMatrix64bySize(1, 1)
	m1.Set(0, 0, 42.)
	m2 := matrix.NewMatrix64bySize(1, 2)
	m2.Set(0, 0, 42.)
	if m1.IsEqual(m2) {
		t.Errorf("Matrixes is equal")
	}
}

func TestEqual4(t *testing.T) {
	m1 := matrix.NewMatrix64bySize(1, 1)
	m1.Set(0, 0, 42.)
	m2 := matrix.NewMatrix64bySize(1, 1)
	m2.Set(0, 0, -42.)
	if m1.IsEqual(m2) {
		t.Errorf("Matrixes is equal")
	}
}

func TestGet1(t *testing.T) {
	m := matrix.NewMatrix64bySize(1, 1)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	// Testing
	_ = m.Get(-1, 0)
}

func TestGet2(t *testing.T) {
	m := matrix.NewMatrix64bySize(1, 1)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	// Testing
	_ = m.Get(0, -1)
}

func TestGet3(t *testing.T) {
	m := matrix.NewMatrix64bySize(1, 1)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	// Testing
	_ = m.Get(1, 0)
}

func TestGet4(t *testing.T) {
	m := matrix.NewMatrix64bySize(1, 1)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	// Testing
	_ = m.Get(0, 1)
}

func TestGet5(t *testing.T) {
	m := *new(matrix.T64)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	// Testing
	_ = m.Get(0, 1)
}

func TestCopyByMatrix(t *testing.T) {
	m1 := matrix.NewMatrix64bySize(2, 1)
	m1.Set(1, 0, 42.)
	m2 := matrix.NewMatrix64byMatrix64(m1)
	if m1.GetRowSize() != m2.GetRowSize() || m1.GetColumnSize() != m2.GetColumnSize() {
		t.Errorf("Cannot create matrix with same size")
	}
}

func TestSet1(t *testing.T) {
	m := matrix.NewMatrix64bySize(1, 1)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	// Testing
	m.Set(-1, 0, 42.)
}

func TestSet2(t *testing.T) {
	m := matrix.NewMatrix64bySize(1, 1)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	// Testing
	m.Set(0, -1, 42.)
}

func TestSet3(t *testing.T) {
	m := matrix.NewMatrix64bySize(1, 1)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	// Testing
	m.Set(1, 0, 42.)
}

func TestSet4(t *testing.T) {
	m := matrix.NewMatrix64bySize(1, 1)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	// Testing
	m.Set(0, 1, 42.)
}

func TestTimes(t *testing.T) {
	A := matrix.NewMatrix64bySize(2, 4)
	A.Set(0, 0, 2.)
	A.Set(0, 2, 4.)
	A.Set(0, 3, -1.)
	A.Set(1, 0, 1.)
	A.Set(1, 1, -1.)
	A.Set(1, 2, 1.)

	B := matrix.NewMatrix64bySize(4, 1)
	B.Set(0, 0, 2.)
	B.Set(1, 0, 1.)
	B.Set(3, 0, -2.)

	C := A.Times(&B)
	if C.Get(0, 0) != 6. || C.Get(1, 0) != 1. {
		t.Errorf("multiplication is wrong. C = %v", C)
	}
}

func TestTimesPanic(t *testing.T) {
	m1 := matrix.NewMatrix64bySize(1, 1)
	m2 := matrix.NewMatrix64bySize(2, 2)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	// Testing
	_ = m1.Times(&m2)

}
