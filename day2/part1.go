package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "strconv"
)

func main() {
    file, _ := os.Open("input.txt")
    defer file.Close()

    reader := bufio.NewReader(file)
    scanner := bufio.NewScanner(reader)

    const max = 3
    safeReports := 0
    reportsScanned := 0
    for scanner.Scan() {
        line := scanner.Text()
        numberStrings := strings.Split(line, " ")
        numbers := make([]int, len(numberStrings))
        for i, v := range numberStrings {
            numberValue, _ := strconv.Atoi(v)
            numbers[i] = numberValue
        }

        var isIncreasing bool
        for i, _ := range numbers {
            if i == len(numbers) - 1 {
                safeReports++
                fmt.Printf("report %d is safe\n", reportsScanned)
                break
            }

            if i == 0 {
                if numbers[0] == numbers[1] {
                    break
                }
                isIncreasing = numbers[0] < numbers[1]
            }

            if isIncreasing {
                if numbers[i] >= numbers[i+1] {
                    break
                } else if numbers[i+1] - numbers[i] > max {
                    break
                }
            }

            if !isIncreasing {
                if numbers[i] <= numbers[i+1] {
                    break
                } else if numbers[i] - numbers[i+1] > max {
                    break
                }
            }
        }
        reportsScanned++
    }

    fmt.Printf("Safe reports: %d", safeReports)
}