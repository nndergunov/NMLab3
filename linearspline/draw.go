package main

import (
	"github.com/Arafatk/glot"
)

func drawTwoDimensionalGraph(x, y []float64) {
	dimensions := 2

	plot, _ := glot.NewPlot(dimensions, false, false)

	plot.AddPointGroup("Spline", "lines", [][]float64{x, y})

	plot.SetTitle("Linear interpolation spline")

	plot.SetXLabel("X")
	plot.SetYLabel("Y")

	// plot.SetXrange(1, 6)
	// plot.SetYrange(-1, 5000000)

	plot.SavePlot("splineGraph.png")
}
