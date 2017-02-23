package chalk

// Base attributes
const (
	spReset attribute = iota

	spBold
	spFaint
	spItalic
	spUnderline
	// spBlinkSlow // TODO
	// spBlinkRapid // TODO
	// spReverseVideo // TODO
	// spConcealed // TODO
	// spCrossedOut // TODO
)

// Bold reports Formatter with Bold as initial format
func Bold() Formatter {
	return Formatter{spBold}
}

// AsBold reports Bold string based on provided content
func AsBold(a ...interface{}) string {
	return Bold().Sprint(a...)
}

// Bold reports Formatter with Bold as additional format
func (f Formatter) Bold() Formatter {
	return append(f, spBold)
}

// Faint reports Formatter with Faint as initial format
func Faint() Formatter {
	return Formatter{spFaint}
}

// AsFaint reports Faint string based on provided content
func AsFaint(a ...interface{}) string {
	return Faint().Sprint(a...)
}

// Faint reports Formatter with Faint as additional format
func (f Formatter) Faint() Formatter {
	return append(f, spFaint)
}

// Italic reports Formatter with Italic as initial format
func Italic() Formatter {
	return Formatter{spItalic}
}

// AsItalic reports Italic string based on provided content
func AsItalic(a ...interface{}) string {
	return Italic().Sprint(a...)
}

// Italic reports Formatter with Italic as additional format
func (f Formatter) Italic() Formatter {
	return append(f, spItalic)
}

// Underline reports Formatter with Underline as initial format
func Underline() Formatter {
	return Formatter{spUnderline}
}

// AsUnderline reports Underline string based on provided content
func AsUnderline(a ...interface{}) string {
	return Underline().Sprint(a...)
}

// Underline reports Formatter with Underline as additional format
func (f Formatter) Underline() Formatter {
	return append(f, spUnderline)
}
