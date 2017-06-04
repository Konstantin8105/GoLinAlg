package linAlg_test

import (
	"math/rand"
	"sync"
	"testing"
)

func a1(A, B, result *[][]float64) {
	n := len(*A)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				(*result)[i][j] += (*A)[i][k] * (*B)[k][j]
			}
		}
	}
}

func a2(A, B, result *[][]float64) {
	n := len(*A)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := 0.0
			for k := 0; k < n; k++ {
				sum += (*A)[i][k] * (*B)[k][j]
			}
			(*result)[i][j] = sum
		}
	}
}

func a3(A, B, result *[][]float64) {
	n := len(*A)
	buffer := make([]float64, n)
	for i := 0; i < n; i++ {
		for k := 0; k < n; k++ {
			buffer[k] = (*A)[i][k]
		}
		for j := 0; j < n; j++ {
			sum := 0.0
			for k := 0; k < n; k++ {
				sum += buffer[k] * (*B)[k][j]
			}
			(*result)[i][j] = sum
		}
	}
}

func a4(A, B, result *[][]float64) {
	n := len(*A)
	buffer1 := make([]float64, n)
	buffer2 := make([]float64, n)
	half := n / 2
	for i := 0; i < n/2; i++ {
		for k := 0; k < n; k++ {
			buffer1[k] = (*A)[i][k]
			buffer2[k] = (*A)[i+half][k]
		}
		for j := 0; j < n; j++ {
			sum1 := 0.0
			sum2 := 0.0
			for k := 0; k < n; k++ {
				sum1 += buffer1[k] * (*B)[k][j]
				sum2 += buffer2[k] * (*B)[k][j]
			}
			(*result)[i][j] = sum1
			(*result)[i+half][j] = sum2
		}
	}
}

func a5(A, B, result *[][]float64) {
	n := len(*A)
	buffer0 := make([]float64, n)
	buffer1 := make([]float64, n)
	buffer2 := make([]float64, n)
	buffer3 := make([]float64, n)
	buffer4 := make([]float64, n)
	part := n / 5
	for i := 0; i < part; i++ {
		for k := 0; k < n; k++ {
			buffer0[k] = (*A)[i*5+0][k]
			buffer1[k] = (*A)[i*5+1][k]
			buffer2[k] = (*A)[i*5+2][k]
			buffer3[k] = (*A)[i*5+3][k]
			buffer4[k] = (*A)[i*5+4][k]
		}
		for j := 0; j < n; j++ {
			sum0 := 0.0
			sum1 := 0.0
			sum2 := 0.0
			sum3 := 0.0
			sum4 := 0.0
			for k := 0; k < n; k++ {
				sum0 += buffer0[k] * (*B)[k][j]
				sum1 += buffer1[k] * (*B)[k][j]
				sum2 += buffer2[k] * (*B)[k][j]
				sum3 += buffer3[k] * (*B)[k][j]
				sum4 += buffer4[k] * (*B)[k][j]
			}
			(*result)[i*5+0][j] = sum0
			(*result)[i*5+1][j] = sum1
			(*result)[i*5+2][j] = sum2
			(*result)[i*5+3][j] = sum3
			(*result)[i*5+4][j] = sum4
		}
	}
}

func a6(A, B, result *[][]float64) {
	n := len(*A)
	buffer1 := make([]float64, n)
	buffer2 := make([]float64, n)
	buffer3 := make([]float64, n)
	buffer4 := make([]float64, n)
	buffer5 := make([]float64, n)
	buffer6 := make([]float64, n)
	buffer7 := make([]float64, n)
	buffer8 := make([]float64, n)
	buffer9 := make([]float64, n)
	buffer10 := make([]float64, n)
	part := n / 10
	for i := 0; i < part; i++ {
		for k := 0; k < n; k++ {
			buffer1[k] = (*A)[i*10+0][k]
			buffer2[k] = (*A)[i*10+1][k]
			buffer3[k] = (*A)[i*10+2][k]
			buffer4[k] = (*A)[i*10+3][k]
			buffer5[k] = (*A)[i*10+4][k]
			buffer6[k] = (*A)[i*10+5][k]
			buffer7[k] = (*A)[i*10+6][k]
			buffer8[k] = (*A)[i*10+7][k]
			buffer9[k] = (*A)[i*10+8][k]
			buffer10[k] = (*A)[i*10+9][k]
		}
		for j := 0; j < n; j++ {
			sum1 := 0.0
			sum2 := 0.0
			sum3 := 0.0
			sum4 := 0.0
			sum5 := 0.0
			sum6 := 0.0
			sum7 := 0.0
			sum8 := 0.0
			sum9 := 0.0
			sum10 := 0.0
			for k := 0; k < n; k++ {
				sum1 += buffer1[k] * (*B)[k][j]
				sum2 += buffer2[k] * (*B)[k][j]
				sum3 += buffer3[k] * (*B)[k][j]
				sum4 += buffer4[k] * (*B)[k][j]
				sum5 += buffer5[k] * (*B)[k][j]
				sum6 += buffer6[k] * (*B)[k][j]
				sum7 += buffer7[k] * (*B)[k][j]
				sum8 += buffer8[k] * (*B)[k][j]
				sum9 += buffer9[k] * (*B)[k][j]
				sum10 += buffer10[k] * (*B)[k][j]
			}
			(*result)[i*10+0][j] = sum1
			(*result)[i*10+1][j] = sum2
			(*result)[i*10+2][j] = sum3
			(*result)[i*10+3][j] = sum4
			(*result)[i*10+4][j] = sum5
			(*result)[i*10+5][j] = sum6
			(*result)[i*10+6][j] = sum7
			(*result)[i*10+7][j] = sum8
			(*result)[i*10+8][j] = sum9
			(*result)[i*10+9][j] = sum10
		}
	}
}

func a7(A, B, result *[][]float64) {
	n := len(*A)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			buffer := make([]float64, n)
			for k := 0; k < n; k++ {
				buffer[k] = (*A)[i][k]
			}
			for j := 0; j < n; j++ {
				sum := 0.0
				for k := 0; k < n; k++ {
					sum += buffer[k] * (*B)[k][j]
				}
				(*result)[i][j] = sum
			}
		}(i)
	}
	wg.Wait()
}

func a8(A, B, result *[][]float64) {
	var wg sync.WaitGroup
	amount := 4
	amountBuffers := 10
	for g := 0; g < amount; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			n := len(*A)
			buffer0 := make([]float64, n)
			buffer1 := make([]float64, n)
			buffer2 := make([]float64, n)
			buffer3 := make([]float64, n)
			buffer4 := make([]float64, n)
			buffer5 := make([]float64, n)
			buffer6 := make([]float64, n)
			buffer7 := make([]float64, n)
			buffer8 := make([]float64, n)
			buffer9 := make([]float64, n)
			part := n / amountBuffers
			for i := g; i < part; i += amount {
				for k := 0; k < n; k++ {
					buffer0[k] = (*A)[i*amountBuffers+0][k]
					buffer1[k] = (*A)[i*amountBuffers+1][k]
					buffer2[k] = (*A)[i*amountBuffers+2][k]
					buffer3[k] = (*A)[i*amountBuffers+3][k]
					buffer4[k] = (*A)[i*amountBuffers+4][k]
					buffer5[k] = (*A)[i*amountBuffers+5][k]
					buffer6[k] = (*A)[i*amountBuffers+6][k]
					buffer7[k] = (*A)[i*amountBuffers+7][k]
					buffer8[k] = (*A)[i*amountBuffers+8][k]
					buffer9[k] = (*A)[i*amountBuffers+9][k]
				}
				for j := 0; j < n; j++ {
					sum0 := 0.0
					sum1 := 0.0
					sum2 := 0.0
					sum3 := 0.0
					sum4 := 0.0
					sum5 := 0.0
					sum6 := 0.0
					sum7 := 0.0
					sum8 := 0.0
					sum9 := 0.0
					for k := 0; k < n; k++ {
						sum0 += buffer0[k] * (*B)[k][j]
						sum1 += buffer1[k] * (*B)[k][j]
						sum2 += buffer2[k] * (*B)[k][j]
						sum3 += buffer3[k] * (*B)[k][j]
						sum4 += buffer4[k] * (*B)[k][j]
						sum5 += buffer5[k] * (*B)[k][j]
						sum6 += buffer6[k] * (*B)[k][j]
						sum7 += buffer7[k] * (*B)[k][j]
						sum8 += buffer8[k] * (*B)[k][j]
						sum9 += buffer9[k] * (*B)[k][j]
					}
					(*result)[i*amountBuffers+0][j] = sum0
					(*result)[i*amountBuffers+1][j] = sum1
					(*result)[i*amountBuffers+2][j] = sum2
					(*result)[i*amountBuffers+3][j] = sum3
					(*result)[i*amountBuffers+4][j] = sum4
					(*result)[i*amountBuffers+5][j] = sum5
					(*result)[i*amountBuffers+6][j] = sum6
					(*result)[i*amountBuffers+7][j] = sum7
					(*result)[i*amountBuffers+8][j] = sum8
					(*result)[i*amountBuffers+9][j] = sum9
				}
			}
		}(g)
	}
	wg.Wait()
}

func a9(A, B, result *[][]float64) {
	var wg sync.WaitGroup
	amount := 4
	for g := 0; g < amount; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			n := len(*A)
			buffer0 := make([]float64, n)
			buffer1 := make([]float64, n)
			buffer2 := make([]float64, n)
			buffer3 := make([]float64, n)
			buffer4 := make([]float64, n)
			part := n / 5
			for i := g; i < part; i += amount {
				for k := 0; k < n; k++ {
					buffer0[k] = (*A)[i*5+0][k]
					buffer1[k] = (*A)[i*5+1][k]
					buffer2[k] = (*A)[i*5+2][k]
					buffer3[k] = (*A)[i*5+3][k]
					buffer4[k] = (*A)[i*5+4][k]
				}
				for j := 0; j < n; j++ {
					sum0 := 0.0
					sum1 := 0.0
					sum2 := 0.0
					sum3 := 0.0
					sum4 := 0.0
					for k := 0; k < n; k++ {
						sum0 += buffer0[k] * (*B)[k][j]
						sum1 += buffer1[k] * (*B)[k][j]
						sum2 += buffer2[k] * (*B)[k][j]
						sum3 += buffer3[k] * (*B)[k][j]
						sum4 += buffer4[k] * (*B)[k][j]
					}
					(*result)[i*5+0][j] = sum0
					(*result)[i*5+1][j] = sum1
					(*result)[i*5+2][j] = sum2
					(*result)[i*5+3][j] = sum3
					(*result)[i*5+4][j] = sum4
				}
			}
		}(g)
	}
	wg.Wait()
}

