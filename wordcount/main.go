package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	//countOnlyOneFile()
	countMultipleFiles()
}

func countMultipleFiles() {
	if len(os.Args) < 2 {
		log.Println("need to provide at least one filename!")
		os.Exit(1)
	}

	// Loop over all arguments from index 1 onwards.
	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)

		// If there's an error, print it and skip to the next file.
		if err != nil {
			log.Printf("%s: %v", filename, err)
			continue
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		var wordCount int

		for scanner.Scan() {
			wordCount++
		}

		if scanner.Err() != nil {
			log.Printf("scan error %s: %v", filename, scanner.Err())
			file.Close()
			continue
		}

		file.Close()
		fmt.Printf("%s: %d\n", filename, wordCount)
	}

}

func countOnlyOneFile() {
	if len(os.Args) < 2 {
		log.Println("need to provide filename!")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	wordCount := 0
	//wordCount = scannerApproach(file)
	wordCount = scannerSplitApproach(file)

	fmt.Println("Found", wordCount, "words")
}

func scannerApproach(file io.Reader) int {

	// Open the first argument as a file.

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

	return wordCount
}

func scannerSplitApproach(file io.Reader) int {
	// Create a scanner to read from the file.
	scanner := bufio.NewScanner(file)

	// Change the split function to split on words instead of lines.
	scanner.Split(bufio.ScanWords)

	// Counter to track the running total.
	var wordCount int

	// Add 1 to the count for each word.
	for scanner.Scan() {
		wordCount++
	}

	if scanner.Err() != nil {
		log.Println(scanner.Err())
		os.Exit(1)
	}

	return wordCount
}
