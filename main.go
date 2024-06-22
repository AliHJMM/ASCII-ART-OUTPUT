package main

import (
	"fmt"
	"os"
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