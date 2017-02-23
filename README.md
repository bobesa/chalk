Chalk
=====
Easy to use terminal coloring module.
It is written as a partial replacement for fmt (you can use `Print`, `Sprint`, `Println`, `Sprintln`, `Printf`, `Sprintf`) and have multiple ways of use.

_This module has been written for Mac & Linux and not tested on Windows._

# Formatter
You can call the name of color (`Black`, `Red`...) or format (`Bold`, `Italic`...) as module function.
This will return a Formatter that has `Print`, `Sprint`, `Println`, `Sprintln`, `Printf`, `Sprintf` methods.

## Basic Formatting
```go
chalk.Bold().Println("bold text")
chalk.Red().Printf("red %d", 42)
```

## Format Chaining

```go
chalk.Bold().Cyan().Println("bold cyan text")
chalk.Red().Italic().Printf("red italic %d", 42)
```

# Shorthand calls
You can call the formatter functions with `As` prefix (`AsBlack`, `AsRed`, `AsBold`, `AsItalic`...) to directly report color formatted string.
This shorthand syntax supports multiple arguments (just as most of `fmt.Print*` functions).
```go
chalk.AsBold("bold text")
chalk.AsRed("red", 42)
```

# List of Colors & Formats

- **Bold**
- Faint
- _Italic_
- Underline

- Black ⚫️ 
- Red 🔴 
- Green 🍏 
- Yellow 💛 
- Blue 🔵 
- Magenta 💜 
- Cyan 🚙 
- White ⚪️ 

- BlackBackground ⚫️ 
- RedBackground 🔴 
- GreenBackground 🍏 
- YellowBackground 💛 
- BlueBackground 🔵 
- MagentaBackground 💜 
- CyanBackground 🚙 
- WhiteBackground ⚪️ 