func a10(A, B, result *[][]float64) {
	if len(*A) < 20 {
		a8(A, B, result)
		return
	}
	var wg sync.WaitGroup
	amount := 4
	amountBuffers := 20
	for g := 0; g < amount; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			n := len(*A)
			buffer0 := make([]float64, n)
			buffer1 := make([]float64, n)
			buffer2 := make([]float64, n)
			buffer3 := make([]float64, n)
			buffer4 := make([]float64, n)
			buffer5 := make([]float64, n)
			buffer6 := make([]float64, n)
			buffer7 := make([]float64, n)
			buffer8 := make([]float64, n)
			buffer9 := make([]float64, n)
			buffer10 := make([]float64, n)
			buffer11 := make([]float64, n)
			buffer12 := make([]float64, n)
			buffer13 := make([]float64, n)
			buffer14 := make([]float64, n)
			buffer15 := make([]float64, n)
			buffer16 := make([]float64, n)
			buffer17 := make([]float64, n)
			buffer18 := make([]float64, n)
			buffer19 := make([]float64, n)
			part := n / amountBuffers
			for i := g; i < part; i += amount {
				for k := 0; k < n; k++ {
					buffer0[k] = (*A)[i*amountBuffers+0][k]
					buffer1[k] = (*A)[i*amountBuffers+1][k]
					buffer2[k] = (*A)[i*amountBuffers+2][k]
					buffer3[k] = (*A)[i*amountBuffers+3][k]
					buffer4[k] = (*A)[i*amountBuffers+4][k]
					buffer5[k] = (*A)[i*amountBuffers+5][k]
					buffer6[k] = (*A)[i*amountBuffers+6][k]
					buffer7[k] = (*A)[i*amountBuffers+7][k]
					buffer8[k] = (*A)[i*amountBuffers+8][k]
					buffer9[k] = (*A)[i*amountBuffers+9][k]
					buffer10[k] = (*A)[i*amountBuffers+10][k]
					buffer11[k] = (*A)[i*amountBuffers+11][k]
					buffer12[k] = (*A)[i*amountBuffers+12][k]
					buffer13[k] = (*A)[i*amountBuffers+13][k]
					buffer14[k] = (*A)[i*amountBuffers+14][k]
					buffer15[k] = (*A)[i*amountBuffers+15][k]
					buffer16[k] = (*A)[i*amountBuffers+16][k]
					buffer17[k] = (*A)[i*amountBuffers+17][k]
					buffer18[k] = (*A)[i*amountBuffers+18][k]
					buffer19[k] = (*A)[i*amountBuffers+19][k]
				}
				for j := 0; j < n; j++ {
					sum0 := 0.0
					sum1 := 0.0
					sum2 := 0.0
					sum3 := 0.0
					sum4 := 0.0
					sum5 := 0.0
					sum6 := 0.0
					sum7 := 0.0
					sum8 := 0.0
					sum9 := 0.0
					sum10 := 0.0
					sum11 := 0.0
					sum12 := 0.0
					sum13 := 0.0
					sum14 := 0.0
					sum15 := 0.0
					sum16 := 0.0
					sum17 := 0.0
					sum18 := 0.0
					sum19 := 0.0
					for k := 0; k < n; k++ {
						sum0 += buffer0[k] * (*B)[k][j]
						sum1 += buffer1[k] * (*B)[k][j]
						sum2 += buffer2[k] * (*B)[k][j]
						sum3 += buffer3[k] * (*B)[k][j]
						sum4 += buffer4[k] * (*B)[k][j]
						sum5 += buffer5[k] * (*B)[k][j]
						sum6 += buffer6[k] * (*B)[k][j]
						sum7 += buffer7[k] * (*B)[k][j]
						sum8 += buffer8[k] * (*B)[k][j]
						sum9 += buffer9[k] * (*B)[k][j]
						sum10 += buffer10[k] * (*B)[k][j]
						sum11 += buffer11[k] * (*B)[k][j]
						sum12 += buffer12[k] * (*B)[k][j]
						sum13 += buffer13[k] * (*B)[k][j]
						sum14 += buffer14[k] * (*B)[k][j]
						sum15 += buffer15[k] * (*B)[k][j]
						sum16 += buffer16[k] * (*B)[k][j]
						sum17 += buffer17[k] * (*B)[k][j]
						sum18 += buffer18[k] * (*B)[k][j]
						sum19 += buffer19[k] * (*B)[k][j]
					}
					(*result)[i*amountBuffers+0][j] = sum0
					(*result)[i*amountBuffers+1][j] = sum1
					(*result)[i*amountBuffers+2][j] = sum2
					(*result)[i*amountBuffers+3][j] = sum3
					(*result)[i*amountBuffers+4][j] = sum4
					(*result)[i*amountBuffers+5][j] = sum5
					(*result)[i*amountBuffers+6][j] = sum6
					(*result)[i*amountBuffers+7][j] = sum7
					(*result)[i*amountBuffers+8][j] = sum8
					(*result)[i*amountBuffers+9][j] = sum9
					(*result)[i*amountBuffers+10][j] = sum10
					(*result)[i*amountBuffers+11][j] = sum11
					(*result)[i*amountBuffers+12][j] = sum12
					(*result)[i*amountBuffers+13][j] = sum13
					(*result)[i*amountBuffers+14][j] = sum14
					(*result)[i*amountBuffers+15][j] = sum15
					(*result)[i*amountBuffers+16][j] = sum16
					(*result)[i*amountBuffers+17][j] = sum17
					(*result)[i*amountBuffers+18][j] = sum18
					(*result)[i*amountBuffers+19][j] = sum19
				}
			}
		}(g)
	}
	wg.Wait()
}

func a11(A, B, result *[][]float64) {
	if len(*A) < 40 {
		a8(A, B, result)
		return
	}
	var wg sync.WaitGroup
	amount := 4
	amountBuffers := 30
	for g := 0; g < amount; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			n := len(*A)
			buffer0 := make([]float64, n)
			buffer1 := make([]float64, n)
			buffer2 := make([]float64, n)
			buffer3 := make([]float64, n)
			buffer4 := make([]float64, n)
			buffer5 := make([]float64, n)
			buffer6 := make([]float64, n)
			buffer7 := make([]float64, n)
			buffer8 := make([]float64, n)
			buffer9 := make([]float64, n)
			buffer10 := make([]float64, n)
			buffer11 := make([]float64, n)
			buffer12 := make([]float64, n)
			buffer13 := make([]float64, n)
			buffer14 := make([]float64, n)
			buffer15 := make([]float64, n)
			buffer16 := make([]float64, n)
			buffer17 := make([]float64, n)
			buffer18 := make([]float64, n)
			buffer19 := make([]float64, n)
			buffer20 := make([]float64, n)
			buffer21 := make([]float64, n)
			buffer22 := make([]float64, n)
			buffer23 := make([]float64, n)
			buffer24 := make([]float64, n)
			buffer25 := make([]float64, n)
			buffer26 := make([]float64, n)
			buffer27 := make([]float64, n)
			buffer28 := make([]float64, n)
			buffer29 := make([]float64, n)
			part := n / amountBuffers
			for i := g; i < part; i += amount {
				for k := 0; k < n; k++ {
					buffer0[k] = (*A)[i*amountBuffers+0][k]
					buffer1[k] = (*A)[i*amountBuffers+1][k]
					buffer2[k] = (*A)[i*amountBuffers+2][k]
					buffer3[k] = (*A)[i*amountBuffers+3][k]
					buffer4[k] = (*A)[i*amountBuffers+4][k]
					buffer5[k] = (*A)[i*amountBuffers+5][k]
					buffer6[k] = (*A)[i*amountBuffers+6][k]
					buffer7[k] = (*A)[i*amountBuffers+7][k]
					buffer8[k] = (*A)[i*amountBuffers+8][k]
					buffer9[k] = (*A)[i*amountBuffers+9][k]
					buffer10[k] = (*A)[i*amountBuffers+10][k]
					buffer11[k] = (*A)[i*amountBuffers+11][k]
					buffer12[k] = (*A)[i*amountBuffers+12][k]
					buffer13[k] = (*A)[i*amountBuffers+13][k]
					buffer14[k] = (*A)[i*amountBuffers+14][k]
					buffer15[k] = (*A)[i*amountBuffers+15][k]
					buffer16[k] = (*A)[i*amountBuffers+16][k]
					buffer17[k] = (*A)[i*amountBuffers+17][k]
					buffer18[k] = (*A)[i*amountBuffers+18][k]
					buffer19[k] = (*A)[i*amountBuffers+19][k]
					buffer20[k] = (*A)[i*amountBuffers+20][k]
					buffer21[k] = (*A)[i*amountBuffers+21][k]
					buffer22[k] = (*A)[i*amountBuffers+22][k]
					buffer23[k] = (*A)[i*amountBuffers+23][k]
					buffer24[k] = (*A)[i*amountBuffers+24][k]
					buffer25[k] = (*A)[i*amountBuffers+25][k]
					buffer26[k] = (*A)[i*amountBuffers+26][k]
					buffer27[k] = (*A)[i*amountBuffers+27][k]
					buffer28[k] = (*A)[i*amountBuffers+28][k]
					buffer29[k] = (*A)[i*amountBuffers+29][k]
				}
				for j := 0; j < n; j++ {
					sum0 := 0.0
					sum1 := 0.0
					sum2 := 0.0
					sum3 := 0.0
					sum4 := 0.0
					sum5 := 0.0
					sum6 := 0.0
					sum7 := 0.0
					sum8 := 0.0
					sum9 := 0.0
					sum10 := 0.0
					sum11 := 0.0
					sum12 := 0.0
					sum13 := 0.0
					sum14 := 0.0
					sum15 := 0.0
					sum16 := 0.0
					sum17 := 0.0
					sum18 := 0.0
					sum19 := 0.0
					sum20 := 0.0
					sum21 := 0.0
					sum22 := 0.0
					sum23 := 0.0
					sum24 := 0.0
					sum25 := 0.0
					sum26 := 0.0
					sum27 := 0.0
					sum28 := 0.0
					sum29 := 0.0
					for k := 0; k < n; k++ {
						sum0 += buffer0[k] * (*B)[k][j]
						sum1 += buffer1[k] * (*B)[k][j]
						sum2 += buffer2[k] * (*B)[k][j]
						sum3 += buffer3[k] * (*B)[k][j]
						sum4 += buffer4[k] * (*B)[k][j]
						sum5 += buffer5[k] * (*B)[k][j]
						sum6 += buffer6[k] * (*B)[k][j]
						sum7 += buffer7[k] * (*B)[k][j]
						sum8 += buffer8[k] * (*B)[k][j]
						sum9 += buffer9[k] * (*B)[k][j]
						sum10 += buffer10[k] * (*B)[k][j]
						sum11 += buffer11[k] * (*B)[k][j]
						sum12 += buffer12[k] * (*B)[k][j]
						sum13 += buffer13[k] * (*B)[k][j]
						sum14 += buffer14[k] * (*B)[k][j]
						sum15 += buffer15[k] * (*B)[k][j]
						sum16 += buffer16[k] * (*B)[k][j]
						sum17 += buffer17[k] * (*B)[k][j]
						sum18 += buffer18[k] * (*B)[k][j]
						sum19 += buffer19[k] * (*B)[k][j]
						sum20 += buffer20[k] * (*B)[k][j]
						sum21 += buffer21[k] * (*B)[k][j]
						sum22 += buffer22[k] * (*B)[k][j]
						sum23 += buffer23[k] * (*B)[k][j]
						sum24 += buffer24[k] * (*B)[k][j]
						sum25 += buffer25[k] * (*B)[k][j]
						sum26 += buffer26[k] * (*B)[k][j]
						sum27 += buffer27[k] * (*B)[k][j]
						sum28 += buffer28[k] * (*B)[k][j]
						sum29 += buffer29[k] * (*B)[k][j]
					}
					(*result)[i*amountBuffers+0][j] = sum0
					(*result)[i*amountBuffers+1][j] = sum1
					(*result)[i*amountBuffers+2][j] = sum2
					(*result)[i*amountBuffers+3][j] = sum3
					(*result)[i*amountBuffers+4][j] = sum4
					(*result)[i*amountBuffers+5][j] = sum5
					(*result)[i*amountBuffers+6][j] = sum6
					(*result)[i*amountBuffers+7][j] = sum7
					(*result)[i*amountBuffers+8][j] = sum8
					(*result)[i*amountBuffers+9][j] = sum9
					(*result)[i*amountBuffers+10][j] = sum10
					(*result)[i*amountBuffers+11][j] = sum11
					(*result)[i*amountBuffers+12][j] = sum12
					(*result)[i*amountBuffers+13][j] = sum13
					(*result)[i*amountBuffers+14][j] = sum14
					(*result)[i*amountBuffers+15][j] = sum15
					(*result)[i*amountBuffers+16][j] = sum16
					(*result)[i*amountBuffers+17][j] = sum17
					(*result)[i*amountBuffers+18][j] = sum18
					(*result)[i*amountBuffers+19][j] = sum19
					(*result)[i*amountBuffers+20][j] = sum20
					(*result)[i*amountBuffers+21][j] = sum21
					(*result)[i*amountBuffers+22][j] = sum22
					(*result)[i*amountBuffers+23][j] = sum23
					(*result)[i*amountBuffers+24][j] = sum24
					(*result)[i*amountBuffers+25][j] = sum25
					(*result)[i*amountBuffers+26][j] = sum26
					(*result)[i*amountBuffers+27][j] = sum27
					(*result)[i*amountBuffers+28][j] = sum28
					(*result)[i*amountBuffers+29][j] = sum29
				}
			}
		}(g)
	}
	wg.Wait()
}

