package solver

import (
	"math"

	"github.com/Konstantin8105/GoLinAlg/matrix"
)

// Eigen - Eigenvalues and eigenvectors of a real matrix.
// If A is symmetric, then A = V*D*V' where the eigenvalue matrix D is
// diagonal and the eigenvector matrix V is orthogonal.
// I.e. A = V.times(D.times(V.transpose())) and
// V.times(V.transpose()) equals the identity matrix.
// If A is not symmetric, then the eigenvalue matrix D is block diagonal
// with the real eigenvalues in 1-by-1 blocks and any complex eigenvalues,
// lambda + i*mu, in 2-by-2 blocks, [lambda, mu; -mu, lambda].  The
// columns of V represent the eigenvectors in the sense that A*V = V*D,
// i.e. A.times(V) equals V.times(D).  The matrix V may be badly
// conditioned, or even singular, so the validity of the equation
// A = V*D*inverse(V) depends upon V.cond().
type Eigen struct {
	n            int         // Row and column dimension (square matrix).
	issymmetric  bool        // Symmetry flag - internal symmetry flag.
	d, e         []float64   // Arrays for internal storage of eigenvalues.
	V            [][]float64 // Array for internal storage of eigenvectors.
	H            [][]float64 // Array for internal storage of nonsymmetric Hessenberg form.
	ort          []float64   // Working storage for nonsymmetric algorithm.
	cdivr, cdivi float64     // Not-serialize. Complex scalar division.
}

// Symmetric Householder reduction to tridiagonal form.
// This is derived from the Algol procedures tred2 by
// Bowdler, Martin, Reinsch, and Wilkinson, Handbook for
// Auto. Comp., Vol.ii-Linear Algebra, and the corresponding
// Fortran subroutine in EISPACK.
func (eig *Eigen) tred2() {
	for j := 0; j < eig.n; j++ {
		eig.d[j] = eig.V[eig.n-1][j]
	}

	// Householder reduction to tridiagonal form.

	for i := eig.n - 1; i > 0; i-- {

		// Scale to avoid under/overflow.

		scale := 0.0
		h := 0.0
		for k := 0; k < i; k++ {
			scale += math.Abs(eig.d[k])
		}
		if scale == 0.0 {
			eig.e[i] = eig.d[i-1]
			for j := 0; j < i; j++ {
				eig.d[j] = eig.V[i-1][j]
				eig.V[i][j] = 0.0
				eig.V[j][i] = 0.0
			}
		} else {

			// Generate Householder vector.

			for k := 0; k < i; k++ {
				eig.d[k] /= scale
				h += eig.d[k] * eig.d[k]
			}
			f := eig.d[i-1]
			g := math.Sqrt(h)
			if f > 0 {
				g = -g
			}
			eig.e[i] = scale * g
			h -= f * g
			eig.d[i-1] = f - g
			for j := 0; j < i; j++ {
				eig.e[j] = 0.0
			}

			// Apply similarity transformation to remaining columns.

			for j := 0; j < i; j++ {
				f = eig.d[j]
				eig.V[j][i] = f
				g = eig.e[j] + eig.V[j][j]*f
				for k := j + 1; k <= i-1; k++ {
					g += eig.V[k][j] * eig.d[k]
					eig.e[k] += eig.V[k][j] * f
				}
				eig.e[j] = g
			}
			f = 0.0
			for j := 0; j < i; j++ {
				eig.e[j] /= h
				f += eig.e[j] * eig.d[j]
			}
			hh := f / (h + h)
			for j := 0; j < i; j++ {
				eig.e[j] -= hh * eig.d[j]
			}
			for j := 0; j < i; j++ {
				f = eig.d[j]
				g = eig.e[j]
				for k := j; k <= i-1; k++ {
					eig.V[k][j] -= (f*eig.e[k] + g*eig.d[k])
				}
				eig.d[j] = eig.V[i-1][j]
				eig.V[i][j] = 0.0
			}
		}
		eig.d[i] = h
	}

	// Accumulate transformations.
	for i := 0; i < eig.n-1; i++ {
		eig.V[eig.n-1][i] = eig.V[i][i]
		eig.V[i][i] = 1.0
		h := eig.d[i+1]
		if h != 0.0 {
			for k := 0; k <= i; k++ {
				eig.d[k] = eig.V[k][i+1] / h
			}
			for j := 0; j <= i; j++ {
				g := 0.0
				for k := 0; k <= i; k++ {
					g += eig.V[k][i+1] * eig.V[k][j]
				}
				for k := 0; k <= i; k++ {
					eig.V[k][j] -= g * eig.d[k]
				}
			}
		}
		for k := 0; k <= i; k++ {
			eig.V[k][i+1] = 0.0
		}
	}
	for j := 0; j < eig.n; j++ {
		eig.d[j] = eig.V[eig.n-1][j]
		eig.V[eig.n-1][j] = 0.0
	}
	eig.V[eig.n-1][eig.n-1] = 1.0
	eig.e[0] = 0.0
}

