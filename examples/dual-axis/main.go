package main

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
)

func main() {
	primary := []float64{1, 2, 3, 4, 3, 2, 1}
	secondary := []float64{10, 20, 40, 30, 20, 15, 10}

	graph := asciigraph.PlotMany([][]float64{primary, secondary},
		asciigraph.DualAxis(),
		asciigraph.SeriesColors(asciigraph.Red, asciigraph.Blue),
		asciigraph.SeriesLegends("Temperature", "Pressure"),
	)

	fmt.Println(graph)
}
