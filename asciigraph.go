package asciigraph

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

// Plot returns ascii graph for a series.
func Plot(series []float64, options ...Option) string {
	return PlotMany([][]float64{series}, options...)
}

// getCharSet returns the CharSet for a given series index, falling back to DefaultCharSet.
func getCharSet(config *config, seriesIndex int) CharSet {
	if seriesIndex < len(config.SeriesChars) {
		charSet := config.SeriesChars[seriesIndex]
		// Fill in any empty fields with defaults
		if charSet.Horizontal == "" {
			charSet.Horizontal = DefaultCharSet.Horizontal
		}
		if charSet.VerticalLine == "" {
			charSet.VerticalLine = DefaultCharSet.VerticalLine
		}
		if charSet.ArcDownRight == "" {
			charSet.ArcDownRight = DefaultCharSet.ArcDownRight
		}
		if charSet.ArcDownLeft == "" {
			charSet.ArcDownLeft = DefaultCharSet.ArcDownLeft
		}
		if charSet.ArcUpRight == "" {
			charSet.ArcUpRight = DefaultCharSet.ArcUpRight
		}
		if charSet.ArcUpLeft == "" {
			charSet.ArcUpLeft = DefaultCharSet.ArcUpLeft
		}
		if charSet.EndCap == "" {
			charSet.EndCap = DefaultCharSet.EndCap
		}
		if charSet.StartCap == "" {
			charSet.StartCap = DefaultCharSet.StartCap
		}
		return charSet
	}
	return DefaultCharSet
}

