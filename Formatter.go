package chalk

import (
	"fmt"
	"strconv"
	"strings"
)

// Formatter is
type Formatter []attribute

// colorFormatString reports basic formatting for string
func (f Formatter) colorFormatString() string {
	format := make([]string, len(f))
	for i, v := range f {
		format[i] = strconv.Itoa(int(v))
	}
	return fmt.Sprintf("%s[%sm", spEscape, strings.Join(format, ";"))
}

// Print is colored equivalent of fmt.Print
func (f Formatter) Print(a ...interface{}) {
	fmt.Print(f.colorFormatString() + fmt.Sprint(a...) + spClear)
}

// Sprint is colored equivalent of fmt.Sprint
func (f Formatter) Sprint(a ...interface{}) string {
	return fmt.Sprint(f.colorFormatString() + fmt.Sprint(a...) + spClear)
}

// Println is colored equivalent of fmt.Println
func (f Formatter) Println(a ...interface{}) {
	fmt.Println(f.colorFormatString() + fmt.Sprint(a...) + spClear)
}

// Sprintln is colored equivalent of fmt.Sprintln
func (f Formatter) Sprintln(a ...interface{}) string {
	return fmt.Sprintln(f.colorFormatString() + fmt.Sprint(a...) + spClear)
}

// Printf is colored equivalent of fmt.Printf
func (f Formatter) Printf(format string, a ...interface{}) {
	fmt.Printf(f.colorFormatString()+format+spClear, a...)
}

// Sprintf is colored equivalent of fmt.Sprintf
func (f Formatter) Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(f.colorFormatString()+format+spClear, a...)
}
