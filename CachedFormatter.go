package chalk

import "fmt"

// CachedFormatter prepares formatting and reuses it
type CachedFormatter string

// Set sets formatting until other formatting function or Clear is called
func (f CachedFormatter) Set() {
	fmt.Printf(string(f))
}

// Print is colored equivalent of fmt.Print
func (f CachedFormatter) Print(a ...interface{}) {
	fmt.Print(string(f) + fmt.Sprint(a...) + spClear)
}

// Sprint is colored equivalent of fmt.Sprint
func (f CachedFormatter) Sprint(a ...interface{}) string {
	return fmt.Sprint(string(f) + fmt.Sprint(a...) + spClear)
}

// Println is colored equivalent of fmt.Println
func (f CachedFormatter) Println(a ...interface{}) {
	fmt.Println(string(f) + fmt.Sprint(a...) + spClear)
}

// Sprintln is colored equivalent of fmt.Sprintln
func (f CachedFormatter) Sprintln(a ...interface{}) string {
	return fmt.Sprintln(string(f) + fmt.Sprint(a...) + spClear)
}

// Printf is colored equivalent of fmt.Printf
func (f CachedFormatter) Printf(format string, a ...interface{}) {
	fmt.Printf(string(f)+format+spClear, a...)
}

// Sprintf is colored equivalent of fmt.Sprintf
func (f CachedFormatter) Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(string(f)+format+spClear, a...)
}
