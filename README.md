# stylecodes
`stylecodes` is a Go package that provides direct control over ANSI escape codes for terminal styling. Unlike 
higher-level styling libraries, stylecodes focuses on a lower-level interface, giving developers the ability to 
understand and explicitly work with ANSI escape codes, which affect the entire terminal and not just a single line of 
text.

## Why stylecodes?
While there are many libraries that provide terminal styling, they often abstract away the underlying ANSI escape codes, 
leaving developers unaware of their broader impact on the terminal. `stylecodes` provides:

- **Direct Control**: Allows developers to work with ANSI escape codes directly, providing clear and precise control over 
terminal output.
- **Transparent Interface**: Provides a transparent interface that helps developers learn and understand how ANSI escape 
codes work and their effect on the terminal.
- **Modular Design**: Offers distinct packages for foreground colors (`color`), background colors (`bgcolor`), and text 
modifiers (`modifier`), allowing developers to import only what they need.
- **Convenience**: Includes a `Style` function which wraps text in a style code and an end code.

## Installation
To install `stylecodes`, use go get:
```bash
go get github.com/avearmin/stylecodes
```

## Usage
### Basic Usage
To use `stylecodes`, you can directly use the provided constants for different styles. For example:
```go
package main

import (
    "fmt"
    "github.com/avearmin/stylecodes/color"
)

func main() {
    fmt.Println(color.Red + "This is red text" + color.End)
}
```
### Using Multiple Styles
```go
package main

import (
	"fmt"
	"github.com/avearmin/stylecodes"
	"github.com/avearmin/stylecodes/color"
	"github.com/avearmin/stylecodes/bgcolor"
	"github.com/avearmin/stylecodes/modifier"
)

func main() {
	fmt.Println(
		color.Black + bgcolor.BrightRed + modifier.Bold + modifier.Underline + "This is important" + stylecodes.EndAll
	)
}
```

### Using Multiple Styles With Fine Control Of Ending Styles
```go
package main

import (
	"fmt"
	"github.com/avearmin/stylecodes/color"
	"github.com/avearmin/stylecodes/modifier"
)

func main() {
	fmt.Println(color.Red + modifier.Bold + modifier.Underline + "This is red, bold, and underlined" + color.End + modifier.EndBold)
	fmt.Println("but this is just underlined" + modifier.EndUnderline)
}
```

### Using the Style Function
```go
package main

import (
    "fmt"
    "github.com/avearmin/stylecodes"
    "github.com/avearmin/stylecodes/color"
    "github.com/avearmin/stylecodes/bgcolor"
    "github.com/avearmin/stylecodes/modifier"
)

func main() {
    result := stylecodes.Style("This is important!", color.Black, bgcolor.BrightRed, modifier.Bold, modifier.Underline)
    fmt.Println(result)
}
```

### Using Hex, RGB, and ANSI 256-Color Functions
```go
package main

import (
    "fmt"
	"github.com/avearmin/stylecodes"
    "github.com/avearmin/stylecodes/color"
    "github.com/avearmin/stylecodes/bgcolor"
)

func main() {
    // Using Hex color
    fmt.Println(color.Hex("#ff5733") + "Hello in Hex Color" + color.End)

    // Using RGB color
    fmt.Println(color.RGB(128, 0, 128) + "Hello in RGB Color" + color.End)

    // Using ANSI 256-Color mode
    fmt.Println(color.Ansi256(34) + "Hello in ANSI 256-Color" + color.End)

    // Using Hex color for background
    fmt.Println(bgcolor.Hex("#87ceeb") + "Hello with Hex Background" + bgcolor.End)
	
	// Using the Style function 
	fmt.Println(stylecodes.Style("Hello from Style() with Hex color", color.Hex("#ff5733")))
}
```