func a12(A, B, result *[][]float64) {
	if len(*A) < 50 {
		a8(A, B, result)
		return
	}
	var wg sync.WaitGroup
	amount := 4
	amountBuffers := 50
	for g := 0; g < amount; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			n := len(*A)
			buffer0 := make([]float64, n)
			buffer1 := make([]float64, n)
			buffer2 := make([]float64, n)
			buffer3 := make([]float64, n)
			buffer4 := make([]float64, n)
			buffer5 := make([]float64, n)
			buffer6 := make([]float64, n)
			buffer7 := make([]float64, n)
			buffer8 := make([]float64, n)
			buffer9 := make([]float64, n)
			buffer10 := make([]float64, n)
			buffer11 := make([]float64, n)
			buffer12 := make([]float64, n)
			buffer13 := make([]float64, n)
			buffer14 := make([]float64, n)
			buffer15 := make([]float64, n)
			buffer16 := make([]float64, n)
			buffer17 := make([]float64, n)
			buffer18 := make([]float64, n)
			buffer19 := make([]float64, n)
			buffer20 := make([]float64, n)
			buffer21 := make([]float64, n)
			buffer22 := make([]float64, n)
			buffer23 := make([]float64, n)
			buffer24 := make([]float64, n)
			buffer25 := make([]float64, n)
			buffer26 := make([]float64, n)
			buffer27 := make([]float64, n)
			buffer28 := make([]float64, n)
			buffer29 := make([]float64, n)
			buffer30 := make([]float64, n)
			buffer31 := make([]float64, n)
			buffer32 := make([]float64, n)
			buffer33 := make([]float64, n)
			buffer34 := make([]float64, n)
			buffer35 := make([]float64, n)
			buffer36 := make([]float64, n)
			buffer37 := make([]float64, n)
			buffer38 := make([]float64, n)
			buffer39 := make([]float64, n)
			buffer40 := make([]float64, n)
			buffer41 := make([]float64, n)
			buffer42 := make([]float64, n)
			buffer43 := make([]float64, n)
			buffer44 := make([]float64, n)
			buffer45 := make([]float64, n)
			buffer46 := make([]float64, n)
			buffer47 := make([]float64, n)
			buffer48 := make([]float64, n)
			buffer49 := make([]float64, n)
			part := n / amountBuffers
			for i := g; i < part; i += amount {
				for k := 0; k < n; k++ {
					buffer0[k] = (*A)[i*amountBuffers+0][k]
					buffer1[k] = (*A)[i*amountBuffers+1][k]
					buffer2[k] = (*A)[i*amountBuffers+2][k]
					buffer3[k] = (*A)[i*amountBuffers+3][k]
					buffer4[k] = (*A)[i*amountBuffers+4][k]
					buffer5[k] = (*A)[i*amountBuffers+5][k]
					buffer6[k] = (*A)[i*amountBuffers+6][k]
					buffer7[k] = (*A)[i*amountBuffers+7][k]
					buffer8[k] = (*A)[i*amountBuffers+8][k]
					buffer9[k] = (*A)[i*amountBuffers+9][k]
					buffer10[k] = (*A)[i*amountBuffers+10][k]
					buffer11[k] = (*A)[i*amountBuffers+11][k]
					buffer12[k] = (*A)[i*amountBuffers+12][k]
					buffer13[k] = (*A)[i*amountBuffers+13][k]
					buffer14[k] = (*A)[i*amountBuffers+14][k]
					buffer15[k] = (*A)[i*amountBuffers+15][k]
					buffer16[k] = (*A)[i*amountBuffers+16][k]
					buffer17[k] = (*A)[i*amountBuffers+17][k]
					buffer18[k] = (*A)[i*amountBuffers+18][k]
					buffer19[k] = (*A)[i*amountBuffers+19][k]
					buffer20[k] = (*A)[i*amountBuffers+20][k]
					buffer21[k] = (*A)[i*amountBuffers+21][k]
					buffer22[k] = (*A)[i*amountBuffers+22][k]
					buffer23[k] = (*A)[i*amountBuffers+23][k]
					buffer24[k] = (*A)[i*amountBuffers+24][k]
					buffer25[k] = (*A)[i*amountBuffers+25][k]
					buffer26[k] = (*A)[i*amountBuffers+26][k]
					buffer27[k] = (*A)[i*amountBuffers+27][k]
					buffer28[k] = (*A)[i*amountBuffers+28][k]
					buffer29[k] = (*A)[i*amountBuffers+29][k]
					buffer30[k] = (*A)[i*amountBuffers+30][k]
					buffer31[k] = (*A)[i*amountBuffers+31][k]
					buffer32[k] = (*A)[i*amountBuffers+32][k]
					buffer33[k] = (*A)[i*amountBuffers+33][k]
					buffer34[k] = (*A)[i*amountBuffers+34][k]
					buffer35[k] = (*A)[i*amountBuffers+35][k]
					buffer36[k] = (*A)[i*amountBuffers+36][k]
					buffer37[k] = (*A)[i*amountBuffers+37][k]
					buffer38[k] = (*A)[i*amountBuffers+38][k]
					buffer39[k] = (*A)[i*amountBuffers+39][k]
					buffer40[k] = (*A)[i*amountBuffers+40][k]
					buffer41[k] = (*A)[i*amountBuffers+41][k]
					buffer42[k] = (*A)[i*amountBuffers+42][k]
					buffer43[k] = (*A)[i*amountBuffers+43][k]
					buffer44[k] = (*A)[i*amountBuffers+44][k]
					buffer45[k] = (*A)[i*amountBuffers+45][k]
					buffer46[k] = (*A)[i*amountBuffers+46][k]
					buffer47[k] = (*A)[i*amountBuffers+47][k]
					buffer48[k] = (*A)[i*amountBuffers+48][k]
					buffer49[k] = (*A)[i*amountBuffers+49][k]
				}
				for j := 0; j < n; j++ {
					sum0 := 0.0
					sum1 := 0.0
					sum2 := 0.0
					sum3 := 0.0
					sum4 := 0.0
					sum5 := 0.0
					sum6 := 0.0
					sum7 := 0.0
					sum8 := 0.0
					sum9 := 0.0
					sum10 := 0.0
					sum11 := 0.0
					sum12 := 0.0
					sum13 := 0.0
					sum14 := 0.0
					sum15 := 0.0
					sum16 := 0.0
					sum17 := 0.0
					sum18 := 0.0
					sum19 := 0.0
					sum20 := 0.0
					sum21 := 0.0
					sum22 := 0.0
					sum23 := 0.0
					sum24 := 0.0
					sum25 := 0.0
					sum26 := 0.0
					sum27 := 0.0
					sum28 := 0.0
					sum29 := 0.0
					sum30 := 0.0
					sum31 := 0.0
					sum32 := 0.0
					sum33 := 0.0
					sum34 := 0.0
					sum35 := 0.0
					sum36 := 0.0
					sum37 := 0.0
					sum38 := 0.0
					sum39 := 0.0
					sum40 := 0.0
					sum41 := 0.0
					sum42 := 0.0
					sum43 := 0.0
					sum44 := 0.0
					sum45 := 0.0
					sum46 := 0.0
					sum47 := 0.0
					sum48 := 0.0
					sum49 := 0.0
					for k := 0; k < n; k++ {
						sum0 += buffer0[k] * (*B)[k][j]
						sum1 += buffer1[k] * (*B)[k][j]
						sum2 += buffer2[k] * (*B)[k][j]
						sum3 += buffer3[k] * (*B)[k][j]
						sum4 += buffer4[k] * (*B)[k][j]
						sum5 += buffer5[k] * (*B)[k][j]
						sum6 += buffer6[k] * (*B)[k][j]
						sum7 += buffer7[k] * (*B)[k][j]
						sum8 += buffer8[k] * (*B)[k][j]
						sum9 += buffer9[k] * (*B)[k][j]
						sum10 += buffer10[k] * (*B)[k][j]
						sum11 += buffer11[k] * (*B)[k][j]
						sum12 += buffer12[k] * (*B)[k][j]
						sum13 += buffer13[k] * (*B)[k][j]
						sum14 += buffer14[k] * (*B)[k][j]
						sum15 += buffer15[k] * (*B)[k][j]
						sum16 += buffer16[k] * (*B)[k][j]
						sum17 += buffer17[k] * (*B)[k][j]
						sum18 += buffer18[k] * (*B)[k][j]
						sum19 += buffer19[k] * (*B)[k][j]
						sum20 += buffer20[k] * (*B)[k][j]
						sum21 += buffer21[k] * (*B)[k][j]
						sum22 += buffer22[k] * (*B)[k][j]
						sum23 += buffer23[k] * (*B)[k][j]
						sum24 += buffer24[k] * (*B)[k][j]
						sum25 += buffer25[k] * (*B)[k][j]
						sum26 += buffer26[k] * (*B)[k][j]
						sum27 += buffer27[k] * (*B)[k][j]
						sum28 += buffer28[k] * (*B)[k][j]
						sum29 += buffer29[k] * (*B)[k][j]
						sum30 += buffer30[k] * (*B)[k][j]
						sum31 += buffer31[k] * (*B)[k][j]
						sum32 += buffer32[k] * (*B)[k][j]
						sum33 += buffer33[k] * (*B)[k][j]
						sum34 += buffer34[k] * (*B)[k][j]
						sum35 += buffer35[k] * (*B)[k][j]
						sum36 += buffer36[k] * (*B)[k][j]
						sum37 += buffer37[k] * (*B)[k][j]
						sum38 += buffer38[k] * (*B)[k][j]
						sum39 += buffer39[k] * (*B)[k][j]
						sum40 += buffer40[k] * (*B)[k][j]
						sum41 += buffer41[k] * (*B)[k][j]
						sum42 += buffer42[k] * (*B)[k][j]
						sum43 += buffer43[k] * (*B)[k][j]
						sum44 += buffer44[k] * (*B)[k][j]
						sum45 += buffer45[k] * (*B)[k][j]
						sum46 += buffer46[k] * (*B)[k][j]
						sum47 += buffer47[k] * (*B)[k][j]
						sum48 += buffer48[k] * (*B)[k][j]
						sum49 += buffer49[k] * (*B)[k][j]
					}
					(*result)[i*amountBuffers+0][j] = sum0
					(*result)[i*amountBuffers+1][j] = sum1
					(*result)[i*amountBuffers+2][j] = sum2
					(*result)[i*amountBuffers+3][j] = sum3
					(*result)[i*amountBuffers+4][j] = sum4
					(*result)[i*amountBuffers+5][j] = sum5
					(*result)[i*amountBuffers+6][j] = sum6
					(*result)[i*amountBuffers+7][j] = sum7
					(*result)[i*amountBuffers+8][j] = sum8
					(*result)[i*amountBuffers+9][j] = sum9
					(*result)[i*amountBuffers+10][j] = sum10
					(*result)[i*amountBuffers+11][j] = sum11
					(*result)[i*amountBuffers+12][j] = sum12
					(*result)[i*amountBuffers+13][j] = sum13
					(*result)[i*amountBuffers+14][j] = sum14
					(*result)[i*amountBuffers+15][j] = sum15
					(*result)[i*amountBuffers+16][j] = sum16
					(*result)[i*amountBuffers+17][j] = sum17
					(*result)[i*amountBuffers+18][j] = sum18
					(*result)[i*amountBuffers+19][j] = sum19
					(*result)[i*amountBuffers+20][j] = sum20
					(*result)[i*amountBuffers+21][j] = sum21
					(*result)[i*amountBuffers+22][j] = sum22
					(*result)[i*amountBuffers+23][j] = sum23
					(*result)[i*amountBuffers+24][j] = sum24
					(*result)[i*amountBuffers+25][j] = sum25
					(*result)[i*amountBuffers+26][j] = sum26
					(*result)[i*amountBuffers+27][j] = sum27
					(*result)[i*amountBuffers+28][j] = sum28
					(*result)[i*amountBuffers+29][j] = sum29
					(*result)[i*amountBuffers+30][j] = sum30
					(*result)[i*amountBuffers+31][j] = sum31
					(*result)[i*amountBuffers+32][j] = sum32
					(*result)[i*amountBuffers+33][j] = sum33
					(*result)[i*amountBuffers+34][j] = sum34
					(*result)[i*amountBuffers+35][j] = sum35
					(*result)[i*amountBuffers+36][j] = sum36
					(*result)[i*amountBuffers+37][j] = sum37
					(*result)[i*amountBuffers+38][j] = sum38
					(*result)[i*amountBuffers+39][j] = sum39
					(*result)[i*amountBuffers+40][j] = sum40
					(*result)[i*amountBuffers+41][j] = sum41
					(*result)[i*amountBuffers+42][j] = sum42
					(*result)[i*amountBuffers+43][j] = sum43
					(*result)[i*amountBuffers+44][j] = sum44
					(*result)[i*amountBuffers+45][j] = sum45
					(*result)[i*amountBuffers+46][j] = sum46
					(*result)[i*amountBuffers+47][j] = sum47
					(*result)[i*amountBuffers+48][j] = sum48
					(*result)[i*amountBuffers+49][j] = sum49
				}
			}
		}(g)
	}
	wg.Wait()
}

