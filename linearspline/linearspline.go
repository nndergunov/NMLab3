package main

func linearSpline(x, y []float64, h float64, X, Y []float64) {
	n := len(x)
	j := 0

	for i := 0; i < n-1; i++ {
		var yn, xn float64

		for xn = x[i]; xn < x[i+1]; xn += h {
			yn = (y[i+1]-y[i])*(xn-x[i])/(x[i+1]-x[i]) + y[i]
			Y[j] = yn
			X[j] = xn

			j++
		}
	}
}
