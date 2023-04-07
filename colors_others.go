//go:build !windows
// +build !windows

package colors

func enableWindowsANSIColor() {
	// No-op for non-Windows platforms
}