func a13(A, B, result *[][]float64) {
	if len(*A) < 60 {
		a8(A, B, result)
		return
	}
	var wg sync.WaitGroup
	amount := 4
	amountBuffers := 60
	for g := 0; g < amount; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			n := len(*A)
			buffer0 := make([]float64, n)
			buffer1 := make([]float64, n)
			buffer2 := make([]float64, n)
			buffer3 := make([]float64, n)
			buffer4 := make([]float64, n)
			buffer5 := make([]float64, n)
			buffer6 := make([]float64, n)
			buffer7 := make([]float64, n)
			buffer8 := make([]float64, n)
			buffer9 := make([]float64, n)
			buffer10 := make([]float64, n)
			buffer11 := make([]float64, n)
			buffer12 := make([]float64, n)
			buffer13 := make([]float64, n)
			buffer14 := make([]float64, n)
			buffer15 := make([]float64, n)
			buffer16 := make([]float64, n)
			buffer17 := make([]float64, n)
			buffer18 := make([]float64, n)
			buffer19 := make([]float64, n)
			buffer20 := make([]float64, n)
			buffer21 := make([]float64, n)
			buffer22 := make([]float64, n)
			buffer23 := make([]float64, n)
			buffer24 := make([]float64, n)
			buffer25 := make([]float64, n)
			buffer26 := make([]float64, n)
			buffer27 := make([]float64, n)
			buffer28 := make([]float64, n)
			buffer29 := make([]float64, n)
			buffer30 := make([]float64, n)
			buffer31 := make([]float64, n)
			buffer32 := make([]float64, n)
			buffer33 := make([]float64, n)
			buffer34 := make([]float64, n)
			buffer35 := make([]float64, n)
			buffer36 := make([]float64, n)
			buffer37 := make([]float64, n)
			buffer38 := make([]float64, n)
			buffer39 := make([]float64, n)
			buffer40 := make([]float64, n)
			buffer41 := make([]float64, n)
			buffer42 := make([]float64, n)
			buffer43 := make([]float64, n)
			buffer44 := make([]float64, n)
			buffer45 := make([]float64, n)
			buffer46 := make([]float64, n)
			buffer47 := make([]float64, n)
			buffer48 := make([]float64, n)
			buffer49 := make([]float64, n)
			buffer50 := make([]float64, n)
			buffer51 := make([]float64, n)
			buffer52 := make([]float64, n)
			buffer53 := make([]float64, n)
			buffer54 := make([]float64, n)
			buffer55 := make([]float64, n)
			buffer56 := make([]float64, n)
			buffer57 := make([]float64, n)
			buffer58 := make([]float64, n)
			buffer59 := make([]float64, n)
			part := n / amountBuffers
			for i := g; i < part; i += amount {
				for k := 0; k < n; k++ {
					buffer0[k] = (*A)[i*amountBuffers+0][k]
					buffer1[k] = (*A)[i*amountBuffers+1][k]
					buffer2[k] = (*A)[i*amountBuffers+2][k]
					buffer3[k] = (*A)[i*amountBuffers+3][k]
					buffer4[k] = (*A)[i*amountBuffers+4][k]
					buffer5[k] = (*A)[i*amountBuffers+5][k]
					buffer6[k] = (*A)[i*amountBuffers+6][k]
					buffer7[k] = (*A)[i*amountBuffers+7][k]
					buffer8[k] = (*A)[i*amountBuffers+8][k]
					buffer9[k] = (*A)[i*amountBuffers+9][k]
					buffer10[k] = (*A)[i*amountBuffers+10][k]
					buffer11[k] = (*A)[i*amountBuffers+11][k]
					buffer12[k] = (*A)[i*amountBuffers+12][k]
					buffer13[k] = (*A)[i*amountBuffers+13][k]
					buffer14[k] = (*A)[i*amountBuffers+14][k]
					buffer15[k] = (*A)[i*amountBuffers+15][k]
					buffer16[k] = (*A)[i*amountBuffers+16][k]
					buffer17[k] = (*A)[i*amountBuffers+17][k]
					buffer18[k] = (*A)[i*amountBuffers+18][k]
					buffer19[k] = (*A)[i*amountBuffers+19][k]
					buffer20[k] = (*A)[i*amountBuffers+20][k]
					buffer21[k] = (*A)[i*amountBuffers+21][k]
					buffer22[k] = (*A)[i*amountBuffers+22][k]
					buffer23[k] = (*A)[i*amountBuffers+23][k]
					buffer24[k] = (*A)[i*amountBuffers+24][k]
					buffer25[k] = (*A)[i*amountBuffers+25][k]
					buffer26[k] = (*A)[i*amountBuffers+26][k]
					buffer27[k] = (*A)[i*amountBuffers+27][k]
					buffer28[k] = (*A)[i*amountBuffers+28][k]
					buffer29[k] = (*A)[i*amountBuffers+29][k]
					buffer30[k] = (*A)[i*amountBuffers+30][k]
					buffer31[k] = (*A)[i*amountBuffers+31][k]
					buffer32[k] = (*A)[i*amountBuffers+32][k]
					buffer33[k] = (*A)[i*amountBuffers+33][k]
					buffer34[k] = (*A)[i*amountBuffers+34][k]
					buffer35[k] = (*A)[i*amountBuffers+35][k]
					buffer36[k] = (*A)[i*amountBuffers+36][k]
					buffer37[k] = (*A)[i*amountBuffers+37][k]
					buffer38[k] = (*A)[i*amountBuffers+38][k]
					buffer39[k] = (*A)[i*amountBuffers+39][k]
					buffer40[k] = (*A)[i*amountBuffers+40][k]
					buffer41[k] = (*A)[i*amountBuffers+41][k]
					buffer42[k] = (*A)[i*amountBuffers+42][k]
					buffer43[k] = (*A)[i*amountBuffers+43][k]
					buffer44[k] = (*A)[i*amountBuffers+44][k]
					buffer45[k] = (*A)[i*amountBuffers+45][k]
					buffer46[k] = (*A)[i*amountBuffers+46][k]
					buffer47[k] = (*A)[i*amountBuffers+47][k]
					buffer48[k] = (*A)[i*amountBuffers+48][k]
					buffer49[k] = (*A)[i*amountBuffers+49][k]
					buffer50[k] = (*A)[i*amountBuffers+50][k]
					buffer51[k] = (*A)[i*amountBuffers+51][k]
					buffer52[k] = (*A)[i*amountBuffers+52][k]
					buffer53[k] = (*A)[i*amountBuffers+53][k]
					buffer54[k] = (*A)[i*amountBuffers+54][k]
					buffer55[k] = (*A)[i*amountBuffers+55][k]
					buffer56[k] = (*A)[i*amountBuffers+56][k]
					buffer57[k] = (*A)[i*amountBuffers+57][k]
					buffer58[k] = (*A)[i*amountBuffers+58][k]
					buffer59[k] = (*A)[i*amountBuffers+59][k]
				}
				for j := 0; j < n; j++ {
					sum0 := 0.0
					sum1 := 0.0
					sum2 := 0.0
					sum3 := 0.0
					sum4 := 0.0
					sum5 := 0.0
					sum6 := 0.0
					sum7 := 0.0
					sum8 := 0.0
					sum9 := 0.0
					sum10 := 0.0
					sum11 := 0.0
					sum12 := 0.0
					sum13 := 0.0
					sum14 := 0.0
					sum15 := 0.0
					sum16 := 0.0
					sum17 := 0.0
					sum18 := 0.0
					sum19 := 0.0
					sum20 := 0.0
					sum21 := 0.0
					sum22 := 0.0
					sum23 := 0.0
					sum24 := 0.0
					sum25 := 0.0
					sum26 := 0.0
					sum27 := 0.0
					sum28 := 0.0
					sum29 := 0.0
					sum30 := 0.0
					sum31 := 0.0
					sum32 := 0.0
					sum33 := 0.0
					sum34 := 0.0
					sum35 := 0.0
					sum36 := 0.0
					sum37 := 0.0
					sum38 := 0.0
					sum39 := 0.0
					sum40 := 0.0
					sum41 := 0.0
					sum42 := 0.0
					sum43 := 0.0
					sum44 := 0.0
					sum45 := 0.0
					sum46 := 0.0
					sum47 := 0.0
					sum48 := 0.0
					sum49 := 0.0
					sum50 := 0.0
					sum51 := 0.0
					sum52 := 0.0
					sum53 := 0.0
					sum54 := 0.0
					sum55 := 0.0
					sum56 := 0.0
					sum57 := 0.0
					sum58 := 0.0
					sum59 := 0.0
					for k := 0; k < n; k++ {
						sum0 += buffer0[k] * (*B)[k][j]
						sum1 += buffer1[k] * (*B)[k][j]
						sum2 += buffer2[k] * (*B)[k][j]
						sum3 += buffer3[k] * (*B)[k][j]
						sum4 += buffer4[k] * (*B)[k][j]
						sum5 += buffer5[k] * (*B)[k][j]
						sum6 += buffer6[k] * (*B)[k][j]
						sum7 += buffer7[k] * (*B)[k][j]
						sum8 += buffer8[k] * (*B)[k][j]
						sum9 += buffer9[k] * (*B)[k][j]
						sum10 += buffer10[k] * (*B)[k][j]
						sum11 += buffer11[k] * (*B)[k][j]
						sum12 += buffer12[k] * (*B)[k][j]
						sum13 += buffer13[k] * (*B)[k][j]
						sum14 += buffer14[k] * (*B)[k][j]
						sum15 += buffer15[k] * (*B)[k][j]
						sum16 += buffer16[k] * (*B)[k][j]
						sum17 += buffer17[k] * (*B)[k][j]
						sum18 += buffer18[k] * (*B)[k][j]
						sum19 += buffer19[k] * (*B)[k][j]
						sum20 += buffer20[k] * (*B)[k][j]
						sum21 += buffer21[k] * (*B)[k][j]
						sum22 += buffer22[k] * (*B)[k][j]
						sum23 += buffer23[k] * (*B)[k][j]
						sum24 += buffer24[k] * (*B)[k][j]
						sum25 += buffer25[k] * (*B)[k][j]
						sum26 += buffer26[k] * (*B)[k][j]
						sum27 += buffer27[k] * (*B)[k][j]
						sum28 += buffer28[k] * (*B)[k][j]
						sum29 += buffer29[k] * (*B)[k][j]
						sum30 += buffer30[k] * (*B)[k][j]
						sum31 += buffer31[k] * (*B)[k][j]
						sum32 += buffer32[k] * (*B)[k][j]
						sum33 += buffer33[k] * (*B)[k][j]
						sum34 += buffer34[k] * (*B)[k][j]
						sum35 += buffer35[k] * (*B)[k][j]
						sum36 += buffer36[k] * (*B)[k][j]
						sum37 += buffer37[k] * (*B)[k][j]
						sum38 += buffer38[k] * (*B)[k][j]
						sum39 += buffer39[k] * (*B)[k][j]
						sum40 += buffer40[k] * (*B)[k][j]
						sum41 += buffer41[k] * (*B)[k][j]
						sum42 += buffer42[k] * (*B)[k][j]
						sum43 += buffer43[k] * (*B)[k][j]
						sum44 += buffer44[k] * (*B)[k][j]
						sum45 += buffer45[k] * (*B)[k][j]
						sum46 += buffer46[k] * (*B)[k][j]
						sum47 += buffer47[k] * (*B)[k][j]
						sum48 += buffer48[k] * (*B)[k][j]
						sum49 += buffer49[k] * (*B)[k][j]
						sum50 += buffer50[k] * (*B)[k][j]
						sum51 += buffer51[k] * (*B)[k][j]
						sum52 += buffer52[k] * (*B)[k][j]
						sum53 += buffer53[k] * (*B)[k][j]
						sum54 += buffer54[k] * (*B)[k][j]
						sum55 += buffer55[k] * (*B)[k][j]
						sum56 += buffer56[k] * (*B)[k][j]
						sum57 += buffer57[k] * (*B)[k][j]
						sum58 += buffer58[k] * (*B)[k][j]
						sum59 += buffer59[k] * (*B)[k][j]
					}
					(*result)[i*amountBuffers+0][j] = sum0
					(*result)[i*amountBuffers+1][j] = sum1
					(*result)[i*amountBuffers+2][j] = sum2
					(*result)[i*amountBuffers+3][j] = sum3
					(*result)[i*amountBuffers+4][j] = sum4
					(*result)[i*amountBuffers+5][j] = sum5
					(*result)[i*amountBuffers+6][j] = sum6
					(*result)[i*amountBuffers+7][j] = sum7
					(*result)[i*amountBuffers+8][j] = sum8
					(*result)[i*amountBuffers+9][j] = sum9
					(*result)[i*amountBuffers+10][j] = sum10
					(*result)[i*amountBuffers+11][j] = sum11
					(*result)[i*amountBuffers+12][j] = sum12
					(*result)[i*amountBuffers+13][j] = sum13
					(*result)[i*amountBuffers+14][j] = sum14
					(*result)[i*amountBuffers+15][j] = sum15
					(*result)[i*amountBuffers+16][j] = sum16
					(*result)[i*amountBuffers+17][j] = sum17
					(*result)[i*amountBuffers+18][j] = sum18
					(*result)[i*amountBuffers+19][j] = sum19
					(*result)[i*amountBuffers+20][j] = sum20
					(*result)[i*amountBuffers+21][j] = sum21
					(*result)[i*amountBuffers+22][j] = sum22
					(*result)[i*amountBuffers+23][j] = sum23
					(*result)[i*amountBuffers+24][j] = sum24
					(*result)[i*amountBuffers+25][j] = sum25
					(*result)[i*amountBuffers+26][j] = sum26
					(*result)[i*amountBuffers+27][j] = sum27
					(*result)[i*amountBuffers+28][j] = sum28
					(*result)[i*amountBuffers+29][j] = sum29
					(*result)[i*amountBuffers+30][j] = sum30
					(*result)[i*amountBuffers+31][j] = sum31
					(*result)[i*amountBuffers+32][j] = sum32
					(*result)[i*amountBuffers+33][j] = sum33
					(*result)[i*amountBuffers+34][j] = sum34
					(*result)[i*amountBuffers+35][j] = sum35
					(*result)[i*amountBuffers+36][j] = sum36
					(*result)[i*amountBuffers+37][j] = sum37
					(*result)[i*amountBuffers+38][j] = sum38
					(*result)[i*amountBuffers+39][j] = sum39
					(*result)[i*amountBuffers+40][j] = sum40
					(*result)[i*amountBuffers+41][j] = sum41
					(*result)[i*amountBuffers+42][j] = sum42
					(*result)[i*amountBuffers+43][j] = sum43
					(*result)[i*amountBuffers+44][j] = sum44
					(*result)[i*amountBuffers+45][j] = sum45
					(*result)[i*amountBuffers+46][j] = sum46
					(*result)[i*amountBuffers+47][j] = sum47
					(*result)[i*amountBuffers+48][j] = sum48
					(*result)[i*amountBuffers+49][j] = sum49
					(*result)[i*amountBuffers+50][j] = sum50
					(*result)[i*amountBuffers+51][j] = sum51
					(*result)[i*amountBuffers+52][j] = sum52
					(*result)[i*amountBuffers+53][j] = sum53
					(*result)[i*amountBuffers+54][j] = sum54
					(*result)[i*amountBuffers+55][j] = sum55
					(*result)[i*amountBuffers+56][j] = sum56
					(*result)[i*amountBuffers+57][j] = sum57
					(*result)[i*amountBuffers+58][j] = sum58
					(*result)[i*amountBuffers+59][j] = sum59
				}
			}
		}(g)
	}
	wg.Wait()
}