// PlotMany returns ascii graph for multiple series.
func PlotMany(data [][]float64, options ...Option) string {
	var logMaximum float64
	config := configure(config{
		Offset:     3,
		Precision:  nil,
		LineEnding: "\n",
	}, options)

	// Create a deep copy of the input data
	dataCopy := make([][]float64, len(data))
	for i, series := range data {
		dataCopy[i] = make([]float64, len(series))
		copy(dataCopy[i], series)
	}
	data = dataCopy

	lenMax := 0
	for i := range data {
		if l := len(data[i]); l > lenMax {
			lenMax = l
		}
	}

	if config.Width > 0 {
		for i := range data {
			for j := len(data[i]); j < lenMax; j++ {
				data[i] = append(data[i], math.NaN())
			}
			data[i] = interpolateArray(data[i], config.Width)
		}

		lenMax = config.Width
	}

	minimum, maximum := math.Inf(1), math.Inf(-1)
	for i := range data {
		minVal, maxVal := minMaxFloat64Slice(data[i])
		if minVal < minimum {
			minimum = minVal
		}
		if maxVal > maximum {
			maximum = maxVal
		}
	}
	if config.LowerBound != nil && *config.LowerBound < minimum {
		minimum = *config.LowerBound
	}
	if config.UpperBound != nil && *config.UpperBound > maximum {
		maximum = *config.UpperBound
	}
	interval := math.Abs(maximum - minimum)

	if config.Height <= 0 {
		config.Height = calculateHeight(interval)
	}

	if config.Offset <= 0 {
		config.Offset = 3
	}

	var ratio float64
	if interval != 0 {
		ratio = float64(config.Height) / interval
	} else {
		ratio = 1
	}
	min2 := round(minimum * ratio)
	max2 := round(maximum * ratio)

	intmin2 := int(min2)
	intmax2 := int(max2)

	rows := int(math.Abs(float64(intmax2 - intmin2)))
	width := lenMax + config.Offset

	type cell struct {
		Text  string
		Color AnsiColor
	}
	plot := make([][]cell, rows+1)

	// initialise empty 2D grid
	for i := 0; i < rows+1; i++ {
		line := make([]cell, width)
		for j := 0; j < width; j++ {
			line[j].Text = " "
			line[j].Color = Default
		}
		plot[i] = line
	}

	var precision uint = 2 //Default precision to maintain backwards compatibility
	if config.Precision != nil {
		precision = *config.Precision
	}

	logMaximum = math.Log10(math.Max(math.Abs(maximum), math.Abs(minimum))) //to find number of zeros after decimal
	if minimum == float64(0) && maximum == float64(0) {
		logMaximum = float64(-1)
	}

	if logMaximum < 0 {
		// negative log
		if math.Mod(logMaximum, 1) != 0 {
			// non-zero digits after decimal
			precision += uint(math.Abs(logMaximum))
		} else {
			precision += uint(math.Abs(logMaximum) - 1.0)
		}
	} else if logMaximum > 2 && config.Precision == nil {
		precision = 0
	}

	maxNumLength := utf8.RuneCountInString(fmt.Sprintf("%0.*f", precision, maximum))
	minNumLength := utf8.RuneCountInString(fmt.Sprintf("%0.*f", precision, minimum))
	magnitudes := make([]float64, 0, rows+1)
	if config.YAxisValueFormatter != nil {
		maxNumLength = 0
	}

	// calculate Y-axis values and the width when formatted using the YAxisValueFormatter
	for y := intmin2; y < intmax2+1; y++ {
		var magnitude float64
		if rows > 0 && interval > 0 {
			magnitude = maximum - (float64(y-intmin2) * interval / float64(rows))
		} else if interval == 0 {
			magnitude = minimum
		} else {
			magnitude = float64(y)
		}
		magnitudes = append(magnitudes, magnitude)

		if config.YAxisValueFormatter != nil {
			l := utf8.RuneCountInString(config.YAxisValueFormatter(magnitude))
			if l > maxNumLength {
				maxNumLength = l
			}
		}
	}
	var maxWidth int
	if config.YAxisValueFormatter != nil {
		maxWidth = maxNumLength
	} else {
		maxWidth = int(math.Max(float64(maxNumLength), float64(minNumLength)))
	}
	leftPad := config.Offset + maxWidth

	// axis and labels reusing the previously calculated values
	for w, magnitude := range magnitudes {
		var label string
		if config.YAxisValueFormatter == nil {
			label = fmt.Sprintf("%*.*f", maxWidth+1, precision, magnitude)
		} else {
			val := config.YAxisValueFormatter(magnitude)
			label = strings.Repeat(" ", maxWidth+1-utf8.RuneCountInString(val)) + val
		}

		h := int(math.Max(float64(config.Offset)-float64(utf8.RuneCountInString(label)), 0))

		plot[w][h].Text = label
		plot[w][h].Color = config.LabelColor
		plot[w][config.Offset-1].Text = "┤"
		plot[w][config.Offset-1].Color = config.AxisColor
	}

	for i := range data {
		series := data[i]

		color := Default
		if i < len(config.SeriesColors) {
			color = config.SeriesColors[i]
		}

		// Get the character set for this series
		charSet := getCharSet(config, i)

		var y0, y1 int

		if !math.IsNaN(series[0]) {
			y0 = int(round(series[0]*ratio) - min2)
			plot[rows-y0][config.Offset-1].Text = "┼" // first value
			plot[rows-y0][config.Offset-1].Color = config.AxisColor
		}

		for x := 0; x < len(series)-1; x++ { // plot the line
			d0 := series[x]
			d1 := series[x+1]

			if math.IsNaN(d0) && math.IsNaN(d1) {
				continue
			}

			if math.IsNaN(d1) && !math.IsNaN(d0) {
				y0 = int(round(d0*ratio) - float64(intmin2))
				plot[rows-y0][x+config.Offset].Text = charSet.EndCap
				plot[rows-y0][x+config.Offset].Color = color
				continue
			}

			if math.IsNaN(d0) && !math.IsNaN(d1) {
				y1 = int(round(d1*ratio) - float64(intmin2))
				plot[rows-y1][x+config.Offset].Text = charSet.StartCap
				plot[rows-y1][x+config.Offset].Color = color
				continue
			}

			y0 = int(round(d0*ratio) - float64(intmin2))
			y1 = int(round(d1*ratio) - float64(intmin2))

			if y0 == y1 {
				plot[rows-y0][x+config.Offset].Text = charSet.Horizontal
			} else {
				if y0 > y1 {
					plot[rows-y1][x+config.Offset].Text = charSet.ArcUpRight
					plot[rows-y0][x+config.Offset].Text = charSet.ArcDownLeft
				} else {
					plot[rows-y1][x+config.Offset].Text = charSet.ArcDownRight
					plot[rows-y0][x+config.Offset].Text = charSet.ArcUpLeft
				}

				start := int(math.Min(float64(y0), float64(y1))) + 1
				end := int(math.Max(float64(y0), float64(y1)))
				for y := start; y < end; y++ {
					plot[rows-y][x+config.Offset].Text = charSet.VerticalLine
				}
			}

			start := int(math.Min(float64(y0), float64(y1)))
			end := int(math.Max(float64(y0), float64(y1)))
			for y := start; y <= end; y++ {
				plot[rows-y][x+config.Offset].Color = color
			}
		}
	}

	// join columns
	var lines bytes.Buffer
	for h, horizontal := range plot {
		if h != 0 {
			lines.WriteString(config.LineEnding)
		}

		// remove trailing spaces
		lastCharIndex := 0
		for i := width - 1; i >= 0; i-- {
			if horizontal[i].Text != " " {
				lastCharIndex = i
				break
			}
		}

		c := Default
		for _, v := range horizontal[:lastCharIndex+1] {
			if v.Color != c {
				c = v.Color
				lines.WriteString(c.String())
			}

			lines.WriteString(v.Text)
		}
		if c != Default {
			lines.WriteString(Default.String())
		}
	}

	// add x-axis if configured
	if config.XAxisRange != nil {
		addXAxis(&lines, config, lenMax, leftPad)
	}

	// add caption if not empty
	if config.Caption != "" {
		lines.WriteString(config.LineEnding)
		lines.WriteString(strings.Repeat(" ", leftPad))
		if len(config.Caption) < lenMax {
			lines.WriteString(strings.Repeat(" ", (lenMax-len(config.Caption))/2))
		}
		if config.CaptionColor != Default {
			lines.WriteString(config.CaptionColor.String())
		}
		lines.WriteString(config.Caption)
		if config.CaptionColor != Default {
			lines.WriteString(Default.String())
		}
	}

	if len(config.SeriesLegends) > 0 {
		addLegends(&lines, config, lenMax, leftPad)
	}

	return lines.String()
}

