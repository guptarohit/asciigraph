package main

import (
	"fmt"
	"math"

	"github.com/guptarohit/asciigraph"
)

func main() {
	data := make([][]float64, 6)

	// concentric semi-circles
	for i := 0; i < 6; i++ {
		for x := -40; x <= 40; x++ {
			v := math.NaN()
			if r := 40 - i; x >= -r && x <= r {
				v = math.Sqrt(math.Pow(float64(r), 2)-math.Pow(float64(x), 2)) / 2
			}
			data[i] = append(data[i], v)
		}
	}
	graph := asciigraph.PlotMany(data, asciigraph.Precision(0), asciigraph.SeriesColors(
		asciigraph.Red,
		asciigraph.Orange,
		asciigraph.Yellow,
		asciigraph.Green,
		asciigraph.Blue,
		asciigraph.Purple,
	))

	fmt.Println(graph)
	// Output:
	//   20 ┤                               ╭───────╭╮───────╮
	//   19 ┤                        ╭──╭───╭───────╭╮───────╮───╮──╮
	//   18 ┤                    ╭─╭──╭─╭───╭───────╭╮───────╮───╮─╮──╮─╮
	//   17 ┤                 ╭─╭─╭─╭─╭──╭──────────╯╰──────────╮──╮─╮─╮─╮─╮
	//   16 ┤              ╭─╭─╭╭─╭─╭────╯                      ╰────╮─╮─╮╮─╮─╮
	//   15 ┤            ╭╭─╭─╭╭─╭──╯                                ╰──╮─╮╮─╮─╮╮
	//   14 ┤          ╭╭─╭╭─╭╭──╯                                      ╰──╮╮─╮╮─╮╮
	//   13 ┤        ╭─╭╭╭─╭╭─╯                                            ╰─╮╮─╮╮╮─╮
	//   12 ┤       ╭╭╭─╭╭╭─╯                                                ╰─╮╮╮─╮╮╮
	//   11 ┤     ╭─╭╭╭╭╭─╯                                                    ╰─╮╮╮╮╮─╮
	//   10 ┤    ╭╭─╭╭╭╭╯                                                        ╰╮╮╮╮─╮╮
	//    9 ┤   ╭╭╯╭╭╭╭╯                                                          ╰╮╮╮╮╰╮╮
	//    8 ┤  ╭╭╯╭╭╭╭╯                                                            ╰╮╮╮╮╰╮╮
	//    7 ┤  ││╭╭╭╭╯                                                              ╰╮╮╮╮││
	//    6 ┤ ╭╭╭╭╭╭╯                                                                ╰╮╮╮╮╮╮
	//    5 ┤ ││││││                                                                  ││││││
	//    4 ┤╭╭╭╭╭╭╯                                                                  ╰╮╮╮╮╮╮
	//    3 ┤││││││                                                                    ││││││
	//    2 ┤││││││                                                                    ││││││
	//    1 ┤││││││                                                                    ││││││
	//    0 ┼╶╶╶╶╶╯                                                                    ╰╴╴╴╴╴
}
