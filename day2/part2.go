package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "strconv"
)

func main() {
    // file, _ := os.Open("example2.txt")
    file, _ := os.Open("input.txt")
    defer file.Close()

    reader := bufio.NewReader(file)
    scanner := bufio.NewScanner(reader)

    safeReports := 0
    reportsScanned := 0
    for scanner.Scan() {
        reportsScanned++
        line := scanner.Text()
        numberStrings := strings.Split(line, " ")

        numbers := make([]int, len(numberStrings))
        for i, v := range numberStrings {
            numberValue, _ := strconv.Atoi(v)
            numbers[i] = numberValue
        }

        safe := isSafe(numbers)
        if safe {
            safeReports++
            continue
        }

        // brute force: try each combination
        for index := range len(numbers) {
            dampenedNumbers := getDampenedNumbers(numbers, index)
            safe := isSafe(dampenedNumbers)
            if safe {
                safeReports++
                break
            }
        }
    }

    fmt.Printf("Safe reports: %d of %d", safeReports, reportsScanned)
}

func getDampenedNumbers(numbers []int, index int) ([]int) {
    var copyNumbers []int
    if index == 0 {
        copyNumbers = append(copyNumbers, numbers[1:]...)
    } else {
        copyNumbers = append(copyNumbers, numbers[:index]...)
        copyNumbers = append(copyNumbers, numbers[index+1:]...)
    }
    return copyNumbers
}

func isSafe(numbers []int) (bool) {
    const max = 3
    isIncreasing := numbers[0] < numbers[1]
    for i := range len(numbers) - 1 {
        if isIncreasing && numbers[i] >= numbers[i + 1] {
            return false
        }
        if isIncreasing && numbers[i + 1] - numbers[i] > max {
            return false
        }
        if !isIncreasing && numbers[i] <= numbers[i + 1] {
            return false
        }
        if !isIncreasing && numbers[i] - numbers[i + 1]  > max {
            return false
        }
    }
    return true
}