func a13a(A, B, result *[][]float64) {
	if len(*A) < 60 {
		a8(A, B, result)
		return
	}
	var wg sync.WaitGroup
	amount := 4
	amountBuffers := 60
	for g := 0; g < amount; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			n := len(*A)
			buffer0 := make([]float64, n)
			buffer1 := make([]float64, n)
			buffer2 := make([]float64, n)
			buffer3 := make([]float64, n)
			buffer4 := make([]float64, n)
			buffer5 := make([]float64, n)
			buffer6 := make([]float64, n)
			buffer7 := make([]float64, n)
			buffer8 := make([]float64, n)
			buffer9 := make([]float64, n)
			buffer10 := make([]float64, n)
			buffer11 := make([]float64, n)
			buffer12 := make([]float64, n)
			buffer13 := make([]float64, n)
			buffer14 := make([]float64, n)
			buffer15 := make([]float64, n)
			buffer16 := make([]float64, n)
			buffer17 := make([]float64, n)
			buffer18 := make([]float64, n)
			buffer19 := make([]float64, n)
			buffer20 := make([]float64, n)
			buffer21 := make([]float64, n)
			buffer22 := make([]float64, n)
			buffer23 := make([]float64, n)
			buffer24 := make([]float64, n)
			buffer25 := make([]float64, n)
			buffer26 := make([]float64, n)
			buffer27 := make([]float64, n)
			buffer28 := make([]float64, n)
			buffer29 := make([]float64, n)
			buffer30 := make([]float64, n)
			buffer31 := make([]float64, n)
			buffer32 := make([]float64, n)
			buffer33 := make([]float64, n)
			buffer34 := make([]float64, n)
			buffer35 := make([]float64, n)
			buffer36 := make([]float64, n)
			buffer37 := make([]float64, n)
			buffer38 := make([]float64, n)
			buffer39 := make([]float64, n)
			buffer40 := make([]float64, n)
			buffer41 := make([]float64, n)
			buffer42 := make([]float64, n)
			buffer43 := make([]float64, n)
			buffer44 := make([]float64, n)
			buffer45 := make([]float64, n)
			buffer46 := make([]float64, n)
			buffer47 := make([]float64, n)
			buffer48 := make([]float64, n)
			buffer49 := make([]float64, n)
			buffer50 := make([]float64, n)
			buffer51 := make([]float64, n)
			buffer52 := make([]float64, n)
			buffer53 := make([]float64, n)
			buffer54 := make([]float64, n)
			buffer55 := make([]float64, n)
			buffer56 := make([]float64, n)
			buffer57 := make([]float64, n)
			buffer58 := make([]float64, n)
			buffer59 := make([]float64, n)
			part := n / amountBuffers
			var sum0, sum1, sum2, sum3, sum4, sum5, sum6, sum7, sum8, sum9 float64
			var sum10, sum11, sum12, sum13, sum14, sum15, sum16, sum17, sum18, sum19 float64
			var sum20, sum21, sum22, sum23, sum24, sum25, sum26, sum27, sum28, sum29 float64
			var sum30, sum31, sum32, sum33, sum34, sum35, sum36, sum37, sum38, sum39 float64
			var sum40, sum41, sum42, sum43, sum44, sum45, sum46, sum47, sum48, sum49 float64
			var sum50, sum51, sum52, sum53, sum54, sum55, sum56, sum57, sum58, sum59 float64
			for i := g; i < part; i += amount {
				for k := 0; k < n; k++ {
					buffer0[k] = (*A)[i*amountBuffers+0][k]
					buffer1[k] = (*A)[i*amountBuffers+1][k]
					buffer2[k] = (*A)[i*amountBuffers+2][k]
					buffer3[k] = (*A)[i*amountBuffers+3][k]
					buffer4[k] = (*A)[i*amountBuffers+4][k]
					buffer5[k] = (*A)[i*amountBuffers+5][k]
					buffer6[k] = (*A)[i*amountBuffers+6][k]
					buffer7[k] = (*A)[i*amountBuffers+7][k]
					buffer8[k] = (*A)[i*amountBuffers+8][k]
					buffer9[k] = (*A)[i*amountBuffers+9][k]
					buffer10[k] = (*A)[i*amountBuffers+10][k]
					buffer11[k] = (*A)[i*amountBuffers+11][k]
					buffer12[k] = (*A)[i*amountBuffers+12][k]
					buffer13[k] = (*A)[i*amountBuffers+13][k]
					buffer14[k] = (*A)[i*amountBuffers+14][k]
					buffer15[k] = (*A)[i*amountBuffers+15][k]
					buffer16[k] = (*A)[i*amountBuffers+16][k]
					buffer17[k] = (*A)[i*amountBuffers+17][k]
					buffer18[k] = (*A)[i*amountBuffers+18][k]
					buffer19[k] = (*A)[i*amountBuffers+19][k]
					buffer20[k] = (*A)[i*amountBuffers+20][k]
					buffer21[k] = (*A)[i*amountBuffers+21][k]
					buffer22[k] = (*A)[i*amountBuffers+22][k]
					buffer23[k] = (*A)[i*amountBuffers+23][k]
					buffer24[k] = (*A)[i*amountBuffers+24][k]
					buffer25[k] = (*A)[i*amountBuffers+25][k]
					buffer26[k] = (*A)[i*amountBuffers+26][k]
					buffer27[k] = (*A)[i*amountBuffers+27][k]
					buffer28[k] = (*A)[i*amountBuffers+28][k]
					buffer29[k] = (*A)[i*amountBuffers+29][k]
					buffer30[k] = (*A)[i*amountBuffers+30][k]
					buffer31[k] = (*A)[i*amountBuffers+31][k]
					buffer32[k] = (*A)[i*amountBuffers+32][k]
					buffer33[k] = (*A)[i*amountBuffers+33][k]
					buffer34[k] = (*A)[i*amountBuffers+34][k]
					buffer35[k] = (*A)[i*amountBuffers+35][k]
					buffer36[k] = (*A)[i*amountBuffers+36][k]
					buffer37[k] = (*A)[i*amountBuffers+37][k]
					buffer38[k] = (*A)[i*amountBuffers+38][k]
					buffer39[k] = (*A)[i*amountBuffers+39][k]
					buffer40[k] = (*A)[i*amountBuffers+40][k]
					buffer41[k] = (*A)[i*amountBuffers+41][k]
					buffer42[k] = (*A)[i*amountBuffers+42][k]
					buffer43[k] = (*A)[i*amountBuffers+43][k]
					buffer44[k] = (*A)[i*amountBuffers+44][k]
					buffer45[k] = (*A)[i*amountBuffers+45][k]
					buffer46[k] = (*A)[i*amountBuffers+46][k]
					buffer47[k] = (*A)[i*amountBuffers+47][k]
					buffer48[k] = (*A)[i*amountBuffers+48][k]
					buffer49[k] = (*A)[i*amountBuffers+49][k]
					buffer50[k] = (*A)[i*amountBuffers+50][k]
					buffer51[k] = (*A)[i*amountBuffers+51][k]
					buffer52[k] = (*A)[i*amountBuffers+52][k]
					buffer53[k] = (*A)[i*amountBuffers+53][k]
					buffer54[k] = (*A)[i*amountBuffers+54][k]
					buffer55[k] = (*A)[i*amountBuffers+55][k]
					buffer56[k] = (*A)[i*amountBuffers+56][k]
					buffer57[k] = (*A)[i*amountBuffers+57][k]
					buffer58[k] = (*A)[i*amountBuffers+58][k]
					buffer59[k] = (*A)[i*amountBuffers+59][k]
				}
				for j := 0; j < n; j++ {
					sum0 = 0.0
					sum1 = 0.0
					sum2 = 0.0
					sum3 = 0.0
					sum4 = 0.0
					sum5 = 0.0
					sum6 = 0.0
					sum7 = 0.0
					sum8 = 0.0
					sum9 = 0.0
					sum10 = 0.0
					sum11 = 0.0
					sum12 = 0.0
					sum13 = 0.0
					sum14 = 0.0
					sum15 = 0.0
					sum16 = 0.0
					sum17 = 0.0
					sum18 = 0.0
					sum19 = 0.0
					sum20 = 0.0
					sum21 = 0.0
					sum22 = 0.0
					sum23 = 0.0
					sum24 = 0.0
					sum25 = 0.0
					sum26 = 0.0
					sum27 = 0.0
					sum28 = 0.0
					sum29 = 0.0
					sum30 = 0.0
					sum31 = 0.0
					sum32 = 0.0
					sum33 = 0.0
					sum34 = 0.0
					sum35 = 0.0
					sum36 = 0.0
					sum37 = 0.0
					sum38 = 0.0
					sum39 = 0.0
					sum40 = 0.0
					sum41 = 0.0
					sum42 = 0.0
					sum43 = 0.0
					sum44 = 0.0
					sum45 = 0.0
					sum46 = 0.0
					sum47 = 0.0
					sum48 = 0.0
					sum49 = 0.0
					sum50 = 0.0
					sum51 = 0.0
					sum52 = 0.0
					sum53 = 0.0
					sum54 = 0.0
					sum55 = 0.0
					sum56 = 0.0
					sum57 = 0.0
					sum58 = 0.0
					sum59 = 0.0
					for k := 0; k < n; k++ {
						sum0 += buffer0[k] * (*B)[k][j]
						sum1 += buffer1[k] * (*B)[k][j]
						sum2 += buffer2[k] * (*B)[k][j]
						sum3 += buffer3[k] * (*B)[k][j]
						sum4 += buffer4[k] * (*B)[k][j]
						sum5 += buffer5[k] * (*B)[k][j]
						sum6 += buffer6[k] * (*B)[k][j]
						sum7 += buffer7[k] * (*B)[k][j]
						sum8 += buffer8[k] * (*B)[k][j]
						sum9 += buffer9[k] * (*B)[k][j]
						sum10 += buffer10[k] * (*B)[k][j]
						sum11 += buffer11[k] * (*B)[k][j]
						sum12 += buffer12[k] * (*B)[k][j]
						sum13 += buffer13[k] * (*B)[k][j]
						sum14 += buffer14[k] * (*B)[k][j]
						sum15 += buffer15[k] * (*B)[k][j]
						sum16 += buffer16[k] * (*B)[k][j]
						sum17 += buffer17[k] * (*B)[k][j]
						sum18 += buffer18[k] * (*B)[k][j]
						sum19 += buffer19[k] * (*B)[k][j]
						sum20 += buffer20[k] * (*B)[k][j]
						sum21 += buffer21[k] * (*B)[k][j]
						sum22 += buffer22[k] * (*B)[k][j]
						sum23 += buffer23[k] * (*B)[k][j]
						sum24 += buffer24[k] * (*B)[k][j]
						sum25 += buffer25[k] * (*B)[k][j]
						sum26 += buffer26[k] * (*B)[k][j]
						sum27 += buffer27[k] * (*B)[k][j]
						sum28 += buffer28[k] * (*B)[k][j]
						sum29 += buffer29[k] * (*B)[k][j]
						sum30 += buffer30[k] * (*B)[k][j]
						sum31 += buffer31[k] * (*B)[k][j]
						sum32 += buffer32[k] * (*B)[k][j]
						sum33 += buffer33[k] * (*B)[k][j]
						sum34 += buffer34[k] * (*B)[k][j]
						sum35 += buffer35[k] * (*B)[k][j]
						sum36 += buffer36[k] * (*B)[k][j]
						sum37 += buffer37[k] * (*B)[k][j]
						sum38 += buffer38[k] * (*B)[k][j]
						sum39 += buffer39[k] * (*B)[k][j]
						sum40 += buffer40[k] * (*B)[k][j]
						sum41 += buffer41[k] * (*B)[k][j]
						sum42 += buffer42[k] * (*B)[k][j]
						sum43 += buffer43[k] * (*B)[k][j]
						sum44 += buffer44[k] * (*B)[k][j]
						sum45 += buffer45[k] * (*B)[k][j]
						sum46 += buffer46[k] * (*B)[k][j]
						sum47 += buffer47[k] * (*B)[k][j]
						sum48 += buffer48[k] * (*B)[k][j]
						sum49 += buffer49[k] * (*B)[k][j]
						sum50 += buffer50[k] * (*B)[k][j]
						sum51 += buffer51[k] * (*B)[k][j]
						sum52 += buffer52[k] * (*B)[k][j]
						sum53 += buffer53[k] * (*B)[k][j]
						sum54 += buffer54[k] * (*B)[k][j]
						sum55 += buffer55[k] * (*B)[k][j]
						sum56 += buffer56[k] * (*B)[k][j]
						sum57 += buffer57[k] * (*B)[k][j]
						sum58 += buffer58[k] * (*B)[k][j]
						sum59 += buffer59[k] * (*B)[k][j]
					}
					(*result)[i*amountBuffers+0][j] = sum0
					(*result)[i*amountBuffers+1][j] = sum1
					(*result)[i*amountBuffers+2][j] = sum2
					(*result)[i*amountBuffers+3][j] = sum3
					(*result)[i*amountBuffers+4][j] = sum4
					(*result)[i*amountBuffers+5][j] = sum5
					(*result)[i*amountBuffers+6][j] = sum6
					(*result)[i*amountBuffers+7][j] = sum7
					(*result)[i*amountBuffers+8][j] = sum8
					(*result)[i*amountBuffers+9][j] = sum9
					(*result)[i*amountBuffers+10][j] = sum10
					(*result)[i*amountBuffers+11][j] = sum11
					(*result)[i*amountBuffers+12][j] = sum12
					(*result)[i*amountBuffers+13][j] = sum13
					(*result)[i*amountBuffers+14][j] = sum14
					(*result)[i*amountBuffers+15][j] = sum15
					(*result)[i*amountBuffers+16][j] = sum16
					(*result)[i*amountBuffers+17][j] = sum17
					(*result)[i*amountBuffers+18][j] = sum18
					(*result)[i*amountBuffers+19][j] = sum19
					(*result)[i*amountBuffers+20][j] = sum20
					(*result)[i*amountBuffers+21][j] = sum21
					(*result)[i*amountBuffers+22][j] = sum22
					(*result)[i*amountBuffers+23][j] = sum23
					(*result)[i*amountBuffers+24][j] = sum24
					(*result)[i*amountBuffers+25][j] = sum25
					(*result)[i*amountBuffers+26][j] = sum26
					(*result)[i*amountBuffers+27][j] = sum27
					(*result)[i*amountBuffers+28][j] = sum28
					(*result)[i*amountBuffers+29][j] = sum29
					(*result)[i*amountBuffers+30][j] = sum30
					(*result)[i*amountBuffers+31][j] = sum31
					(*result)[i*amountBuffers+32][j] = sum32
					(*result)[i*amountBuffers+33][j] = sum33
					(*result)[i*amountBuffers+34][j] = sum34
					(*result)[i*amountBuffers+35][j] = sum35
					(*result)[i*amountBuffers+36][j] = sum36
					(*result)[i*amountBuffers+37][j] = sum37
					(*result)[i*amountBuffers+38][j] = sum38
					(*result)[i*amountBuffers+39][j] = sum39
					(*result)[i*amountBuffers+40][j] = sum40
					(*result)[i*amountBuffers+41][j] = sum41
					(*result)[i*amountBuffers+42][j] = sum42
					(*result)[i*amountBuffers+43][j] = sum43
					(*result)[i*amountBuffers+44][j] = sum44
					(*result)[i*amountBuffers+45][j] = sum45
					(*result)[i*amountBuffers+46][j] = sum46
					(*result)[i*amountBuffers+47][j] = sum47
					(*result)[i*amountBuffers+48][j] = sum48
					(*result)[i*amountBuffers+49][j] = sum49
					(*result)[i*amountBuffers+50][j] = sum50
					(*result)[i*amountBuffers+51][j] = sum51
					(*result)[i*amountBuffers+52][j] = sum52
					(*result)[i*amountBuffers+53][j] = sum53
					(*result)[i*amountBuffers+54][j] = sum54
					(*result)[i*amountBuffers+55][j] = sum55
					(*result)[i*amountBuffers+56][j] = sum56
					(*result)[i*amountBuffers+57][j] = sum57
					(*result)[i*amountBuffers+58][j] = sum58
					(*result)[i*amountBuffers+59][j] = sum59
				}
			}
		}(g)
	}
	wg.Wait()
}

