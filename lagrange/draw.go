package main

import (
	"github.com/Arafatk/glot"
	"log"
)

func drawTwoDimensionalGraph(x, y []float64) {
	dimensions := 2

	plot, _ := glot.NewPlot(dimensions, false, false)

	err := plot.AddPointGroup("Lagrange approximation", "lines", [][]float64{x, y})
	if err != nil {
		log.Println(err)
	}

	plot.SetTitle("Lagrange approximation")

	plot.SetXLabel("X")
	plot.SetYLabel("Y")

	err = plot.SavePlot("lagrange.png")
	if err != nil {
		log.Println(err)
	}
}
