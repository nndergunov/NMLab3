package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func f(x float64) float64 {
	return 3*math.Pow(x, 17) + 5*math.Pow(x, 16)
}

func main() {
	N := 45
	resNumber := 16

	x := make([]float64, resNumber)
	y := make([]float64, resNumber)

	for i := 1; i <= resNumber; i++ {
		x[i-1] = float64(i)
		y[i-1] = f(float64(i))
	}

	pol := buildPolynomial(x, y)

	for i := resNumber + 1; i <= N; i++ {
		y = append(y, lagrange(x, y, float64(i)))
		x = append(x, float64(i))
	}

	fp, err := os.Create("lagrange.txt")
	if err != nil {
		log.Println(err)
	}

	fp.WriteString(pol)

	for i := 0; i < N; i++ {
		_, err = fp.WriteString(fmt.Sprintf("%f\t%f\n", x[i], y[i]))
		if err != nil {
			log.Println(err)
		}
	}

	drawTwoDimensionalGraph(x, y)
}
