package chalk

import "fmt"

type attribute int // SGR Code

var (
	coloringEnabled = true

	spEscape = "\x1b"
	spClear  = fmt.Sprintf("%s[%dm", spEscape, spReset)
)

// Disable disables coloring
func Disable() {
	coloringEnabled = false
	spClear = ""
}

// Enable enables coloring
func Enable() {
	coloringEnabled = true
	spClear = fmt.Sprintf("%s[%dm", spEscape, spReset)
}
