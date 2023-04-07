//go:build windows
// +build windows

package colors

import (
	"os"

	"golang.org/x/sys/windows"
)

func enableWindowsANSIColor() {
	stdout := windows.Handle(os.Stdout.Fd())
	var originalMode uint32

	windows.GetConsoleMode(stdout, &originalMode)
	windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}
