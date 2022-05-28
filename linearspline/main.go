package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	f := func(x float64) float64 { return 2*math.Pow(x, 8) + 3*math.Pow(x, 7) + 5*math.Pow(x, 5) - 2 }
	var N int

	h := 0.01

	n := 6

	x := make([]float64, n)
	y := make([]float64, n)

	x = []float64{1, 2, 3, 4, 5, 6}
	y = []float64{1, 2, 3, 4, 5, 6}

	for i := 0; i < n-1; i++ {
		N += int((x[i+1] - x[i]) / h)
	}

	N = N + 2 + (n - 4)

	fmt.Printf("\nThe no. of interpolated values N= %d\n", N)

	Y := make([]float64, N)
	X := make([]float64, N)

	linearSpline(x, y, h, X, Y, f)

	fp, err := os.Create("linearSpline.txt")
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < N; i++ {
		_, err = fp.WriteString(fmt.Sprintf("%f\t%f\n", X[i], Y[i]))
		if err != nil {
			log.Println(err)
		}
	}

	drawTwoDimensionalGraph(X, Y)
}
