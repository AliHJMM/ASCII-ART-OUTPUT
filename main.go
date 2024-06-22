package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 2 {
		ascii(args)
	} else if len(args) == 4 {
		ascii_output(args)
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
	}
}

func ascii(args []string) {
	if len(args) != 2 {
		fmt.Println("Error : invalid arguments")
		os.Exit(1)
	}
	txt := args[1]

	textSlice := strings.Split(txt, "\\n")

	if !charValidation(txt) {
		fmt.Println("Error : invalid char")
		os.Exit(1)
	}
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error : reading file")
		os.Exit(1)
	}
	slice := strings.Split(string(file), "\n")
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
