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

func TestHexGradient(t *testing.T) {
	tests := []struct {
		name       string
		startColor string
		endColor   string
		input      string
		expected   string
	}{
		{
			name:       "Valid colors and input",
			startColor: "#FF0000",
			endColor:   "#00FF00",
			input:      "Gradient",
			expected:   "\033[38;2;255;0;0mG\033[38;2;213;42;0mr\033[38;2;170;85;0ma\033[38;2;128;128;0md\033[38;2;85;170;0mi\033[38;2;42;213;0me\033[38;2;0;255;0mnt\033[0m",
		},
		{
			name:       "Empty input",
			startColor: "#FF0000",
			endColor:   "#00FF00",
			input:      "",
			expected:   "",
		},
		{
			name:       "Invalid start color",
			startColor: "#FF0Z00",
			endColor:   "#00FF00",
			input:      "Gradient",
			expected:   "Gradient",
		},
		{
			name:       "Invalid end color",
			startColor: "#FF0000",
			endColor:   "#0ZFF00",
			input:      "Gradient",
			expected:   "Gradient",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HexGradient(tt.startColor, tt.endColor, tt.input)
			if result != tt.expected {
				t.Errorf("HexGradient(%q, %q, %q) = %q, expected %q", tt.startColor, tt.endColor, tt.input, result, tt.expected)
			}
		})
	}
}
