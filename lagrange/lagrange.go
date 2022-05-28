package main

func lagrange(x, y []float64, xn float64) float64 {
	var result float64

	n := len(x)

	for i := 0; i < n; i++ {
		term := y[i]

		for j := 0; j < n; j++ {
			if j != i {
				term = term * (xn - x[j]) / (x[i] - x[j])
			}
		}

		result += term
	}

	return result
}
