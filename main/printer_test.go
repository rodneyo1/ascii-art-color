package main

import (
	"bytes"
	"os"
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
	{
		input:      "hello",
		bannerFile: "shadow",
		expectedOutput: `                                 
_|                _| _|          
_|_|_|     _|_|   _| _|   _|_|   
_|    _| _|_|_|_| _| _| _|    _| 
_|    _| _|       _| _| _|    _| 
_|    _|   _|_|_| _| _|   _|_|   
								 
								 `,
	},
	{
		input:      "hello",
		bannerFile: "thinkertoy",
		expectedOutput: `   
o        o o     
|        | |     
O--o o-o | | o-o 
|  | |-' | | | | 
o  o o-o o o o-o
`,
	},
	// Add more test cases as needed
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
		// Compare output with expected string (trimmed)
		expectedOutputTrimmed := strings.TrimSpace(testCase.expectedOutput)
		outputTrimmed := strings.TrimSpace(output.String())

		if outputTrimmed != expectedOutputTrimmed {
			t.Errorf("Output doesn't match for input '%s' with banner '%s'.\nExpected:\n%s\nGot:\n%s", testCase.input, testCase.bannerFile, expectedOutputTrimmed, outputTrimmed)
		}
	}
}
