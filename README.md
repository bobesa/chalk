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

## As Variable
You can also "cook" **Formatters** to variable and reuse.
```go
cyan := chalk.Cyan().Bold()
boldi := chalk.Bold().Italic()

for index, msg := range messages {
    boldi.Println(index)
    cyan.Println(msg)
}
```

## Caching
If you are using just certain formatting all the time, it is better to use **Cached Formatter** instead. 
**Cached Formatter** has reduced overhead, but you will lose the ability to change it any further.
Following methods are available on this formatter: `Print`, `Sprint`, `Println`, `Sprintln`, `Printf`, `Sprintf` & `Set`.
```go
var (
    cyan = chalk.Cyan.Cached() // No longer editable, but faster
)

func PrintVeryLongListOfMessages() {
    for _, msg := range veryLongListOfMessages {
        cyan.Println(msg) // Same calls as on regular Formatter
    }
}
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

# JSON Color formatting (Experimental)
If you are brave enough you can play with `JSON(v interface{})` and `JSON(v interface{}, prefix, indent string)` function (similar signature to `json.Marshal` & `json.MarshalIndent`) that parses given value and reports new color formatted json.
_Output should be similar to that of `json.Marshal*`, but as stated before; This functionality is experimental and should mostly be used for debugging purposes._

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