package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
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
	seriesColors       []asciigraph.AnsiColor
	seriesLegends      []string
	captionColor       asciigraph.AnsiColor
	axisColor          asciigraph.AnsiColor
	labelColor         asciigraph.AnsiColor
	lowerBound              = math.Inf(1)
	upperBound              = math.Inf(-1)
	delimiter               = ","
	seriesNum          uint = 1
	customChar         string
	lastGraphLines     int // Track last graph height for clearing in realtime mode
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
	flag.Func("sc", "comma-separated `series colors` corresponding to each series", func(str string) error {
		if colors, ok := parseColors(str); !ok {
			return errors.New("unrecognized color, check available color names at https://www.w3.org/TR/SVG11/types.html#ColorKeywords")
		} else {
			seriesColors = colors
			return nil
		}
	})

	flag.Func("cc", "`caption color` of the plot", func(str string) error {
		if c, ok := parseColor(str); !ok {
			return errors.New("unrecognized color, check available color names at https://www.w3.org/TR/SVG11/types.html#ColorKeywords")
		} else {
			captionColor = c
			return nil
		}
	})

	flag.Func("ac", "y-`axis color` of the plot", func(str string) error {
		if c, ok := parseColor(str); !ok {
			return errors.New("unrecognized color, check available color names at https://www.w3.org/TR/SVG11/types.html#ColorKeywords")
		} else {
			axisColor = c
			return nil
		}
	})

	flag.Func("lc", "y-axis `label color` of the plot", func(str string) error {
		if c, ok := parseColor(str); !ok {
			return errors.New("unrecognized color, check available color names at https://www.w3.org/TR/SVG11/types.html#ColorKeywords")
		} else {
			labelColor = c
			return nil
		}
	})
	flag.Func("sl", "comma-separated `series legends` corresponding to each series", func(str string) error {
		for _, legend := range strings.Split(str, ",") {
			seriesLegends = append(seriesLegends, strings.TrimSpace(legend))
		}
		return nil
	})

	flag.Float64Var(&lowerBound, "lb", lowerBound, "`lower bound` set the minimum value for the vertical axis (ignored if series contains lower values)")
	flag.Float64Var(&upperBound, "ub", upperBound, "`upper bound` set the maximum value for the vertical axis (ignored if series contains larger values)")
	flag.StringVar(&delimiter, "d", delimiter, "data `delimiter` for splitting data points in the input stream")
	flag.UintVar(&seriesNum, "sn", seriesNum, "`number of series` (columns) in the input data")
	flag.StringVar(&customChar, "x", customChar, "`character` to use for plotting (e.g., *, #, •). Use comma-separated for multiple series (e.g., \"*,#\")")

	flag.Parse()

	series := make([][]float64, seriesNum)

	var seriesCharsOption asciigraph.Option
	if customChar != "" {
		chars := strings.Split(customChar, ",")
		charSets := make([]asciigraph.CharSet, len(chars))
		for i, c := range chars {
			charSets[i] = asciigraph.CreateCharSet(strings.TrimSpace(c))
		}
		seriesCharsOption = asciigraph.SeriesChars(charSets...)
	}

	if enableRealTime && realTimeDataBuffer == 0 {
		realTimeDataBuffer = int(width)
	}

	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanLines)

	nextFlushTime := time.Now()

	flushInterval := time.Duration(float64(time.Second) / fps)

	for s.Scan() {
		line := s.Text()
		points := strings.Split(line, delimiter)

		if uint(len(points)) < seriesNum {
			log.Fatal("number of series in the input data stream is less than the specified series number")
		} else if uint(len(points)) > seriesNum {
			points = points[:seriesNum]
		}

		for i, point := range points {
			p, err := strconv.ParseFloat(strings.TrimSpace(point), 64)
			if err != nil {
				log.Printf("ignore %q: cannot parse value", point)
				p = math.NaN()
			}
			series[i] = append(series[i], p)
		}

		if enableRealTime {
			if realTimeDataBuffer > 0 && len(series[0]) > realTimeDataBuffer {
				for i := range series {
					seriesLength := len(series[i])
					series[i] = series[i][seriesLength-realTimeDataBuffer:]
				}
			}

			if currentTime := time.Now(); currentTime.After(nextFlushTime) || currentTime.Equal(nextFlushTime) {
				seriesCopy := append([][]float64(nil), series...)
				opts := []asciigraph.Option{
					asciigraph.Height(int(height)),
					asciigraph.Width(int(width)),
					asciigraph.Offset(int(offset)),
					asciigraph.Precision(precision),
					asciigraph.Caption(caption),
					asciigraph.SeriesColors(seriesColors...),
					asciigraph.SeriesLegends(seriesLegends...),
					asciigraph.CaptionColor(captionColor),
					asciigraph.AxisColor(axisColor),
					asciigraph.LabelColor(labelColor),
					asciigraph.LowerBound(lowerBound),
					asciigraph.UpperBound(upperBound),
				}
				if seriesCharsOption != nil {
					opts = append(opts, seriesCharsOption)
				}
				plot := asciigraph.PlotMany(seriesCopy, opts...)
				// Clear previous graph in realtime mode
				clearPreviousGraph()
				fmt.Println(plot)
				// Record the number of lines for next clearing
				lines := strings.Split(plot, "\n")
				lastGraphLines = len(lines)
				nextFlushTime = time.Now().Add(flushInterval)
			}
		}
	}
	if !enableRealTime {
		if err := s.Err(); err != nil {
			log.Fatal(err)
		}

		if len(series) == 0 {
			log.Fatal("no data")
		}

		opts := []asciigraph.Option{
			asciigraph.Height(int(height)),
			asciigraph.Width(int(width)),
			asciigraph.Offset(int(offset)),
			asciigraph.Precision(precision),
			asciigraph.Caption(caption),
			asciigraph.SeriesColors(seriesColors...),
			asciigraph.SeriesLegends(seriesLegends...),
			asciigraph.CaptionColor(captionColor),
			asciigraph.AxisColor(axisColor),
			asciigraph.LabelColor(labelColor),
			asciigraph.LowerBound(lowerBound),
			asciigraph.UpperBound(upperBound),
		}
		if seriesCharsOption != nil {
			opts = append(opts, seriesCharsOption)
		}
		plot := asciigraph.PlotMany(series, opts...)

		fmt.Println(plot)
	}
}

func parseColors(colors string) ([]asciigraph.AnsiColor, bool) {
	colorList := strings.Split(colors, ",")
	parsedColors := make([]asciigraph.AnsiColor, len(colorList))

	for i, color := range colorList {
		parsedColor, ok := parseColor(strings.TrimSpace(color))
		if !ok {
			return parsedColors, ok
		}
		parsedColors[i] = parsedColor
	}

	return parsedColors, true
}

// clearPreviousGraph clears the previous graph in realtime mode
// by moving cursor up and clearing the lines
func clearPreviousGraph() {
	if lastGraphLines > 0 {
		// Move cursor up by lastGraphLines
		fmt.Printf("\033[%dA", lastGraphLines)
		// Clear each line
		for i := 0; i < lastGraphLines; i++ {
			fmt.Print("\033[2K\r")
			if i < lastGraphLines-1 {
				fmt.Print("\033[B") // Move down one line
			}
		}
		// Move cursor back to start position
		fmt.Printf("\033[%dA", lastGraphLines)
	}
}

func parseColor(color string) (asciigraph.AnsiColor, bool) {
	parsedColor, ok := asciigraph.ColorNames[color]
	return parsedColor, ok
}
