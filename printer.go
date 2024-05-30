package asciiart

import (
	"fmt"
	"strings"
)

// Printer prints the ASCII art string with the specified color
func Printer(s string, b [][]string, color *string, tocolor string) {
	// ANSI reset code
	var Reset = "\033[0m"

	// Default to no color if the specified color is not found in the map
	chosenColor := ""
	colorCode, exists := colorPicker(strings.ToLower(*color))

	if exists {
		chosenColor = colorCode
	}

	// Map each character of the string to its ASCII character in the set
	if tocolor == "" { // if character to color is not provided color the whole text
		for i := 0; i < 9; i++ {
			for _, char := range s {
				toPrint := char - 32
				fmt.Print(chosenColor + b[toPrint][i] + Reset)
			}
			fmt.Println()
		}
	} else {
		for i := 0; i < 9; i++ {
			for _, char := range s {
				toPrint := char - 32
				if strings.Contains(tocolor, string(char)) {
					fmt.Print(chosenColor + b[toPrint][i] + Reset)
				} else {
					fmt.Print(b[toPrint][i])
				}
			}
			fmt.Println()
		}
	}
}
