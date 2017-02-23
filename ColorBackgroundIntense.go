package chalk

// Hi-Intensity Background text colors
const (
	bgHiBlack attribute = iota + 100
	bgHiRed
	bgHiGreen
	bgHiYellow
	bgHiBlue
	bgHiMagenta
	bgHiCyan
	bgHiWhite
)

// BlackIntenseBackground reports Formatter with Black as initial format
func BlackIntenseBackground() Formatter {
	return Formatter{bgHiBlack}
}

// AsBlackIntenseBackground reports Black string based on provided content
func AsBlackIntenseBackground(a ...interface{}) string {
	return BlackIntenseBackground().Sprint(a...)
}

// BlackIntenseBackground reports Formatter with Black as additional format
func (f Formatter) BlackIntenseBackground() Formatter {
	return append(f, bgHiBlack)
}

// RedIntenseBackground reports Formatter with Red as initial format
func RedIntenseBackground() Formatter {
	return Formatter{bgHiRed}
}

// AsRedIntenseBackground reports Red string based on provided content
func AsRedIntenseBackground(a ...interface{}) string {
	return Red().Sprint(a...)
}

// RedIntenseBackground reports Formatter with Red as additional format
func (f Formatter) RedIntenseBackground() Formatter {
	return append(f, bgHiRed)
}

// GreenIntenseBackground reports Formatter with Green as initial format
func GreenIntenseBackground() Formatter {
	return Formatter{bgHiGreen}
}

// AsGreenIntenseBackground reports Green string based on provided content
func AsGreenIntenseBackground(a ...interface{}) string {
	return Green().Sprint(a...)
}

// GreenIntenseBackground reports Formatter with Green as additional format
func (f Formatter) GreenIntenseBackground() Formatter {
	return append(f, bgHiGreen)
}

// YellowIntenseBackground reports Formatter with Yellow as initial format
func YellowIntenseBackground() Formatter {
	return Formatter{bgHiYellow}
}

// AsYellowIntenseBackground reports Yellow string based on provided content
func AsYellowIntenseBackground(a ...interface{}) string {
	return Yellow().Sprint(a...)
}

// YellowIntenseBackground reports Formatter with Yellow as additional format
func (f Formatter) YellowIntenseBackground() Formatter {
	return append(f, bgHiYellow)
}

// BlueIntenseBackground reports Formatter with Blue as initial format
func BlueIntenseBackground() Formatter {
	return Formatter{bgHiBlue}
}

// AsBlueIntenseBackground reports Blue string based on provided content
func AsBlueIntenseBackground(a ...interface{}) string {
	return Blue().Sprint(a...)
}

// BlueIntenseBackground reports Formatter with Blue as additional format
func (f Formatter) BlueIntenseBackground() Formatter {
	return append(f, bgHiBlue)
}

// MagentaIntenseBackground reports Formatter with Magenta as initial format
func MagentaIntenseBackground() Formatter {
	return Formatter{bgHiMagenta}
}

// AsMagentaIntenseBackground reports Magenta string based on provided content
func AsMagentaIntenseBackground(a ...interface{}) string {
	return Magenta().Sprint(a...)
}

// MagentaIntenseBackground reports Formatter with Magenta as additional format
func (f Formatter) MagentaIntenseBackground() Formatter {
	return append(f, bgHiMagenta)
}

// CyanIntenseBackground reports Formatter with Cyan as initial format
func CyanIntenseBackground() Formatter {
	return Formatter{bgHiCyan}
}

// AsCyanIntenseBackground reports Cyan string based on provided content
func AsCyanIntenseBackground(a ...interface{}) string {
	return Cyan().Sprint(a...)
}

// CyanIntenseBackground reports Formatter with Cyan as additional format
func (f Formatter) CyanIntenseBackground() Formatter {
	return append(f, bgHiCyan)
}

// WhiteIntenseBackground reports Formatter with White as initial format
func WhiteIntenseBackground() Formatter {
	return Formatter{bgHiWhite}
}

// AsWhiteIntenseBackground reports White string based on provided content
func AsWhiteIntenseBackground(a ...interface{}) string {
	return White().Sprint(a...)
}

// WhiteIntenseBackground reports Formatter with White as additional format
func (f Formatter) WhiteIntenseBackground() Formatter {
	return append(f, bgHiWhite)
}
