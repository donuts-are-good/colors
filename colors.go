package colors

import (
	"runtime"
)

// Colors is the default no-dlc color pack. you got your regular colors and your bright colors. You also get an eraser free with every pack, folks. You can't beat it.
type Colors struct {
	Nc            string
	BrightBlack   string
	BrightRed     string
	BrightGreen   string
	BrightYellow  string
	BrightPurple  string
	BrightMagenta string
	BrightCyan    string
	BrightWhite   string
	Black         string
	Red           string
	Green         string
	Yellow        string
	Purple        string
	Magenta       string
	Cyan          string
	White         string
}

var (
	// c is a pointer to an instance of Colors
	c *Colors
)

// init sets up the Colors struct with the corresponding escape codes for each color
func init() {
	c = &Colors{
		Nc:            "\033[0m",
		BrightBlack:   "\033[1;30m",
		BrightRed:     "\033[1;31m",
		BrightGreen:   "\033[1;32m",
		BrightYellow:  "\033[1;33m",
		BrightPurple:  "\033[1;34m",
		BrightMagenta: "\033[1;35m",
		BrightCyan:    "\033[1;36m",
		BrightWhite:   "\033[1;37m",
		Black:         "\033[0;30m",
		Red:           "\033[0;31m",
		Green:         "\033[0;32m",
		Yellow:        "\033[0;33m",
		Purple:        "\033[0;34m",
		Magenta:       "\033[0;35m",
		Cyan:          "\033[0;36m",
		White:         "\033[0;37m",
	}
	// if running on Windows, set all colors to an empty string to disable color
	if runtime.GOOS == "windows" {
		c.Nc = ""
		c.BrightBlack = ""
		c.BrightRed = ""
		c.BrightGreen = ""
		c.BrightYellow = ""
		c.BrightPurple = ""
		c.BrightMagenta = ""
		c.BrightCyan = ""
		c.BrightWhite = ""
		c.Black = ""
		c.Red = ""
		c.Green = ""
		c.Yellow = ""
		c.Purple = ""
		c.Magenta = ""
		c.Cyan = ""
		c.White = ""
	}
}

// reset sets all the colors in the Colors struct back to their original values
func (c *Colors) Reset() {
	c.Nc = "\033[0m"
	c.BrightBlack = "\033[1;30m"
	c.BrightRed = "\033[1;31m"
	c.BrightGreen = "\033[1;32m"
	c.BrightYellow = "\033[1;33m"
	c.BrightPurple = "\033[1;34m"
	c.BrightMagenta = "\033[1;35m"
	c.BrightCyan = "\033[1;36m"
	c.BrightWhite = "\033[1;37m"
	c.Black = "\033[0;30m"
	c.Red = "\033[0;31m"
	c.Green = "\033[0;32m"
	c.Yellow = "\033[0;33m"
	c.Purple = "\033[0;34m"
	c.Magenta = "\033[0;35m"
	c.Cyan = "\033[0;36m"
	c.White = "\033[0;37m"
}
