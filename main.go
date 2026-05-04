package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	// check the number of args (must be 2)
	if len(os.Args) < 2 {
		input = "No Arguments..."
	} else {
		if len(os.Args) > 2 {
			fmt.Println("Too Many Arguments. Printing Only the 1st")
		}
		input = os.Args[1]
	}
	// if the argument is "" exit without printing anything
	if input == "" {
		return
	}
	// read the file with the ascii art
	lines, err := readBanner("standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// split input in parts by newline characters
	parts := strings.Split(input, "\\n")

	// if the input is only "\n" print newlines
	onlyNewlines := true
	for _, part := range parts {
		if part != "" {
			onlyNewlines = false
			break
		}
	}
	if onlyNewlines {
		for i := 0; i < len(parts)-1; i++ {
			fmt.Println()
		}
		return
	}
	for _, part := range parts {
		if part == "" {
			fmt.Println()
			continue
		}

		// print 8 vertical layers for the current part
		for i := 0; i < 8; i++ {
			for _, char := range part {
				// find the starting line index in the file for this character
				// ASCII space is 32. Each char block is 9 lines (8 art + 1 empty).
				startLine := int(char-32)*9 + 1

				// print the current line (i) of the character's art
				if startLine+i < len(lines) {
					fmt.Print(lines[startLine+i])
				}
			}
			fmt.Println() // move to the next vertical line of the ASCII art
		}
	}
}

func readBanner(bannerName string) ([]string, error) {
	file, err := os.Open(bannerName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// create an empty slice of strings (lines) to store the content of the file
	var lines []string
	// read through the file line by line
	scanner := bufio.NewScanner(file)

	// add the text of every line to our lines slice
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	// return the full list of lines
	return lines, scanner.Err()
}
