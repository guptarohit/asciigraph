package main

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
)

func main() {
	// CPU usage readings (%) over time, with a couple of spikes and dips.
	data := []float64{42, 48, 55, 78, 91, 85, 60, 38, 22, 15, 25, 50, 65, 50, 45}

	// ColorAbove and ColorBelow highlight points that breach a threshold,
	// regardless of which series they belong to. Points in between keep the
	// series' own color (or the default if none is set).
	graph := asciigraph.Plot(data,
		asciigraph.Height(10),
		asciigraph.Caption("CPU usage % (red: critical, yellow: idle)"),
		asciigraph.ColorAbove(80, asciigraph.Red),
		asciigraph.ColorBelow(20, asciigraph.Yellow),
	)

	fmt.Println(graph)
}
