package main

import "fmt"

func PrintAsciiArtRecursive(sentences []string, textFile []string) {
	// Base case: If no sentences left, return
	if len(sentences) == 0 {
		return
	}

	// Process the first sentence
	if sentences[0] != "" {
		printSentenceAsciiRecursive(sentences[0], textFile, 1)
		fmt.Print()
	} else {
		// If the sentence is empty (i.e., there was a '\n'), print an empty line
		fmt.Println()
	}

	// Recursive call to process the remaining sentences
	PrintAsciiArtRecursive(sentences[1:], textFile)
}

func printSentenceAsciiRecursive(word string, textFile []string, h int) {
	// Base case: If we've processed all the lines for this character height, return
	if h > 8 {
		return
	}

	// Print the current line for each character in the word
	for i := 0; i < len(word); i++ {
		for lineIndex, line := range textFile {
			if lineIndex == (int(word[i])-32)*9+h {
				// Print the line corresponding to the current character and line number
				fmt.Print(line)
			}
		}
	}
	fmt.Println()

	// Recursive call for the next line height
	printSentenceAsciiRecursive(word, textFile, h+1)
}
