package asciigraph

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

// Create legend item as a colored box and text
func createLegendItem(text string, color AnsiColor) (string, int) {
	return fmt.Sprintf(
			"%s■%s %s",
			color.String(),
			Default.String(),
			text,
		),
		// Can't use len() because of AnsiColor, add 2 for box and space
		utf8.RuneCountInString(text) + 2
}

// Add legend for each series added to the graph
func addLegends(lines *bytes.Buffer, config *config, lenMax int, leftPad int) {
	lines.WriteString("\n\n")
	lines.WriteString(strings.Repeat(" ", leftPad))

	var legendsText string
	var legendsTextLen int
	rightPad := 3
	for i, text := range config.SeriesLegends {
		item, itemLen := createLegendItem(text, config.SeriesColors[i])
		legendsText += item
		legendsTextLen += itemLen

		if i < len(config.SeriesLegends)-1 {
			legendsText += strings.Repeat(" ", rightPad)
			legendsTextLen += rightPad
		}
	}

	if legendsTextLen < lenMax {
		lines.WriteString(strings.Repeat(" ", (lenMax-legendsTextLen)/2))
	}
	lines.WriteString(legendsText)
}
