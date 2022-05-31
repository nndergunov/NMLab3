package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

func buildPolynomial(x, y []float64) string {
	var pol string = "y = "

	n := len(x)

	for i := 0; i < n; i++ {
		pol += formatFloat(y[i])

		for j := 0; j < n; j++ {
			if j != i {
				pol += fmt.Sprintf(" * (x - %s) / (%s - %s)",
					formatFloat(x[j]), formatFloat(x[i]), formatFloat(x[j]))
			}
		}

		pol += " + "
	}

	pol = strings.TrimSuffix(pol, " + ")

	return pol
}

func formatFloat(f float64) string {
	w := strconv.FormatFloat(f, 'f', -1, 64)

	return w
}