// Symmetric tridiagonal QL algorithm.
// This is derived from the Algol procedures tql2, by
// Bowdler, Martin, Reinsch, and Wilkinson, Handbook for
// Auto. Comp., Vol.ii-Linear Algebra, and the corresponding
// Fortran subroutine in EISPACK.
func (eig *Eigen) tql2() {
	for i := 1; i < eig.n; i++ {
		eig.e[i-1] = eig.e[i]
	}
	eig.e[eig.n-1] = 0.0

	f := 0.0
	tst1 := 0.0
	eps := math.Pow(2.0, -52.0)
	for l := 0; l < eig.n; l++ {

		// Find small subdiagonal element
		tst1 = math.Max(tst1, math.Abs(eig.d[l])+math.Abs(eig.e[l]))
		m := l
		for m < eig.n {
			if math.Abs(eig.e[m]) <= eps*tst1 {
				break
			}
			m++
		}

		// If m == l, d[l] is an eigenvalue,
		// otherwise, iterate.
		if m > l {
			iter := 0
			for {
				iter = iter + 1 // (Could check iteration count here.)

				// Compute implicit shift
				g := eig.d[l]
				p := (eig.d[l+1] - g) / (2.0 * eig.e[l])
				r := hypot(p, 1.0)
				if p < 0 {
					r = -r
				}
				eig.d[l] = eig.e[l] / (p + r)
				eig.d[l+1] = eig.e[l] * (p + r)
				dl1 := eig.d[l+1]
				h := g - eig.d[l]
				for i := l + 2; i < eig.n; i++ {
					eig.d[i] -= h
				}
				f += h

				// Implicit QL transformation.
				p = eig.d[m]
				c := 1.0
				c2 := c
				c3 := c
				el1 := eig.e[l+1]
				s := 0.0
				s2 := 0.0
				for i := m - 1; i >= l; i-- {
					c3 = c2
					c2 = c
					s2 = s
					g = c * eig.e[i]
					h = c * p
					r = hypot(p, eig.e[i])
					eig.e[i+1] = s * r
					s = eig.e[i] / r
					c = p / r
					p = c*eig.d[i] - s*g
					eig.d[i+1] = h + s*(c*g+s*eig.d[i])

					// Accumulate transformation.
					for k := 0; k < eig.n; k++ {
						h = eig.V[k][i+1]
						eig.V[k][i+1] = s*eig.V[k][i] + c*h
						eig.V[k][i] = c*eig.V[k][i] - s*h
					}
				}
				p = -s * s2 * c3 * el1 * eig.e[l] / dl1
				eig.e[l] = s * p
				eig.d[l] = c * p

				// Check for convergence.
				if math.Abs(eig.e[l]) > eps*tst1 {
					break
				}
			}
		}
		eig.d[l] = eig.d[l] + f
		eig.e[l] = 0.0
	}

	// Sort eigenvalues and corresponding vectors.
	for i := 0; i < eig.n-1; i++ {
		k := i
		p := eig.d[i]
		for j := i + 1; j < eig.n; j++ {
			if eig.d[j] < p {
				k = j
				p = eig.d[j]
			}
		}
		if k != i {
			eig.d[k] = eig.d[i]
			eig.d[i] = p
			for j := 0; j < eig.n; j++ {
				p = eig.V[j][i]
				eig.V[j][i] = eig.V[j][k]
				eig.V[j][k] = p
			}
		}
	}
}

