package colors

import (
	"testing"
)

func TestHex(t *testing.T) {
	tests := []struct {
		name     string
		hexColor string
		expected string
	}{
		{"Valid 6-character color", "FF12AB", "\033[38;2;255;18;171m"},
		{"Valid 6-character color with hash", "#FF12AB", "\033[38;2;255;18;171m"},
		{"Valid 3-character color", "F1A", "\033[38;2;255;17;170m"},
		{"Valid 3-character color with hash", "#F1A", "\033[38;2;255;17;170m"},
		{"Invalid length", "FF12", ""},
		{"Invalid characters", "FF12Z9", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Hex(tt.hexColor)
			if result != tt.expected {
				t.Errorf("Hex(%q) = %q, expected %q", tt.hexColor, result, tt.expected)
			}
		})
	}
}
