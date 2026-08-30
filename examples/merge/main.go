package main

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
)

func main() {
	data := [][]float64{{0, 1, 2, 3, 3, 3, 2, 0}, {5, 4, 2, 1, 4, 6, 6}}
	graph := asciigraph.PlotMany(data, asciigraph.MergeSeries())

	fmt.Println(graph)
	// Output:
	//  6.00 ┤    ╭─
	//  5.00 ┼╮   │
	//  4.00 ┤╰╮ ╭╯
	//  3.00 ┤ │╭┼─╮
	//  2.00 ┤ ╰╮│ ╰╮
	//  1.00 ┤╭╯╰╯  │
	//  0.00 ┼╯     ╰
}