func a14(A, B, result *[][]float64) {
	n := len(*A)
	amountBuffers := 10
	if n < amountBuffers {
		a6(A, B, result)
	}
	buffer := make([]float64, n*amountBuffers)
	for i := 0; i < n; i += amountBuffers {
		for b := 0; b < amountBuffers; b++ {
			for k := 0; k < n; k++ {
				buffer[k+b*n] = (*A)[i+b][k]
			}
		}

		sum := make([]float64, amountBuffers)
		for j := 0; j < n; j++ {
			for b := 0; b < amountBuffers; b++ {
				for k := 0; k < n; k++ {
					sum[b] += buffer[k] * (*B)[k][j]
				}
				(*result)[i+b][j] = sum[b]
			}
		}
	}
}

func a15(A, B, result *[][]float64) {
	n := len(*A)
	amountBuffers := 10
	if n < amountBuffers {
		a6(A, B, result)
	}
	buffer := make([]float64, n*amountBuffers)
	for i := 0; i < n; i += amountBuffers {
		var wg sync.WaitGroup

		for b := 0; b < amountBuffers; b++ {
			wg.Add(1)
			go func(b, n int) {
				for k := 0; k < n; k++ {
					buffer[k+b*n] = (*A)[i+b][k]
				}
				wg.Done()
			}(b, n)
		}
		wg.Wait()

		sum := make([]float64, amountBuffers)
		for j := 0; j < n; j++ {
			for b := 0; b < amountBuffers; b++ {
				for k := 0; k < n; k++ {
					sum[b] += buffer[k] * (*B)[k][j]
				}
				(*result)[i+b][j] = sum[b]
			}
		}
	}
}

