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

func TestRemoveRowAndColumn(t *testing.T) {
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
