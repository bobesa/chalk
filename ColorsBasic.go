package chalk

// Base attributes
const (
	spReset attribute = iota

	spBold
	spFaint
	spItalic
	spUnderline
	spBlinkSlow
	spBlinkRapid
	spReverseVideo
	spConcealed
	spCrossedOut
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

// BlinkSlow reports Formatter with BlinkSlow as initial format
func BlinkSlow() Formatter {
	return Formatter{spBlinkSlow}
}

// AsBlinkSlow reports BlinkSlow string based on provided content
func AsBlinkSlow(a ...interface{}) string {
	return BlinkSlow().Sprint(a...)
}

// BlinkSlow reports Formatter with BlinkSlow as additional format
func (f Formatter) BlinkSlow() Formatter {
	return append(f, spBlinkSlow)
}

// BlinkRapid reports Formatter with BlinkRapid as initial format
func BlinkRapid() Formatter {
	return Formatter{spBlinkRapid}
}

// AsBlinkRapid reports BlinkRapid string based on provided content
func AsBlinkRapid(a ...interface{}) string {
	return BlinkRapid().Sprint(a...)
}

// BlinkRapid reports Formatter with BlinkRapid as additional format
func (f Formatter) BlinkRapid() Formatter {
	return append(f, spBlinkRapid)
}

// ReverseVideo reports Formatter with ReverseVideo as initial format
func ReverseVideo() Formatter {
	return Formatter{spReverseVideo}
}

// AsReverseVideo reports ReverseVideo string based on provided content
func AsReverseVideo(a ...interface{}) string {
	return ReverseVideo().Sprint(a...)
}

// ReverseVideo reports Formatter with ReverseVideo as additional format
func (f Formatter) ReverseVideo() Formatter {
	return append(f, spReverseVideo)
}

// Concealed reports Formatter with Concealed as initial format
func Concealed() Formatter {
	return Formatter{spConcealed}
}

// AsConcealed reports Concealed string based on provided content
func AsConcealed(a ...interface{}) string {
	return Concealed().Sprint(a...)
}

// Concealed reports Formatter with Concealed as additional format
func (f Formatter) Concealed() Formatter {
	return append(f, spConcealed)
}

// CrossedOut reports Formatter with CrossedOut as initial format
func CrossedOut() Formatter {
	return Formatter{spCrossedOut}
}

// AsCrossedOut reports CrossedOut string based on provided content
func AsCrossedOut(a ...interface{}) string {
	return CrossedOut().Sprint(a...)
}

// CrossedOut reports Formatter with CrossedOut as additional format
func (f Formatter) CrossedOut() Formatter {
	return append(f, spCrossedOut)
}
