package main

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
)

func main() {
	data := []float64{3, 4, 9, 6, 2, 4, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data,
		asciigraph.XAxisRange(0, 14),
		asciigraph.XAxisTickCount(3),
	)

	fmt.Println(graph)
	// Output:
	//  10.00 ┤        ╭╮
	//   9.00 ┤ ╭╮     ││
	//   8.00 ┤ ││   ╭╮││
	//   7.00 ┤ ││   ││││╭╮
	//   6.00 ┤ │╰╮  ││││││ ╭
	//   5.00 ┤ │ │ ╭╯╰╯│││╭╯
	//   4.00 ┤╭╯ │╭╯   ││││
	//   3.00 ┼╯  ││    ││││
	//   2.00 ┤   ╰╯    ╰╯╰╯
	//        └┬──────┬──────┬
	//         0      7     14
}
