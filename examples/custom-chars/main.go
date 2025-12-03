package main

import (
	"fmt"
	"math"

	"github.com/guptarohit/asciigraph"
)

func main() {
	// Create some sample data
	data1 := make([]float64, 0)
	data2 := make([]float64, 0)
	data3 := make([]float64, 0)

	for i := 0; i < 30; i++ {
		data1 = append(data1, math.Sin(float64(i)*0.2)*5+10)
		data2 = append(data2, math.Cos(float64(i)*0.2)*3+10)
		data3 = append(data3, math.Sin(float64(i)*0.3)*4+8)
	}

	// Example 1: Using default box-drawing characters
	fmt.Println("Example 1: Default box-drawing characters")
	fmt.Println("==========================================")
	graph1 := asciigraph.PlotMany([][]float64{data1},
		asciigraph.Height(10),
		asciigraph.Caption("Default Characters"),
	)
	fmt.Println(graph1)
	fmt.Println()

	// Example 2: Using asterisks for one series
	fmt.Println("Example 2: Custom asterisk characters")
	fmt.Println("======================================")
	graph2 := asciigraph.PlotMany([][]float64{data1},
		asciigraph.Height(10),
		asciigraph.SeriesChars(asciigraph.CreateCharSet("*")),
		asciigraph.Caption("Asterisk Characters"),
	)
	fmt.Println(graph2)
	fmt.Println()

	// Example 3: Using dots for one series
	fmt.Println("Example 3: Custom dot characters")
	fmt.Println("=================================")
	graph3 := asciigraph.PlotMany([][]float64{data1},
		asciigraph.Height(10),
		asciigraph.SeriesChars(asciigraph.CreateCharSet("•")),
		asciigraph.SeriesColors(asciigraph.Green),
		asciigraph.Caption("Dot Characters (Green)"),
	)
	fmt.Println(graph3)
	fmt.Println()

	// Example 4: Multiple series with different characters
	fmt.Println("Example 4: Multiple series with different characters")
	fmt.Println("=====================================================")
	graph4 := asciigraph.PlotMany([][]float64{data1, data2, data3},
		asciigraph.Height(12),
		asciigraph.SeriesChars(
			asciigraph.CreateCharSet("*"),
			asciigraph.CreateCharSet("#"),
			asciigraph.CreateCharSet("+"),
		),
		asciigraph.SeriesColors(asciigraph.Red, asciigraph.Green, asciigraph.Blue),
		asciigraph.SeriesLegends("Series 1 (*)", "Series 2 (#)", "Series 3 (+)"),
		asciigraph.Caption("Three Series with Different Characters"),
	)
	fmt.Println(graph4)
	fmt.Println()

	// Example 5: Partial character set (some fields use defaults)
	fmt.Println("Example 5: Partial character set (mixed with defaults)")
	fmt.Println("========================================================")
	partialSet := asciigraph.CharSet{
		Horizontal:   "=",
		VerticalLine: "|",
		// Other fields will use defaults
	}
	graph5 := asciigraph.PlotMany([][]float64{data1},
		asciigraph.Height(10),
		asciigraph.SeriesChars(partialSet),
		asciigraph.SeriesColors(asciigraph.Cyan),
		asciigraph.Caption("Partial CharSet (= and | with default corners)"),
	)
	fmt.Println(graph5)
	fmt.Println()

	// Example 6: Using simple ASCII characters
	fmt.Println("Example 6: Simple ASCII-only characters")
	fmt.Println("========================================")
	asciiSet := asciigraph.CharSet{
		Horizontal:      "-",
		VerticalLine:    "|",
		CornerUpLeft:    "/",
		CornerUpRight:   "\\",
		CornerDownLeft:  "\\",
		CornerDownRight: "/",
		EndCap:          "-",
		StartCap:        "-",
	}
	graph6 := asciigraph.PlotMany([][]float64{data1},
		asciigraph.Height(10),
		asciigraph.SeriesChars(asciiSet),
		asciigraph.Caption("ASCII-only Characters"),
	)
	fmt.Println(graph6)
}
