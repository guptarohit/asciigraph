package asciigraph

import (
	"strings"
)

// CharSet defines the characters used for plotting a series.
type CharSet struct {
	Horizontal     string // Horizontal line character (default: ─)
	VerticalLine   string // Vertical line character (default: │)
	ArcDownRight   string // Arc character going down and right (default: ╭)
	ArcDownLeft    string // Arc character going down and left (default: ╮)
	ArcUpRight     string // Arc character going up and right (default: ╰)
	ArcUpLeft      string // Arc character going up and left (default: ╯)
	EndCap         string // End cap character (default: ╴)
	StartCap       string // Start cap character (default: ╶)
	UpRight        string // Axis corner character (default: └)
	DownHorizontal string // X-axis tick mark character (default: ┬)
}

// DefaultCharSet provides the default box-drawing characters.
var DefaultCharSet = CharSet{
	Horizontal:     "─",
	VerticalLine:   "│",
	ArcDownRight:   "╭",
	ArcDownLeft:    "╮",
	ArcUpRight:     "╰",
	ArcUpLeft:      "╯",
	EndCap:         "╴",
	StartCap:       "╶",
	UpRight:        "└",
	DownHorizontal: "┬",
}

// CreateCharSet is a helper function that creates a CharSet with all fields set to the same character.
// This is useful for simple uniform character sets like "*", "•", "#", etc.
func CreateCharSet(char string) CharSet {
	return CharSet{
		Horizontal:     char,
		VerticalLine:   char,
		ArcDownRight:   char,
		ArcDownLeft:    char,
		ArcUpRight:     char,
		ArcUpLeft:      char,
		EndCap:         char,
		StartCap:       char,
		UpRight:        char,
		DownHorizontal: char,
	}
}

// Option represents a configuration setting.
type Option interface {
	apply(c *config)
}

// config holds various graph options
type config struct {
	Width, Height          int
	LowerBound, UpperBound *float64
	Offset                 int
	Caption                string
	Precision              *uint
	CaptionColor           AnsiColor
	AxisColor              AnsiColor
	LabelColor             AnsiColor
	SeriesColors           []AnsiColor
	SeriesLegends          []string
	LineEnding             string
	SeriesChars            []CharSet
	YAxisValueFormatter    YAxisValueFormatterFunc
	XAxisRange             *[2]float64
	XAxisTickCount         int
	XAxisValueFormatter    XAxisValueFormatterFunc
}

// YAxisValueFormatterFunc formats a single Y-axis value.
type YAxisValueFormatterFunc func(float64) string

// XAxisValueFormatterFunc formats a single X-axis tick value.
type XAxisValueFormatterFunc func(float64) string

// An optionFunc applies an option.
type optionFunc func(*config)

// apply implements the Option interface.
func (of optionFunc) apply(c *config) { of(c) }

func configure(defaults config, options []Option) *config {
	for _, o := range options {
		o.apply(&defaults)
	}
	if defaults.LineEnding == "" {
		defaults.LineEnding = "\n"
	}
	return &defaults
}

// Width sets the graphs width. By default, the width of the graph is
// determined by the number of data points. If the value given is a
// positive number, the data points are interpolated on the x axis.
// Values <= 0 reset the width to the default value.
func Width(w int) Option {
	return optionFunc(func(c *config) {
		if w > 0 {
			c.Width = w
		} else {
			c.Width = 0
		}
	})
}

// Height sets the graphs height.
func Height(h int) Option {
	return optionFunc(func(c *config) {
		if h > 0 {
			c.Height = h
		} else {
			c.Height = 0
		}
	})
}

// LowerBound sets the graph's minimum value for the vertical axis. It will be ignored
// if the series contains a lower value.
func LowerBound(min float64) Option {
	return optionFunc(func(c *config) { c.LowerBound = &min })
}

// UpperBound sets the graph's maximum value for the vertical axis. It will be ignored
// if the series contains a bigger value.
func UpperBound(max float64) Option {
	return optionFunc(func(c *config) { c.UpperBound = &max })
}

// Offset sets the graphs offset.
func Offset(o int) Option {
	return optionFunc(func(c *config) { c.Offset = o })
}

// Precision sets the graphs precision.
func Precision(p uint) Option {
	return optionFunc(func(c *config) { c.Precision = &p })
}

// Caption sets the graphs caption.
func Caption(caption string) Option {
	return optionFunc(func(c *config) {
		c.Caption = strings.TrimSpace(caption)
	})
}

// CaptionColor sets the caption color.
func CaptionColor(ac AnsiColor) Option {
	return optionFunc(func(c *config) {
		c.CaptionColor = ac
	})
}

// AxisColor sets the axis color.
func AxisColor(ac AnsiColor) Option {
	return optionFunc(func(c *config) {
		c.AxisColor = ac
	})
}

// LabelColor sets the axis label color.
func LabelColor(ac AnsiColor) Option {
	return optionFunc(func(c *config) {
		c.LabelColor = ac
	})
}

// SeriesColors sets the series colors.
func SeriesColors(ac ...AnsiColor) Option {
	return optionFunc(func(c *config) {
		c.SeriesColors = ac
	})
}

// SeriesLegends sets the legend text for the corresponding series.
func SeriesLegends(text ...string) Option {
	return optionFunc(func(c *config) {
		c.SeriesLegends = text
	})
}

// LineEnding sets the line ending sequence. Use "\r\n" for raw terminals
// (e.g., Windows terminals) or "\n" for standard Unix-style output.
// Defaults to "\n".
func LineEnding(ending string) Option {
	return optionFunc(func(c *config) {
		c.LineEnding = ending
	})
}

// SeriesChars sets the character sets for each series.
// If fewer CharSets are provided than series, DefaultCharSet is used for remaining series.
func SeriesChars(charSets ...CharSet) Option {
	return optionFunc(func(c *config) {
		c.SeriesChars = charSets
	})
}

// YAxisValueFormatter formats values printed on the Y-axis.
func YAxisValueFormatter(f YAxisValueFormatterFunc) Option {
	return optionFunc(func(c *config) {
		c.YAxisValueFormatter = f
	})
}

// XAxisRange enables the X-axis and maps the given domain [min, max] onto the plot width.
func XAxisRange(min, max float64) Option {
	return optionFunc(func(c *config) {
		c.XAxisRange = &[2]float64{min, max}
	})
}

// XAxisTickCount sets the number of ticks on the X-axis. Default is 5, minimum is 2.
func XAxisTickCount(n int) Option {
	return optionFunc(func(c *config) {
		if n >= 2 {
			c.XAxisTickCount = n
		}
	})
}

// XAxisValueFormatter formats values printed on the X-axis.
func XAxisValueFormatter(f XAxisValueFormatterFunc) Option {
	return optionFunc(func(c *config) {
		c.XAxisValueFormatter = f
	})
}
