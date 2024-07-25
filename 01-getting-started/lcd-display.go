package gettingstarted

import (
	"fmt"
	"strings"
)

func appendDigit(digitString []string, digit string, size int) []string {
	// If its not the first digit insert space
	for i, row := range digitString {
		if row != "" {
			digitString[i] += " "
		}
	}
	// Generate the types of lines
	horizontalBlank := fmt.Sprintf(" %s ", strings.Repeat(" ", size))
	horizontalEdge := fmt.Sprintf(" %s ", strings.Repeat("-", size))
	verticalLeftEdge := fmt.Sprintf("|%s ", strings.Repeat(" ", size))
	verticalRightEdge := fmt.Sprintf(" %s|", strings.Repeat(" ", size))
	verticalDoubleEdge := fmt.Sprintf("|%s|", strings.Repeat(" ", size))

	// For each part we change the type of edge we draw depending on the digit
	var typeOfLine string
	// Mark the start of the part we are working on
	var start int

	// TOP, index 0
	switch digit {
	case "1", "4":
		typeOfLine = horizontalBlank
	default:
		typeOfLine = horizontalEdge
	}
	digitString[0] += typeOfLine

	// UPPER-MIDDLE, index: [1, size+1]
	switch digit {
	case "5", "6": // Left side only
		typeOfLine = verticalLeftEdge
	case "1", "2", "3", "7": // Right side only
		typeOfLine = verticalRightEdge
	default:
		typeOfLine = verticalDoubleEdge
	}
	start = 1
	for i := 0; i < size; i++ {
		digitString[start+i] += typeOfLine
	}

	// MIDDLE, index: size + 1
	switch digit {
	case "1", "7", "0":
		digitString[size+1] += horizontalBlank
	default:
		digitString[size+1] += horizontalEdge
	}

	// LOWER-MIDDLE, index: [size + 2, 2*size + 1]
	switch digit {
	case "2": // Left side only
		typeOfLine = verticalLeftEdge
	case "6", "8", "0": // Double
		typeOfLine = verticalDoubleEdge
	default:
		typeOfLine = verticalRightEdge
	}
	start = size + 2
	for i := 0; i < size; i++ {
		digitString[start+i] += typeOfLine
	}

	// BOTTOM, index: 2*size + 2
	switch digit {
	case "1", "4", "7":
		typeOfLine = horizontalBlank
	default:
		typeOfLine = horizontalEdge
	}
	digitString[2*size+2] += typeOfLine
	return digitString
}

func LCDDisplay() {
	for {
		var s int
		var n string
		_, err := fmt.Scanf("%d %s\n", &s, &n)
		if err != nil {
			return
		}
		if s == 0 && n == "0" {
			return
		}
		if len(n) > 8 {
			continue
		}

		if s <= 0 || s > 10 {
			return
		}

		// Calculate size of matrix
		var outputString []string
		rows := 2*s + 3
		outputString = make([]string, rows)

		// Append digits 1 by 1
		for _, digit := range n {
			outputString = appendDigit(outputString, string(digit), s)
		}

		// Print numbers
		for _, row := range outputString {
			fmt.Println(row)
		}
		fmt.Println()
	}
}
