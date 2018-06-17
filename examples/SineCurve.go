package main

import (
	"fmt"
	"github.com/guptarohit/asciigraph"
	"math"
)

func main() {
	var data []float64

	// sine curve
	for i := 0; i < 105; i++ {
		data = append(data, 15*math.Sin(float64(i)*((math.Pi*4)/120.0)))
	}

	conf := map[string]interface{}{"height": 10}

	graph := asciigraph.Plot(data, conf)

	fmt.Println(graph)
}