//  Nonsymmetric reduction to Hessenberg form.
//  This is derived from the Algol procedures orthes and ortran,
//  by Martin and Wilkinson, Handbook for Auto. Comp.,
//  Vol.ii-Linear Algebra, and the corresponding
//  Fortran subroutines in EISPACK.
func (eig *Eigen) orthes() {
	low := 0
	high := eig.n - 1

	for m := low + 1; m <= high-1; m++ {

		// Scale column.

		scale := 0.0
		for i := m; i <= high; i++ {
			scale += math.Abs(eig.H[i][m-1])
		}
		if scale != 0.0 {

			// Compute Householder transformation.
			h := 0.0
			for i := high; i >= m; i-- {
				eig.ort[i] = eig.H[i][m-1] / scale
				h += eig.ort[i] * eig.ort[i]
			}
			g := math.Sqrt(h)
			if eig.ort[m] > 0 {
				g = -g
			}
			h = h - eig.ort[m]*g
			eig.ort[m] = eig.ort[m] - g

			// Apply Householder similarity transformation
			// H = (I-u*u'/h)*H*(I-u*u')/h)
			for j := m; j < eig.n; j++ {
				f := 0.0
				for i := high; i >= m; i-- {
					f += eig.ort[i] * eig.H[i][j]
				}
				f /= h
				for i := m; i <= high; i++ {
					eig.H[i][j] -= f * eig.ort[i]
				}
			}

			for i := 0; i <= high; i++ {
				f := 0.0
				for j := high; j >= m; j-- {
					f += eig.ort[j] * eig.H[i][j]
				}
				f /= h
				for j := m; j <= high; j++ {
					eig.H[i][j] -= f * eig.ort[j]
				}
			}
			eig.ort[m] = scale * eig.ort[m]
			eig.H[m][m-1] = scale * g
		}
	}

	// Accumulate transformations (Algol's ortran).
	for i := 0; i < eig.n; i++ {
		for j := 0; j < eig.n; j++ {
			if i == j {
				eig.V[i][j] = 1.0
			} else {
				eig.V[i][j] = 0.0
			}
		}
	}

	for m := high - 1; m >= low+1; m-- {
		if eig.H[m][m-1] != 0.0 {
			for i := m + 1; i <= high; i++ {
				eig.ort[i] = eig.H[i][m-1]
			}
			for j := m; j <= high; j++ {
				g := 0.0
				for i := m; i <= high; i++ {
					g += eig.ort[i] * eig.V[i][j]
				}
				// Double division avoids possible underflow
				g /= eig.ort[m]
				g /= eig.H[m][m-1]
				for i := m; i <= high; i++ {
					eig.V[i][j] += g * eig.ort[i]
				}
			}
		}
	}
}

// Complex scalar division.
func (eig *Eigen) cdiv(xr, xi, yr, yi float64) {
	var r, d float64
	if math.Abs(yr) > math.Abs(yi) {
		r = yi / yr
		d = yr + r*yi
		eig.cdivr = (xr + r*xi) / d
		eig.cdivi = (xi - r*xr) / d
	} else {
		r = yr / yi
		d = yi + r*yr
		eig.cdivr = (r*xr + xi) / d
		eig.cdivi = (r*xi - xr) / d
	}
}

