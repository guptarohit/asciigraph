package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/guptarohit/asciigraph"
)

var (
	height  uint
	width   uint
	offset  uint = 3
	caption string
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "%s expects data points from stdin. Invalid values are logged to stderr.\n", os.Args[0])
	}
	flag.UintVar(&height, "h", height, "`height` in text rows, 0 for auto-scaling")
	flag.UintVar(&width, "w", width, "`width` in columns, 0 for auto-scaling")
	flag.UintVar(&offset, "o", offset, "`offset` in columns, for the label")
	flag.StringVar(&caption, "c", caption, "`caption` for the graph")
	flag.Parse()

	data := make([]float64, 0, 64)

	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		word := s.Text()
		p, err := strconv.ParseFloat(word, 64)
		if err != nil {
			log.Printf("ignore %q: cannot parse value", word)
			continue
		}
		data = append(data, p)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	if len(data) == 0 {
		log.Fatal("no data")
	}

	plot := asciigraph.Plot(data,
		asciigraph.Height(int(height)),
		asciigraph.Width(int(width)),
		asciigraph.Offset(int(offset)),
		asciigraph.Caption(caption))

	fmt.Println(plot)
}
