package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    	file, _ := os.Open("input.txt")
    	defer file.Close()

    	result := 0
        scanner := bufio.NewScanner(bufio.NewReader(file))

        // first scan 3 times, assume at least 3 lines
        lines := [3]string{}
        for i := range 3 {
            scanner.Scan()
            lines[i] = scanner.Text()
        }

        result += countX(lines)

        for scanner.Scan() {
            lines[0] = lines[1]
            lines[1] = lines[2]
            lines[2] = scanner.Text()
            result += countX(lines)
        }

    	fmt.Printf("Result: %d\n", result)
}

func countX(lines [3]string) (int) {
    var forwardDiagonalLines []string
    var backwardDiagonalLines []string
    result := 0

    for index := range 3 {
        forwardDiagonalLines = collectForwardDiagonalLines(forwardDiagonalLines, lines[index], index)
        backwardDiagonalLines = collectBackwardDiagonalLines(backwardDiagonalLines, lines[index], index)
    }

    for i, diagonal := range forwardDiagonalLines {
        if diagonal == "SAM" || diagonal == "MAS" {
            otherDiagonalIndex := len(backwardDiagonalLines) - 1 - i
            if len(backwardDiagonalLines) <= otherDiagonalIndex {
                continue
            }
            otherDiagonal := backwardDiagonalLines[otherDiagonalIndex]
            if otherDiagonal == "SAM" || otherDiagonal == "MAS" {
                result++
            }
        }
    }

    return result
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