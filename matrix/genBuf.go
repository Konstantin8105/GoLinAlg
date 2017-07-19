package main

template = `

func timesAlgorithmBuffer<N>(a, b, c *[][]float64, m, n, h int) {
	var sum<N> float64
	buf<N> := make([]float64, n, n)
	for j := 0; j < h; j++ {
		for k := 0; k < n; k++ {
			buf<N>[k] = (*b)[k][j]
		}
		for i := 0; i < m; i++ {
			sum = 0
			for k := 0; k < n; k++ {
				sum += (*a)[i][k] * buf[k]
			}
			(*c)[i][j] = sum
		}
	}
}`

func main() {

}
