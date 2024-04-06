package asciigraph

// Plot returns ascii graph for a series.
// Deprecated: Use Graph struct and its Plot method instead.
func Plot(series []float64, options ...Option) string {
	return PlotMany([][]float64{series}, options...)
}

// PlotMany returns ascii graph for multiple series.
// Deprecated: Use Graph struct and its Plot method instead.
func PlotMany(data [][]float64, options ...Option) string {
	g := NewGraph(data, options...)
	return g.Plot()
}
