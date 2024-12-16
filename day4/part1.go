package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
		result += count(row)

		verticalLines = collectVerticalLines(verticalLines, row)
		forwardDiagonalLines = collectForwardDiagonalLines(forwardDiagonalLines, row, rowIndex)
		backwardDiagonalLines = collectBackwardDiagonalLines(backwardDiagonalLines, row, rowIndex)

		rowIndex++
	}

	for _, line := range verticalLines {
		result += count(line)
	}

	for _, line := range forwardDiagonalLines {
		result += count(line)
	}

	for _, line := range backwardDiagonalLines {
		result += count(line)
	}

	fmt.Printf("Result: %d\n", result)

}

func count(line string) int {
	return strings.Count(line, "XMAS") + strings.Count(line, "SAMX")
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