// defaultXAxisFormatter formats X-axis tick values using %g.
var defaultXAxisFormatter XAxisValueFormatterFunc = func(v float64) string {
	return fmt.Sprintf("%g", v)
}

// addXAxis appends an X-axis line and tick labels below the plot body.
func addXAxis(lines *bytes.Buffer, config *config, lenMax int, leftPad int) {
	if lenMax <= 0 {
		return
	}

	xMin := config.XAxisRange[0]
	xMax := config.XAxisRange[1]

	tickCount := config.XAxisTickCount
	if lenMax == 1 {
		tickCount = 1
	} else if tickCount < 2 {
		tickCount = 5
	}
	if tickCount > lenMax {
		tickCount = lenMax
	}

	formatter := config.XAxisValueFormatter
	if formatter == nil {
		formatter = defaultXAxisFormatter
	}

	// compute tick column positions and labels
	type tick struct {
		col   int
		label string
	}
	ticks := make([]tick, tickCount)
	for i := 0; i < tickCount; i++ {
		var col int
		var value float64
		if tickCount == 1 {
			col = 0
			value = xMin
		} else {
			col = i * (lenMax - 1) / (tickCount - 1)
			value = xMin + float64(col)/float64(lenMax-1)*(xMax-xMin)
		}
		ticks[i] = tick{col: col, label: formatter(value)}
	}

	// axis line: leftPad-1 spaces + └ + ─/┬ characters
	totalWidth := leftPad + lenMax
	axisLine := make([]rune, totalWidth)
	for i := range axisLine {
		axisLine[i] = ' '
	}
	axisLine[leftPad-1] = '└'
	for i := 0; i < lenMax; i++ {
		axisLine[leftPad+i] = '─'
	}
	for _, tk := range ticks {
		axisLine[leftPad+tk.col] = '┬'
	}

	// write axis line with colors
	lines.WriteString(config.LineEnding)
	axisStr := strings.TrimRight(string(axisLine), " ")
	if config.AxisColor != Default {
		lines.WriteString(config.AxisColor.String())
	}
	lines.WriteString(axisStr)
	if config.AxisColor != Default {
		lines.WriteString(Default.String())
	}

	// label line: place each label centered on its tick column
	maxRightExtent := totalWidth
	for _, tk := range ticks {
		labelLen := utf8.RuneCountInString(tk.label)
		endCol := leftPad + tk.col + (labelLen - labelLen/2)
		if endCol > maxRightExtent {
			maxRightExtent = endCol
		}
	}
	labelLine := make([]rune, maxRightExtent)
	for i := range labelLine {
		labelLine[i] = ' '
	}

	lastEnd := -1 // tracks the rightmost column used by the previous label
	for _, tk := range ticks {
		labelRunes := []rune(tk.label)
		labelLen := len(labelRunes)

		// center the label on the tick column
		startCol := leftPad + tk.col - labelLen/2
		if startCol < 0 {
			startCol = 0
		}

		// skip if this label would overlap the previous one (need 1-space gap)
		if startCol <= lastEnd {
			continue
		}

		for j, r := range labelRunes {
			pos := startCol + j
			if pos < len(labelLine) {
				labelLine[pos] = r
			}
		}
		lastEnd = startCol + labelLen
	}

	// trim and write label line
	labelStr := strings.TrimRight(string(labelLine), " ")
	if labelStr != "" {
		lines.WriteString(config.LineEnding)
		if config.LabelColor != Default {
			lines.WriteString(config.LabelColor.String())
		}
		lines.WriteString(labelStr)
		if config.LabelColor != Default {
			lines.WriteString(Default.String())
		}
	}
}
