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
	// out := flag.String("output", "banner.txt", "output file") // create a flag
	// flag.Parse()
	// outputfile, _ := os.Create(*out) // create a file from the name collected by the flag

	color := flag.String("color", "Red", "text color")// create color flag
	flag.Parse()

	// if len(os.Args) != 4 {
	// 	fmt.Println("More or less arguments passed! Please refer to README.md and try again.")
	// 	os.Exit(0)
	// }
	// file, err := os.Open(flag.Args()[1] + ".txt")
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Unable to open file!")
		os.Exit(0)
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

	r := flag.Args()[1]
	tocolor := flag.Args()[0]
	// Format ("/n") in input string
	s := strings.Replace(r, "\\n", "\n", -1)
	s = strings.Replace(s, "\\t", "    ", -1)
	asciiart.HandleLn(s, result, color, tocolor)
	//fmt.Println("Art has been created successfully!")
}
