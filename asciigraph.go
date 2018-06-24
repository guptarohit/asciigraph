package asciigraph

import (
	"fmt"
	"math"
	"strings"
)

// Plot returns ascii graph for a series.
func Plot(series []float64, config map[string]interface{}) string {

	var height, offset int
	var caption string

	if val, ok := config["width"].(int); ok {
		series = interpolateArray(series, val)
	}

	minimum, maximum := minMaxFloat64Slice(series)

	interval := math.Abs(maximum - minimum)

	if val, ok := config["height"].(int); ok {
		height = val
	} else {
		if int(interval) <= 0 {
			height = int(interval * math.Pow10(int(math.Ceil(-math.Log10(interval)))))
		} else {
			height = int(interval)
		}
	}

	if val, ok := config["offset"].(int); ok {
		offset = val
	} else {
		offset = 3
	}

	if val, ok := config["caption"].(string); ok {
		caption = val
	} else {
		caption = ""
	}

	ratio := float64(height) / interval

	min2 := math.Floor(minimum * ratio)
	max2 := math.Ceil(maximum * ratio)

	intmin2 := int(min2)
	intmax2 := int(max2)

	rows := int(math.Abs(float64(intmax2 - intmin2)))
	width := len(series) + offset

	var plot [][]string

	// initialise empty 2D grid
	for i := 0; i < rows+1; i++ {
		var line []string
		for j := 0; j < width; j++ {
			line = append(line, " ")
		}
		plot = append(plot, line)
	}

	precision := 2
	logMaximum := math.Log10(math.Abs(maximum)) //to find number of zeros after decimal

	if logMaximum < 0 {
		// negative log
		if math.Mod(logMaximum, 1) != 0 {
			// non-zero digits after decimal
			precision = precision + int(math.Abs(logMaximum))
		} else {
			precision = precision + int(math.Abs(logMaximum)-1.0)
		}

	} else if logMaximum > 2 {
		precision = 0
	}

	maxNumLength := len(fmt.Sprintf("%0.*f", precision, maximum))
	minNumLength := len(fmt.Sprintf("%0.*f", precision, minimum))

	maxWidth := int(math.Max(float64(maxNumLength), float64(minNumLength)))

	// axis and labels
	for y := intmin2; y < intmax2+1; y++ {
		label := fmt.Sprintf("%*.*f", maxWidth+1, precision, maximum-(float64(y-intmin2)*interval/float64(rows)))
		w := y - intmin2
		h := int(math.Max(float64(offset)-float64(len(label)), 0))

		plot[w][h] = label
		if y == 0 {
			plot[w][offset-1] = "┼"
		} else {
			plot[w][offset-1] = "┤"
		}
	}

	y0 := int(series[0]*ratio - min2)

	var y1 int

	plot[rows-y0][offset-1] = "┼" // first value

	for x := 0; x < len(series)-1; x++ { // plot the line
		y0 = int(round(series[x+0]*ratio) - float64(intmin2))
		y1 = int(round(series[x+1]*ratio) - float64(intmin2))
		if y0 == y1 {
			plot[rows-y0][x+offset] = "─"
		} else {
			if y0 > y1 {
				plot[rows-y1][x+offset] = "╰"
			} else {
				plot[rows-y1][x+offset] = "╭"
			}
			if y0 > y1 {
				plot[rows-y0][x+offset] = "╮"
			} else {
				plot[rows-y0][x+offset] = "╯"
			}

			start := int(math.Min(float64(y0), float64(y1))) + 1
			end := int(math.Max(float64(y0), float64(y1)))
			for y := start; y < end; y++ {
				plot[rows-y][x+offset] = "│"
			}

		}

	}

	var lines []string

	// join columns
	for _, v := range plot {
		lines = append(lines, strings.Join(v, ""))
	}

	// add caption if not empty
	if caption != "" {
		lines = append(lines, fmt.Sprintf("%s", strings.Repeat(" ", offset+maxWidth+2)+caption))
	}

	return strings.Join(lines, "\n") // join rows
}
