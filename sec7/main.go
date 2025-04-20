package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	wordCount := 0 

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		wordCount += len(words)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("total words: %d\n", wordCount) 
}