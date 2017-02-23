package chalk

// Foreground text colors
const (
	fgBlack attribute = iota + 30
	fgRed
	fgGreen
	fgYellow
	fgBlue
	fgMagenta
	fgCyan
	fgWhite
)

// Black reports Formatter with Black as initial format
func Black() Formatter {
	return Formatter{fgBlack}
}

// AsBlack reports Black string based on provided content
func AsBlack(a ...interface{}) string {
	return Black().Sprint(a...)
}

// Black reports Formatter with Black as additional format
func (f Formatter) Black() Formatter {
	return append(f, fgBlack)
}

// Red reports Formatter with Red as initial format
func Red() Formatter {
	return Formatter{fgRed}
}

// AsRed reports Red string based on provided content
func AsRed(a ...interface{}) string {
	return Red().Sprint(a...)
}

// Red reports Formatter with Red as additional format
func (f Formatter) Red() Formatter {
	return append(f, fgRed)
}

// Green reports Formatter with Green as initial format
func Green() Formatter {
	return Formatter{fgGreen}
}

// AsGreen reports Green string based on provided content
func AsGreen(a ...interface{}) string {
	return Green().Sprint(a...)
}

// Green reports Formatter with Green as additional format
func (f Formatter) Green() Formatter {
	return append(f, fgGreen)
}

// Yellow reports Formatter with Yellow as initial format
func Yellow() Formatter {
	return Formatter{fgYellow}
}

// AsYellow reports Yellow string based on provided content
func AsYellow(a ...interface{}) string {
	return Yellow().Sprint(a...)
}

// Yellow reports Formatter with Yellow as additional format
func (f Formatter) Yellow() Formatter {
	return append(f, fgYellow)
}

// Blue reports Formatter with Blue as initial format
func Blue() Formatter {
	return Formatter{fgBlue}
}

// AsBlue reports Blue string based on provided content
func AsBlue(a ...interface{}) string {
	return Blue().Sprint(a...)
}

// Blue reports Formatter with Blue as additional format
func (f Formatter) Blue() Formatter {
	return append(f, fgBlue)
}

// Magenta reports Formatter with Magenta as initial format
func Magenta() Formatter {
	return Formatter{fgMagenta}
}

// AsMagenta reports Magenta string based on provided content
func AsMagenta(a ...interface{}) string {
	return Magenta().Sprint(a...)
}

// Magenta reports Formatter with Magenta as additional format
func (f Formatter) Magenta() Formatter {
	return append(f, fgMagenta)
}

// Cyan reports Formatter with Cyan as initial format
func Cyan() Formatter {
	return Formatter{fgCyan}
}

// AsCyan reports Cyan string based on provided content
func AsCyan(a ...interface{}) string {
	return Cyan().Sprint(a...)
}

// Cyan reports Formatter with Cyan as additional format
func (f Formatter) Cyan() Formatter {
	return append(f, fgCyan)
}

// White reports Formatter with White as initial format
func White() Formatter {
	return Formatter{fgWhite}
}

// AsWhite reports White string based on provided content
func AsWhite(a ...interface{}) string {
	return White().Sprint(a...)
}

// White reports Formatter with White as additional format
func (f Formatter) White() Formatter {
	return append(f, fgWhite)
}
