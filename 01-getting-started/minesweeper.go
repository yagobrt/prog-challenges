package gettingstarted

import (
	"fmt"
	"strconv"
	"strings"
)

func updateCount(row, col int, minePositions [][]string) {
	numRows := len(minePositions)
	numCols := len(minePositions[row])
	for r := -1; r <= 1; r++ {
		if row+r >= numRows || row+r < 0 {
			continue
		}
		for c := -1; c <= 1; c++ {
			if col+c >= numCols || col+c < 0 {
				continue
			}
			s := minePositions[row+r][col+c]
			if s == "*" {
				continue
			}
			i, _ := strconv.Atoi(s)
			minePositions[row+r][col+c] = strconv.Itoa(i + 1)
		}
	}
	return
}

func generateOutput(fieldNum int, minePositions [][]string) {
	fmt.Printf("Field #%d:\n", fieldNum)
	for _, rowChars := range minePositions {
		for _, char := range rowChars {
			fmt.Printf("%s", char)
		}
		fmt.Println()
	}
	return
}

func Minesweeper() {
	numOfFields := 0
	for {
		var rows, cols int
		_, err := fmt.Scanf("%d %d\n", &rows, &cols)
		if err != nil {
			return
		}
		if rows == 0 && cols == 0 {
			return
		}

		// read input
		minePositions := make([][]string, rows)
		for r := 0; r < rows; r++ {
			var rowChars string
			fmt.Scanf("%s\n", &rowChars)

			rowChars = strings.ReplaceAll(rowChars, ".", "0")
			minePositions[r] = strings.Split(rowChars, "")
		}

		// count mines
		for rowNum, rowChars := range minePositions {
			for col, char := range rowChars {
				if char == "*" {
					updateCount(rowNum, col, minePositions)
				}
			}
		}

		// Print resutls
		numOfFields += 1
		if numOfFields > 1 {
			fmt.Println()
		}
		generateOutput(numOfFields, minePositions)
	}
}
