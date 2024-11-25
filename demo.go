package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Open the log file
	file, err := os.Open("example.log")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Define log classification
	infoRegex := regexp.MustCompile(`(?i)\binfo\b`)
	warnRegex := regexp.MustCompile(`(?i)\bwarn\b`)
	errorRegex := regexp.MustCompile(`(?i)\berror\b`)

	// Read and classify logs
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case errorRegex.MatchString(line):
			fmt.Println("[ERROR]", line)
		case warnRegex.MatchString(line):
			fmt.Println("[WARN]", line)
		case infoRegex.MatchString(line):
			fmt.Println("[INFO]", line)
		default:
			fmt.Println("[OTHER]", line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
