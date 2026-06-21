package main

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
)

func main() {
	// CPU usage readings (%) over time, with a couple of spikes and dips.
	data := []float64{42, 48, 55, 81, 85, 91, 87, 34, 12, 17, 10, 18, 55, 50}

	// ColorAbove and ColorBelow highlight points that breach a threshold,
	// regardless of which series they belong to. Points in between keep the
	// series' own color (or the default if none is set).
	graph := asciigraph.Plot(data,
		asciigraph.Height(10),
		asciigraph.Width(25),
		asciigraph.LowerBound(0),
		asciigraph.UpperBound(100),
		asciigraph.Caption("CPU usage % (red: critical, green: idle)"),
		asciigraph.ColorAbove(asciigraph.Red, 80),
		asciigraph.ColorBelow(asciigraph.Green, 25),
	)

	fmt.Println(graph)
}
