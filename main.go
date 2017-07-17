package main

import (
	"fmt"

	"github.com/Konstantin8105/GoLinAlg/linAlg"
	"github.com/Konstantin8105/GoLinAlg/linAlg/solver"
)

func main() {
	A := linAlg.NewMatrix64bySize(2, 2)
	A.Set(0, 0, -1.0)
	A.Set(1, 0, 2.0)

	A.Set(0, 1, -6.0)
	A.Set(1, 1, 6.0)

	e := solver.NewEigen(A)

	fmt.Println("A = ", A)
	fmt.Println("getV = ", e.GetV())
	fmt.Println("getD = ", e.GetD())
	fmt.Println("eigenvalue = ", e.GetRealEigenvalues())
	fmt.Println("eigenImag  = ", e.GetImagEigenvalues())
}
