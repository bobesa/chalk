package chalk

const (
	bgBlack attribute = iota + 40
	bgRed
	bgGreen
	bgYellow
	bgBlue
	bgMagenta
	bgCyan
	bgWhite
)

// BlackBackground reports Formatter with Black as initial format
func BlackBackground() Formatter {
	return Formatter{bgBlack}
}

// AsBlackBackground reports Black string based on provided content
func AsBlackBackground(a ...interface{}) string {
	return BlackBackground().Sprint(a...)
}

// BlackBackground reports Formatter with Black as additional format
func (f Formatter) BlackBackground() Formatter {
	return append(f, bgBlack)
}

// RedBackground reports Formatter with Red as initial format
func RedBackground() Formatter {
	return Formatter{bgRed}
}

// AsRedBackground reports Red string based on provided content
func AsRedBackground(a ...interface{}) string {
	return Red().Sprint(a...)
}

// RedBackground reports Formatter with Red as additional format
func (f Formatter) RedBackground() Formatter {
	return append(f, bgRed)
}

// GreenBackground reports Formatter with Green as initial format
func GreenBackground() Formatter {
	return Formatter{bgGreen}
}

// AsGreenBackground reports Green string based on provided content
func AsGreenBackground(a ...interface{}) string {
	return Green().Sprint(a...)
}

// GreenBackground reports Formatter with Green as additional format
func (f Formatter) GreenBackground() Formatter {
	return append(f, bgGreen)
}

// YellowBackground reports Formatter with Yellow as initial format
func YellowBackground() Formatter {
	return Formatter{bgYellow}
}

// AsYellowBackground reports Yellow string based on provided content
func AsYellowBackground(a ...interface{}) string {
	return Yellow().Sprint(a...)
}

// YellowBackground reports Formatter with Yellow as additional format
func (f Formatter) YellowBackground() Formatter {
	return append(f, bgYellow)
}

// BlueBackground reports Formatter with Blue as initial format
func BlueBackground() Formatter {
	return Formatter{bgBlue}
}

// AsBlueBackground reports Blue string based on provided content
func AsBlueBackground(a ...interface{}) string {
	return Blue().Sprint(a...)
}

// BlueBackground reports Formatter with Blue as additional format
func (f Formatter) BlueBackground() Formatter {
	return append(f, bgBlue)
}

// MagentaBackground reports Formatter with Magenta as initial format
func MagentaBackground() Formatter {
	return Formatter{bgMagenta}
}

// AsMagentaBackground reports Magenta string based on provided content
func AsMagentaBackground(a ...interface{}) string {
	return Magenta().Sprint(a...)
}

// MagentaBackground reports Formatter with Magenta as additional format
func (f Formatter) MagentaBackground() Formatter {
	return append(f, bgMagenta)
}

// CyanBackground reports Formatter with Cyan as initial format
func CyanBackground() Formatter {
	return Formatter{bgCyan}
}

// AsCyanBackground reports Cyan string based on provided content
func AsCyanBackground(a ...interface{}) string {
	return Cyan().Sprint(a...)
}

// CyanBackground reports Formatter with Cyan as additional format
func (f Formatter) CyanBackground() Formatter {
	return append(f, bgCyan)
}

// WhiteBackground reports Formatter with White as initial format
func WhiteBackground() Formatter {
	return Formatter{bgWhite}
}

// AsWhiteBackground reports White string based on provided content
func AsWhiteBackground(a ...interface{}) string {
	return White().Sprint(a...)
}

// WhiteBackground reports Formatter with White as additional format
func (f Formatter) WhiteBackground() Formatter {
	return append(f, bgWhite)
}
