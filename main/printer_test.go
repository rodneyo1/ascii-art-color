package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"
)

type TestCase struct {
	input          string
	bannerFile     string
	expectedOutput string
}

var testCases = []TestCase{
	{
		input:      "hello",
		bannerFile: "standard",
		expectedOutput: ` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
							   
							   `,
	},
}

var ansiEscapePattern = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func removeANSIEscapeCodes(input string) string {
	return ansiEscapePattern.ReplaceAllString(input, "")
}

func TestMainFunction(t *testing.T) {
	for _, testCase := range testCases {
		// Set command line arguments
		os.Args = []string{"program_name", testCase.input, testCase.bannerFile}
		// Redirect stdout to capture output
		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		// Run the main function
		main()
		// Reset stdout
		w.Close()
		os.Stdout = old
		// Read captured output
		var output bytes.Buffer
		buf := make([]byte, 1024)
		for {
			n, err := r.Read(buf)
			output.Write(buf[:n])
			if err != nil {
				break
			}
		}
		// Remove ANSI escape codes from output
		outputString := output.String()
		cleanedOutput := removeANSIEscapeCodes(outputString)
		// Compare output with expected string (trimmed)
		expectedOutputTrimmed := strings.TrimSpace(testCase.expectedOutput)
		outputTrimmed := strings.TrimSpace(cleanedOutput)

		// Debug prints
		fmt.Printf("Expected Output:\n%s\n", expectedOutputTrimmed)
		fmt.Printf("Actual Output:\n%s\n", outputTrimmed)

		if outputTrimmed != expectedOutputTrimmed {
			t.Errorf("Output doesn't match for input '%s' with banner '%s'.\nExpected:\n%s\nGot:\n%s", testCase.input, testCase.bannerFile, expectedOutputTrimmed, outputTrimmed)
		}
	}
}