//  Nonsymmetric reduction from Hessenberg to real Schur form.
//  This is derived from the Algol procedure hqr2,
//  by Martin and Wilkinson, Handbook for Auto. Comp.,
//  Vol.ii-Linear Algebra, and the corresponding
//  Fortran subroutine in EISPACK.
func (eig *Eigen) hqr2() {

	// Initialize

	nn := eig.n
	n := nn - 1
	low := 0
	high := nn - 1
	eps := math.Pow(2.0, -52.0)
	exshift := 0.0
	p := 0.0
	q := 0.0
	r := 0.0
	s := 0.0
	z := 0.0
	var t, w, x, y float64

	// Store roots isolated by balanc and compute matrix norm

	norm := 0.0
	for i := 0; i < nn; i++ {
		if i < low || i > high {
			eig.d[i] = eig.H[i][i]
			eig.e[i] = 0.0
		}
		j := 0
		if i-1 > j {
			j = i - 1
		}
		for ; /* j := math.Max(i-1, 0) */ j < nn; j++ {
			norm += math.Abs(eig.H[i][j])
		}
	}

	// Outer loop over eigenvalue index
	iter := 0
	for n >= low {
		// Look for single small sub-diagonal element
		l := n
		for l > low {
			s = math.Abs(eig.H[l-1][l-1]) + math.Abs(eig.H[l][l])
			if s == 0.0 {
				s = norm
			}
			if math.Abs(eig.H[l][l-1]) < eps*s {
				break
			}
			l--
		}

		// Check for convergence
		if l == n {
			// One root found
			eig.H[n][n] = eig.H[n][n] + exshift
			eig.d[n] = eig.H[n][n]
			eig.e[n] = 0.0
			n--
			iter = 0
		} else if l == n-1 {
			// Two roots found
			w = eig.H[n][n-1] * eig.H[n-1][n]
			p = (eig.H[n-1][n-1] - eig.H[n][n]) / 2.0
			q = p*p + w
			z = math.Sqrt(math.Abs(q))
			eig.H[n][n] = eig.H[n][n] + exshift
			eig.H[n-1][n-1] = eig.H[n-1][n-1] + exshift
			x = eig.H[n][n]

			if q >= 0 {
				// Real pair
				if p >= 0 {
					z = p + z
				} else {
					z = p - z
				}
				eig.d[n-1] = x + z
				eig.d[n] = eig.d[n-1]
				if z != 0.0 {
					eig.d[n] = x - w/z
				}
				eig.e[n-1] = 0.0
				eig.e[n] = 0.0
				x = eig.H[n][n-1]
				s = math.Abs(x) + math.Abs(z)
				p = x / s
				q = z / s
				r = math.Sqrt(p*p + q*q)
				p /= r
				q /= r

				// Row modification
				for j := n - 1; j < nn; j++ {
					z = eig.H[n-1][j]
					eig.H[n-1][j] = q*z + p*eig.H[n][j]
					eig.H[n][j] = q*eig.H[n][j] - p*z
				}

				// Column modification
				for i := 0; i <= n; i++ {
					z = eig.H[i][n-1]
					eig.H[i][n-1] = q*z + p*eig.H[i][n]
					eig.H[i][n] = q*eig.H[i][n] - p*z
				}

				// Accumulate transformations
				for i := low; i <= high; i++ {
					z = eig.V[i][n-1]
					eig.V[i][n-1] = q*z + p*eig.V[i][n]
					eig.V[i][n] = q*eig.V[i][n] - p*z
				}

			} else {
				// Complex pair
				eig.d[n-1] = x + p
				eig.d[n] = x + p
				eig.e[n-1] = z
				eig.e[n] = -z
			}
			n -= 2
			iter = 0

			// No convergence yet

		} else {
			// Form shift
			x = eig.H[n][n]
			y = 0.0
			w = 0.0
			if l < n {
				y = eig.H[n-1][n-1]
				w = eig.H[n][n-1] * eig.H[n-1][n]
			}

			// Wilkinson's original ad hoc shift
			if iter == 10 {
				exshift += x
				for i := low; i <= n; i++ {
					eig.H[i][i] -= x
				}
				s = math.Abs(eig.H[n][n-1]) + math.Abs(eig.H[n-1][n-2])
				y = 0.75 * s
				x = y
				w = -0.4375 * s * s
			}

			// MATLAB's new ad hoc shift
			if iter == 30 {
				s = (y - x) / 2.0
				s = s*s + w
				if s > 0 {
					s = math.Sqrt(s)
					if y < x {
						s = -s
					}
					s = x - w/((y-x)/2.0+s)
					for i := low; i <= n; i++ {
						eig.H[i][i] -= s
					}
					exshift += s
					w = 0.964
					y = w
					x = y
				}
			}

			iter++ // (Could check iteration count here.)

			// Look for two consecutive small sub-diagonal elements
			m := n - 2
			for m >= l {
				z = eig.H[m][m]
				r = x - z
				s = y - z
				p = (r*s-w)/eig.H[m+1][m] + eig.H[m][m+1]
				q = eig.H[m+1][m+1] - z - r - s
				r = eig.H[m+2][m+1]
				s = math.Abs(p) + math.Abs(q) + math.Abs(r)
				p /= s
				q /= s
				r /= s
				if m == l {
					break
				}
				if math.Abs(eig.H[m][m-1])*(math.Abs(q)+math.Abs(r)) <
					eps*(math.Abs(p)*(math.Abs(eig.H[m-1][m-1])+math.Abs(z)+
						math.Abs(eig.H[m+1][m+1]))) {
					break
				}
				m--
			}

			for i := m + 2; i <= n; i++ {
				eig.H[i][i-2] = 0.0
				if i > m+2 {
					eig.H[i][i-3] = 0.0
				}
			}

			// Double QR step involving rows l:n and columns m:n

			for k := m; k <= n-1; k++ {
				notlast := (k != n-1)
				if k != m {
					p = eig.H[k][k-1]
					q = eig.H[k+1][k-1]
					if notlast {
						r = eig.H[k+2][k-1]
					} else {
						r = 0.0
					}
					x = math.Abs(p) + math.Abs(q) + math.Abs(r)
					if x == 0.0 {
						continue
					}
					p /= x
					q /= x
					r /= x
				}

				s = math.Sqrt(p*p + q*q + r*r)
				if p < 0 {
					s = -s
				}
				if s != 0 {
					if k != m {
						eig.H[k][k-1] = -s * x
					} else if l != m {
						eig.H[k][k-1] = -eig.H[k][k-1]
					}
					p += s
					x = p / s
					y = q / s
					z = r / s
					q /= p
					r /= p

					// Row modification
					for j := k; j < nn; j++ {
						p = eig.H[k][j] + q*eig.H[k+1][j]
						if notlast {
							p = p + r*eig.H[k+2][j]
							eig.H[k+2][j] = eig.H[k+2][j] - p*z
						}
						eig.H[k][j] -= p * x
						eig.H[k+1][j] -= p * y
					}

					// Column modification
					size := n
					if k+3 < size {
						size = k + 3
					}
					for i := 0; i <= size; i++ {
						p = x*eig.H[i][k] + y*eig.H[i][k+1]
						if notlast {
							p += z * eig.H[i][k+2]
							eig.H[i][k+2] -= p * r
						}
						eig.H[i][k] -= p
						eig.H[i][k+1] -= p * q
					}

					// Accumulate transformations
					for i := low; i <= high; i++ {
						p = x*eig.V[i][k] + y*eig.V[i][k+1]
						if notlast {
							p += z * eig.V[i][k+2]
							eig.V[i][k+2] -= p * r
						}
						eig.V[i][k] -= p
						eig.V[i][k+1] -= p * q
					}
				} // (s != 0)
			} // k loop
		} // check convergence
	} // while (n >= low)

	// Backsubstitute to find vectors of upper triangular form

	if norm == 0.0 {
		return
	}

	for n := nn - 1; n >= 0; n-- {
		p = eig.d[n]
		q = eig.e[n]

		// Real vector

		if q == 0 {
			l := n
			eig.H[n][n] = 1.0
			for i := n - 1; i >= 0; i-- {
				w = eig.H[i][i] - p
				r = 0.0
				for j := l; j <= n; j++ {
					r += eig.H[i][j] * eig.H[j][n]
				}
				if eig.e[i] < 0.0 {
					z = w
					s = r
				} else {
					l = i
					if eig.e[i] == 0.0 {
						if w != 0.0 {
							eig.H[i][n] = -r / w
						} else {
							eig.H[i][n] = -r / (eps * norm)
						}

						// Solve real equations
					} else {
						x = eig.H[i][i+1]
						y = eig.H[i+1][i]
						q = (eig.d[i]-p)*(eig.d[i]-p) + eig.e[i]*eig.e[i]
						t = (x*s - z*r) / q
						eig.H[i][n] = t
						if math.Abs(x) > math.Abs(z) {
							eig.H[i+1][n] = (-r - w*t) / x
						} else {
							eig.H[i+1][n] = (-s - y*t) / z
						}
					}

					// Overflow control
					t = math.Abs(eig.H[i][n])
					if (eps*t)*t > 1 {
						for j := i; j <= n; j++ {
							eig.H[j][n] = eig.H[j][n] / t
						}
					}
				}
			}

			// Complex vector
		} else if q < 0 {
			l := n - 1

			// Last vector component imaginary so matrix is triangular
			if math.Abs(eig.H[n][n-1]) > math.Abs(eig.H[n-1][n]) {
				eig.H[n-1][n-1] = q / eig.H[n][n-1]
				eig.H[n-1][n] = -(eig.H[n][n] - p) / eig.H[n][n-1]
			} else {
				eig.cdiv(0.0, -eig.H[n-1][n], eig.H[n-1][n-1]-p, q)
				eig.H[n-1][n-1] = eig.cdivr
				eig.H[n-1][n] = eig.cdivi
			}
			eig.H[n][n-1] = 0.0
			eig.H[n][n] = 1.0
			for i := n - 2; i >= 0; i-- {
				ra := 0.0
				sa := 0.0
				var vr, vi float64
				for j := l; j <= n; j++ {
					ra += eig.H[i][j] * eig.H[j][n-1]
					sa += eig.H[i][j] * eig.H[j][n]
				}
				w = eig.H[i][i] - p

				if eig.e[i] < 0.0 {
					z = w
					r = ra
					s = sa
				} else {
					l = i
					if eig.e[i] == 0 {
						eig.cdiv(-ra, -sa, w, q)
						eig.H[i][n-1] = eig.cdivr
						eig.H[i][n] = eig.cdivi
					} else {
						// Solve complex equations
						x = eig.H[i][i+1]
						y = eig.H[i+1][i]
						vr = (eig.d[i]-p)*(eig.d[i]-p) + eig.e[i]*eig.e[i] - q*q
						vi = (eig.d[i] - p) * 2.0 * q
						if vr == 0.0 && vi == 0.0 {
							vr = eps * norm * (math.Abs(w) + math.Abs(q) +
								math.Abs(x) + math.Abs(y) + math.Abs(z))
						}
						eig.cdiv(x*r-z*ra+q*sa, x*s-z*sa-q*ra, vr, vi)
						eig.H[i][n-1] = eig.cdivr
						eig.H[i][n] = eig.cdivi
						if math.Abs(x) > (math.Abs(z) + math.Abs(q)) {
							eig.H[i+1][n-1] = (-ra - w*eig.H[i][n-1] + q*eig.H[i][n]) / x
							eig.H[i+1][n] = (-sa - w*eig.H[i][n] - q*eig.H[i][n-1]) / x
						} else {
							eig.cdiv(-r-y*eig.H[i][n-1], -s-y*eig.H[i][n], z, q)
							eig.H[i+1][n-1] = eig.cdivr
							eig.H[i+1][n] = eig.cdivi
						}
					}

					// Overflow control
					t = math.Max(math.Abs(eig.H[i][n-1]), math.Abs(eig.H[i][n]))
					if (eps*t)*t > 1 {
						for j := i; j <= n; j++ {
							eig.H[j][n-1] = eig.H[j][n-1] / t
							eig.H[j][n] = eig.H[j][n] / t
						}
					}
				}
			}
		}
	}

	// Vectors of isolated roots
	for i := 0; i < nn; i++ {
		if i < low || i > high {
			for j := i; j < nn; j++ {
				eig.V[i][j] = eig.H[i][j]
			}
		}
	}

	// Back transformation to get eigenvectors of original matrix
	for j := nn - 1; j >= low; j-- {
		for i := low; i <= high; i++ {
			z = 0.0
			for k := low; k <= j || k <= high; /* math.Min(j, high) */ k++ {
				z += eig.V[i][k] * eig.H[k][j]
			}
			eig.V[i][j] = z
		}
	}
}

