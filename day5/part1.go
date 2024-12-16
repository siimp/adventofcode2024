package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "slices"
    "strconv"
)

func main() {
    file, _ := os.Open("input.txt")
	defer file.Close()

	result := 0
    pagesBefore := make(map[string][]string)

    rulesSection := true
    scanner := bufio.NewScanner(bufio.NewReader(file))
    for scanner.Scan() {
        line := scanner.Text()

        if line == "" {
            rulesSection = false
            continue
        }

        if rulesSection {
            parseRule(line, pagesBefore)
        } else {
            pages := strings.Split(line, ",")
            isValid := isValidLine(pages, pagesBefore)
            if isValid {
                pageNumber, _ := strconv.Atoi(pages[len(pages)/2])
                result += pageNumber
            }
        }
    }

	fmt.Printf("Result: %d\n", result)
}

func parseRule(line string, pagesBefore map[string][]string) {
    numbers := strings.Split(line, "|")
    pagesBefore[numbers[0]] = append(pagesBefore[numbers[0]], numbers[1])
}

func isValidLine(pages []string, pagesBefore map[string][]string) (bool) {
    for index, page := range pages {
        for i := 0; i < index; i++ {
            if slices.Contains(pagesBefore[page], pages[i]) {
                return false
            }
        }
    }
    return true
}