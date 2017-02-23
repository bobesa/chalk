package chalk

// Hi-Intensity Foreground text colors
const (
	fgHiBlack attribute = iota + 90
	fgHiRed
	fgHiGreen
	fgHiYellow
	fgHiBlue
	fgHiMagenta
	fgHiCyan
	fgHiWhite
)

// BlackIntense reports Formatter with Black as initial format
func BlackIntense() Formatter {
	return Formatter{fgHiBlack}
}

// AsBlackIntense reports Black string based on provided content
func AsBlackIntense(a ...interface{}) string {
	return BlackIntense().Sprint(a...)
}

// BlackIntense reports Formatter with Black as additional format
func (f Formatter) BlackIntense() Formatter {
	return append(f, fgHiBlack)
}

// RedIntense reports Formatter with Red as initial format
func RedIntense() Formatter {
	return Formatter{fgHiRed}
}

// AsRedIntense reports Red string based on provided content
func AsRedIntense(a ...interface{}) string {
	return Red().Sprint(a...)
}

// RedIntense reports Formatter with Red as additional format
func (f Formatter) RedIntense() Formatter {
	return append(f, fgHiRed)
}

// GreenIntense reports Formatter with Green as initial format
func GreenIntense() Formatter {
	return Formatter{fgHiGreen}
}

// AsGreenIntense reports Green string based on provided content
func AsGreenIntense(a ...interface{}) string {
	return Green().Sprint(a...)
}

// GreenIntense reports Formatter with Green as additional format
func (f Formatter) GreenIntense() Formatter {
	return append(f, fgHiGreen)
}

// YellowIntense reports Formatter with Yellow as initial format
func YellowIntense() Formatter {
	return Formatter{fgHiYellow}
}

// AsYellowIntense reports Yellow string based on provided content
func AsYellowIntense(a ...interface{}) string {
	return Yellow().Sprint(a...)
}

// YellowIntense reports Formatter with Yellow as additional format
func (f Formatter) YellowIntense() Formatter {
	return append(f, fgHiYellow)
}

// BlueIntense reports Formatter with Blue as initial format
func BlueIntense() Formatter {
	return Formatter{fgHiBlue}
}

// AsBlueIntense reports Blue string based on provided content
func AsBlueIntense(a ...interface{}) string {
	return Blue().Sprint(a...)
}

// BlueIntense reports Formatter with Blue as additional format
func (f Formatter) BlueIntense() Formatter {
	return append(f, fgHiBlue)
}

// MagentaIntense reports Formatter with Magenta as initial format
func MagentaIntense() Formatter {
	return Formatter{fgHiMagenta}
}

// AsMagentaIntense reports Magenta string based on provided content
func AsMagentaIntense(a ...interface{}) string {
	return Magenta().Sprint(a...)
}

// MagentaIntense reports Formatter with Magenta as additional format
func (f Formatter) MagentaIntense() Formatter {
	return append(f, fgHiMagenta)
}

// CyanIntense reports Formatter with Cyan as initial format
func CyanIntense() Formatter {
	return Formatter{fgHiCyan}
}

// AsCyanIntense reports Cyan string based on provided content
func AsCyanIntense(a ...interface{}) string {
	return Cyan().Sprint(a...)
}

// CyanIntense reports Formatter with Cyan as additional format
func (f Formatter) CyanIntense() Formatter {
	return append(f, fgHiCyan)
}

// WhiteIntense reports Formatter with White as initial format
func WhiteIntense() Formatter {
	return Formatter{fgHiWhite}
}

// AsWhiteIntense reports White string based on provided content
func AsWhiteIntense(a ...interface{}) string {
	return White().Sprint(a...)
}

// WhiteIntense reports Formatter with White as additional format
func (f Formatter) WhiteIntense() Formatter {
	return append(f, fgHiWhite)
}