// NewEigen - Check for symmetry, then construct the eigenvalue decomposition. Structure to access D and V.
func NewEigen(A matrix.T64) (e Eigen) {
	e.n = A.GetColumnSize()
	e.V = make([][]float64, e.n)
	for i := 0; i < e.n; i++ {
		e.V[i] = make([]float64, e.n)
	}
	e.d = make([]float64, e.n)
	e.e = make([]float64, e.n)
	// var V = new double[n][n];
	// var d = new double[n];
	// var e = new double[n];

	issymmetric := true
	for j := 0; j < e.n; j++ {
		for i := 0; i < e.n; i++ {
			if A.Get(i, j) != A.Get(j, i) {
				issymmetric = false
				goto BREAK
			}
		}
	}
BREAK:
	if issymmetric {
		for i := 0; i < e.n; i++ {
			for j := 0; j < e.n; j++ {
				e.V[i][j] = A.Get(i, j)
			}
		}

		// Tridiagonalize.
		e.tred2()

		// Diagonalize.
		e.tql2()

	} else {
		//H = new double[n][n];
		//ort = new double[n];
		e.H = make([][]float64, e.n)
		for i := 0; i < e.n; i++ {
			e.H[i] = make([]float64, e.n)
		}
		e.ort = make([]float64, e.n)

		for j := 0; j < e.n; j++ {
			for i := 0; i < e.n; i++ {
				e.H[i][j] = A.Get(i, j)
			}
		}

		// Reduce to Hessenberg form.
		e.orthes()

		// Reduce Hessenberg to real Schur form.
		e.hqr2()
	}
	return e
}

