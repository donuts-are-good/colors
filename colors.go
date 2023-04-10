package colors

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

var (
	NC = "\033[0m"

	Bold      = "\033[1m"
	Faint     = "\033[2m"
	Italic    = "\033[3m"
	Underline = "\033[4m"

	Black   = "\033[0;30m"
	Red     = "\033[0;31m"
	Green   = "\033[0;32m"
	Yellow  = "\033[0;33m"
	Purple  = "\033[0;34m"
	Magenta = "\033[0;35m"
	Cyan    = "\033[0;36m"
	White   = "\033[0;37m"

	BrightBlack   = "\033[1;30m"
	BrightRed     = "\033[1;31m"
	BrightGreen   = "\033[1;32m"
	BrightYellow  = "\033[1;33m"
	BrightPurple  = "\033[1;34m"
	BrightMagenta = "\033[1;35m"
	BrightCyan    = "\033[1;36m"
	BrightWhite   = "\033[1;37m"

	BgBlack   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"

	BgBrightBlack   = "\033[100m"
	BgBrightRed     = "\033[101m"
	BgBrightGreen   = "\033[102m"
	BgBrightYellow  = "\033[103m"
	BgBrightBlue    = "\033[104m"
	BgBrightMagenta = "\033[105m"
	BgBrightCyan    = "\033[106m"
	BgBrightWhite   = "\033[107m"
)

func init() {
	if runtime.GOOS == "windows" {
		enableWindowsANSIColor()
	}
}

func Hex(hexColor string) string {
	hexColor = strings.TrimPrefix(hexColor, "#")

	if len(hexColor) != 3 && len(hexColor) != 6 {
		return ""
	}

	if len(hexColor) == 3 {
		hexColor = strings.Repeat(string(hexColor[0]), 2) + strings.Repeat(string(hexColor[1]), 2) + strings.Repeat(string(hexColor[2]), 2)
	}

	r, err := strconv.ParseInt(hexColor[0:2], 16, 64)
	if err != nil {
		return ""
	}

	g, err := strconv.ParseInt(hexColor[2:4], 16, 64)
	if err != nil {
		return ""
	}

	b, err := strconv.ParseInt(hexColor[4:6], 16, 64)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}
