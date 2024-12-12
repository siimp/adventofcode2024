package main

import (
"fmt"
"os"
"bufio"
"strings"
"strconv"
"regexp"
)

func main() {
    file, _ := os.Open("input.txt")
    defer file.Close()

    result := 0

    regex := regexp.MustCompile("mul\\(\\d*,\\d*\\)|do\\(\\)|don't\\(\\)")
    scanner := bufio.NewScanner(bufio.NewReader(file))
    isEnabled := true
    for scanner.Scan() {
        line := scanner.Text()
        instructions := regex.FindAllString(line, -1)
        for _, instruction := range instructions {
            if instruction == "do()" {
                isEnabled = true
            } else if instruction == "don't()" {
                isEnabled = false
            }

            if strings.HasPrefix(instruction, "mul") && isEnabled {
                result += calculateMul(instruction)
            }
        }
    }

    fmt.Printf("Result: %d\n", result)

}

func calculateMul(mulInstruction string) (int) {
            mulArguments := strings.Split(mulInstruction, ",")

            if len(mulArguments) < 2 {
                fmt.Println("Less than two arguments")
                return 0
            }

            firstArgumentString := mulArguments[0]
            firstArgumentStart := strings.Index(firstArgumentString, "(")
            if firstArgumentStart == -1 {
                fmt.Println("First argument not found")
                return 0
            }

            firstArgument, err := strconv.Atoi(firstArgumentString[firstArgumentStart+1:])
            if err != nil {
                fmt.Println("First argument invalid number")
                return 0
            }

            secondArgumentString := mulArguments[1]
            secondArgumentEnd := strings.Index(secondArgumentString, ")")
            if secondArgumentEnd == -1 {
                fmt.Println("Second argument not found")
                return 0
            }


            secondArgument, err := strconv.Atoi(secondArgumentString[:secondArgumentEnd])
            if err != nil {
                fmt.Println("Second argument invalid number")
                return 0
            }

            return firstArgument * secondArgument
}