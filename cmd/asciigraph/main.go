package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/guptarohit/asciigraph"
)

var (
	height             uint
	width              uint
	offset             uint = 3
	caption            string
	enableRealTime     bool
	realTimeDataBuffer int
	fps                float64 = 24
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
	flag.BoolVar(&enableRealTime, "r", enableRealTime, "enables `realtime` graph for data stream")
	flag.IntVar(&realTimeDataBuffer, "b", realTimeDataBuffer, "data points `buffer` when realtime graph enabled, default equal to `width`")
	flag.Float64Var(&fps, "f", fps, "set `fps` to control how frequently graph to be rendered when realtime graph enabled")
	flag.Parse()

	data := make([]float64, 0, 64)

	if realTimeDataBuffer == 0 {
		realTimeDataBuffer = int(width)
	}

	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	nextFlushTime := time.Now()

	flushInterval := time.Duration(float64(time.Second) / fps)

	for s.Scan() {
		word := s.Text()
		p, err := strconv.ParseFloat(word, 64)
		if err != nil {
			log.Printf("ignore %q: cannot parse value", word)
			continue
		}
		data = append(data, p)
		if enableRealTime {
			if realTimeDataBuffer > 0 && len(data) > realTimeDataBuffer {
				data = data[len(data)-realTimeDataBuffer:]
			}

			if currentTime := time.Now(); currentTime.After(nextFlushTime) || currentTime.Equal(nextFlushTime) {
				plot := asciigraph.Plot(data,
					asciigraph.Height(int(height)),
					asciigraph.Width(int(width)),
					asciigraph.Offset(int(offset)),
					asciigraph.Caption(caption))
				asciigraph.Clear()
				fmt.Println(plot)
				nextFlushTime = time.Now().Add(flushInterval)
			}
		}
	}
	if !enableRealTime {
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
}