func a16(A, B, result *[][]float64) {
	n := len(*A)

	ch := make(chan int)

	for b := 0; b < 40; b++ {
		go func() {
			buffer := make([]float64, n)
			var sum float64
			for i := range ch {
				for k := 0; k < n; k++ {
					buffer[k] = (*A)[i][k]
				}
				sum = 0.0
				for j := 0; j < n; j++ {
					for k := 0; k < n; k++ {
						sum += buffer[k] * (*B)[k][j]
					}
					(*result)[i][j] = sum
				}
			}
		}()
	}

	for i := 0; i < n; i++ {
		ch <- i
	}
	close(ch)

	return
}

// getTest - test simple test data
func getenerateMatrix(n int) (A, B, C [][]float64) {
	A = make([][]float64, n)
	B = make([][]float64, n)
	C = make([][]float64, n)

	for i := 0; i < n; i++ {
		A[i] = make([]float64, n)
		B[i] = make([]float64, n)
		C[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			A[i][j] = 4.0 * rand.Float64() * float64(j-i+n*2)
			B[i][j] = 4.0 * rand.Float64() * float64(j-i+n*2)
		}
	}
	return
}

func bench(b *testing.B, f func(*[][]float64, *[][]float64, *[][]float64), n int) {
	b.StopTimer()
	A, B, C := getenerateMatrix(n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		f(&A, &B, &C)
	}
	b.ReportAllocs()
}

/*
func TestA12(t *testing.T) {
	n := 500
	A, B, C := getenerateMatrix(n)
	fmt.Println("Start calculate a1")
	a1(&A, &B, &C)
	C2 := make([][]float64, n)
	for i := 0; i < n; i++ {
		C2[i] = make([]float64, n)
	}
	fmt.Println("Start calculate a12")
	a12(&A, &B, &C2)
	good := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if C[i][j] != C2[i][j] {
				good = false
				t.Errorf("FAIL")
				break
			}
		}
	}
	if good {
		fmt.Println("All is ok")
	}
}


////
func BenchmarkA1_10(b *testing.B) {
	bench(b, a1, 10)
}
func BenchmarkA1_40(b *testing.B) {
	bench(b, a1, 40)
}
func BenchmarkA1_160(b *testing.B) {
	bench(b, a1, 160)
}
func BenchmarkA1_640(b *testing.B) {
	bench(b, a1, 640)
}
func BenchmarkA1_1280(b *testing.B) {
	bench(b, a1, 1280)
}

////
func BenchmarkA2_10(b *testing.B) {
	bench(b, a2, 10)
}
func BenchmarkA2_40(b *testing.B) {
	bench(b, a2, 40)
}
func BenchmarkA2_160(b *testing.B) {
	bench(b, a2, 160)
}
func BenchmarkA2_640(b *testing.B) {
	bench(b, a2, 640)
}
func BenchmarkA2_1280(b *testing.B) {
	bench(b, a2, 1280)
}

////
func BenchmarkA3_10(b *testing.B) {
	bench(b, a3, 10)
}
func BenchmarkA3_40(b *testing.B) {
	bench(b, a3, 40)
}
func BenchmarkA3_160(b *testing.B) {
	bench(b, a3, 160)
}
func BenchmarkA3_640(b *testing.B) {
	bench(b, a3, 640)
}
func BenchmarkA3_1280(b *testing.B) {
	bench(b, a3, 1280)
}

////
func BenchmarkA4_10(b *testing.B) {
	bench(b, a4, 10)
}
func BenchmarkA4_40(b *testing.B) {
	bench(b, a4, 40)
}
func BenchmarkA4_160(b *testing.B) {
	bench(b, a4, 160)
}
func BenchmarkA4_640(b *testing.B) {
	bench(b, a4, 640)
}
func BenchmarkA4_1280(b *testing.B) {
	bench(b, a4, 1280)
}

////
func BenchmarkA5_10(b *testing.B) {
	bench(b, a5, 10)
}
func BenchmarkA5_40(b *testing.B) {
	bench(b, a5, 40)
}
func BenchmarkA5_160(b *testing.B) {
	bench(b, a5, 160)
}
func BenchmarkA5_640(b *testing.B) {
	bench(b, a5, 640)
}
func BenchmarkA5_1280(b *testing.B) {
	bench(b, a5, 1280)
}

////
func BenchmarkA6_10(b *testing.B) {
	bench(b, a6, 10)
}
func BenchmarkA6_40(b *testing.B) {
	bench(b, a6, 40)
}
func BenchmarkA6_160(b *testing.B) {
	bench(b, a6, 160)
}
func BenchmarkA6_640(b *testing.B) {
	bench(b, a6, 640)
}
func BenchmarkA6_1280(b *testing.B) {
	bench(b, a6, 1280)
}

////
func BenchmarkA7_10(b *testing.B) {
	bench(b, a7, 10)
}
func BenchmarkA7_40(b *testing.B) {
	bench(b, a7, 40)
}
func BenchmarkA7_160(b *testing.B) {
	bench(b, a7, 160)
}
func BenchmarkA7_640(b *testing.B) {
	bench(b, a7, 640)
}
func BenchmarkA7_1280(b *testing.B) {
	bench(b, a7, 1280)
}

////
func BenchmarkA8_10(b *testing.B) {
	bench(b, a8, 10)
}
func BenchmarkA8_40(b *testing.B) {
	bench(b, a8, 40)
}
func BenchmarkA8_160(b *testing.B) {
	bench(b, a8, 160)
}
func BenchmarkA8_640(b *testing.B) {
	bench(b, a8, 640)
}
func BenchmarkA8_1280(b *testing.B) {
	bench(b, a8, 1280)
}

////
func BenchmarkA9_10(b *testing.B) {
	bench(b, a9, 10)
}
func BenchmarkA9_40(b *testing.B) {
	bench(b, a9, 40)
}
func BenchmarkA9_160(b *testing.B) {
	bench(b, a9, 160)
}
func BenchmarkA9_640(b *testing.B) {
	bench(b, a9, 640)
}
func BenchmarkA9_1280(b *testing.B) {
	bench(b, a9, 1280)
}

////
func BenchmarkA10_10(b *testing.B) {
	bench(b, a10, 10)
}
func BenchmarkA10_40(b *testing.B) {
	bench(b, a10, 40)
}
func BenchmarkA10_160(b *testing.B) {
	bench(b, a10, 160)
}
func BenchmarkA10_640(b *testing.B) {
	bench(b, a10, 640)
}
func BenchmarkA10_1280(b *testing.B) {
	bench(b, a10, 1280)
}

////
func BenchmarkA11_10(b *testing.B) {
	bench(b, a11, 10)
}
func BenchmarkA11_40(b *testing.B) {
	bench(b, a11, 40)
}
func BenchmarkA11_160(b *testing.B) {
	bench(b, a11, 160)
}
func BenchmarkA11_640(b *testing.B) {
	bench(b, a11, 640)
}
func BenchmarkA11_1280(b *testing.B) {
	bench(b, a11, 1280)
}

////
func BenchmarkA12_10(b *testing.B) {
	bench(b, a12, 10)
}
func BenchmarkA12_40(b *testing.B) {
	bench(b, a12, 40)
}
func BenchmarkA12_160(b *testing.B) {
	bench(b, a12, 160)
}
func BenchmarkA12_640(b *testing.B) {
	bench(b, a12, 640)
}
func BenchmarkA12_1280(b *testing.B) {
	bench(b, a12, 1280)
}
*/
////
func BenchmarkA13_10(b *testing.B) {
	bench(b, a13, 10)
}
func BenchmarkA13_40(b *testing.B) {
	bench(b, a13, 40)
}
func BenchmarkA13_160(b *testing.B) {
	bench(b, a13, 160)
}
func BenchmarkA13_640(b *testing.B) {
	bench(b, a13, 640)
}
func BenchmarkA13_1280(b *testing.B) {
	bench(b, a13, 1280)
}

