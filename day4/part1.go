package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const XMAS = "XMAS"
const SAMX = "SAMX"

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	result := 0

	var verticalLines []string
	var forwardDiagonalLines []string
	var backwardDiagonalLines []string

    rowIndex := 0
	scanner := bufio.NewScanner(bufio.NewReader(file))
	for scanner.Scan() {
		row := scanner.Text()
		// calculate horizontal lines
		result += strings.Count(row, XMAS)
		result += strings.Count(row, SAMX)

		verticalLines = collectVerticalLines(verticalLines, row)
		forwardDiagonalLines = collectForwardDiagonalLines(forwardDiagonalLines, row, rowIndex)
		backwardDiagonalLines = collectBackwardDiagonalLines(backwardDiagonalLines, row, rowIndex)

		rowIndex++
	}

	for _, line := range verticalLines {
		result += strings.Count(line, XMAS)
		result += strings.Count(line, SAMX)
	}

	for _, line := range forwardDiagonalLines {
		result += strings.Count(line, XMAS)
		result += strings.Count(line, SAMX)
	}

	for _, line := range backwardDiagonalLines {
		result += strings.Count(line, XMAS)
		result += strings.Count(line, SAMX)
	}

	fmt.Printf("Result: %d\n", result)

}

func collectVerticalLines(verticalLines []string, row string) []string {
	for i, c := range row {
		if len(verticalLines)-1 < i {
			verticalLines = append(verticalLines, "")
		}
		verticalLines[i] = verticalLines[i] + string(c)
	}

	return verticalLines
}

func collectForwardDiagonalLines(diagonalLines []string, row string, rowIndex int) []string {
	/*
	rowIndex=0, row-->  M  M  M
	                    M  S  A
	                    A  M  X

	diagonalLines = [M, MM, MSA, AM, A]
	*/

	for i, c := range row {
	    index := i + rowIndex
	    if len(diagonalLines) <= index {
	        diagonalLines = append(diagonalLines, "")
	    }
        diagonalLines[index] = diagonalLines[index] + string(c)
    }

	return diagonalLines
}

func collectBackwardDiagonalLines(diagonalLines []string, row string, rowIndex int) []string {
	/*
	rowIndex=0, row-->  M  M  M
	                    M  S  A
	                    A  M  X

	diagonalLines = [M, MA, MSX, MM, A]
	*/

	for i, c := range row {
	    index := len(row) - 1 - i + rowIndex
	    for len(diagonalLines) <= index {
	        diagonalLines = append(diagonalLines, "")
	    }

        diagonalLines[index] = diagonalLines[index] + string(c)
    }

	return diagonalLines
}
