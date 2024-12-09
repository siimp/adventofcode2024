package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "strconv"
)

func main() {
    file, _ := os.Open("example2.txt")
    // file, _ := os.Open("input.txt")
    defer file.Close()

    reader := bufio.NewReader(file)
    scanner := bufio.NewScanner(reader)

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

        isSafe, unsafeIndex := isSafeNumbers(numbers)
        if isSafe {
            // fmt.Printf("report %d is safe\n", reportsScanned)
            safeReports++
        } else {
            // problem dampener
            // try removing index and index+1
            var dampenedNumbers []int
            dampenedNumbers = getDampenedNumbers(numbers, unsafeIndex)
            isSafeDampened, dampenedUnsafeIndex := isSafeNumbers(dampenedNumbers)
            if isSafeDampened {
                fmt.Printf("report %d is safe (problem dampened via index %d)\n", reportsScanned, unsafeIndex)
                safeReports++
            } else {
                dampenedNumbers = getDampenedNumbers(numbers, unsafeIndex + 1)
                isSafeDampened, dampenedUnsafeIndex = isSafeNumbers(dampenedNumbers)
                if isSafeDampened {
                    fmt.Printf("report %d is safe (problem dampened via index %d)\n", reportsScanned, unsafeIndex)
                    safeReports++
                } else {
                    fmt.Printf("report %d is unsafe %d. index=%d; %d dampenedUnsafeIndex=%d\n", reportsScanned, numbers, unsafeIndex, dampenedNumbers, dampenedUnsafeIndex)
                }
            }
        }

        reportsScanned++
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

func isSafeNumbers(numbers []int) (bool, int) {
    const max = 3
    var isIncreasing bool
    isSafe := true
    index := -1
    for i, _ := range numbers {
        if i == len(numbers) - 1 || !isSafe {
            break
        }
        if i == 0 {
            if numbers[0] == numbers[1] {
                isSafe = false
                index = i
                break
            }
            isIncreasing = numbers[0] < numbers[1]
        }

        if isIncreasing {
            if numbers[i] >= numbers[i+1] {
                isSafe = false
                index = i
            } else if numbers[i+1] - numbers[i] > max {
                isSafe = false
                index = i
            }
        }

        if !isIncreasing {
            if numbers[i] <= numbers[i+1] {
                isSafe = false
                index = i
            } else if numbers[i] - numbers[i+1] > max {
                isSafe = false
                index = i
            }
        }
    }
    return isSafe, index
}