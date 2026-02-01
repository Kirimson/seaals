package image

type Opts struct {
	Position   TextPosition
	Align      TextAlignment
	Filter     string
	Size       int
	Colour     string
	Stroke     string
	StrokeSize int
	Width      int
}

type (
	TextAlignment string
	TextPosition  string
)

var (
	TextCenter TextAlignment = "middle"
	TextLeft   TextAlignment = "start"
	TextRight  TextAlignment = "end"
)

var (
	TextTop    TextPosition = "top"
	TextMiddle TextPosition = "middle"
	TextBottom TextPosition = "bottom"
)

func defaultInt(value int, defVal int) int {
	if value == 0 {
		return defVal
	}
	return value
}

func defaultString(value string, defVal string) string {
	if value == "" {
		return defVal
	}
	return value
}

func DefaultOpts(opts *Opts) *Opts {
	opts.Size = defaultInt(opts.Size, 48)
	opts.Colour = defaultString(opts.Colour, "white")
	opts.Stroke = defaultString(opts.Stroke, "black")
	opts.StrokeSize = defaultInt(opts.StrokeSize, 2)
	opts.Position = TextPosition(defaultString(string(opts.Position), "bottom"))
	opts.Align = TextAlignment(defaultString(string(opts.Align), "middle"))
	return opts
}
