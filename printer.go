package asciiart

import (
	"fmt"
	"strings"
)

func Printer(s string, b [][]string, color *string, tocolor string) {
	var Reset = "\033[0m" 
	var Red = "\033[31m" 
	var Green = "\033[32m" 
	var Yellow = "\033[33m" 
	var Blue = "\033[34m" 
	var Magenta = "\033[35m" 
	var Cyan = "\033[36m" 
	var Gray = "\033[37m" 
	var White = "\033[97m"
	var chosenColor string
	switch *color {
	case "Red":
		chosenColor = Red
	case "Green":
		chosenColor = Green 
	case "Yellow":
		chosenColor = Yellow
	case "Blue":
		chosenColor = Blue
	case "Magenta":
		chosenColor = Magenta
	case "Cyan":
		chosenColor = Cyan
	case "Gray":
		chosenColor = Gray
	case "White":
		chosenColor = White 
	default:
		// fmt.Println("Invalid color! Refer to README.md")
		// os.Exit(0)
	}
	
	// map each character of a string to its ascii character in our set
	if tocolor == "" { // if character to color is not provided color the whole text
		for i := 0; i < 9; i++ {
			for _, char := range s {
				toPrint := char - 32
				fmt.Print(chosenColor+b[toPrint][i]+Reset)
			}
			fmt.Println()
		}
	} else {
		for i := 0; i < 9; i++ {
			for _, char := range s {
				if strings.Contains(tocolor, string(char)) {
					toPrint := char - 32
					fmt.Print(chosenColor+b[toPrint][i]+Reset)
					continue
				}
				toPrint := char - 32
				fmt.Print(b[toPrint][i]) // write the art onto the created file
			}
			fmt.Println()
		}
	}
	
}