// GetV - Return the eigenvector matrix
func (eig *Eigen) GetV() matrix.T64 {
	m := matrix.NewMatrix64bySize(eig.n, eig.n)
	for i := 0; i < eig.n; i++ {
		for j := 0; j < eig.n; j++ {
			m.Set(i, j, eig.V[i][j])
		}
	}
	return m
}

// GetRealEigenvalues - Return the real parts of the eigenvalues
func (eig *Eigen) GetRealEigenvalues() []float64 {
	return eig.d
}

// GetImagEigenvalues - Return the imaginary parts of the eigenvalues
func (eig *Eigen) GetImagEigenvalues() []float64 {
	return eig.e
}

// GetD - Return the block diagonal eigenvalue matrix
func (eig *Eigen) GetD() matrix.T64 {
	D := matrix.NewMatrix64bySize(eig.n, eig.n)
	for i := 0; i < eig.n; i++ {
		for j := 0; j < eig.n; j++ {
			D.Set(i, j, 0.0)
		}
		D.Set(i, i, eig.d[i])
		if eig.e[i] > 0 {
			D.Set(i, i+1, eig.e[i])
		} else if eig.e[i] < 0 {
			D.Set(i, i-1, eig.e[i])
		}
	}
	return D
}

func hypot(a, b float64) (r float64) {
	if math.Abs(a) > math.Abs(b) {
		r = b / a
		r = math.Abs(a) * math.Sqrt(1+r*r)
	} else if b != 0 {
		r = a / b
		r = math.Abs(b) * math.Sqrt(1+r*r)
	} else {
		r = 0.0
	}
	return r
}
