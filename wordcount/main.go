package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("need to provide filename!")
		os.Exit(1)
	}

	// Open the first argument as a file.
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Create a scanner to read from the file.
	scanner := bufio.NewScanner(file)

	// Counter to track the running total.
	var wordCount int

	// For each line of the file, get the number of words, and add it to
	// the total.
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		wordCount += len(words)
	}

	if scanner.Err() != nil {
		log.Println(scanner.Err())
		os.Exit(1)
	}

	fmt.Println("Found", wordCount, "words")
}
