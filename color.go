package asciigraph

import "fmt"

type AnsiColor byte

var (
	Default              AnsiColor = 0
	AliceBlue            AnsiColor = 255
	AntiqueWhite         AnsiColor = 255
	Aqua                 AnsiColor = 14
	Aquamarine           AnsiColor = 122
	Azure                AnsiColor = 15
	Beige                AnsiColor = 230
	Bisque               AnsiColor = 224
	Black                AnsiColor = 188 // dummy value
	BlanchedAlmond       AnsiColor = 230
	Blue                 AnsiColor = 12
	BlueViolet           AnsiColor = 92
	Brown                AnsiColor = 88
	BurlyWood            AnsiColor = 180
	CadetBlue            AnsiColor = 73
	Chartreuse           AnsiColor = 118
	Chocolate            AnsiColor = 166
	Coral                AnsiColor = 209
	CornflowerBlue       AnsiColor = 68
	Cornsilk             AnsiColor = 230
	Crimson              AnsiColor = 161
	Cyan                 AnsiColor = 14
	DarkBlue             AnsiColor = 18
	DarkCyan             AnsiColor = 30
	DarkGoldenrod        AnsiColor = 136
	DarkGray             AnsiColor = 248
	DarkGreen            AnsiColor = 22
	DarkKhaki            AnsiColor = 143
	DarkMagenta          AnsiColor = 90
	DarkOliveGreen       AnsiColor = 59
	DarkOrange           AnsiColor = 208
	DarkOrchid           AnsiColor = 134
	DarkRed              AnsiColor = 88
	DarkSalmon           AnsiColor = 173
	DarkSeaGreen         AnsiColor = 108
	DarkSlateBlue        AnsiColor = 60
	DarkSlateGray        AnsiColor = 238
	DarkTurquoise        AnsiColor = 44
	DarkViolet           AnsiColor = 92
	DeepPink             AnsiColor = 198
	DeepSkyBlue          AnsiColor = 39
	DimGray              AnsiColor = 242
	DodgerBlue           AnsiColor = 33
	Firebrick            AnsiColor = 124
	FloralWhite          AnsiColor = 15
	ForestGreen          AnsiColor = 28
	Fuchsia              AnsiColor = 13
	Gainsboro            AnsiColor = 253
	GhostWhite           AnsiColor = 15
	Gold                 AnsiColor = 220
	Goldenrod            AnsiColor = 178
	Gray                 AnsiColor = 8
	Green                AnsiColor = 2
	GreenYellow          AnsiColor = 155
	Honeydew             AnsiColor = 15
	HotPink              AnsiColor = 205
	IndianRed            AnsiColor = 167
	Indigo               AnsiColor = 54
	Ivory                AnsiColor = 15
	Khaki                AnsiColor = 222
	Lavender             AnsiColor = 254
	LavenderBlush        AnsiColor = 255
	LawnGreen            AnsiColor = 118
	LemonChiffon         AnsiColor = 230
	LightBlue            AnsiColor = 152
	LightCoral           AnsiColor = 210
	LightCyan            AnsiColor = 195
	LightGoldenrodYellow AnsiColor = 230
	LightGray            AnsiColor = 252
	LightGreen           AnsiColor = 120
	LightPink            AnsiColor = 217
	LightSalmon          AnsiColor = 216
	LightSeaGreen        AnsiColor = 37
	LightSkyBlue         AnsiColor = 117
	LightSlateGray       AnsiColor = 103
	LightSteelBlue       AnsiColor = 152
	LightYellow          AnsiColor = 230
	Lime                 AnsiColor = 10
	LimeGreen            AnsiColor = 77
	Linen                AnsiColor = 255
	Magenta              AnsiColor = 13
	Maroon               AnsiColor = 1
	MediumAquamarine     AnsiColor = 79
	MediumBlue           AnsiColor = 20
	MediumOrchid         AnsiColor = 134
	MediumPurple         AnsiColor = 98
	MediumSeaGreen       AnsiColor = 72
	MediumSlateBlue      AnsiColor = 99
	MediumSpringGreen    AnsiColor = 48
	MediumTurquoise      AnsiColor = 80
	MediumVioletRed      AnsiColor = 162
	MidnightBlue         AnsiColor = 17
	MintCream            AnsiColor = 15
	MistyRose            AnsiColor = 224
	Moccasin             AnsiColor = 223
	NavajoWhite          AnsiColor = 223
	Navy                 AnsiColor = 4
	OldLace              AnsiColor = 230
	Olive                AnsiColor = 3
	OliveDrab            AnsiColor = 64
	Orange               AnsiColor = 214
	OrangeRed            AnsiColor = 202
	Orchid               AnsiColor = 170
	PaleGoldenrod        AnsiColor = 223
	PaleGreen            AnsiColor = 120
	PaleTurquoise        AnsiColor = 159
	PaleVioletRed        AnsiColor = 168
	PapayaWhip           AnsiColor = 230
	PeachPuff            AnsiColor = 223
	Peru                 AnsiColor = 173
	Pink                 AnsiColor = 218
	Plum                 AnsiColor = 182
	PowderBlue           AnsiColor = 152
	Purple               AnsiColor = 5
	Red                  AnsiColor = 9
	RosyBrown            AnsiColor = 138
	RoyalBlue            AnsiColor = 63
	SaddleBrown          AnsiColor = 94
	Salmon               AnsiColor = 210
	SandyBrown           AnsiColor = 215
	SeaGreen             AnsiColor = 29
	SeaShell             AnsiColor = 15
	Sienna               AnsiColor = 131
	Silver               AnsiColor = 7
	SkyBlue              AnsiColor = 117
	SlateBlue            AnsiColor = 62
	SlateGray            AnsiColor = 66
	Snow                 AnsiColor = 15
	SpringGreen          AnsiColor = 48
	SteelBlue            AnsiColor = 67
	Tan                  AnsiColor = 180
	Teal                 AnsiColor = 6
	Thistle              AnsiColor = 182
	Tomato               AnsiColor = 203
	Turquoise            AnsiColor = 80
	Violet               AnsiColor = 213
	Wheat                AnsiColor = 223
	White                AnsiColor = 15
	WhiteSmoke           AnsiColor = 255
	Yellow               AnsiColor = 11
	YellowGreen          AnsiColor = 149
)

func (c AnsiColor) String() string {
	if c == Default {
		return "\x1b[0m"
	}
	if c == Black {
		c = 0
	}
	if c <= Silver {
		// 3-bit color
		return fmt.Sprintf("\x1b[%dm", 30+byte(c))
	}
	if c <= White {
		// 4-bit color
		return fmt.Sprintf("\x1b[%dm", 82+byte(c))
	}
	// 8-bit color
	return fmt.Sprintf("\x1b[38;5;%dm", byte(c))
}
