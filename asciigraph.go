package asciigraph

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

type Graph struct {
	series []float64
	config *config

	maxWidth  int
	intmin2   int
	intmax2   int
	rows      int
	plot      [][]string
	maximum   float64
	minimum   float64
	ratio     float64
	interval  float64
	min2      float64
	precision int
}

func NewGraph(series []float64, options ...Option) *Graph {
	g := new(Graph)

	g.series = series
	g.config = configure(config{
		Offset: 3,
	}, options)

	var logMaximum float64

	if g.config.Width > 0 {
		g.series = interpolateArray(g.series, g.config.Width)
	}
	g.minimum, g.maximum = minMaxFloat64Slice(g.series)
	g.interval = math.Abs(g.maximum - g.minimum)

	if g.config.Height <= 0 {
		if int(g.interval) <= 0 {
			g.config.Height = int(g.interval * math.Pow10(int(math.Ceil(-math.Log10(g.interval)))))
		} else {
			g.config.Height = int(g.interval)
		}
	}

	if g.config.Offset <= 0 {
		g.config.Offset = 3
	}

	if g.interval != 0 {
		g.ratio = float64(g.config.Height) / g.interval
	} else {
		g.ratio = 1
	}
	g.min2 = round(g.minimum * g.ratio)
	max2 := round(g.maximum * g.ratio)

	g.intmin2 = int(g.min2)
	g.intmax2 = int(max2)

	g.rows = int(math.Abs(float64(g.intmax2 - g.intmin2)))
	width := len(g.series) + g.config.Offset

	// initialise empty 2D grid
	for i := 0; i < g.rows+1; i++ {
		var line []string
		for j := 0; j < width; j++ {
			line = append(line, " ")
		}
		g.plot = append(g.plot, line)
	}

	g.precision = 2
	logMaximum = math.Log10(math.Max(math.Abs(g.maximum), math.Abs(g.minimum))) //to find number of zeros after decimal
	if g.minimum == float64(0) && g.maximum == float64(0) {
		logMaximum = float64(-1)
	}

	if logMaximum < 0 {
		// negative log
		if math.Mod(logMaximum, 1) != 0 {
			// non-zero digits after decimal
			g.precision = g.precision + int(math.Abs(logMaximum))
		} else {
			g.precision = g.precision + int(math.Abs(logMaximum)-1.0)
		}
	} else if logMaximum > 2 {
		g.precision = 0
	}

	maxNumLength := int(len(fmt.Sprintf("%0.*f", g.precision, g.maximum)))
	minNumLength := int(len(fmt.Sprintf("%0.*f", g.precision, g.minimum)))
	g.maxWidth = int(math.Max(float64(maxNumLength), float64(minNumLength)))

	return g
}

func (g *Graph) Plot() string {

	// axis and labels
	for y := g.intmin2; y < g.intmax2+1; y++ {
		var magnitude float64
		if g.rows > 0 {
			magnitude = g.maximum - (float64(y-g.intmin2) * g.interval / float64(g.rows))
		} else {
			magnitude = float64(y)
		}

		label := fmt.Sprintf("%*.*f", g.maxWidth+1, g.precision, magnitude)
		w := y - g.intmin2
		h := int(math.Max(float64(g.config.Offset)-float64(len(label)), 0))

		g.plot[w][h] = label
		if y == 0 {
			g.plot[w][g.config.Offset-1] = "┼"
		} else {
			g.plot[w][g.config.Offset-1] = "┤"
		}
	}

	y0 := int(round(g.series[0]*g.ratio) - g.min2)
	var y1 int

	g.plot[g.rows-y0][g.config.Offset-1] = "┼" // first value

	for x := 0; x < len(g.series)-1; x++ { // g.plot the line
		y0 = int(round(g.series[x+0]*g.ratio) - float64(g.intmin2))
		y1 = int(round(g.series[x+1]*g.ratio) - float64(g.intmin2))
		if y0 == y1 {
			g.plot[g.rows-y0][x+g.config.Offset] = "─"
		} else {
			if y0 > y1 {
				g.plot[g.rows-y1][x+g.config.Offset] = "╰"
				g.plot[g.rows-y0][x+g.config.Offset] = "╮"
			} else {
				g.plot[g.rows-y1][x+g.config.Offset] = "╭"
				g.plot[g.rows-y0][x+g.config.Offset] = "╯"
			}

			start := int(math.Min(float64(y0), float64(y1))) + 1
			end := int(math.Max(float64(y0), float64(y1)))
			for y := start; y < end; y++ {
				g.plot[g.rows-y][x+g.config.Offset] = "│"
			}
		}
	}

	// join columns
	var lines bytes.Buffer
	for h, horizontal := range g.plot {
		if h != 0 {
			lines.WriteRune('\n')
		}
		for _, v := range horizontal {
			lines.WriteString(v)
		}
	}

	// add caption if not empty
	if g.config.Caption != "" {
		lines.WriteRune('\n')
		lines.WriteString(strings.Repeat(" ", g.config.Offset+g.maxWidth+2))
		lines.WriteString(g.config.Caption)
	}
	fmt.Printf("\n")
	fmt.Printf(lines.String())
	fmt.Printf("\n")
	return lines.String()
}

// Plot returns ascii graph for a series.
func Plot(series []float64, options ...Option) string {
	g := NewGraph(series, options...)
	return g.Plot()
}
