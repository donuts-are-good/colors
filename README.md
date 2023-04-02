![github-banner](https://user-images.githubusercontent.com/96031819/229327991-15ff4ce6-11ef-48eb-bc64-c5e28b46e885.png)
![donuts-are-good's followers](https://img.shields.io/github/followers/donuts-are-good?&color=555&style=for-the-badge&label=followers) ![donuts-are-good's stars](https://img.shields.io/github/stars/donuts-are-good?affiliations=OWNER%2CCOLLABORATOR&color=555&style=for-the-badge) ![donuts-are-good's visitors](https://komarev.com/ghpvc/?username=donuts-are-good&color=555555&style=for-the-badge&label=visitors)

# 🎨 Colors

`colors` is a Go package for formatting text with ANSI escape codes. It provides a simple interface for coloring text in the terminal, with support for 16 colors.

## Usage

To use `colors`, import the package and create an instance of the `Colors` struct:

```go
package main

// import the `colors` package
import (
	"fmt"
	"github.com/donuts-are-good/colors"
)

func main() {
	// create a new instance of the "Colors" struct
	c := &colors.Colors{}

	// print the test string "This text is red." in red color using the "Red" color code and reset code "Nc" of the "Colors" struct
	fmt.Println(c.Red + "This text is red." + c.Nc)
}

```
The Colors struct provides a set of color codes that can be used to format text. To color text, simply concatenate the desired color code with the text, and then concatenate the reset code c.Nc to reset the color.

## Available Colors
**colors supports 16 colors:**

- Black
- Red
- Green
- Yellow
- Blue
- Magenta
- Cyan
- White
- Bright Black
- Bright Red
- Bright Green
- Bright Yellow
- Bright Purple
- Bright Magenta
- Bright Cyan
- Bright White

Each color is available in both normal and bright variations.

## Operating System Support
`colors` uses ANSI escape codes to format text, which are not supported by all operating systems. On Windows, ANSI escape codes are not supported by default, but can be enabled by installing a compatible terminal emulator. `colors` automatically disables escape codes on Windows to avoid compatibility issues.

## Contributing
If you find a bug or have a feature request, please open an issue on GitHub. Pull requests are also welcome.

## License
`colors` is licensed under the MIT license. See the LICENSE file for more information.