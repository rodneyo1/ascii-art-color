package asciiart

import (
	"testing"
)

func TestColorPicker(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
		expectedExists bool
	}{
		{"red", "\033[31m", true},
		{"blue", "\033[34m", true},
		{"#FF5733", "\033[38;5;203m", true}, // orange hex code
		{"255,0,0", "\033[38;5;196m", true}, // RGB for red
		{"unknown", "", false},
		{"#invalid", "", false},
		{"256,256,256", "", false},
	}

	for _, test := range tests {
		output, exists := colorPicker(test.input)
		if output != test.expectedOutput || exists != test.expectedExists {
			t.Errorf("colorPicker(%q) = %q, %v; want %q, %v", test.input, output, exists, test.expectedOutput, test.expectedExists)
		}
	}
}
