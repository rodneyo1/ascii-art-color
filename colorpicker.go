package asciiart

import (
	"fmt"
	"strconv"
	"strings"
)

// colorPicker returns the ANSI escape code for a given color name or RGB string.
func colorPicker(color string) (string, bool) {
	var colorMap = map[string]string{
		"black":          "\033[30m",
		"red":            "\033[31m",
		"green":          "\033[32m",
		"yellow":         "\033[33m",
		"blue":           "\033[34m",
		"magenta":        "\033[35m",
		"cyan":           "\033[36m",
		"white":          "\033[37m",
	}

	// Check if the color exists in the map
	colorCode, exists := colorMap[color]
	if exists {
		return colorCode, true
	}

	// Check if the input is an RGB string
	if strings.Contains(color, ",") {
		rgb := strings.Split(color, ",")
		if len(rgb) == 3 {
			r, err1 := strconv.Atoi(rgb[0])
			g, err2 := strconv.Atoi(rgb[1])
			b, err3 := strconv.Atoi(rgb[2])
			if err1 == nil && err2 == nil && err3 == nil && r >= 0 && r <= 255 && g >= 0 && g <= 255 && b >= 0 && b <= 255 {
				// Convert RGB to ANSI 256-color code
				ansiColorCode := rgbToAnsi(r, g, b)
				return fmt.Sprintf("\033[38;5;%dm", ansiColorCode), true
			}
		}
	}

	// Color not found
	return "", false
}

// rgbToAnsi256 converts RGB values to ANSI 256-color code
func rgbToAnsi(r, g, b int) int {
	if r == g && g == b {
		// Grayscale mode
		if r < 8 {
			return 16
		}
		if r > 248 {
			return 231
		}
		return 232 + (r-8)/10
	}

	// Color mode
	return 16 + (36*(r/51)) + (6*(g/51)) + (b/51)
}