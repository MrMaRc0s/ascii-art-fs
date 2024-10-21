package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Ensure correct argument usage
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [STRING] [BANNER] EX: go run . something standard")
		os.Exit(1)
	}

	var banner string
	if len(os.Args) == 3 {
		banner = os.Args[2]
		switch banner {
		case "standard":
			banner = "banner/standard.txt"
		case "shadow":
			banner = "banner/shadow.txt"
		case "thinkertoy":
			banner = "banner/thinkertoy.txt"
		default:
			fmt.Println("Use a valid banner template")
			os.Exit(2)
		}
	}

	if len(os.Args) == 2 {
		banner = "banner/standard.txt"
	}

	// Read the file
	bannerFile, err := os.ReadFile(banner)
	if err != nil {
		fmt.Println("There was a problem reading the banner file template")
		os.Exit(3)
	}

	// Split the lines of the file into a slice
	line := strings.Replace(string(bannerFile), "\r\n", "\n", -1)
	lines := strings.Split(line, "\n")

	// Split the input string into lines, based on literal newline characters
	sepText := strings.Split(os.Args[1], "\\n") // Handling literal "\n" from the input

	// Print the ASCII Art based on the input
	PrintAsciiArtRecursive(sepText, lines)
}
