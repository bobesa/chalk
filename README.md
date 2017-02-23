Chalk
=====
Easy to use terminal coloring module.
It is written as a partial replacement for fmt (you can use `Print`, `Sprint`, `Println`, `Sprintln`, `Printf`, `Sprintf`) and have multiple ways of use.

_This module has been written for Mac & Linux and not tested on Windows._

# Formatter
You can call the name of color (`Black`, `Red`...) or format (`Bold`, `Italic`...) as module function.
This will return a **Formatter** that has `Print`, `Sprint`, `Println`, `Sprintln`, `Printf`, `Sprintf` methods.

## Basic Formatting
```go
chalk.Bold().Println("bold text")
chalk.Red().Printf("red %d", 42)
chalk.GreenBackground().Print("green", "text")
```

## Format Chaining
```go
chalk.Bold().Cyan().Println("bold cyan text")
chalk.Red().Italic().Printf("red italic %d", 42)
chalk.GreenBackground().Bold().Print("green", "bold", "text")
```

## Set & Clear
You can call `Set` function on **Formatter** to set formatter's color format active until `Clear` or other formatting function is called.
```go
// Set bold + italic formatting
chalk.Bold().Italic().Set()

// Print formatted messages
log.Println("Bold Italic formatted message 1")
log.Println("Bold Italic formatted message 2")
log.Println("Bold Italic formatted message 3")

// Add cyan formatting to previous formatting
// This becomes Bold+Italic+Cyan formatting in terminal
chalk.Cyan().Set() 

// Print newly formatted messages
log.Println("Bold Italic Cyan formatted message 1")

 // Clear all formatting
 // This sets terminal to default color formatting
chalk.Clear()

// Print non-formatted (or default formatted) message
log.Println("Non formatted message")
```

# Disable/Enable coloring
You can temporarily disable/enable coloring with `Disable()` or `Enable()` functions.
This only overrides behaviour and you can still use formatting functions as is, but they won't do any formating.

```go
// Print bold text
chalk.Bold().Println("bold text")

// Disable formatting
chalk.Disable()

// Print non-bold text (because of disabled formatting)
chalk.Bold().Println("not-bold text")

// Enable formatting again
chalk.Enable()

// Print bold text
chalk.Bold().Println("bold text")
```

# Shorthand calls
You can call the **Formatter** functions with `As` prefix (`AsBlack`, `AsRed`, `AsBold`, `AsItalic`...) to directly report color formatted string.
This shorthand syntax supports multiple arguments (just as most of `fmt.Print*` functions).
```go
chalk.AsBold("bold text")
chalk.AsRed("red", 42)
chalk.AsGreenBackground("green", "text")
```

# List of Colors & Formats

- **Bold**
- Faint
- _Italic_
- Underline
- BlinkSlow
- BlinkRapid
- ReverseVideo
- Concealed
- CrossedOut

---

- Black ⚫️ 
- Red 🔴 
- Green 🍏 
- Yellow 💛 
- Blue 🔵 
- Magenta 💜 
- Cyan 🚙 
- White ⚪️ 

---

- BlackBackground ⚫️ 
- RedBackground 🔴 
- GreenBackground 🍏 
- YellowBackground 💛 
- BlueBackground 🔵 
- MagentaBackground 💜 
- CyanBackground 🚙 
- WhiteBackground ⚪️ 

---

- BlackIntense ⚫️ 
- RedIntense 🔴 
- GreenIntense 🍏 
- YellowIntense 💛 
- BlueIntense 🔵 
- MagentaIntense 💜 
- CyanIntense 🚙 
- WhiteIntense ⚪️ 

---

- BlackIntenseBackground ⚫️ 
- RedIntenseBackground 🔴 
- GreenIntenseBackground 🍏 
- YellowIntenseBackground 💛 
- BlueIntenseBackground 🔵 
- MagentaIntenseBackground 💜 
- CyanIntenseBackground 🚙 
- WhiteIntenseBackground ⚪️ 