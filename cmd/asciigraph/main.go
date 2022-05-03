package main

import (
	"bufio"
	"errors"
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
	precision          uint = 2
	caption            string
	enableRealTime     bool
	realTimeDataBuffer int
	fps                float64 = 24
	seriesColor        asciigraph.AnsiColor
	captionColor       asciigraph.AnsiColor
	axisColor          asciigraph.AnsiColor
	labelColor         asciigraph.AnsiColor
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
	flag.UintVar(&precision, "p", precision, "`precision` of data point labels along the y-axis")
	flag.StringVar(&caption, "c", caption, "`caption` for the graph")
	flag.BoolVar(&enableRealTime, "r", enableRealTime, "enables `realtime` graph for data stream")
	flag.IntVar(&realTimeDataBuffer, "b", realTimeDataBuffer, "data points `buffer` when realtime graph enabled, default equal to `width`")
	flag.Float64Var(&fps, "f", fps, "set `fps` to control how frequently graph to be rendered when realtime graph enabled")
	flag.Func("sc", "`series color` of the plot", func(str string) error {
		if c, ok := asciigraph.ColorNames[str]; !ok {
			return errors.New("unrecognized color, check available color names at https://www.w3.org/TR/SVG11/types.html#ColorKeywords")
		} else {
			seriesColor = c
			return nil
		}
	})

	flag.Func("cc", "`caption color` of the plot", func(str string) error {
		if c, ok := asciigraph.ColorNames[str]; !ok {
			return errors.New("unrecognized color, check available color names at https://www.w3.org/TR/SVG11/types.html#ColorKeywords")
		} else {
			captionColor = c
			return nil
		}
	})

	flag.Func("ac", "y-`axis color` of the plot", func(str string) error {
		if c, ok := asciigraph.ColorNames[str]; !ok {
			return errors.New("unrecognized color, check available color names at https://www.w3.org/TR/SVG11/types.html#ColorKeywords")
		} else {
			axisColor = c
			return nil
		}
	})

	flag.Func("lc", "y-axis `label color` of the plot", func(str string) error {
		if c, ok := asciigraph.ColorNames[str]; !ok {
			return errors.New("unrecognized color, check available color names at https://www.w3.org/TR/SVG11/types.html#ColorKeywords")
		} else {
			labelColor = c
			return nil
		}
	})

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
					asciigraph.Precision(precision),
					asciigraph.Caption(caption),
					asciigraph.SeriesColors(seriesColor),
					asciigraph.CaptionColor(captionColor),
					asciigraph.AxisColor(axisColor),
					asciigraph.LabelColor(labelColor),
				)
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
			asciigraph.Precision(precision),
			asciigraph.Caption(caption),
			asciigraph.SeriesColors(seriesColor),
			asciigraph.CaptionColor(captionColor),
			asciigraph.AxisColor(axisColor),
			asciigraph.LabelColor(labelColor),
		)

		fmt.Println(plot)
	}
}
