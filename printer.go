package asciiart

import (
	"fmt"
	"os"
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
	var chosencolor string
	switch *color {
	case "Red":
		chosencolor = Red
	case "Green":
		chosencolor = Green 
	case "Yellow":
		chosencolor = Yellow
	case "Blue":
		chosencolor = Blue
	case "Magenta":
		chosencolor = Magenta
	case "Cyan":
		chosencolor = Cyan
	case "Gray":
		chosencolor = Gray
	case "White":
		chosencolor = White 
	}
	
	// map each character of a string to its ascii character in our set
	if tocolor == "" { // if character to color is not provided color the whole text
		for i := 0; i < 9; i++ {
			for _, char := range s {
				toPrint := char - 32
				fmt.Print(chosencolor+b[toPrint][i]+Reset)
			}
			fmt.Println()
		}
	// } else if *output != "" {
	// 	outputfile, err := os.Create(*output)
	// 	if err != nil {
	// 		fmt.Println("Ooops! Unable to create file ", *output)
	// 	}
	// 	for i := 0; i < 9; i++ {
	// 		for _, char := range s {
	// 				toPrint := char - 32
	// 				outputfile.Write([]byte(b[toPrint][i]))
	// 			}
	// 		outputfile.Write([]byte("\n"))
	// 	}
	// 	fmt.Println("Art created successfully!")
	} else {
		for i := 0; i < 9; i++ {
			for _, char := range s {
				if strings.Contains(tocolor, string(char)) {
					toPrint := char - 32
					fmt.Print(chosencolor+b[toPrint][i]+Reset)
					continue
				}
				toPrint := char - 32
				fmt.Print(b[toPrint][i]) // write the art onto the created file
			}
			fmt.Println()
		}
	}
}

func PrintToBanner(s string, b [][]string, color *string, output *string, tocolor string) {
	var Reset = "\033[0m" 
	var Red = "\033[31m" 
	var Green = "\033[32m" 
	var Yellow = "\033[33m" 
	var Blue = "\033[34m" 
	var Magenta = "\033[35m" 
	var Cyan = "\033[36m" 
	var Gray = "\033[37m" 
	var White = "\033[97m"
	var chosencolor string
	switch *color {
	case "Red":
		chosencolor = Red
	case "Green":
		chosencolor = Green 
	case "Yellow":
		chosencolor = Yellow
	case "Blue":
		chosencolor = Blue
	case "Magenta":
		chosencolor = Magenta
	case "Cyan":
		chosencolor = Cyan
	case "Gray":
		chosencolor = Gray
	case "White":
		chosencolor = White 
	}
	outputfile, err := os.Create(*output)
	if err != nil {
		fmt.Println("Ooops! Unable to create file ", *output)
	}
	// map each character of a string to its ascii character in our set
	if tocolor == "" { // if character to color is not provided color the whole text
		for i := 0; i < 9; i++ {
			for _, char := range s {
				toPrint := char - 32
				outputfile.Write([]byte(chosencolor+b[toPrint][i]+Reset))
			}
			outputfile.Write([]byte("\n"))
		}
	} else if tocolor == "" && *color == "" {
		for i := 0; i < 9; i++ {
			for _, char := range s {
					toPrint := char - 32
					outputfile.Write([]byte(b[toPrint][i]))
					continue
			}
			outputfile.Write([]byte("\n"))
		}
	} else {
		for i := 0; i < 9; i++ {
			for _, char := range s {
				if strings.Contains(tocolor, string(char)) {
					toPrint := char - 32
					outputfile.Write([]byte(chosencolor+b[toPrint][i]+Reset))
					continue
				}
				toPrint := char - 32
				outputfile.Write([]byte(b[toPrint][i])) // write the art onto the created file
			}
			outputfile.Write([]byte("\n"))
		}
	}
	fmt.Println("Art created successfully!")
}
