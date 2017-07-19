package main

template = `

func timesAlgorithmSimple(a, b, c *[][]float64, m, n, h int) {
	var sum float64
	buf := make([]float64, n, n)
	for j := 0; j < h; j++ {
		for k := 0; k < n; k++ {
			buf[k] = (*b)[k][j]
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
