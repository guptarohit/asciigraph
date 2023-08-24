package asciigraph

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

// Get legend item as a colored box and text
func getLegendItem(text string, color AnsiColor) (string, int) {
	return fmt.Sprintf(
		"%s■%s %s",
		color.String(),
		Default.String(),
		text,
	// Can't use len() because of AnsiColor, add 2 for box and space
	), utf8.RuneCountInString(text) + 2
}

// Add legend for each series added to the graph
func addLegend(lines* bytes.Buffer, config* config, lenMax int, leftPad int) {
	lines.WriteString("\n\n")
	lines.WriteString(strings.Repeat(" ", leftPad))
	
	var legendText string
	var legendLen int
	rightPad := 3
	for i, text := range config.LegendText {
		item, itemLen := getLegendItem(text, config.SeriesColors[i])
		legendText += item
		legendLen += itemLen

		if i < len(config.LegendText) - 1 {
			legendText += strings.Repeat(" ", rightPad)
			legendLen += rightPad
		}
	}
	
	if legendLen < lenMax {
		lines.WriteString(strings.Repeat(" ", (lenMax-legendLen)/2))
	}
	lines.WriteString(legendText)
}
