package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"asciiart"
)

func main() {
	output := flag.String("output", "banner.txt", "output file") // create an output flag
	color := flag.String("color", "", "text color")// create color flag
	flag.Parse()

	// if len(os.Args) != 4 {
	// 	fmt.Println("More or less arguments passed! Please refer to README.md and try again.")
	// 	os.Exit(0)
	// }
	// file, err := os.Open(flag.Args()[1] + ".txt")
	var r string
	var tocolor string
	var file *os.File
	var err error
	if len(os.Args) < 2 || len(os.Args) > 6  {
		fmt.Print("Ooops! few or too many arguments. Usages are:\ngo run . [string]\ngo run . [color option] [string to color] [string]\ngo run . [color option] [string]\ngo run . [color option] [output option] [string to color] [string] [banner]\n")
		os.Exit(0)
	}
	if len(os.Args) == 2 {
		r = os.Args[1]
		tocolor = ""
		*color = ""
		*output = ""

	}
	if len(os.Args) == 3 {
		r = flag.Args()[0]
		tocolor = ""
		*output = ""
	}
	if len(os.Args) == 4 {
		if *color == "" && *output != "" {
			file, err = os.Open(flag.Args()[1]+".txt")
			if err != nil {
				fmt.Println("Unable to open file!")
				os.Exit(0)
			}
			r = flag.Args()[0]
			tocolor = ""
		}
		if *color != "" && *output == "" {
			file, err = os.Open(flag.Args()[1]+".txt")
			if err != nil {
				fmt.Println("Unable to open file!")
				os.Exit(0)
			}
			r = flag.Args()[0]
			tocolor = ""
		}
		// r = flag.Args()[1]
		// tocolor = flag.Args()[0]
		// *output = ""
	}
	if len(os.Args) == 5 ||  len(os.Args) == 6{
		file, err = os.Open(flag.Args()[2]+".txt")
		if err != nil {
			fmt.Println("Unable to open file!")
			os.Exit(0)
		}
		r = flag.Args()[1]
		tocolor = flag.Args()[0]
	} else {
		file, err = os.Open("standard.txt")
		if err != nil {
		fmt.Println("Unable to open file!")
		os.Exit(0)
		}
	}
	
	if *color == "" && *output != "" {
		file, err = os.Open(flag.Args()[1]+".txt")
		if err != nil {
			fmt.Println("Unable to open file!")
			os.Exit(0)
		}
		r = flag.Args()[0]
		tocolor = ""
	}
	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) == 0 {
		fmt.Println("Oops! The file is empty")
		os.Exit(0)
	}
	// group lines read into 9 to display each character individually
	character := make([]string, 9)
	var result [][]string
	for i := 1; i < len(lines); i += 9 {
		end := i + 9
		if end > len(lines) {
			end = len(lines)
		}
		character = lines[i:end]
		result = append(result, character)
	}
	
	// Format ("/n") in input string
	s := strings.Replace(r, "\\n", "\n", -1)
	s = strings.Replace(s, "\\t", "    ", -1)
	asciiart.HandleLn(s, result, color, output, tocolor)
	//fmt.Println("Art has been created successfully!")
}
