package chalk

import "fmt"

type attribute int // SGR Code

var (
	spEscape = "\x1b"
	spClear  = fmt.Sprintf("%s[%dm", spEscape, spReset)
)
