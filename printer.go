package asciiart

import (
	"fmt"
	"strings"
)

// Color map to hold ANSI escape codes
var colorMap = map[string]string{
	"Red":     "\033[31m",
	"Green":   "\033[32m",
	"Yellow":  "\033[33m",
	"Blue":    "\033[34m",
	"Magenta": "\033[35m",
	"Cyan":    "\033[36m",
	"Gray":    "\033[37m",
	"White":   "\033[97m",
}

// Printer prints the ASCII art string with the specified color
func Printer(s string, b [][]string, color *string, tocolor string) {
	// ANSI reset code
	var Reset = "\033[0m"

	// Default to no color if the specified color is not found in the map
	chosenColor := ""
	if val, ok := colorMap[strings.Title(*color)]; ok {
		chosenColor = val
	} else {
		fmt.Println("Invalid color! Refer to README.md")
		return
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
