package colors

import (
	"runtime"
)

// 16 ANSI colors, use NC to reset the color.
var (
	NC            = "\033[0m"
	BrightBlack   = "\033[1;30m"
	BrightRed     = "\033[1;31m"
	BrightGreen   = "\033[1;32m"
	BrightYellow  = "\033[1;33m"
	BrightPurple  = "\033[1;34m"
	BrightMagenta = "\033[1;35m"
	BrightCyan    = "\033[1;36m"
	BrightWhite   = "\033[1;37m"
	Black         = "\033[0;30m"
	Red           = "\033[0;31m"
	Green         = "\033[0;32m"
	Yellow        = "\033[0;33m"
	Purple        = "\033[0;34m"
	Magenta       = "\033[0;35m"
	Cyan          = "\033[0;36m"
	White         = "\033[0;37m"
)

func init() {
	if runtime.GOOS == "windows" {
		enableWindowsANSIColor()
	}
}
