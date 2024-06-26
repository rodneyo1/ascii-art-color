package asciiart

import (
	"fmt"
	"strconv"
	"strings"
)

// Color map to hold basic ANSI escape codes
var colorMap = map[string]string{
	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
	"orange":  "\033[255;2;24m",
}

func HexToRGB(hex string) (int, int, int) {
	hex = strings.TrimPrefix(hex, "#") // Remove #

	var r, g, b int

	// Parse the hex string into RGB values
	n, err := fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b)
	if err != nil || n != 3 {
		return 0, 0, 0 // Return default
	}
	return r, g, b
}

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
	return 16 + (36 * (r / 51)) + (6 * (g / 51)) + (b / 51)
}

// colorPicker returns the ANSI escape code for a given color string
func colorPicker(color string) (string, bool) {
	// Check if the color exists in the map
	colorCode, exists := colorMap[strings.ToLower(color)]
	if exists {
		return colorCode, true
	}

	// Check if the input is a hex string (with or without #)
	if len(color) == 7 && strings.HasPrefix(color, "#") {
		r, g, b := HexToRGB(color)
		ansiColorCode := rgbToAnsi(r, g, b)
		return fmt.Sprintf("\033[38;5;%dm", ansiColorCode), true
	}

	// Check if the input is an RGB string
	if strings.Contains(color, "rgb") {
		rgbValue := color[4 : len(color)-1]
		rgb := strings.Split(rgbValue, ",")
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