////
func BenchmarkA13a_10(b *testing.B) {
	bench(b, a13a, 10)
}
func BenchmarkA13a_40(b *testing.B) {
	bench(b, a13a, 40)
}
func BenchmarkA13a_160(b *testing.B) {
	bench(b, a13a, 160)
}
func BenchmarkA13a_640(b *testing.B) {
	bench(b, a13a, 640)
}
func BenchmarkA13a_1280(b *testing.B) {
	bench(b, a13a, 1280)
}

/*
////
func BenchmarkA14_10(b *testing.B) {
	bench(b, a14, 10)
}
func BenchmarkA14_40(b *testing.B) {
	bench(b, a14, 40)
}
func BenchmarkA14_160(b *testing.B) {
	bench(b, a14, 160)
}
func BenchmarkA14_640(b *testing.B) {
	bench(b, a14, 640)
}
func BenchmarkA14_1280(b *testing.B) {
	bench(b, a14, 1280)
}

////
func BenchmarkA15_10(b *testing.B) {
	bench(b, a15, 10)
}
func BenchmarkA15_40(b *testing.B) {
	bench(b, a15, 40)
}
func BenchmarkA15_160(b *testing.B) {
	bench(b, a15, 160)
}
func BenchmarkA15_640(b *testing.B) {
	bench(b, a15, 640)
}
func BenchmarkA15_1280(b *testing.B) {
	bench(b, a15, 1280)
}

////
func BenchmarkA16_10(b *testing.B) {
	bench(b, a16, 10)
}
func BenchmarkA16_40(b *testing.B) {
	bench(b, a16, 40)
}
func BenchmarkA16_160(b *testing.B) {
	bench(b, a16, 160)
}
func BenchmarkA16_640(b *testing.B) {
	bench(b, a16, 640)
}
func BenchmarkA16_1280(b *testing.B) {
	bench(b, a16, 1280)
}

func BenchmarkSingleArray(b *testing.B) {
	b.StopTimer()
	n := 5000
	array := make([]float64, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			array[i*n+j] = float64(i + j)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < n; k++ {
			array[2*n+k] = array[2*n+k] + math.Sin(array[2*n+k]/array[k*n+2])
		}
	}
	b.ReportAllocs()
}

func BenchmarkSingleArray_2(b *testing.B) {
	b.StopTimer()
	n := 5000
	array := make([]float64, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			array[i*n+j] = float64(i + j)
		}
	}
	b.StartTimer()
	n = n / 2
	for i := 0; i < b.N; i++ {
		for k := 0; k < n; k++ {
			array[2*(n*2+0)+k] = array[2*(n*2+0)+k] + math.Sin(array[2*(n*2+0)+k]/array[k*(n*2+0)+2])
			array[2*(n*2+1)+k] = array[2*(n*2+1)+k] + math.Sin(array[2*(n*2+1)+k]/array[k*(n*2+1)+2])
		}
	}
	b.ReportAllocs()
}

func BenchmarkSingleArray_5(b *testing.B) {
	b.StopTimer()
	n := 5000
	array := make([]float64, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			array[i*n+j] = float64(i + j)
		}
	}
	b.StartTimer()
	n = n / 5
	N := n * 5
	for i := 0; i < b.N; i++ {
		for k := 0; k < n; k++ {
			array[2*(N+0)+k] = array[2*(N+0)+k] + math.Sin(array[2*(N+0)+k]/array[k*(N+0)+2])
			array[2*(N+1)+k] = array[2*(N+1)+k] + math.Sin(array[2*(N+1)+k]/array[k*(N+1)+2])
			array[2*(N+2)+k] = array[2*(N+2)+k] + math.Sin(array[2*(N+2)+k]/array[k*(N+2)+2])
			array[2*(N+3)+k] = array[2*(N+3)+k] + math.Sin(array[2*(N+3)+k]/array[k*(N+3)+2])
			array[2*(N+4)+k] = array[2*(N+4)+k] + math.Sin(array[2*(N+4)+k]/array[k*(N+4)+2])
		}
	}
	b.ReportAllocs()
}

func BenchmarkSingleArray_10(b *testing.B) {
	b.StopTimer()
	n := 5000
	array := make([]float64, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			array[i*n+j] = float64(i + j)
		}
	}
	b.StartTimer()
	n = n / 10
	N := n * 10
	for i := 0; i < b.N; i++ {
		for k := 0; k < n; k++ {
			array[2*(N+0)+k] = array[2*(N+0)+k] + math.Sin(array[2*(N+0)+k]/array[k*(N+0)+2])
			array[2*(N+1)+k] = array[2*(N+1)+k] + math.Sin(array[2*(N+1)+k]/array[k*(N+1)+2])
			array[2*(N+2)+k] = array[2*(N+2)+k] + math.Sin(array[2*(N+2)+k]/array[k*(N+2)+2])
			array[2*(N+3)+k] = array[2*(N+3)+k] + math.Sin(array[2*(N+3)+k]/array[k*(N+3)+2])
			array[2*(N+4)+k] = array[2*(N+4)+k] + math.Sin(array[2*(N+4)+k]/array[k*(N+4)+2])
			array[2*(N+5)+k] = array[2*(N+5)+k] + math.Sin(array[2*(N+5)+k]/array[k*(N+5)+2])
			array[2*(N+6)+k] = array[2*(N+6)+k] + math.Sin(array[2*(N+6)+k]/array[k*(N+6)+2])
			array[2*(N+7)+k] = array[2*(N+7)+k] + math.Sin(array[2*(N+7)+k]/array[k*(N+7)+2])
			array[2*(N+8)+k] = array[2*(N+8)+k] + math.Sin(array[2*(N+8)+k]/array[k*(N+8)+2])
			array[2*(N+9)+k] = array[2*(N+9)+k] + math.Sin(array[2*(N+9)+k]/array[k*(N+9)+2])
		}
	}
	b.ReportAllocs()
}

func BenchmarkDoubleArray(b *testing.B) {
	b.StopTimer()
	n := 5000
	array := make([][]float64, n)
	for i := 0; i < n; i++ {
		array[i] = make([]float64, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			array[i][j] = float64(i + j)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < n; k++ {
			array[2][k] = array[2][k] + math.Sin(array[2][k]/array[k][2])
		}
	}
	b.ReportAllocs()
}

func BenchmarkSingleEqualArray(b *testing.B) {
	b.StopTimer()
	n := 500000
	array := make([]float64, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < n; k++ {
			array[k] = 1.2222
		}
	}
	b.ReportAllocs()
}

func BenchmarkSingleEqualArray_2_N0(b *testing.B) {
	b.StopTimer()
	n := 500000
	array := make([]float64, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		N := n / 2
		for k := 0; k < N; k++ {
			array[k*2+0] = 1.2222
			array[k*2+1] = 1.2222
		}
		if n-N*2 == 1 {
			array[n-1] = 1.2222
		}
	}
	b.ReportAllocs()
}

func BenchmarkSingleEqualArray_2_N1(b *testing.B) {
	b.StopTimer()
	n := 500001
	array := make([]float64, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		N := n / 2
		for k := 0; k < N; k++ {
			array[k*2+0] = 1.2222
			array[k*2+1] = 1.2222
		}
		if n-N*2 == 1 {
			array[n-1] = 1.2222
		}
	}
	b.ReportAllocs()
}

func BenchmarkSingleEqualArray_2_N2(b *testing.B) {
	b.StopTimer()
	n := 500002
	array := make([]float64, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		N := n / 2
		for k := 0; k < N; k++ {
			array[k*2+0] = 1.2222
			array[k*2+1] = 1.2222
		}
		if n-N*2 == 1 {
			array[n-1] = 1.2222
		}
	}
	b.ReportAllocs()
}

func BenchmarkSingleEqualArray_5(b *testing.B) {
	b.StopTimer()
	n := 500000
	array := make([]float64, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		N := n / 5
		for k := 0; k < N; k++ {
			array[k*2+0] = 1.2222
			array[k*2+1] = 1.2222
			array[k*2+2] = 1.2222
			array[k*2+3] = 1.2222
			array[k*2+4] = 1.2222
		}
	}
	b.ReportAllocs()
}


*/
