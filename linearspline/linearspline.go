package main

/* Function to perform piecewise linear spline interpolation
Parameters:
x[n]: x-axis points
y[n]: y-axis points
X[n]: array that stores the interpolated x-axis points
Y[n]: array that stores the interpolated y-axis points.
*/
func linearSpline(x, y []float64, h float64, X, Y []float64, f func(x float64) float64) {
	n := len(x)
	j := 0

	for i := 0; i < n-1; i++ {
		var yn, xn float64

		for xn = x[i]; xn < x[i+1]; xn += h {
			yn = f(xn)
			Y[j] = yn
			X[j] = xn

			j++
		}
	}
}
