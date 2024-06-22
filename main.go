package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Get command-line arguments
	args := os.Args

	// Check the number of arguments
	if len(args) == 2 {
		ascii(args) // Handle standard ASCII conversion
	} else if len(args) == 4 {
		ascii_output(args) // Handle ASCII conversion with output to file
	} else {
		// Print usage instructions if the number of arguments is incorrect
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
	}
}

func ascii(args []string) {
	if len(args) != 2 {
		fmt.Println("Error: invalid arguments")
		os.Exit(1)
	}
	txt := args[1]

	// Split input text by newline character
	textSlice := strings.Split(txt, "\\n")

	// Validate characters in the input text
	if !charValidation(txt) {
		fmt.Println("Error: invalid char")
		os.Exit(1)
	}

	// Read the standard banner file
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error: reading file")
		os.Exit(1)
	}

	// Split file content by newline
	slice := strings.Split(string(file), "\n")

	// Process each line of the input text
	for j, txt := range textSlice {
		if txt != "" {
			for i := 0; i < 8; i++ {
				for _, v := range txt {
					firstLine := int(v-32)*9 + 1 + i
					fmt.Print(slice[firstLine])
				}
				fmt.Println()
			}
		} else if j != len(textSlice)-1 {
			fmt.Println("")
		}
	}
}

func charValidation(str string) bool {
	// Validate that all characters in the string are printable ASCII characters
	slice := []rune(str)
	for _, char := range slice {
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}

func ascii_output(args []string) {
	if len(args) != 4 {
		fmt.Println("Error: invalid arguments")
		os.Exit(1)
	}
	txt := args[2]
	format := args[3]
	str := ""

	// Define and parse the output flag
	outputPtr := flag.String("output", "", "Output file name")
	flag.Parse()

	if *outputPtr == "" {
		fmt.Println("Usage: go run . --output=<filename>")
		os.Exit(1)
	}

	// Split input text by newline character
	textSlice := strings.Split(txt, "\\n")

	// Validate characters in the input text
	if !charValidation(txt) {
		fmt.Println("Error: invalid char")
		os.Exit(1)
	}

	// Read the specified banner file
	file, err := os.ReadFile(format + ".txt")
	if err != nil {
		fmt.Println("Error: reading file")
		os.Exit(1)
	}

	// Split file content by newline
	slice := strings.Split(string(file), "\n")

	// Process each line of the input text and generate the ASCII art
	for j, txt := range textSlice {
		if txt != "" {
			for i := 0; i < 8; i++ {
				for _, v := range txt {
					firstLine := int(v-32)*9 + 1 + i
					str += slice[firstLine]
				}
				str += "\n"
			}
		} else if j != len(textSlice)-1 {
			str += "\n"
		}
	}

	// Write the generated ASCII art to the specified output file
	os.WriteFile(*outputPtr, []byte(str), 0o644)
}
